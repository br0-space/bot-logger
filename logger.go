package logger

import (
	"flag"
	"sync"

	"github.com/br0-space/bot-logger/live"
	"github.com/br0-space/bot-logger/null"
)

type Interface interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	SetExtraCallDepth(depth int)
	ResetExtraCallDepth()
}

var (
	instance Interface
	lock     = &sync.Mutex{}
)

func New() Interface {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		if runsAsTest() {
			instance = null.New()
		} else {
			instance = live.New()
		}
	}

	return instance
}

func runsAsTest() bool {
	return flag.Lookup("test.v") != nil
}
