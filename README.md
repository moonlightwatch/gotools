# gotools

## 项目结构

- `func PickReaderPrefix(r io.Reader, l int) ([]byte, io.Reader, error)`: PickReaderPrefix 从reader中读取指定长度的数据，并返回读取的数据和新的reader

- collections
    - `type TTLMap struct`: TTLMap 是一个带有过期时间的map
    - `type List[T any] struct`: List 是一个泛型实现的列表，但并非协程安全的
    - `type SafeList[T any] struct`: SafeList 是一个泛型实现的列表，是协程安全的
    - `type SuperMap[K int | string, V any] struct`: SuperMap 是一个泛型实现的协程安全的Map,提供多种实用的操作函数，如：过滤、更新、比较 等
- stream
    - `func Replace(reader io.Reader, writer io.Writer, old, new []byte) error `: Replace 从reader中读取数据，替换old为new，并写入writer（流式处理，不会全部读到内存里）
    - `func ReplaceGzip(reader io.Reader, writer io.Writer, old, new []byte) error `: ReplaceGzip 从reader中读取数据，替换old为new，并写入writer，reader和writer内容都是gzip格式（流式处理，不会全部读到内存里）
- strtools
    - `ParseList(s string) []string`: ParseList 解析字符串为字符串列表