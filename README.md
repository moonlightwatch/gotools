# gotools

## 项目结构

- `func PickReaderPrefix(r io.Reader, l int) ([]byte, io.Reader, error)`: PickReaderPrefix 从reader中读取指定长度的数据，并返回读取的数据和新的reader

- collections
    - `type TTLMap struct`: TTLMap 是一个带有过期时间的map
- strtools
    - `ParseList(s string) []string`: ParseList 解析字符串为字符串列表