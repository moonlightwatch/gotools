package strtools

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
