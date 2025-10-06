package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := range 3 {
		// Prior to Go 1.22 this was a bug
		// go always need a func call, use an anonymous func call
		go func() {
			fmt.Println("goroutine", i)
		}()
	}
	time.Sleep(1 * time.Millisecond)
}
