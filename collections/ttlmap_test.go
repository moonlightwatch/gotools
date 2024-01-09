package collections

import (
	"testing"
	"time"
)

func BenchmarkTTLMap(b *testing.B) {
	m := NewTTLMap(time.Second)
	for i := 0; i < b.N; i++ {
		m.Set("key", "value", time.Second)
		m.Get("key")
	}
}

func TestTTLMap(t *testing.T) {
	m := NewTTLMap(time.Second)
	m.Set("key", "value", 2*time.Second)
	if v, ok := m.Get("key"); !ok || v != "value" {
		t.Errorf("expected %v, actual %v", "value", v)
	}
	time.Sleep(2 * time.Second)
	if _, ok := m.Get("key"); ok {
		t.Errorf("expected %v, actual %v", false, ok)
	}
}

func TestTTLMapGet(t *testing.T) {
	m := NewTTLMap(10 * time.Second)
	m.Set("key", "value", 20*time.Second)
	m.Set("key01", "value", 2*time.Second)
	if v, ok := m.Get("key"); !ok || v != "value" {
		t.Errorf("expected %v, actual %v", "value", v)
	}
	if _, ok := m.Get("key2"); ok {
		t.Errorf("expected %v, actual %v", false, ok)
	}
	if v, ok := m.Get("key"); !ok || v != "value" {
		t.Errorf("expected %v, actual %v", false, ok)
	}
	time.Sleep(2 * time.Second)
	if _, ok := m.Get("key01"); ok {
		t.Errorf("expected %v, actual %v", false, ok)
	}
}

func TestTTLMapDel(t *testing.T) {
	m := NewTTLMap(time.Second)
	m.Set("key", "value", 20*time.Second)
	if v, ok := m.Get("key"); !ok || v != "value" {
		t.Errorf("expected %v, actual %v", "value", v)
	}
	if m.Len() != 1 {
		t.Errorf("expected %v, actual %v", 1, m.Len())
	}
	m.Delete("key")
	if _, ok := m.Get("key"); ok {
		t.Errorf("expected %v, actual %v", false, ok)
	}
	if m.Len() != 0 {
		t.Errorf("expected %v, actual %v", 0, m.Len())
	}
}
