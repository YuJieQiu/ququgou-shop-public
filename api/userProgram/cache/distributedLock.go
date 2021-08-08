package cache

import (
	"time"
)

//锁管理
func Lock(lockKey string, val interface{}, timeout time.Duration) bool {
	//timeoutDuration := 30 * time.Second
	if WebCache.Enabled {
		b := WebCache.SetNX(DistributedLock+lockKey, val, timeout)

		return b
	} else {
		return true
	}

}

//取消锁
func UnLock(lockKey string) {
	WebCache.Del(DistributedLock + lockKey)
	return
}
