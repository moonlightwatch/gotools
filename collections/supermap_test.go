package collections_test

import (
	"testing"

	"github.com/moonlightwatch/gotools/collections"
)

func TestSuperMap(t *testing.T) {

	{ // 测试SuperMap IntComparable
		m := collections.NewSuperMap[int, collections.IntComparable]()
		m.Add(1, 1)
		m.Add(2, 2)
		m.Add(3, 3)
		if m.Size() != 3 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}
		if !m.Contains(1) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(1): %t", m.Contains(1))
		}
		if !m.Contains(2) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(2): %t", m.Contains(2))
		}
		if !m.Contains(3) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(3): %t", m.Contains(3))
		}
		if m.Contains(4) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
		}
		m.Remove(1)
		if m.Size() != 2 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}
		if m.Contains(1) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(1): %t", m.Contains(1))
		}
		if !m.Contains(2) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(2): %t", m.Contains(2))
		}
		if !m.Contains(3) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(3): %t", m.Contains(3))
		}
		if m.Contains(4) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
		}
		m.Remove(2)
		if m.Size() != 1 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}
		if m.Contains(1) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(1): %t", m.Contains(1))
		}
		if m.Contains(2) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(2): %t", m.Contains(2))
		}
		if !m.Contains(3) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(3): %t", m.Contains(3))
		}
		if m.Contains(4) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
		}
	}
}

func TestSuperMapUpdate(t *testing.T) {

	{ // 测试SuperMap IntComparable
		m := collections.NewSuperMap[int, collections.IntComparable]()
		m.Add(1, 1)
		m.Add(2, 2)
		m.Add(3, 3)
		if m.Size() != 3 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}

		updateMap := collections.NewSuperMap[int, collections.IntComparable]()
		updateMap.Add(1, 2)
		updateMap.Add(4, 4)

		newMap := m.Update(updateMap)
		if m.Size() != 4 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}
		if !newMap.Contains(1) || newMap.Get(1) != 2 {
			t.Errorf("SuperMap.Contains() failed. m.Contains(1): %t", m.Contains(1))
		}
		if !newMap.Contains(4) || newMap.Get(4) != 4 {
			t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
		}
		if newMap.Contains(5) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(5): %t", m.Contains(5))
		}
		if !m.Contains(4) {
			t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
		}
	}
}
