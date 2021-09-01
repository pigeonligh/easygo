# ELog

## Usage

```golang
package main

import (
    log "gopkg.pigeonligh.com/easygo/elog"
)

func main() {
    log.Default()
    log.Info("Hello world!")
}
```

Output:

```
{"action":"Info","message":"Hello world!","source":"/path/to/file.go:9"},
```