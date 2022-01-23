# first-error

## Example

```go
package main

import "github.com/wirekang/first-error"

func main(){
	defer ferr.Recover(func(s){
		fmt.Pritln(s)
		os.Exit(1)
    })

	err:= fmt.Errorf("some error")
	if err != nil {
		err = errors.WithStack(err)
		panic(err)
	}
}
```