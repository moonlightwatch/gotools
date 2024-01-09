package collections_test

import (
	"testing"

	"github.com/moonlightwatch/gotools/collections"
)

func TestListAdd(t *testing.T) {
	l := collections.NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Size() != 3 {
		t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
	}
	if !l.Contains(1) {
		t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if !l.Contains(2) {
		t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if !l.Contains(3) {
		t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
}

func TestListRemove(t *testing.T) {
	l := collections.NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
	}
	if l.Contains(1) {
		t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if !l.Contains(2) {
		t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if !l.Contains(3) {
		t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
	l.Remove(2)
	if l.Size() != 1 {
		t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
	}
	if l.Contains(1) {
		t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if l.Contains(2) {
		t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if !l.Contains(3) {
		t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
	l.Remove(3)
	if l.Size() != 0 {
		t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
	}
	if l.Contains(1) {
		t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if l.Contains(2) {
		t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if l.Contains(3) {
		t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
	l.Remove(4)
}

func TestListElements(t *testing.T) {
	l := collections.NewList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Add(7)
	l.Add(8)
	l.Add(9)
	l.Add(10)
	if len(l.Elements()) != 10 {
		t.Errorf("List.Elements() failed. l.Elements(): %d", l.Elements())
	}
}
func TestDiff(t *testing.T) {
	l1 := collections.NewList[int]()
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)
	l2 := collections.NewList[int]()
	l2.Add(2)
	l2.Add(3)
	l2.Add(4)
	l3 := l1.Difference(l2)
	if l3.Size() != 1 {
		t.Errorf("List.Size() failed. l3.Size(): %d", l3.Size())
	}
	if !l3.Contains(1) {
		t.Errorf("List.Contains() failed. l3.Contains(1): %t", l3.Contains(1))
	}
	if l3.Contains(2) {
		t.Errorf("List.Contains() failed. l3.Contains(2): %t", l3.Contains(2))
	}
	if l3.Contains(3) {
		t.Errorf("List.Contains() failed. l3.Contains(3): %t", l3.Contains(3))
	}
	if l3.Contains(4) {
		t.Errorf("List.Contains() failed. l3.Contains(4): %t", l3.Contains(4))
	}
}

func TestSafeListAdd(t *testing.T) {
	l := collections.NewSafeList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Size() != 3 {
		t.Errorf("SafeList.Size() failed. l.Size(): %d", l.Size())
	}
	if !l.Contains(1) {
		t.Errorf("SafeList.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if !l.Contains(2) {
		t.Errorf("SafeList.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if !l.Contains(3) {
		t.Errorf("SafeList.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("SafeList.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
}

func TestSafeListRemove(t *testing.T) {
	l := collections.NewSafeList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Remove(1)
	if l.Size() != 2 {
		t.Errorf("SafeList.Size() failed. l.Size(): %d", l.Size())
	}
	if l.Contains(1) {
		t.Errorf("SafeList.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if !l.Contains(2) {
		t.Errorf("SafeList.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if !l.Contains(3) {
		t.Errorf("SafeList.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("SafeList.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
	l.Remove(2)
	if l.Size() != 1 {
		t.Errorf("SafeList.Size() failed. l.Size(): %d", l.Size())
	}
	if l.Contains(1) {
		t.Errorf("SafeList.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if l.Contains(2) {
		t.Errorf("SafeList.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if !l.Contains(3) {
		t.Errorf("SafeList.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("SafeList.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
	l.Remove(3)
	if l.Size() != 0 {
		t.Errorf("SafeList.Size() failed. l.Size(): %d", l.Size())
	}
	if l.Contains(1) {
		t.Errorf("SafeList.Contains() failed. l.Contains(1): %t", l.Contains(1))
	}
	if l.Contains(2) {
		t.Errorf("SafeList.Contains() failed. l.Contains(2): %t", l.Contains(2))
	}
	if l.Contains(3) {
		t.Errorf("SafeList.Contains() failed. l.Contains(3): %t", l.Contains(3))
	}
	if l.Contains(4) {
		t.Errorf("SafeList.Contains() failed. l.Contains(4): %t", l.Contains(4))
	}
	l.Remove(4)
}

func TestSafeListElements(t *testing.T) {
	l := collections.NewSafeList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Add(7)
	l.Add(8)
	l.Add(9)
	l.Add(10)
	if len(l.Elements()) != 10 {
		t.Errorf("SafeList.Elements() failed. l.Elements(): %d", l.Elements())
	}
}

func TestSafeDiff(t *testing.T) {
	l1 := collections.NewSafeList[int]()
	l1.Add(1)
	l1.Add(2)
	l1.Add(3)
	l2 := collections.NewSafeList[int]()
	l2.Add(2)
	l2.Add(3)
	l2.Add(4)
	l3 := l1.Difference(l2)
	if l3.Size() != 1 {
		t.Errorf("SafeList.Size() failed. l3.Size(): %d", l3.Size())
	}
	if !l3.Contains(1) {
		t.Errorf("SafeList.Contains() failed. l3.Contains(1): %t", l3.Contains(1))
	}
	if l3.Contains(2) {
		t.Errorf("SafeList.Contains() failed. l3.Contains(2): %t", l3.Contains(2))
	}
	if l3.Contains(3) {
		t.Errorf("SafeList.Contains() failed. l3.Contains(3): %t", l3.Contains(3))
	}
	if l3.Contains(4) {
		t.Errorf("SafeList.Contains() failed. l3.Contains(4): %t", l3.Contains(4))
	}
}
