package logger

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Logging level.
const (
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)
const defaultIo = 100
const timeLine = "2006-01-02 15:04:05"

// Logger represents a simple logger with level.
// The underlying logger is the standard Go logging "log".
type Logger struct {
	level         int
	isOutFile     bool
	outFile       map[int]io.Writer
	closer        []func()
	DefaultWriter io.Writer
}

func NewFileLogger(logLevel int, filename string, outLevel []int) (*Logger, error) {
	err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("MkdirAll %s faild ", filename))
	}

	// init file
	outLevel = append(outLevel, defaultIo)
	l := len(outLevel)
	w := make(map[int]io.Writer, l)
	c := make([]func(), 0, l)
	for _, lv := range outLevel {
		newFileName := filename
		if lv != defaultIo {
			newFileName = fmt.Sprintf("%s.%s.log", filename, getLevelString(lv))
		}
		f, err := os.OpenFile(newFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("create %s faild ", filename))
		}
		w[lv] = f
		c = append(c, func() {
			_ = f.Close()
		})
	}
	// new
	ret := &Logger{level: logLevel, isOutFile: true, outFile: w}
	return ret, nil
}

func (l *Logger) Close() {
	for _, fn := range l.closer {
		fn()
	}
}

// getLevel gets logging level int value corresponding to the specified level.
func getLevel(level string) int {
	level = strings.ToLower(level)

	switch level {
	case "off":
		return Off
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}

func getLevelString(level int) string {
	levelName := "default"
	switch level {
	case Trace:
		levelName = "trace"
		break
	case Debug:
		levelName = "debug"
		break
	case Info:
		levelName = "info"
		break
	case Warn:
		levelName = "warn"
		break
	case Error:
		levelName = "error"
		break
	case Fatal:
		levelName = "fatal"
		break
	default:
		levelName = "default"
	}
	return levelName
}

func (l *Logger) Format(level int, v ...interface{}) {
	f := fmt.Sprintf("[%s][%s] ", strings.ToUpper(getLevelString(level)), time.Now().Format(timeLine))
	nv := append([]interface{}{f}, v)
	if !l.isOutFile {
		_, _ = fmt.Fprintf(l.DefaultWriter,"%s %+v \n",  nv...)
		return
	}
	w, ok := l.outFile[level]
	if !ok {
		w = l.outFile[defaultIo]
	}
	_, _ = fmt.Fprintf(w,"%s %+v \n",  nv...)
}

func (l *Logger) Formatf(level int, f string, v ...interface{}) {
	f = fmt.Sprintf("[%s][%s] ", strings.ToUpper(getLevelString(level)), time.Now().Format(timeLine)) + f + "\n"
	if !l.isOutFile {
		_, _ = fmt.Fprintf(l.DefaultWriter, f, v...)
		return
	}
	w, ok := l.outFile[level]
	if !ok {
		w = l.outFile[defaultIo]
	}
	_, _ = fmt.Fprintf(w, f, v...)
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level string) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

// Info prints info level message.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}
	l.Format(Info, v...)
}

// Infof prints info level message with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}
	l.Formatf(Info, format, v...)
}

// Error prints error level message.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}
	l.Format(Error, v...)
}

// Errorf prints error level message with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}
	l.Formatf(Error, format, v...)
}
