package lock

import (
	"fmt"
	"time"
)

// 使用示例
func Example() {
	// 初始化分布式锁工具
	redLock := NewRedLock("localhost:6379")

	// 示例1: 非阻塞加锁
	lockName := "user:profile:update:123"
	acquired, err := redLock.TryLock(
		lockName,
		0,                    // 不阻塞
		30*time.Second,       // 锁的过期时间30秒
		100*time.Millisecond, // 重试间隔
		false,                // 不使用自动续期
	)

	if err != nil {
		fmt.Printf("获取锁失败: %v\n", err)
		return
	}

	if acquired {
		fmt.Println("成功获取锁，执行业务逻辑")
		// 执行需要加锁的业务逻辑

		// 业务逻辑完成后释放锁
		if ok, err := redLock.Unlock(lockName); !ok || err != nil {
			fmt.Printf("释放锁失败: %v\n", err)
		} else {
			fmt.Println("成功释放锁")
		}
	} else {
		fmt.Println("获取锁失败，锁已被其他进程持有")
	}

	// 示例2: 阻塞加锁
	blockingLockName := "user:order:create:456"
	acquired, err = redLock.TryLock(
		blockingLockName,
		5*time.Second,        // 最多阻塞5秒
		30*time.Second,       // 锁的过期时间30秒
		100*time.Millisecond, // 重试间隔
		false,                // 不使用自动续期
	)

	if err != nil {
		fmt.Printf("获取锁超时: %v\n", err)
		return
	}

	if acquired {
		fmt.Println("成功获取锁，执行业务逻辑")
		// 执行需要加锁的业务逻辑

		// 业务逻辑完成后释放锁
		if ok, err := redLock.Unlock(blockingLockName); !ok || err != nil {
			fmt.Printf("释放锁失败: %v\n", err)
		} else {
			fmt.Println("成功释放锁")
		}
	}

	// 示例3: 使用自动续期
	autoRenewLockName := "user:long_running_task:789"
	acquired, err = redLock.TryLock(
		autoRenewLockName,
		2*time.Second,        // 最多阻塞2秒
		10*time.Second,       // 锁的过期时间10秒
		100*time.Millisecond, // 重试间隔
		true,                 // 使用自动续期
	)

	if err != nil {
		fmt.Printf("获取锁失败: %v\n", err)
		return
	}

	if acquired {
		fmt.Println("成功获取锁，执行长时间运行的业务逻辑")

		// 模拟长时间运行的任务
		time.Sleep(30 * time.Second)

		// 业务逻辑完成后释放锁（自动停止续期）
		if ok, err := redLock.Unlock(autoRenewLockName); !ok || err != nil {
			fmt.Printf("释放锁失败: %v\n", err)
		} else {
			fmt.Println("成功释放锁")
		}
	}
}
