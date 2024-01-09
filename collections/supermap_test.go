package collections_test

import (
	"testing"

	"github.com/moonlightwatch/gotools/collections"
)

func TestSuperMap(t *testing.T) {

	{ // 测试SuperMap IntComparable
		m := collections.NewSuperMap[int, int]()
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

func TestSuperMapKeys(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

	if len(m.Keys()) != 3 {
		t.Errorf("SuperMap.Keys() failed. len(m.Keys()): %d", len(m.Keys()))
	}
}

func TestSuperMapValues(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

	if len(m.Values()) != 3 {
		t.Errorf("SuperMap.Values() failed. len(m.Values()): %d", len(m.Values()))
	}
}

func TestSuperMapGet(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

	if m.Get(1) != 1 {
		t.Errorf("SuperMap.Get() failed. m.Get(1): %d", m.Get(1))
	}
	if m.Get(2) != 2 {
		t.Errorf("SuperMap.Get() failed. m.Get(2): %d", m.Get(2))
	}
	if m.Get(3) != 3 {
		t.Errorf("SuperMap.Get() failed. m.Get(3): %d", m.Get(3))
	}
	if m.Get(4) != 0 {
		t.Errorf("SuperMap.Get() failed. m.Get(4): %d", m.Get(4))
	}
}

func TestSuperMapRemove(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

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
	m.Remove(3)
	if m.Size() != 0 {
		t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
	}
	if m.Contains(1) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(1): %t", m.Contains(1))
	}
	if m.Contains(2) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(2): %t", m.Contains(2))
	}
	if m.Contains(3) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(3): %t", m.Contains(3))
	}
	if m.Contains(4) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
	}
}

func TestSuperMapClear(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

	if m.IsEmpty() != false {
		t.Errorf("SuperMap.IsEmpty() failed. m.IsEmpty(): %t", m.IsEmpty())

	}

	m.Clear()
	if m.Size() != 0 {
		t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
	}
	if m.Contains(1) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(1): %t", m.Contains(1))
	}
	if m.Contains(2) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(2): %t", m.Contains(2))
	}
	if m.Contains(3) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(3): %t", m.Contains(3))
	}
	if m.Contains(4) {
		t.Errorf("SuperMap.Contains() failed. m.Contains(4): %t", m.Contains(4))
	}

	if m.IsEmpty() != true {
		t.Errorf("SuperMap.IsEmpty() failed. m.IsEmpty(): %t", m.IsEmpty())
	}
}

func TestSuperMapUpdate(t *testing.T) {

	{ // 测试SuperMap IntComparable
		m := collections.NewSuperMap[int, int]()
		m.Add(1, 1)
		m.Add(2, 2)
		m.Add(3, 3)
		if m.Size() != 3 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}

		updateMap := collections.NewSuperMap[int, int]()
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

func TestSuperMapEqual(t *testing.T) {

	{ //
		m := collections.NewSuperMap[int, int]()
		m.Add(1, 1)
		m.Add(2, 2)
		m.Add(3, 3)
		if m.Size() != 3 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}

		other := collections.NewSuperMap[int, int]()
		other.Add(1, 1)
		other.Add(2, 2)
		other.Add(3, 3)

		if !m.Equal(other) {
			t.Errorf("SuperMap.Equal() failed. m.Equal(other): %t", m.Equal(other))
		}
	}
	{
		m := collections.NewSuperMap[int, int]()
		m.Add(1, 1)
		m.Add(2, 2)
		if m.Size() != 2 {
			t.Errorf("SuperMap.Size() failed. m.Size(): %d", m.Size())
		}

		other := collections.NewSuperMap[int, int]()
		other.Add(1, 1)
		other.Add(2, 2)
		other.Add(3, 3)

		if m.Equal(other) {
			t.Errorf("SuperMap.Equal() failed. m.Equal(other): %t", m.Equal(other))
		}
	}

}

func TestSuperMapForeach(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

	m.Foreach(func(key int, value int) {
		if key != value {
			t.Errorf("SuperMap.Foreach() failed. key != value: %t", key != value)
		}
	})
}

func TestSuperMapFilter(t *testing.T) {

	m := collections.NewSuperMap[int, int]()
	m.Add(1, 1)
	m.Add(2, 2)
	m.Add(3, 3)

	newMap := m.Filter(func(key int, value int) bool {
		return key > 1
	})

	if newMap.Size() != 2 {
		t.Errorf("SuperMap.Filter() failed. newMap.Size(): %d", newMap.Size())
	}
	if !newMap.Contains(2) {
		t.Errorf("SuperMap.Filter() failed. newMap.Contains(2): %t", newMap.Contains(2))
	}
	if !newMap.Contains(3) {
		t.Errorf("SuperMap.Filter() failed. newMap.Contains(3): %t", newMap.Contains(3))
	}
	if newMap.Contains(1) {
		t.Errorf("SuperMap.Filter() failed. newMap.Contains(1): %t", newMap.Contains(1))
	}
}
