package collections_test

import (
	"testing"

	"github.com/moonlightwatch/gotools/collections"
)

func TestList(t *testing.T) {

	// 测试List
	{
		l := collections.NewList[collections.IntComparable]()
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
}
