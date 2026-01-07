package nlp_test

import (
	"fmt"

	"github.com/ardanlabs/nlp"
)

func ExampleTokenize() {
	tokens := nlp.Tokenize("who's on first?")
	fmt.Println(tokens)

	// Output:
	// [who on first]
}
