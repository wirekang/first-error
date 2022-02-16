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

// StackTrace is alias for StackTraceRange(err,0,0)
func StackTrace(err interface{}) (str string) {
	return StackTraceRange(err, 0, 0)
}

// StackTraceRange returns stack trace string with skip count.
func StackTraceRange(err interface{}, skipNewest, skipOldest int) (str string) {
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
					start := skipNewest
					end := len(st) - skipOldest
					if len(st) < start || len(st) < end || start > end {
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
