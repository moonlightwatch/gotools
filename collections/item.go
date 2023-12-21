package collections

// Comparable 是一个可比较的接口
type Comparable interface {
	Euqal(Comparable) bool
}

// bool|int8|int16|int32|int64|uint8|byte|rune|uint16|uint32|uint64|float32|float64|string|complex64|complex128

type BoolComparable bool

func (b BoolComparable) Euqal(c Comparable) bool {
	return b == c
}

type Int8Comparable int8

func (i Int8Comparable) Euqal(c Comparable) bool {
	return i == c
}

type Int16Comparable int16

func (i Int16Comparable) Euqal(c Comparable) bool {
	return i == c
}

type Int32Comparable int32

func (i Int32Comparable) Euqal(c Comparable) bool {
	return i == c
}

type IntComparable int

func (i IntComparable) Euqal(c Comparable) bool {
	return i == c
}

type Int64Comparable int64

func (i Int64Comparable) Euqal(c Comparable) bool {
	return i == c
}

type Uint8Comparable uint8

func (u Uint8Comparable) Euqal(c Comparable) bool {
	return u == c
}

type ByteComparable byte

func (b ByteComparable) Euqal(c Comparable) bool {
	return b == c
}

type RuneComparable rune

func (r RuneComparable) Euqal(c Comparable) bool {
	return r == c
}

type Uint16Comparable uint16

func (u Uint16Comparable) Euqal(c Comparable) bool {
	return u == c
}

type Uint32Comparable uint32

func (u Uint32Comparable) Euqal(c Comparable) bool {
	return u == c
}

type UintComparable uint

func (u UintComparable) Euqal(c Comparable) bool {
	return u == c
}

type Uint64Comparable uint64

func (u Uint64Comparable) Euqal(c Comparable) bool {
	return u == c
}

type Float32Comparable float32

func (f Float32Comparable) Euqal(c Comparable) bool {
	return f == c
}

type Float64Comparable float64

func (f Float64Comparable) Euqal(c Comparable) bool {
	return f == c
}

type StringComparable string

func (s StringComparable) Euqal(c Comparable) bool {
	return s == c
}

type Complex64Comparable complex64

func (c Complex64Comparable) Euqal(cc Comparable) bool {
	return c == cc
}

type Complex128Comparable complex128

func (c Complex128Comparable) Euqal(cc Comparable) bool {
	return c == cc
}
