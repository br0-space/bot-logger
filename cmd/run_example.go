package main

import (
	"fmt"

	"github.com/br0-space/bot-logger"
	"github.com/spf13/pflag"
)

func main() {
	// Flags understood by the live logger
	pflag.Bool("verbose", false, "enable verbose (debug) logging")
	pflag.Bool("quiet", false, "only log errors")

	// Optional flags to demonstrate Panic/Fatal without enabling them by default
	doPanic := pflag.Bool("do-panic", false, "actually call Panic/Panicf (this will panic)")
	doFatal := pflag.Bool("do-fatal", false, "actually call Fatal/Fatalf (this will exit the program)")

	pflag.Parse()

	log := logger.New()

	// Basic examples
	log.Debug("example debug")
	log.Debugf("example %s", "debugf")

	log.Info("example info")
	log.Infof("example %s", "infof")

	log.Warning("example warning")
	log.Warningf("example %s", "warningf")

	log.Error("example error")
	log.Errorf("example %s", "errorf")

	// Demonstrate extra call depth wrappers
	log.SetExtraCallDepth(1)
	wrappedInfo(log)
	log.ResetExtraCallDepth()

	// Panic examples (opt-in)
	if *doPanic {
		fmt.Println("-- triggering Panic examples --")
		// Use a recover wrapper so we can optionally continue to Fatal examples
		func() {
			defer func() { _ = recover() }()
			log.Panic("example panic")
		}()
		func() {
			defer func() { _ = recover() }()
			log.Panicf("example %s", "panicf")
		}()
	}

	// Fatal examples (opt-in) — note: this will terminate the program
	if *doFatal {
		fmt.Println("-- triggering Fatal examples (program will exit) --")
		log.Fatal("example fatal")
		log.Fatalf("example %s", "fatalf")
	}
}

func wrappedInfo(log logger.Interface) {
	log.Info("example info with extra call depth from wrapper")
}
