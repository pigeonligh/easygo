# Errors

## Usage

```golang
package main

import (
    "fmt"
    
    "github.com/pigeonligh/easygo/errors"
)

func main() {
    a := errors.New("hello")
    b := errors.New("world")
    err = errors.Append(a, b)
    fmt.Println(err)
}
```
