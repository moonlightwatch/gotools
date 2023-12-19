package stream_test

import (
	"bytes"
	"compress/gzip"
	"io"
	"strings"
	"testing"

	"github.com/moonlightwatch/gotools/stream"
)

func TestReplace(t *testing.T) {
	var tests = []struct {
		input    string
		old      []byte
		new      []byte
		expected string
	}{
		{"hello", []byte("l"), []byte("L"), "heLLo"},
		{"abcdefghigklmnopqrstuvwxyz", []byte("l"), []byte("L"), "abcdefghigkLmnopqrstuvwxyz"},
		{"hello", []byte("l"), []byte("LL"), "heLLLLo"},
		{"hello", []byte("admin"), []byte("LL"), "hello"},
		{"hello", []byte("llllllllll"), []byte("LL"), "hello"},
		{"hello", []byte("ho"), []byte("LL"), "hello"},
		{"hello", []byte("ll"), []byte("LL"), "heLLo"},
		{"heallo", []byte("ll"), []byte("LL"), "heaLLo"},
		{"healllo", []byte("ll"), []byte("LL"), "heaLLlo"},
		{"heallllo", []byte("ll"), []byte("LL"), "heaLLLLo"},
		{"heallllo", []byte("lll"), []byte("LL"), "heaLLlo"},
		{"hello", []byte("lo"), []byte("LL"), "helLL"},
		{"helloooo", []byte("oo"), []byte("LL"), "hellLLLL"},
		{"hellooo", []byte("oo"), []byte("LL"), "hellLLo"},
		{"hellooooo", []byte("oo"), []byte("LL"), "hellLLLLo"},
		{"hello", []byte("hello"), []byte("admin"), "admin"},
		{"hello", []byte("h"), []byte("admin"), "adminello"},
		{"hello", []byte("hellohello"), []byte("admin"), "hello"},
		{"hellohello", []byte("hello"), []byte("admin"), "adminadmin"},
	}

	for _, test := range tests {
		reader := bytes.NewBufferString(test.input)
		writer := strings.Builder{}
		err := stream.Replace(reader, &writer, []byte(test.old), []byte(test.new))
		if err != nil {
			t.Errorf("Replace(%s, %s, %s): expected %v, actual %v", test.input, test.old, test.new, nil, err)
		}
		if writer.String() != test.expected {
			t.Errorf("Replace(%s, %s, %s): expected %v, actual %v", test.input, test.old, test.new, test.expected, writer.String())
		}
	}
}

func TestReplaceGzip(t *testing.T) {
	var tests = []struct {
		input    string
		old      []byte
		new      []byte
		expected string
	}{
		{"hello", []byte("l"), []byte("L"), "heLLo"},
		{"abcdefghigklmnopqrstuvwxyz", []byte("l"), []byte("L"), "abcdefghigkLmnopqrstuvwxyz"},
		{"hello", []byte("l"), []byte("LL"), "heLLLLo"},
		{"hello", []byte("admin"), []byte("LL"), "hello"},
		{"hello", []byte("llllllllll"), []byte("LL"), "hello"},
		{"hello", []byte("ho"), []byte("LL"), "hello"},
		{"hello", []byte("ll"), []byte("LL"), "heLLo"},
		{"heallo", []byte("ll"), []byte("LL"), "heaLLo"},
		{"healllo", []byte("ll"), []byte("LL"), "heaLLlo"},
		{"heallllo", []byte("ll"), []byte("LL"), "heaLLLLo"},
		{"heallllo", []byte("lll"), []byte("LL"), "heaLLlo"},
		{"hello", []byte("lo"), []byte("LL"), "helLL"},
		{"helloooo", []byte("oo"), []byte("LL"), "hellLLLL"},
		{"hellooo", []byte("oo"), []byte("LL"), "hellLLo"},
		{"hellooooo", []byte("oo"), []byte("LL"), "hellLLLLo"},
		{"hello", []byte("hello"), []byte("admin"), "admin"},
		{"hello", []byte("h"), []byte("admin"), "adminello"},
		{"hello", []byte("hellohello"), []byte("admin"), "hello"},
		{"hellohello", []byte("hello"), []byte("admin"), "adminadmin"},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		gw := gzip.NewWriter(&buf)
		_, err := gw.Write([]byte(test.input))
		if err != nil {
			t.Errorf("gzip write error: %v", err)
		}
		gw.Flush()
		gw.Close()
		reader := bytes.NewReader(buf.Bytes())
		writer := bytes.Buffer{}
		err = stream.ReplaceGzip(reader, &writer, []byte(test.old), []byte(test.new))
		if err != nil {
			t.Errorf("ReplaceGzip(%s, %s, %s): expected %v, actual %v", test.input, test.old, test.new, nil, err)
		}

		r, err := gzip.NewReader(bytes.NewReader(writer.Bytes()))
		if err != nil {
			t.Errorf("gzip read error: %v", err)
		}
		data, err := io.ReadAll(r)
		if err != nil {
			t.Errorf("gzip read error: %v", err)
		}
		if string(data) != test.expected {
			t.Errorf("ReplaceGzip(%s, %s, %s): expected %v, actual %v", test.input, test.old, test.new, test.expected, writer.String())
		}
	}
}
