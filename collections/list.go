package collections

import (
	"reflect"
	"sync"
)

// List 是一个泛型实现的列表，但并非协程安全的
// 如需协程安全，请使用SafeList
type List[T any] struct {
	elements []T
}

func NewList[T any]() *List[T] {
	return &List[T]{
		elements: make([]T, 0),
	}
}

// Add 添加一个元素
func (l *List[T]) Add(element T) {
	l.elements = append(l.elements, element)
}

// Remove 删除一个元素
func (l *List[T]) Remove(element T) {
	for i, e := range l.elements {
		if reflect.DeepEqual(e, element) {
			l.elements = append(l.elements[:i], l.elements[i+1:]...)
			return
		}
	}
}

// Get 获取指定位置的元素
func (l *List[T]) Get(index int) T {
	return l.elements[index]
}

// Size 获取列表长度
func (l *List[T]) Size() int {
	return len(l.elements)
}

// Contains 判断列表是否包含某个元素
func (l *List[T]) Contains(element T) bool {
	for _, e := range l.elements {
		if reflect.DeepEqual(e, element) {
			return true
		}
	}
	return false
}

// Elements 获取列表中的所有元素
func (l *List[T]) Elements() []T {
	return l.elements
}

// Difference 取差集，返回一个新的列表，其中元素是l中存在，但other中不存在的元素
func (l *List[T]) Difference(other *List[T]) *List[T] {
	result := NewList[T]()
	for _, e := range l.elements {
		if !other.Contains(e) {
			result.Add(e)
		}
	}
	return result
}

// SafeList 是一个泛型实现的列表，是协程安全的
type SafeList[T any] struct {
	elements []T
	mux      *sync.RWMutex
}

func NewSafeList[T any]() *SafeList[T] {
	return &SafeList[T]{
		elements: make([]T, 0),
		mux:      &sync.RWMutex{},
	}
}

// Add 添加一个元素
func (l *SafeList[T]) Add(element T) {
	l.mux.Lock()
	defer l.mux.Unlock()
	l.elements = append(l.elements, element)
}

// Remove 删除一个元素
func (l *SafeList[T]) Remove(element T) {
	l.mux.Lock()
	defer l.mux.Unlock()
	for i, e := range l.elements {
		if reflect.DeepEqual(e, element) {
			l.elements = append(l.elements[:i], l.elements[i+1:]...)
			return
		}
	}
}

// Get 获取指定位置的元素
func (l *SafeList[T]) Get(index int) T {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return l.elements[index]
}

// Size 获取列表长度
func (l *SafeList[T]) Size() int {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return len(l.elements)
}

// Contains 判断列表是否包含某个元素
func (l *SafeList[T]) Contains(element T) bool {
	l.mux.RLock()
	defer l.mux.RUnlock()
	for _, e := range l.elements {
		if reflect.DeepEqual(e, element) {
			return true
		}
	}
	return false
}

// Elements 获取列表中的所有元素
func (l *SafeList[T]) Elements() []T {
	l.mux.RLock()
	defer l.mux.RUnlock()
	return l.elements
}

// Difference 取差集，返回一个新的列表，其中元素l中存在，但other中不存在的元素
func (l *SafeList[T]) Difference(other *SafeList[T]) *SafeList[T] {
	l.mux.RLock()
	defer l.mux.RUnlock()
	other.mux.RLock()
	defer other.mux.RUnlock()
	result := NewSafeList[T]()
	for _, e := range l.elements {
		if !other.Contains(e) {
			result.Add(e)
		}
	}
	return result
}
