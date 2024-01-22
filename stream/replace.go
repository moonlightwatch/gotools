package stream

import (
	"bytes"
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

// MaskKeywords 从reader中读取数据，将words中的关键字替换为mask，并写入writer
func MaskKeywords(reader io.Reader, writer io.Writer, words [][]byte, mask []byte) error {
	if len(words) == 0 {
		_, err := io.Copy(writer, reader)
		return err
	}
	maxLen := 0
	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	localCache := make([]byte, maxLen)
	n, err := reader.Read(localCache)
	if err == io.EOF {
		_, err = writer.Write(MaskBytes(localCache[:n], words, mask))
		return err
	}
	if err != nil {
		return err
	}

	buf := make([]byte, 1)
	for {
		localCache = MaskBytes(localCache, words, mask)
		_, err := reader.Read(buf)
		if err == io.EOF {
			_, err = writer.Write(MaskBytes(localCache, words, mask))
			return err
		}
		if err != nil {
			return err
		}
		firstByte := localCache[0]
		copy(localCache, localCache[1:])
		localCache[len(localCache)-1] = buf[0]
		_, err = writer.Write([]byte{firstByte})
		if err != nil {
			return err
		}
	}
}

// MaskBytes 将byteArray中的关键字替换为mask
func MaskBytes(byteArray []byte, keywords [][]byte, mask []byte) []byte {
	for _, keyword := range keywords {
		byteArray = bytes.ReplaceAll(byteArray, keyword, bytes.Repeat(mask, len(keyword)))
	}
	return byteArray
}
