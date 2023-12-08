package strtools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseList 解析字符串为字符串列表
// 1. 如果字符串以 [ 开头，则按照json解析
// 2. 否则按照逗号分隔
// 3. 如果解析失败，则返回空列表
func ParseList(s string) []string {
	if s == "" {
		return []string{}
	}
	result := []string{}
	if s[0] == '[' {
		temps := []interface{}{}
		err := json.Unmarshal([]byte(s), &temps)
		if err != nil {
			return result
		}
		for _, item := range temps {
			result = append(result, fmt.Sprintf("%v", item))
		}
	} else {
		for _, item := range strings.Split(s, ",") {
			if item == "" {
				continue
			}
			result = append(result, item)
		}
	}
	return result
}
