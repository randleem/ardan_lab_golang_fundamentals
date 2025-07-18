package main

import (
	"fmt"
	"slices"
)

type Item struct {
	X int
	Y int
}

type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

type Player struct {
	Name string
	Item // Player embeds Item
	Keys []Key
}

// Interface set of methods,(and types), we define as "what you need" not "what you provide"
// Interfaces are small (stdlib average ~2 methods per interface)
// if interface with more than 4 methods - think again
// start with concrete types, discover interfaces
type Mover interface {
	Move(int, int)
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

//go install golang.org/x/tools/cmd/stringer@latest
// In ~/.zshrc
//export PATH="$(go env GOPATH)/bin:${PATH}"

// String implement the fmt.Stringer interface.
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	default:
		return fmt.Sprintf("<Key %d>", k)
	}
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
		Keys: []Key{},
	}
	err := p2.Found(Copper)
	if err != nil {
		fmt.Println("jade: ", err)
	}
	err = p2.Found(Key(7))
	if err != nil {
		fmt.Println(": ", err)
	}
	fmt.Println("keys:", p2.Keys)

	ms := []Mover{
		&i,
		&p1,
	}

	MoveAll(ms, 50, 7)
	for _, m := range ms {
		fmt.Println(m)
	}
}

/*Exercise:
- Add a "Keys" field to player with a slice of strings
- Add a Found(key string) method to player
	- It should err if key is not one of :jade, copper, crystal
	- It should add a key only once
*/

func (p *Player) Found(key Key) error {
	k := []Key{Jade, Copper, Crystal}

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

func MoveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

/*
Thought exercise: Sorting Interfaces

func Sort(s Sortable){
//...
}

// What do you need to be able to sort this?
// data type
// data

//all you need are 3 methods below
type Sortable interface{
 Less(i, j int)bool
 Swap(i,j int)
 len()int

}
*/
