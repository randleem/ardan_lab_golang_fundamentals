package main

import (
	"fmt"
	"slices"
)

type Item struct {
	X int
	Y int
}
type Player struct {
	Name string
	Item // Player embeds Item
	Keys []string
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

	i.Move(10, 20)
	fmt.Printf("i (move): %#v\n", i)

	p1 := Player{
		Name: "Emma",
	}
	fmt.Printf("p1: %+v\n", p1)
	fmt.Println("p1.X:", p1.Item.X)

	// p1 is Player - but have Item in Player so can use Move - composition or inheritance
	p1.Move(100, 200)
	fmt.Printf("p1: (move) %+v\n", p1)

	p2 := Player{
		Name: "Ernie",
		Keys: []string{},
	}
	err := p2.Found("jade")
	if err != nil {
		fmt.Println("jade: ", err)
	}
	err = p2.Found("nickel")
	if err != nil {
		fmt.Println("nickel: ", err)
	}
	fmt.Println("keys:", p2.Keys)
}

/*Exercise:
- Add a "Keys" field to player with a slice of strings
- Add a Found(key string) method to player
	- It should err if key is not one of :jade, copper, crystal
	- It should add a key only once
*/

func (p *Player) Found(key string) error {
	k := []string{"jade", "copper", "crystal"}

	f := slices.Contains(k, key)
	if f {
		if len(p.Keys) < 1 {
			p.Keys = append(p.Keys, key)
		}
		return nil
	}
	return fmt.Errorf("key not found")
}

/* ASIDE Use $#v for debugging/logging
a, b := 1, "1"
fmt.Printf("a: %v, b: %v\n", a, b)
fmt.Printf("a: %#v, b: %#v\n", a, b)
*/

// https://go-proverbs.github.io/

// Move moves i by delta X & delta y.
// "i is called the reciever"
// i is a pointer receiver - actual value not a copy
/* Value vs pointer receiver
- In general use value semantics
- Try to keep same sematics on all methos
- some cases when you have to use pointer semantics
	- if you have a locked field
	- if you need to mutate the struct - obcs send pointer ro change original value
	- Decoding/unmarshalling
*/

func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}
