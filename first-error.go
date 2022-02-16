package ferr

import (
	"fmt"
	"os"
	"runtime/debug"

	errors "github.com/pkg/errors"
)

// RecoverAndPrintAndExit prints stack trace and exit with exitCode if recover() returns non-nil value.
func RecoverAndPrintAndExit(exitCode int) {
	cause := recover()
	if cause != nil {
		fmt.Println(StackTrace(cause))
		os.Exit(exitCode)
	}
}

// RecoverCallback calls cb if recover() returns non-nil value.
func RecoverCallback(cb func(string)) {
	cause := recover()
	if cause != nil {
		cb(StackTrace(cause))
	}
}

// StackTrace is alias for StackTraceRange(err,0,100)
func StackTrace(err interface{}) (str string) {
	return StackTraceRange(err, 0, 100)
}

// StackTraceRange returns stack trace with start/max frame range.
func StackTraceRange(err interface{}, start, max int) (str string) {
	if err != nil {
		message := fmt.Sprintf("%v", err)
		err, ok := err.(error)
		stString := ""
		if ok {
			for {
				stackTracer, ok := err.(interface {
					StackTrace() errors.StackTrace
				})
				if ok {
					st := stackTracer.StackTrace()
					if start >= len(st) {
						start = len(st) - 1
					}

					end := start + max
					if end > len(st) {
						end = len(st)
					}
					if start > end {
						st = nil
					} else {
						st = st[start:end]
					}

					stString = fmt.Sprintf("%+v", st)
					err = errors.Unwrap(err)
					if err != nil {
						continue
					}
				}

				break
			}
		}
		if stString == "" {
			stString = string(debug.Stack())
		}

		str = fmt.Sprintf("%s\n%s\n", message, stString)
	}
	return
}
