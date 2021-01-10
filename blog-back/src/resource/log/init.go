package log

import (
	"github.com/go-jar/golog"

	"blog/conf"
)

var accessLogWriter golog.IWriter
var AccessLogger golog.ILogger

var traceLogWriter golog.IWriter
var TraceLogger golog.ILogger

var NoopLogger golog.ILogger = new(golog.NoopLogger)

var TestLogger, _ = golog.NewConsoleLogger(golog.LevelDebug)

func InitLog(systemName string) error {
	if conf.CommonConf.IsDev {
		accessLogWriter = golog.NewConsoleWriter()
	} else {
		fw, err := golog.NewFileWriter(conf.LogConf.RootPath+"/"+systemName+"_access.log", conf.LogConf.BufSize)
		if err != nil {
			return err
		}
		accessLogWriter = golog.NewAsyncWriter(fw, conf.LogConf.AsyncQueueSize)
	}
	AccessLogger = NewLogger(accessLogWriter)

	fw, err := golog.NewFileWriter(conf.LogConf.RootPath+"/"+systemName+"_trace.log", conf.LogConf.BufSize)
	if err != nil {
		return err
	}
	traceLogWriter = golog.NewAsyncWriter(fw, conf.LogConf.AsyncQueueSize)
	TraceLogger = NewLogger(traceLogWriter)

	return nil
}

func NewLogger(writer golog.IWriter) golog.ILogger {
	return golog.NewSimpleLogger(writer, golog.NewFileInfoFormat(2)).SetLevel(conf.LogConf.Level)
}

func FreeLog() {
	accessLogWriter.Free()
	traceLogWriter.Free()
}
