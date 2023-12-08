package collections

import (
	"runtime"
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expireTime int64
}

// TTLMap 是一个带有过期时间的map
type TTLMap struct {
	data map[string]item
	mux  *sync.RWMutex
	stop chan bool
}

// NewTTLMap 创建一个TTLMap
// cleanTick 清理过期数据的时间间隔
func NewTTLMap(cleanTick time.Duration) *TTLMap {
	m := &TTLMap{
		data: make(map[string]item),
		mux:  &sync.RWMutex{},
		stop: make(chan bool),
	}
	go m.clean(cleanTick)
	runtime.SetFinalizer(m, func(m *TTLMap) { m.stop <- true })
	return m
}

func (m *TTLMap) clean(tick time.Duration) {
	ticker := time.NewTicker(tick)
	for {
		select {
		case <-ticker.C:
			now := time.Now().UnixNano()
			m.mux.Lock()
			for k, v := range m.data {
				if v.expireTime < now {
					delete(m.data, k)
				}
			}
			m.mux.Unlock()
		case <-m.stop:
			return
		}
	}

}

// Set 设置一个key-value对，并指定过期时间
func (m *TTLMap) Set(key string, value interface{}, expire time.Duration) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.data[key] = item{value: value, expireTime: time.Now().Add(expire).UnixNano()}
}

// Get 获取一个key-value对
func (m *TTLMap) Get(key string) (interface{}, bool) {
	now := time.Now().UnixNano()
	m.mux.RLock()
	defer m.mux.RUnlock()
	if v, ok := m.data[key]; ok {
		if v.expireTime > now {
			return v.value, true
		}
	}
	return nil, false
}

// Delete 删除一个key-value对
func (m *TTLMap) Delete(key string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	delete(m.data, key)
}

// Len 返回map的长度
func (m *TTLMap) Len() int {
	m.mux.RLock()
	defer m.mux.RUnlock()
	return len(m.data)
}
