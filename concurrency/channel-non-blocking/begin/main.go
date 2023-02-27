// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	timeChan := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)

	for {
		select {
		case <-timeChan:
			fmt.Println("timeChan1")
			return
		case <-timeChan2:
			fmt.Println("timeChan2")
			return
		default:
			fmt.Println("default")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
