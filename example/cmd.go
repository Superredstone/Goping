package main

import (
	"fmt"
	"goping"
	"os"
	"time"
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
