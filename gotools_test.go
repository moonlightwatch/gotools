package gotools

import (
	"bytes"
	"io"
	"testing"
)

func TestPickReaderPrefix(t *testing.T) {
	var tests = []struct {
		input    string
		l        int
		expected string
	}{
		{"hello", 3, "hel"},
		{"hello", 5, "hello"},
		{"hello", 6, "hello"},
		{"hello", 0, ""},
		{"hello", -1, ""},
	}

	for _, test := range tests {
		buf, r, err := PickReaderPrefix(bytes.NewBufferString(test.input), test.l)
		if err != nil {
			t.Errorf("PickReaderPrefix(%s, %d): expected %v, actual %v", test.input, test.l, test.expected, err)
		}
		if string(buf) != test.expected {
			t.Errorf("PickReaderPrefix(%s, %d): expected %v, actual %v", test.input, test.l, test.expected, string(buf))
		}
		if r == nil {
			t.Errorf("PickReaderPrefix(%s, %d): expected %v, actual %v", test.input, test.l, true, false)
		}
		content, err := io.ReadAll(r)
		if err != nil {
			t.Errorf("PickReaderPrefix(%s, %d): expected %v, actual %v", test.input, test.l, nil, err)
		}
		if string(content) != test.input {
			t.Errorf("PickReaderPrefix(%s, %d): expected %v, actual %v", test.input, test.l, test.input, string(content))
		}
	}
}
