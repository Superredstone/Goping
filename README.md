# Goping
A simple library to ping a certain ip.
Besides that it can also perform port scanning (not very fast) 

## Install
```bash
go get github.com/Superredstone/goping
```

## Example
```go
package main

import (
	"fmt"
	"os"
	"time"

    "github.com/Superredstone/goping"
)

func main() {
	Opened, err := goping.Ping("127.0.0.1", "22", time.Duration(time.Duration(10)*time.Millisecond))
	if err != nil {
		os.Exit(1)
	}
	if Opened {
		fmt.Println("PORT 22 IS OPENED")
	} else {
		fmt.Println("PORT 22 IS CLOSED")
	}
}

```