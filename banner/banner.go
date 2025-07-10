package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main(){
	banner("Go", 6)

	s:= "G♡"
	banner(s, 6)
	fmt.Println("len:", len(s))
	fmt.Println("s[1]:", s[1])
	fmt.Printf("s[1]: %c\n", s[1])

	for i, c := range s{
		fmt.Printf("%c at %d\n", c, i)
	}

	/*
	strings are UTF-8 encoded
	len, s[]: byte (unit8)
	for : rune (int)
	*/
}

func banner(text string, width int){
	padding:=(width-utf8.RuneCountInString(text))/2
	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-",width))
}
