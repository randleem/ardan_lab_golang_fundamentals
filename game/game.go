package main

import "fmt"

type Item struct {
	X int
	Y int
}

const (
	maxX = 600
	maxY = 400
)

/*
	Types of "new" or factory functions

func NewItem(x,y int)Item{}
func NewItem(x,y int)*Item{}
func NewItem(x,y int)(Item, error){}
func NewItem(x,y int)(*Item, error){}

Value semantics: everyone has thier own copy
Pointer semantics: everyone shares the same copy (heap, lock)

*/

/* Value semantics - not a pointer - returns a new Item.
func NewItem(x, y int) (Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return Item{}, fmt.Errorf("%d/%d of bounds %d/%d", x, y, maxX, maxY)
	}
	return Item{
		X: x,
		Y: y,
	}, nil
}
*/

// Pointer Item adds Item to heap. memory allcoation costly
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return &Item{}, fmt.Errorf("%d/%d of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	// The go compiler does escape analysis and will allocate i on the heap
	return &i, nil
}

func main() {
	var i Item
	fmt.Printf("i: %#v\n", i)

	i = Item{10, 20} // Must specify/add all fields
	fmt.Printf("i: %#v\n", i)

	// can be in any order if you specify field names
	i = Item{Y: 22, X: 11}
	fmt.Printf("i: %#v\n", i)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 2000))
}

/* ASIDE Use $#v for debugging/logging
a, b := 1, "1"
fmt.Printf("a: %v, b: %v\n", a, b)
fmt.Printf("a: %#v, b: %#v\n", a, b)
*/

// https://go-proverbs.github.io/
