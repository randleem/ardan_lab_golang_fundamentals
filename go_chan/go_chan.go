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

	ch := make(chan int)
	go func() {
		ch <- 7 // send
	}()
	v := <-ch // receive
	fmt.Println(v)

	fmt.Println(sleepSort([]int{20, 30, 10})) //[10,20,30]
}

/*
	Algorithm

- For every value "n" in values, spin a go routine that
  - sleep for "n" milliseocnds
  - sends "n over a channel"

- collect all values from the channel to a slice and return it
*/
func sleepSort(values []int) []int {
	slice := []int{}
	ch := make(chan int)
	for _, v := range values {
		fmt.Println(v)
		go func() {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}()

	}
	for range values {
		r := <-ch
		slice = append(slice, r)
	}

	return slice
}

/* Channel Semantics
- send/receive to/from a channel will block until opposite opporation(*) (until something receives from the channel)
	- guarantee of delivery - receive set up before the send
*/
