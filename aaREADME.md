# first-error
Get stack trace where the error first wrapped with github.com/pkg/errors .WithStack()

## Usage

Warp every error with [WithStack](https://github.com/pkg/errors).

```go
if err != nil {
  err = errors.WithStack(err)
}

```

You can get stack trace of first wrapped error.

```go

ferr.StackTrace(nil) // ""
ferr.StackTrace(err) // error1234 \n\n StackTrace: \n...

```

## Example

### error reporting with ferr.RecoverCallback()

https://github.com/wirekang/mouseable/blob/d3c97367718aef3f17de4a38412d8dc7e9a88f4c/internal/logic/logic.go#L80

```
defer ferr.Recover(
		func(msg string) {
			// Do something with msg.
		},
	)
```
