package strtools

import (
	"testing"
)

func TestParseList(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"", []string{}},
		{"[]", []string{}},
		{"[1,2,3]", []string{"1", "2", "3"}},
		{"[1,2,3,]", []string{}}, // 错误的json
		{`["1","2","3"]`, []string{"1", "2", "3"}},
		{`["1","2","3",]`, []string{}}, // 错误的json
		{"1,2,3", []string{"1", "2", "3"}},
		{"1,2,3,", []string{"1", "2", "3"}},
	}

	for _, test := range tests {
		actual := ParseList(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("ParseList(%s): expected %v, actual %v", test.input, test.expected, actual)
		}
		for i, v := range actual {
			if v != test.expected[i] {
				t.Errorf("ParseList(%s): expected %v, actual %v", test.input, test.expected, actual)
			}
		}
	}
}
