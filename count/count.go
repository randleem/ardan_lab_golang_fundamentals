package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	/* Solution 1 mutex
	var mu sync.Mutex // place over variables the mutex is guarding
	count := 0
	*/

	count := int64(0)
	nGR, nIter := 10, 1000

	var wg sync.WaitGroup

	wg.Add(nGR)
	for range nGR {
		go func() {
			defer wg.Done()
			for range nIter {
				// Soluton 2 sync/atomic
				atomic.AddInt64(&count, 1)
				/* Solution 1 mutex
				mu.Lock()
				count++
				mu.Unlock()
				*/
				/*
					count ++ is translated to:
						fetch count
						increment count
						store count
						race condition - several go routines are accesing the same resource
				*/
				time.Sleep(time.Microsecond)
			}
		}()
	}
	wg.Wait()
	fmt.Println("count", count)
}

/*
go run -race count.go
"-race" is supported by
- run
- build
 -test

 Rule of thumb: use "go test -race"
*/

/*
Wait group allows us to wait for all go routines to
finish - but we can have a race condition if the
goroutines are accesing and changing the same resource
- i.e. the count variable.

There are a few different ways to combat this:

1 - use a mutex - a mutex (short for mutual exclusion)
is a synchronization primitive that prevents multiple
goroutines from accessing a shared resource at the same time.
A mutex ensures that only one goroutine can access a
piece of code or data at any moment.

2 -
*/
