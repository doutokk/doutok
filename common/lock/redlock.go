package lock

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

// RedLock 分布式锁工具
type RedLock struct {
	rs      *redsync.Redsync
	mutexes map[string]*redsync.Mutex
	// 新增字段，用于跟踪自动续期的上下文和取消函数
	renewCtx    map[string]context.Context
	renewCancel map[string]context.CancelFunc
	mu          sync.Mutex // 保护并发访问maps
}

// NewRedLock 创建一个新的分布式锁工具实例
func NewRedLock(redisAddr string) *RedLock {
	// 创建Redis客户端
	client := goredislib.NewClient(&goredislib.Options{
		Addr: redisAddr,
	})

	// 创建连接池
	pool := goredis.NewPool(client)

	// 创建redsync实例
	rs := redsync.New(pool)

	return &RedLock{
		rs:          rs,
		mutexes:     make(map[string]*redsync.Mutex),
		renewCtx:    make(map[string]context.Context),
		renewCancel: make(map[string]context.CancelFunc),
	}
}

// TryLock 尝试获取锁，可以设置是否阻塞等待
// lockName: 锁的名称
// timeout: 阻塞等待超时时间，如果为0则表示不阻塞
// expiry: 锁的过期时间，防止死锁
// retryDelay: 重试间隔
// autoRenew: 是否启用自动续期
// 返回值: 是否获取成功，错误信息
func (r *RedLock) TryLock(lockName string, timeout, expiry, retryDelay time.Duration, autoRenew bool) (bool, error) {
	mutex := r.rs.NewMutex(lockName,
		redsync.WithExpiry(expiry),
		redsync.WithTries(1), // 非阻塞模式下只尝试一次
	)

	// 如果不需要阻塞，直接尝试获取锁
	if timeout == 0 {
		err := mutex.Lock()
		if err != nil {
			return false, err
		}

		// 存储获取的锁以便后续解锁
		r.mu.Lock()
		r.mutexes[lockName] = mutex
		r.mu.Unlock()

		// 如果需要自动续期，启动自动续期协程
		if autoRenew {
			r.StartAutoRenew(lockName, expiry)
		}
		return true, nil
	}

	// 阻塞模式下，在超时范围内多次尝试
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			// 超时退出
			return false, ctx.Err()
		default:
			err := mutex.Lock()
			if err == nil {
				// 获取锁成功
				r.mu.Lock()
				r.mutexes[lockName] = mutex
				r.mu.Unlock()

				// 如果需要自动续期，启动自动续期协程
				if autoRenew {
					r.StartAutoRenew(lockName, expiry)
				}
				return true, nil
			}

			// 获取失败，等待一段时间后重试
			time.Sleep(retryDelay)
		}
	}
}

// StartAutoRenew 启动自动续期
// lockName: 锁的名称
// expiry: 锁的过期时间
func (r *RedLock) StartAutoRenew(lockName string, expiry time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 如果已经有续期协程在运行，先停止它
	if cancel, exists := r.renewCancel[lockName]; exists {
		cancel()
	}

	// 创建新的上下文和取消函数
	ctx, cancel := context.WithCancel(context.Background())
	r.renewCtx[lockName] = ctx
	r.renewCancel[lockName] = cancel

	mutex, exists := r.mutexes[lockName]
	if !exists {
		return
	}

	// 启动协程进行自动续期
	go func() {
		// 设置续期周期为过期时间的一半，确保有足够时间续期
		interval := expiry / 2
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				// 上下文被取消，停止续期
				return
			case <-ticker.C:
				// 尝试续期
				if ok, err := mutex.Extend(); err != nil || !ok {
					// 续期失败，可能锁已经丢失
					fmt.Printf("自动续期失败: %v\n", err)
					return
				}
				// 续期成功
			}
		}
	}()
}

// StopAutoRenew 停止自动续期
// lockName: 锁的名称
func (r *RedLock) StopAutoRenew(lockName string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cancel, exists := r.renewCancel[lockName]; exists {
		cancel()
		delete(r.renewCtx, lockName)
		delete(r.renewCancel, lockName)
	}
}

// Unlock 释放指定名称的锁
// lockName: 锁的名称
// 返回值: 是否释放成功，错误信息
func (r *RedLock) Unlock(lockName string) (bool, error) {
	r.mu.Lock()
	mutex, exists := r.mutexes[lockName]
	r.mu.Unlock()

	if !exists {
		return false, nil
	}

	// 停止自动续期
	r.StopAutoRenew(lockName)

	success, err := mutex.Unlock()
	if success || err == nil {
		r.mu.Lock()
		delete(r.mutexes, lockName)
		r.mu.Unlock()
	}

	return success, err
}

// GetLock 根据锁名称获取已经存在的锁对象
func (r *RedLock) GetLock(lockName string) *redsync.Mutex {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.mutexes[lockName]
}
