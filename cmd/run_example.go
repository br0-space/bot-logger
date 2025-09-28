package main

import (
	logger "github.com/br0-space/bot-logger"
	"github.com/spf13/pflag"
)

func main() {
	// Flags understood by the live logger
	pflag.Bool("verbose", false, "enable verbose (debug) logging")
	pflag.Bool("quiet", false, "only log errors")

	// Optional flags to demonstrate Panic/Fatal without enabling them by default
	doFatal := pflag.Bool("do-fatal", false, "actually call Fatal (this will exit the program)")
	doFatalf := pflag.Bool("do-fatalf", false, "actually call Fatalf (this will exit the program)")

	pflag.Parse()

	log := logger.New()

	log.Debug("example debug")
	log.Debugf("example %s", "debugf")

	log.Info("example info")
	log.Infof("example %s", "infof")

	log.Warning("example warning")
	log.Warningf("example %s", "warningf")

	log.Error("example error")
	log.Errorf("example %s", "errorf")

	log.SetExtraCallDepth(1)
	wrappedInfo(log)
	log.ResetExtraCallDepth()

	func() {
		defer func() { _ = recover() }()

		log.Panic("example panic")
	}()
	func() {
		defer func() { _ = recover() }()

		log.Panicf("example %s", "panicf")
	}()

	if *doFatal {
		log.Fatal("example fatal")
	}

	if *doFatalf {
		log.Fatalf("example %s", "fatalf")
	}
}

func wrappedInfo(log logger.Interface) {
	log.Info("example info from wrapper (using extra call depth)")
}
