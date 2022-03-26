package ferr

import (
	"fmt"
	"os"
	"runtime/debug"

	errors "github.com/pkg/errors"
)

// RecoverAndPrintAndExit prints stack trace and exit with exitCode if recover() returns non-nil value.
func RecoverAndPrintAndExit(errExitCode int) {
	cause := recover()
	if cause != nil {
		fmt.Println(cause)
		fmt.Println(StackTrace(cause))
		os.Exit(errExitCode)
	}
}

// StackTrace is alias for StackTraceRange(err,0,100)
func StackTrace(err interface{}) (str string) {
	return StackTraceRange(err, 0, 100)
}

// StackTraceRange returns stack trace with start/max frame range.
func StackTraceRange(err interface{}, start, max int) (str string) {
	if err != nil {
		err, ok := err.(error)
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

					str = fmt.Sprintf("%+v", st)
					err = errors.Unwrap(err)
					if err != nil {
						continue
					}
				}

				break
			}
		}
		if str == "" {
			str = string(debug.Stack())
		}
	}
	return
}
