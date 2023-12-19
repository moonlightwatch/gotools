package stream

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Replace 从reader中读取数据，替换old为new，并写入writer
func Replace(reader io.Reader, writer io.Writer, old, new []byte) error {
	oldLen := len(old)
	// 长度为0，说明无需替换
	if oldLen == 0 {
		_, err := io.Copy(writer, reader)
		return err
	}

	buf := make([]byte, 1)

	for {
		_, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if buf[0] == old[0] {
			oldBuf := make([]byte, oldLen-1)
			oldBufLen, err := reader.Read(oldBuf)
			oldBuf = append(buf, oldBuf[:oldBufLen]...)
			if err != nil {
				if err == io.EOF {
					_, err = writer.Write(oldBuf)
					if err != nil {
						return err
					}
					break
				}
				return err
			}
			if bytes.EqualFold(oldBuf, old) {
				_, err = writer.Write(new)
				if err != nil {
					return err
				}
			} else {
				reader = io.MultiReader(bytes.NewReader(oldBuf[1:]), reader)
				_, err = writer.Write(buf)
				if err != nil {
					return err
				}
			}
		} else {
			_, err = writer.Write(buf)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// ReplaceGzip 从reader中读取数据，替换old为new，并写入writer，reader和writer都是gzip格式
func ReplaceGzip(reader io.Reader, writer io.Writer, old, new []byte) error {
	oldLen := len(old)
	var r io.Reader
	r, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	gw := gzip.NewWriter(writer)
	defer gw.Close()

	buf := make([]byte, 1)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if buf[0] == old[0] {
			oldBuf := make([]byte, oldLen-1)
			oldBufLen, err := r.Read(oldBuf)
			oldBuf = append(buf, oldBuf[:oldBufLen]...)
			if err != nil {
				if err == io.EOF {
					_, err = gw.Write(oldBuf)
					if err != nil {
						return err
					}
					break
				}
				return err
			}
			if bytes.EqualFold(oldBuf, old) {
				_, err = gw.Write(new)
				if err != nil {
					return err
				}
			} else {
				r = io.MultiReader(bytes.NewReader(oldBuf[1:]), r)
				_, err = gw.Write(buf)
				if err != nil {
					return err
				}
			}
		} else {
			_, err = gw.Write(buf)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
