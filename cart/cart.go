package main

import (
	"fmt"
	"sort"
)

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println("len: ", len(cart))
	fmt.Println("cart[1]: ", cart[1])

	// indexes
	for i := range cart {
		fmt.Println(i)
	}
	// indexes and values
	for i, c := range cart {
		fmt.Println(i, c)
	}

	// values
	for _, c := range cart {
		fmt.Println(c)
	}

	cart = append(cart, "milk")
	fmt.Println(cart)

	// slicing operator, half-open
	fruit := cart[:3]
	fmt.Println(fruit)
	fruit = append(fruit, "lemon")
	fmt.Println("fruit", fruit)
	fmt.Println("cart", cart)

	var s []int
	for i := range 10_000 {
		s = appnendInt(s, i)
	}
	fmt.Println(s[:10])

	// Exercise: concat without using a for loop
	out := concat([]string{"A", "B"}, []string{"C"})
	fmt.Println(out) // [A B C]

	/*Exercise: Median
	- sort values
	- if odd number of values: return middle
	- return average of middles
	*/
	values := []float64{3, 1, 2} // 2
	fmt.Println(median(values))
	values = []float64{3, 1, 2, 4} // 2.5
	fmt.Println(median(values))
	fmt.Println("values: ", values)

	players := []Player{
		{"Rick", 10_000},
		{"Morty", 11},
	}

	// Add a bonus

	// Value semantics "for" loop
	for _, p := range players {
		p.Score += 100
	}
	fmt.Println(players)

	// "Pointer" semantics "for" loop
	for i := range players {
		players[i].Score += 100
	}
	fmt.Println(players)
}

type Player struct {
	Name  string
	Score int
}

func median(values []float64) float64 {
	nv := make([]float64, len(values))
	copy(nv, values)
	sort.Float64s(nv)
	if len(nv)%2 != 0 {
		return nv[(len(nv)-1)/2]
	}
	return (nv[len(nv)/2] + nv[(len(values)-1)/2]) / 2
}

func concat(s1, s2 []string) []string {
	// return slices.Concat(s1, s2)
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func appnendInt(s []int, v int) []int {
	i := len(s)
	if len(s) == cap(s) {
		// nore more space in underlying array
		// NEED TO REALLOCATE AND COPY
		size := 2 * (len(s) + 1)
		fmt.Println(cap(s), "->", size)
		ns := make([]int, size)
		copy(ns, s)
		s = ns[:len(s)]
	}
	s = s[:len(s)+1]
	s[i] = v
	return s
}
