# Errors

## Usage

```golang
package main

import (
	"fmt"

	"gopkg.pigeonligh.com/easygo/errors"
)

func main() {
	var err error
	a := errors.New("hello")
	b := errors.New("world")
	err = errors.Merge(a, b)
	fmt.Println(err)
}
```

Output:

```
[hello, world]
```
