package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger interface {
	Log(message interface{})
}
type Info struct {
	Method    string
	Url       string
	BytesSize int
}

func (info *Info) Format() string {
	return fmt.Sprintf("METHOD: %s, URL: %s, BYTES SIZE: %d", info.Method, info.Url, info.BytesSize)
}

type stdoutLogger struct {
	Out io.Writer
}

func NewLogger() Logger {
	return &stdoutLogger{Out: os.Stdout}
}

func (l *stdoutLogger) Log(message interface{}) {

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(l.Out, "[%s] %v\n", timestamp, message)

}
