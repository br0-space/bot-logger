# bot-logger

A tiny logging helper for Go that gives you a sane default logger in applications and a no-op logger in tests.

- Live logger uses github.com/op/go-logging with colorized output and timestamped levels
- Test runs automatically receive a null logger (no output) so your test output stays clean
- Verbosity controlled via flags: --verbose for debug-level and extended format, --quiet for error-only
- Drop-in interface with familiar methods: Debug, Info, Warning, Error, Panic, Fatal and their formatted variants
- Extra call depth helpers so log lines point to your code, not the wrapper

## Status

[![Build](https://github.com/br0-space/bot-logger/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-logger/actions/workflows/build.yml)
[![Test](https://github.com/br0-space/bot-logger/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-logger/actions/workflows/test.yml)
[![Lint](https://github.com/br0-space/bot-logger/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-logger/actions/workflows/lint.yml)
[![Staticcheck](https://github.com/br0-space/bot-logger/actions/workflows/staticcheck.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-logger/actions/workflows/staticcheck.yml)
[![Vet](https://github.com/br0-space/bot-logger/actions/workflows/vet.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-logger/actions/workflows/vet.yml)
[![CodeQL](https://github.com/br0-space/bot-logger/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/br0-space/bot-logger/actions/workflows/codeql-analysis.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/br0-space/bot-logger.svg)](https://pkg.go.dev/github.com/br0-space/bot-logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/br0-space/bot-logger)](https://goreportcard.com/report/github.com/br0-space/bot-logger)

## Install

```
go get github.com/br0-space/bot-logger
```

## Quick start

```go
package main

import (
    "github.com/br0-space/bot-logger"
)

func main() {
    log := logger.New()
    log.Infof("hello %s", "world")
}
```

## CLI flags: verbose and quiet

The live logger adapts to two command-line flags using spf13/pflag:
- --verbose enables DEBUG level and a more verbose format (includes file and function)
- --quiet sets the level to ERROR

To make those flags available, define them in your program before you parse flags. If you already use pflag, it’s just:

```go
import (
    "github.com/spf13/pflag"
    "github.com/br0-space/bot-logger"
)

func main() {
    // Define flags the logger understands
    pflag.Bool("verbose", false, "enable verbose (debug) logging")
    pflag.Bool("quiet", false, "only log errors")
    pflag.Parse()

    log := logger.New() // picks up the parsed flags

    log.Debug("debug message")
    log.Info("info message")
    log.Warning("warning message")
    log.Error("error message")
}
```

If you don’t use pflag, the defaults apply: INFO level with the standard format.

## Behavior in tests

When your code runs under "go test", logger.New() returns a null logger that discards output. This keeps your test logs clean without additional setup.

## Runnable example

A small example program is included to try the logger quickly from the command line.

Run it with:

```
go run ./cmd/run_example.go
```

Useful flags:
- --verbose to enable DEBUG level and extended format
- --quiet to only show errors
- --do-fatal to actually invoke Fatal (will exit the program)
- --do-fatalf to actually invoke Fatalf (will exit the program)

Notes:
- The example triggers Panic and Panicf but immediately recovers so the program continues.
- Fatal and Fatalf will terminate the program if enabled via the flags above.

## API

The logger.Interface exposes the following methods:

- Debug(args ...any)
- Debugf(format string, args ...any)
- Info(args ...any)
- Infof(format string, args ...any)
- Warning(args ...any)
- Warningf(format string, args ...any)
- Error(args ...any)
- Errorf(format string, args ...any)
- Panic(args ...any)
- Panicf(format string, args ...any)
- Fatal(args ...any)
- Fatalf(format string, args ...any)
- SetExtraCallDepth(depth int)
- ResetExtraCallDepth()

Use SetExtraCallDepth when you wrap the logger and want log lines to reference your caller instead of your wrapper.

Example:

```go
func doSomething(log logger.Interface) {
    // We are adding a wrapper here; bump the call depth so file/line shows the caller of doSomething
    log.SetExtraCallDepth(1)
    defer log.ResetExtraCallDepth()

    log.Info("processing...")
}
```

## Formatting

Default format:
- "2006-02-01 15:04:05.000 LEVEL: message" with colors

Verbose format (when --verbose is set):
- Adds long file path and short function name before the level

## Versioning

This library follows Go module versioning via tags on the main branch. Import path is:

```
github.com/br0-space/bot-logger
```

## License

MIT – see LICENSE.
