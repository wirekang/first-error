package ferr

import (
	"fmt"
	"os"
	"runtime/debug"

	errors2 "github.com/pkg/errors"
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

// StackTrace returns stack trace when error first wrapped, or returns "" if err is nil.
func StackTrace(err interface{}) (str string) {
	if err != nil {
		message := fmt.Sprintf("%v", err)
		err, ok := err.(error)
		st := ""
		if ok {
			for {
				stackTracer, ok := err.(interface {
					StackTrace() errors2.StackTrace
				})
				if ok {
					st = fmt.Sprintf("%+v\n", stackTracer.StackTrace())
					err = errors2.Unwrap(err)
					if err != nil {
						continue
					}
				}

				break
			}
		}
		if st == "" {
			st = string(debug.Stack())
		}

		str = fmt.Sprintf("%s\n\nStackTrace:\n%s", message, st)
	}
	return
}
