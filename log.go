package logger

import (
	"io/ioutil"
	"log"
	"os"
	"time"
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
)

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}

type Config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      LogLevel
}

// Interface logger interface
type Interface interface {
	LogMode(LogLevel) Interface
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

var (
	Discard = New(log.New(ioutil.Discard, "", log.LstdFlags), Config{})
	Default = New(log.New(os.Stdout, "\r\n", log.LstdFlags), Config{
		LogLevel: Info,
		Colorful: true,
	})
)

func New(writer Writer, config Config) Interface {
	var (
		infoStr = "%s\n[info] "
		warnStr = "%s\n[warn] "
		errStr  = "%s\n[error] "
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr = BlueBold + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
	}

	return &logger{
		Writer:  writer,
		Config:  config,
		infoStr: infoStr,
		warnStr: warnStr,
		errStr:  errStr,
	}
}

type logger struct {
	Writer
	Config
	infoStr, warnStr, errStr string
}

// LogMode log mode
func (l *logger) LogMode(level LogLevel) Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l logger) Infof(msg string, data ...interface{}) {
	if l.LogLevel >= Info {
		l.Printf(l.infoStr+msg, data...)
	}
}

// Warn print warn messages
func (l logger) Warnf(msg string, data ...interface{}) {
	if l.LogLevel >= Warn {
		l.Printf(l.warnStr+msg, data...)
	}
}

// Error print error messages
func (l logger) Errorf(msg string, data ...interface{}) {
	if l.LogLevel >= Error {
		l.Printf(l.errStr+msg, data...)
	}
}
