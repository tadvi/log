package log

import (
	"log"
	"os"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// LogLevel
type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
	Debug
)

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}

type Config struct {
	Colorful bool
	LogLevel LogLevel
}

var (
	Default = New(log.New(os.Stderr, "\n", log.LstdFlags), Config{
		LogLevel: Debug,
		Colorful: true,
	})
)

func New(writer Writer, config Config) *Logger {
	var (
		debugStr = "%s\n[debug] "
		infoStr  = "%s\n[info]  "
		warnStr  = "%s\n[warn]  "
		errStr   = "%s\n[error] "
	)

	if config.Colorful {
		debugStr = "%s\n" + "[debug] "
		infoStr = Green + "%s\n" + Reset + Green + "[info]  " + Reset
		warnStr = BlueBold + "%s\n" + Reset + Magenta + "[warn]  " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
	}

	return &Logger{
		Writer:   writer,
		Config:   config,
		debugStr: debugStr,
		infoStr:  infoStr,
		warnStr:  warnStr,
		errStr:   errStr,
	}
}

type Logger struct {
	Writer
	Config
	debugStr, infoStr, warnStr, errStr string
}

// LogMode log mode.
func (l *Logger) LogMode(level LogLevel) *Logger {
	l.LogLevel = level
	return l
}

// Debug print debug messages.
func (l Logger) Debugf(msg string, data ...interface{}) {
	if l.LogLevel >= Debug {
		l.Printf(l.debugStr+msg, data...)
	}
}

// Info print info.
func (l Logger) Infof(msg string, data ...interface{}) {
	if l.LogLevel >= Info {
		l.Printf(l.infoStr+msg, data...)
	}
}

// Warn print warn messages.
func (l Logger) Warnf(msg string, data ...interface{}) {
	if l.LogLevel >= Warn {
		l.Printf(l.warnStr+msg, data...)
	}
}

// Error print error messages.
func (l Logger) Errorf(msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		l.Printf(l.errStr+msg, data...)
	}
}

// Fatalf print error messages and exit.
func (l Logger) Fatalf(msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		l.Printf(l.errStr+msg, data...)
	}
	os.Exit(1)
}

func Debugf(msg string, data ...interface{}) {
	if Default.LogLevel >= Debug {
		Default.Printf(Default.debugStr+msg, data...)
	}
}

// Info print info.
func Infof(msg string, data ...interface{}) {
	if Default.LogLevel >= Info {
		Default.Printf(Default.infoStr+msg, data...)
	}
}

// Warn print warn messages.
func Warnf(msg string, data ...interface{}) {
	if Default.LogLevel >= Warn {
		Default.Printf(Default.warnStr+msg, data...)
	}
}

// Error print error messages.
func Errorf(msg string, data ...interface{}) {
	if Default.LogLevel >= Error {
		Default.Printf(Default.errStr+msg, data...)
	}
}

// Fatalf print error messages and exit.
func Fatalf(msg string, data ...interface{}) {
	if Default.LogLevel >= Error {
		Default.Printf(Default.errStr+msg, data...)
	}
	os.Exit(1)
}
