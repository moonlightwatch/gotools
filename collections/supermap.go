package collections

import "sync"

type SuperMap[K int | string, V Comparable] struct {
	elements map[K]V
	rwlock   *sync.RWMutex
}

func NewSuperMap[K int | string, V Comparable]() *SuperMap[K, V] {
	return &SuperMap[K, V]{
		elements: make(map[K]V),
		rwlock:   &sync.RWMutex{},
	}
}

// Add 添加一个元素
func (m *SuperMap[K, V]) Add(key K, value V) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.elements[key] = value
}

// Remove 删除一个元素
func (m *SuperMap[K, V]) Remove(key K) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.elements, key)
}

// Get 获取指定位置的元素
func (m *SuperMap[K, V]) Get(key K) V {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return m.elements[key]
}

// Size 获取列表长度
func (m *SuperMap[K, V]) Size() int {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return len(m.elements)
}

// Contains 判断列表是否包含某个元素
func (m *SuperMap[K, V]) Contains(key K) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	_, ok := m.elements[key]
	return ok
}

// Keys 获取所有的key
func (m *SuperMap[K, V]) Keys() []K {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	keys := make([]K, 0)
	for k := range m.elements {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取所有的value
func (m *SuperMap[K, V]) Values() []V {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	values := make([]V, 0)
	for _, v := range m.elements {
		values = append(values, v)
	}
	return values
}

// Clear 清空列表
func (m *SuperMap[K, V]) Clear() {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.elements = make(map[K]V)
}

// IsEmpty 判断列表是否为空
func (m *SuperMap[K, V]) IsEmpty() bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	return len(m.elements) == 0
}

// ForEach 遍历列表
func (m *SuperMap[K, V]) ForEach(fn func(key K, value V)) {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	for k, v := range m.elements {
		fn(k, v)
	}
}

// Filter 过滤列表
func (m *SuperMap[K, V]) Filter(fn func(key K, value V) bool) *SuperMap[K, V] {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	newMap := NewSuperMap[K, V]()
	for k, v := range m.elements {
		if fn(k, v) {
			newMap.Add(k, v)
		}
	}
	return newMap
}

// Equal 判断两个列表是否相等
func (m *SuperMap[K, V]) Equal(other *SuperMap[K, V]) bool {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	if len(m.elements) != len(other.elements) {
		return false
	}
	for k, v := range m.elements {
		if other.Get(k).Euqal(v) {
			return false
		}
	}
	return true
}

// Update 更新列表，将other中的元素添加到m中，并返回已更新的kv。Key不存在或者Value不相等的情况下才会更新
func (m *SuperMap[K, V]) Update(other *SuperMap[K, V]) *SuperMap[K, V] {
	newMap := NewSuperMap[K, V]()
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	for k, v := range other.elements {
		value, contains := m.elements[k]
		if !contains || !value.Euqal(v) {
			m.elements[k] = v
			newMap.Add(k, v)
		}
	}
	return newMap
}
