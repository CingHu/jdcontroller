package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type priority int

const (
	//RFC3339Nano string = "2006-01-02T15:04:05.000000000Z07:00"
	RFC3339Nano string = "2006-01-02 15:04:05"
	errorFormat        = "[%s] [%s] %s:%d %s\n"
	logFormat          = "[%s] [%s] %s\n"

	fatalPriority priority = iota
	errorPriority
	infoPriority
	debugPriority
)

// A common interface to access the Fatal method of
// both testing.B and testing.T.
type Fataler interface {
	Fatal(args ...interface{})
}

func (p priority) String() string {
	switch p {
	case fatalPriority:
		return "fatal"
	case errorPriority:
		return "error"
	case infoPriority:
		return "info"
	case debugPriority:
		return "debug"
	}

	return ""
}

var DefaultLogger = Logger{Out: os.Stdout, Err: os.Stderr}

// Debug function, if the debug flag is set, then display. Do nothing otherwise
// If Docker is in damon mode, also send the debug info on the socket
func Debugf(format string, a ...interface{}) (int, error) {
	//	fmt.Print(format, " ")
	//	fmt.Println(a...)
	return DefaultLogger.Debugf(format, a...)
}

func Infof(format string, a ...interface{}) (int, error) {
	return DefaultLogger.Infof(format, a...)
}

func Errorf(format string, a ...interface{}) (int, error) {
	return DefaultLogger.Errorf(format, a...)
}

func Fatal(a ...interface{}) {
	DefaultLogger.Fatalf("%s", a...)
}

func Fatalf(format string, a ...interface{}) {
	DefaultLogger.Fatalf(format, a...)
}

type Logger struct {
	Err io.Writer
	Out io.Writer
}

func (l Logger) Debugf(format string, a ...interface{}) (int, error) {
	if os.Getenv("DEBUG") != "" {
		return l.logf(l.Err, debugPriority, format, a...)
	}
	return 0, nil
}

func (l Logger) Infof(format string, a ...interface{}) (int, error) {
	return l.logf(l.Out, infoPriority, format, a...)
}

func (l Logger) Errorf(format string, a ...interface{}) (int, error) {
	return l.logf(l.Err, errorPriority, format, a...)
}

func (l Logger) Fatalf(format string, a ...interface{}) {
	l.logf(l.Err, fatalPriority, format, a...)
	os.Exit(1)
}

func (l Logger) logf(stream io.Writer, level priority, format string, a ...interface{}) (int, error) {
	var prefix string

	if level <= errorPriority || level == debugPriority {
		// Retrieve the stack infos
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "<unknown>"
			line = -1
		} else {
			file = file[strings.LastIndex(file, "/")+1:]
		}
		prefix = fmt.Sprintf(errorFormat, time.Now().Format(RFC3339Nano), level.String(), file, line, format)
	} else {
		prefix = fmt.Sprintf(logFormat, time.Now().Format(RFC3339Nano), level.String(), format)
	}

	return fmt.Fprintf(stream, prefix, a...)
}
