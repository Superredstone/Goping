# Goping
A simple library to ping a certain ip.
Besides that it can also perform port scanning (not very fast) 

## Example
```go
package main

import (
    "time"

    "github.com/Superredstone/Goping"
)

func main() {
    fmt.Println(Goping.Ping("127.0.0.1", "22", time.Duration(time.Duration(10) * time.Millisecond)))
}
```