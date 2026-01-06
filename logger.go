package logger

import (
	"flag"
	"sync"

	"github.com/br0-space/bot-logger/live"
	"github.com/br0-space/bot-logger/null"
)

type Interface interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warning(args ...any)
	Warningf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
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
