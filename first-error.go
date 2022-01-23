package ferr

import (
	"fmt"
	"runtime/debug"

	"github.com/pkg/errors"
)

func Recover(cb func(string)) {
	cause := recover()
	cb(StackTrace(cause))
}

func StackTrace(err interface{}) (str string) {
	if err != nil {
		message := fmt.Sprintf("%v", err)
		err, ok := err.(error)
		st := ""
		if ok {
			for {
				stackTracer, ok := err.(interface {
					StackTrace() errors.StackTrace
				})
				if ok {
					st = fmt.Sprintf("%+v\n", stackTracer.StackTrace())
					err = errors.Unwrap(err)
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
