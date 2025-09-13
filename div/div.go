package main

import "fmt"

func main() {
	fmt.Println(div(7, 3))
	// fmt.Println(div(7, 0))
	fmt.Println(safeDiv(7, 0))
}

func div(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func safeDiv(a, b int) (q int, err error) {
	// q & err are variables inside safeDiv
	// Just like a & b are
	defer func() {
		if e := recover(); e != nil {
			// fmt.Println("ERROR: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	/*
	   Using named return values should be used in a handful of cases:
	   - defer/recover to change return error values
	   - documentation
	*/
	/*
		Below is valid code but not optimal
		q = div(a,b)
		return
	*/

	return div(a, b), nil
}
