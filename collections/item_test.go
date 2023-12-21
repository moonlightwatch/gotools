package collections_test

import (
	"testing"

	"github.com/moonlightwatch/gotools/collections"
)

func TestCompareable(t *testing.T) {
	{ // BoolComparable
		var i collections.BoolComparable = true
		var j collections.BoolComparable = true
		if !i.Euqal(j) {
			t.Errorf("BoolComparable.Euqal() failed. i: %t, j: %t", i, j)
		}
		if i != j {
			t.Errorf("BoolComparable.Euqal() failed. i: %t, j: %t", i, j)
		}
		j = false
		if i.Euqal(j) {
			t.Errorf("BoolComparable.Euqal() failed. i: %t, j: %t", i, j)
		}
		if i == j {
			t.Errorf("BoolComparable.Euqal() failed. i: %t, j: %t", i, j)
		}
	}
	{ // Int8Comparable
		var i collections.Int8Comparable = 1
		var j collections.Int8Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Int8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Int8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Int8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Int8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Int16Comparable
		var i collections.Int16Comparable = 1
		var j collections.Int16Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Int16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Int16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Int16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Int16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Int32Comparable
		var i collections.Int32Comparable = 1
		var j collections.Int32Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Int32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Int32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Int32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Int32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // IntComparable
		var i collections.IntComparable = 1
		var j collections.IntComparable = 1
		if !i.Euqal(j) {
			t.Errorf("IntComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("IntComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("IntComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("IntComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Int64Comparable
		var i collections.Int64Comparable = 1
		var j collections.Int64Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Int64Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Int64Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Int64Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Int64Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Uint8Comparable
		var i collections.Uint8Comparable = 1
		var j collections.Uint8Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Uint8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Uint8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Uint8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Uint8Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Uint16Comparable
		var i collections.Uint16Comparable = 1
		var j collections.Uint16Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Uint16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Uint16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Uint16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Uint16Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Uint32Comparable
		var i collections.Uint32Comparable = 1
		var j collections.Uint32Comparable = 1
		if !i.Euqal(j) {
			t.Errorf("Uint32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("Uint32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("Uint32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("Uint32Comparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // UintComparable
		var i collections.UintComparable = 1
		var j collections.UintComparable = 1
		if !i.Euqal(j) {
			t.Errorf("UintComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("UintComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = 2
		if i.Euqal(j) {
			t.Errorf("UintComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("UintComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // StringComparable
		var i collections.StringComparable = "1"
		var j collections.StringComparable = "1"
		if !i.Euqal(j) {
			t.Errorf("StringComparable.Euqal() failed. i: %s, j: %s", i, j)
		}
		if i != j {
			t.Errorf("StringComparable.Euqal() failed. i: %s, j: %s", i, j)
		}
		j = "2"
		if i.Euqal(j) {
			t.Errorf("StringComparable.Euqal() failed. i: %s, j: %s", i, j)
		}
		if i == j {
			t.Errorf("StringComparable.Euqal() failed. i: %s, j: %s", i, j)
		}
	}
	{ // Float32Comparable
		var i collections.Float32Comparable = 1.0
		var j collections.Float32Comparable = 1.0
		if !i.Euqal(j) {
			t.Errorf("Float32Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i != j {
			t.Errorf("Float32Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		j = 2.0
		if i.Euqal(j) {
			t.Errorf("Float32Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i == j {
			t.Errorf("Float32Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
	}
	{ // Float64Comparable
		var i collections.Float64Comparable = 1.0
		var j collections.Float64Comparable = 1.0
		if !i.Euqal(j) {
			t.Errorf("Float64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i != j {
			t.Errorf("Float64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		j = 2.0
		if i.Euqal(j) {
			t.Errorf("Float64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i == j {
			t.Errorf("Float64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
	}
	{ // ByteComparable
		var i collections.ByteComparable = '1'
		var j collections.ByteComparable = '1'
		if !i.Euqal(j) {
			t.Errorf("ByteComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("ByteComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = '2'
		if i.Euqal(j) {
			t.Errorf("ByteComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("ByteComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // RuneComparable
		var i collections.RuneComparable = '1'
		var j collections.RuneComparable = '1'
		if !i.Euqal(j) {
			t.Errorf("RuneComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i != j {
			t.Errorf("RuneComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		j = '2'
		if i.Euqal(j) {
			t.Errorf("RuneComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
		if i == j {
			t.Errorf("RuneComparable.Euqal() failed. i: %d, j: %d", i, j)
		}
	}
	{ // Complex64Comparable
		var i collections.Complex64Comparable = 1.0
		var j collections.Complex64Comparable = 1.0
		if !i.Euqal(j) {
			t.Errorf("Complex64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i != j {
			t.Errorf("Complex64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		j = 2.0
		if i.Euqal(j) {
			t.Errorf("Complex64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i == j {
			t.Errorf("Complex64Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
	}
	{ // Complex128Comparable
		var i collections.Complex128Comparable = 1.0
		var j collections.Complex128Comparable = 1.0
		if !i.Euqal(j) {
			t.Errorf("Complex128Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i != j {
			t.Errorf("Complex128Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		j = 2.0
		if i.Euqal(j) {
			t.Errorf("Complex128Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
		if i == j {
			t.Errorf("Complex128Comparable.Euqal() failed. i: %f, j: %f", i, j)
		}
	}

}
