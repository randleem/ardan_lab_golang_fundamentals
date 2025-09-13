package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// What are the N most common words in Sherlock.txt
// a `a` is a "raw string", eg at \ is just a \
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Code that runs before main, must not cotain errors
// - var expressions
// - init function

func main() {
	// mapDemo()
	file, err := os.Open("sherlock.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	// Iterates over line by line
	s := bufio.NewScanner(file)

	freq := make(map[string]int) // word->count
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, word := range words {
			freq[strings.ToLower(word)]++
		}
	}

	if err := s.Err(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	top := topN(freq, 10)
	fmt.Println(top)
}

func topN(freq map[string]int, n int) []string {
	// FIX: Your code goes here
	words := slices.Collect(maps.Keys(freq))
	sort.Slice(words, func(i, j int) bool {
		wi, wj := words[i], words[j]
		// sort words in reverse order
		return freq[wi] > freq[wj]
	})
	n = min(n, len(words))
	return words[:n]
}

func mapDemo() {
	heros := map[string]string{ // hero->name
		"Superman":     "Clark",
		"Wonder Woman": "Diana",
		"Batman":       "Bruce",
	}
	for k := range heros {
		fmt.Println(k)
	}
	// Key + Value
	for k, v := range heros {
		fmt.Println(v, "is", k)
	}

	// for values, use _
	for _, v := range heros {
		fmt.Println(v)
	}

	n := heros["Batman"]
	fmt.Println(n)

	// returns empty string, when acesing non-existing key returns zero value
	n = heros["Aquaman"]
	fmt.Println(n)
	fmt.Printf("%q\n", n)

	// Use comma, ok to find if key is in map
	n, ok := heros["Spiderman"]
	if ok {
		fmt.Printf("%q\n", n)
	} else {
		fmt.Println("Spiderman not found")
	}

	delete(heros, "Batman")
	fmt.Println(heros)
}
