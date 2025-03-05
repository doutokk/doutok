package lock

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

// RedLock 分布式锁工具
type RedLock struct {
	rs      *redsync.Redsync
	mutexes map[string]*redsync.Mutex
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
		rs:      rs,
		mutexes: make(map[string]*redsync.Mutex),
	}
}

// TryLock 尝试获取锁，可以设置是否阻塞等待
// lockName: 锁的名称
// timeout: 阻塞等待超时时间，如果为0则表示不阻塞
// expiry: 锁的过期时间，防止死锁
// retryDelay: 重试间隔
// 返回值: 是否获取成功，错误信息
func (r *RedLock) TryLock(lockName string, timeout, expiry, retryDelay time.Duration) (bool, error) {
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
		r.mutexes[lockName] = mutex
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
				r.mutexes[lockName] = mutex
				return true, nil
			}

			// 获取失败，等待一段时间后重试
			time.Sleep(retryDelay)
		}
	}
}

// Unlock 释放指定名称的锁
// lockName: 锁的名称
// 返回值: 是否释放成功，错误信息
func (r *RedLock) Unlock(lockName string) (bool, error) {
	mutex, exists := r.mutexes[lockName]
	if !exists {
		return false, nil
	}

	success, err := mutex.Unlock()
	if success || err == nil {
		delete(r.mutexes, lockName)
	}

	return success, err
}

// GetLock 根据锁名称获取已经存在的锁对象
func (r *RedLock) GetLock(lockName string) *redsync.Mutex {
	return r.mutexes[lockName]
}
