package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1 means its a buffered channel, this avoids goroutine leaks - where the go routines are running but the programme has finished.
	ch1, ch2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "two"
	}()
	// context is usually created at http handler, and passed all around - so is useful for adding values to be passed around.
	// Has 2 main uses - setting timeouts and cancellations, and adding meta data- authentication, request id - key/value
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	// manually cancel ctx
	defer cancel()
	// whichever channel receives first is printed - ie ch1
	select {
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	// case <-time.After(10 * time.Millisecond):
	// 	fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
