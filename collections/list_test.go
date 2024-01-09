package collections_test

import (
	"testing"

	"github.com/moonlightwatch/gotools/collections"
)

func TestList(t *testing.T) {

	{ // 测试List IntComparable
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

	}
	{ // 测试List StringComparable
		l := collections.NewList[string]()
		l.Add("1")
		l.Add("2")
		l.Add("3")
		if l.Size() != 3 {
			t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
		}
		if !l.Contains("1") {
			t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains("1"))
		}
		if !l.Contains("2") {
			t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains("2"))
		}
		if !l.Contains("3") {
			t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains("3"))
		}
		if l.Contains("4") {
			t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains("4"))
		}
		l.Remove("1")
		if l.Size() != 2 {
			t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
		}
		if l.Contains("1") {
			t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains("1"))
		}
		if !l.Contains("2") {
			t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains("2"))
		}
		if !l.Contains("3") {
			t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains("3"))
		}
		if l.Contains("4") {
			t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains("4"))
		}
		l.Remove("2")
		if l.Size() != 1 {
			t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
		}
		if l.Contains("1") {
			t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains("1"))
		}
		if l.Contains("2") {
			t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains("2"))
		}
		if !l.Contains("3") {
			t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains("3"))
		}
		if l.Contains("4") {
			t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains("4"))
		}
		l.Remove("3")
		if l.Size() != 0 {
			t.Errorf("List.Size() failed. l.Size(): %d", l.Size())
		}
		if l.Contains("1") {
			t.Errorf("List.Contains() failed. l.Contains(1): %t", l.Contains("1"))
		}
		if l.Contains("2") {
			t.Errorf("List.Contains() failed. l.Contains(2): %t", l.Contains("2"))
		}
		if l.Contains("3") {
			t.Errorf("List.Contains() failed. l.Contains(3): %t", l.Contains("3"))
		}
		if l.Contains("4") {
			t.Errorf("List.Contains() failed. l.Contains(4): %t", l.Contains("4"))
		}

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
