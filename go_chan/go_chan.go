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

	fmt.Println(sleepSort([]int{20, 30, 10}))                //[10,20,30]
	fmt.Println("sleepSort2", sleepSort2([]int{20, 30, 10})) //[10,20,30]
	go func() {
		for i := range 4 {
			ch <- i
			// fmt.Println("closed:", ch)
		}
		close(ch)
	}()
	// When you iteratre over a channel you dont know how many things are coming in so will error unless you close the channel
	for v := range ch {
		fmt.Println(">>", v)
	}

	v = <-ch // ch is closed
	fmt.Println("closed:", v)
	v, ok := <-ch // ch is closed, ok false as channel closed
	fmt.Println("closed:", v, "ok", ok)

	/* The "for range" above does
	for {
		v, ok:= <- ch
		if !ok{
			break
		}
		fmt.Println(">>", v)
	}
		So syntactic sugar - shortcut to not needig the above code, breaks behind the scene if the channel is closed.
	*/
	// var ch chan int // ch is nil
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
- receive from a closed channel will return zero without blocking
	- use "comma ok" to check if channel was closed
- send to a closed channel - will panic.
- dont have to close channels in go - not a memory leak
- closing a closed or nil channel will panic
- send/receive to a nil channel will block forever.
*/

func sleepSort2(values []int) []int {
	sortedValues := []int{}
	ch := make(chan int)
	for _, v := range values {
		go func() {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}()
	}
	for range values {
		value := <-ch
		sortedValues = append(sortedValues, value)
	}

	return sortedValues
}
