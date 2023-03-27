package live

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/spf13/pflag"
	"os"
)

type Logger struct {
	wrappedLogger *logging.Logger
	prefix        string
}

var defaultFormat = logging.MustStringFormatter(`%{color}%{time:2006-02-01 15:04:05.000} %{level}:%{color:reset} %{message}`)
var verboseFormat = logging.MustStringFormatter(`%{color}%{time:2006-02-01 15:04:05.000} %{longfile} %{shortfunc} %{level}:%{color:reset} %{message}`)

func New() *Logger {
	wrappedLogger := logging.MustGetLogger("")
	wrappedLogger.ExtraCalldepth = 1

	defaultLogger := &Logger{
		wrappedLogger: wrappedLogger,
		prefix:        "",
	}

	// Backend settings depend on the presence of command line flags --verbose and --quiet
	// To avoid an import cycle, we can NOT read the config from the container, but have to ask pflag directly
	verbose, _ := pflag.CommandLine.GetBool("verbose")
	quiet, _ := pflag.CommandLine.GetBool("quiet")

	// Create one backend (for now) that writes to os.Stderr
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	backendFormatter := logging.NewBackendFormatter(backend, defaultFormat)
	// If command line flag --verbose is set, an extended output format will be used
	if verbose {
		backendFormatter = logging.NewBackendFormatter(backend, verboseFormat)
	}

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	// Set default log level to INFO
	backendLeveled.SetLevel(logging.INFO, "")
	// If command line flag --verbose is set, log level will be DEBUG instead
	if verbose {
		backendLeveled.SetLevel(logging.DEBUG, "")
	}
	// If command line flag --quiet is set, log level will be ERROR instead
	if quiet {
		backendLeveled.SetLevel(logging.ERROR, "")
	}

	logging.SetBackend(backendLeveled)

	return defaultLogger
}

func (l *Logger) Debug(args ...interface{}) {
	l.wrappedLogger.Debug(l.addPrefixToSlice(args...)...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.wrappedLogger.Debugf(l.addPrefixToString(format), args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.wrappedLogger.Info(l.addPrefixToSlice(args...)...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.wrappedLogger.Infof(l.addPrefixToString(format), args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.wrappedLogger.Warning(l.addPrefixToSlice(args...)...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	l.wrappedLogger.Warningf(l.addPrefixToString(format), args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.wrappedLogger.Error(l.addPrefixToSlice(args...)...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.wrappedLogger.Errorf(l.addPrefixToString(format), args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.wrappedLogger.Panic(l.addPrefixToSlice(args...)...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.wrappedLogger.Panicf(l.addPrefixToString(format), args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.wrappedLogger.Fatal(l.addPrefixToSlice(args...)...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.wrappedLogger.Fatalf(l.addPrefixToString(format), args...)
}

func (l *Logger) SetExtraCallDepth(depth int) {
	l.wrappedLogger.ExtraCalldepth = 1 + depth
}

func (l *Logger) ResetExtraCallDepth() {
	l.wrappedLogger.ExtraCalldepth = 1
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) ResetPrefix() {
	l.prefix = ""
}

func (l *Logger) getPrefix() string {
	if l.prefix == "" {
		return ""
	}
	return fmt.Sprintf("%s:", l.prefix)
}

func (l *Logger) addPrefixToSlice(args ...interface{}) []interface{} {
	if l.getPrefix() == "" {
		return args
	}

	x := args[0]
	switch x.(type) {
	case string:
		v := l.getPrefix()
		args = append([]interface{}{v}, args...)
	}
	return args
}

func (l *Logger) addPrefixToString(string string) string {
	if l.getPrefix() == "" {
		return string
	}

	return fmt.Sprintf("%s %s", l.getPrefix(), string)
}
