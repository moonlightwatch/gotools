package tokenbucket

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity        int64
	availableTokens int64
	fillInterval    time.Duration
	lastTokenTime   time.Time

	tokenLock sync.RWMutex
}

// NewTokenBucket 创建一个令牌桶，桶内自带锁，协程安全。capacity为桶的容量，fillInterval为每次填充桶的时间间隔
func NewTokenBucket(capacity int64, fillInterval time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:        capacity,
		availableTokens: capacity,
		fillInterval:    fillInterval,
		lastTokenTime:   time.Now(),
		tokenLock:       sync.RWMutex{},
	}
}

// Take 从桶中取出一个令牌，如果桶中没有令牌，则阻塞直到桶中有令牌
func (tb *TokenBucket) Take() {
	tb.TakeNum(1)
}

// TakeNum 从桶中取出num个令牌，如果桶中没有足够的令牌，则阻塞直到桶中有足够的令牌，若获取的令牌数大于桶的容量，则获取桶的容量个令牌。返回实际获取的令牌数
func (tb *TokenBucket) TakeNum(num int64) int64 {
	tb.refill()
	if num > tb.capacity {
		num = tb.capacity
	}
	for tb.AvailableTokens() < num {
		time.Sleep(tb.fillInterval)
		tb.refill()
	}
	tb.tokenLock.Lock()
	defer tb.tokenLock.Unlock()
	tb.availableTokens -= num
	return num
}

// TryTake 从桶中取出一个令牌，如果桶中没有令牌，则返回false
func (tb *TokenBucket) TryTake() bool {
	return tb.TryTakeNum(1)
}

// TryTakeNum 从桶中取出num个令牌，如果桶中没有足够的令牌，则返回false
func (tb *TokenBucket) TryTakeNum(num int64) bool {
	tb.refill()
	tb.tokenLock.Lock()
	defer tb.tokenLock.Unlock()
	if tb.availableTokens < num {
		return false
	}
	tb.availableTokens -= num
	return true
}

// refill 填充桶
func (tb *TokenBucket) refill() {
	now := time.Now()
	tb.tokenLock.Lock()
	defer tb.tokenLock.Unlock()
	// 若桶中令牌数小于桶的容量，则填充桶
	if tb.availableTokens < tb.capacity {
		// 计算从上次填充桶到现在的时间间隔内，桶中应该填充的令牌数
		tb.availableTokens += (now.UnixNano() - tb.lastTokenTime.UnixNano()) / tb.fillInterval.Nanoseconds()
		// 若桶中令牌数大于桶的容量，则将桶中令牌数设置为桶的容量
		if tb.availableTokens > tb.capacity {
			tb.availableTokens = tb.capacity
		}
	}
	tb.lastTokenTime = now
	fmt.Println(tb.availableTokens)
}

// AvailableTokens 返回桶中可用的令牌数量
func (tb *TokenBucket) AvailableTokens() int64 {
	tb.refill()
	tb.tokenLock.RLock()
	defer tb.tokenLock.RUnlock()
	return tb.availableTokens
}

// Capacity 返回桶的容量
func (tb *TokenBucket) Capacity() int64 {
	return tb.capacity
}

// FillInterval 返回填充桶的时间间隔
func (tb *TokenBucket) FillInterval() time.Duration {
	return tb.fillInterval
}

// LastTokenTime 返回上次填充桶的时间
func (tb *TokenBucket) LastTokenTime() time.Time {
	return tb.lastTokenTime
}

// Reset 重置桶，清空桶中的令牌
func (tb *TokenBucket) Reset() {
	tb.tokenLock.Lock()
	defer tb.tokenLock.Unlock()
	tb.availableTokens = tb.capacity
	tb.lastTokenTime = time.Now()
}
