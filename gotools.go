package gotools

import (
	"bytes"
	"io"
)

// PickReaderPrefix 从reader中读取指定长度的数据，并返回读取的数据和新的reader
func PickReaderPrefix(r io.Reader, l int) ([]byte, io.Reader, error) {
	if l <= 0 {
		return []byte{}, r, nil
	}
	buf := make([]byte, l)
	n, err := r.Read(buf)
	if err == io.EOF {
		return buf[:n], bytes.NewReader(buf[:n]), nil
	}
	newReader := io.MultiReader(bytes.NewReader(buf[:n]), r)
	return buf[:n], newReader, err
}
