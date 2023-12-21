package stream

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Replace 从reader中读取数据，替换old为new，并写入writer
func Replace(reader io.Reader, writer io.Writer, old, new []byte) error {
	oldLen := len(old)
	// 长度为0，无需替换
	if oldLen == 0 {
		_, err := io.Copy(writer, reader)
		return err
	}
	// 为读取操作初始化缓存，使用流式处理，一次读一个就行
	buf := make([]byte, 1)

	for {
		// 读取一个字节
		_, err := reader.Read(buf)
		// 读取完毕，退出
		if err == io.EOF {
			break
		}
		// 读取出错，退出
		if err != nil {
			return err
		}
		// 如果读取的字节和old的第一个字节相同
		if buf[0] == old[0] {
			// 读取old的剩余字节
			oldBuf := make([]byte, oldLen-1)
			oldBufLen, err := reader.Read(oldBuf)
			oldBuf = append(buf, oldBuf[:oldBufLen]...)
			// 读取出错，退出
			if err != nil {
				if err == io.EOF {
					// 如果读取出错是因为读取完毕，将oldBuf写入writer，退出
					_, err = writer.Write(oldBuf)
					if err != nil {
						return err
					}
					break
				}
				return err
			}
			// 如果读取的oldBuf和old相同，将new写入writer，即执行替换操作
			if bytes.EqualFold(oldBuf, old) {
				_, err = writer.Write(new)
				if err != nil {
					return err
				}
			} else { // 否则将oldBuf的第一个字节写入writer，将oldBuf剩余字节和reader组成新的reader供下次读取
				reader = io.MultiReader(bytes.NewReader(oldBuf[1:]), reader)
				_, err = writer.Write(buf)
				if err != nil {
					return err
				}
			}
		} else { // 如果读取的字节和old的第一个字节不同，将读取的字节写入writer，进行下次循环
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
