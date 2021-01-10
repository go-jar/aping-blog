package log

import (
	"os"
	"testing"
	"time"

	"blog/conf"
)

func TestFormatLog(t *testing.T) {
	_ = conf.Init(os.Getenv("GOPATH"))

	msg := FormatAccessLog([]byte("1"), []byte("test.format.access"), []byte("test format access log"))
	t.Log(string(msg))

	msg = FormatTraceLog(&TraceLogArgs{
		TraceId:   []byte("abc"),
		Point:     []byte("test.format.trace"),
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Millisecond * 10),
		Msg:       []byte("test format trace log"),
	})
	t.Log(string(msg))
}
