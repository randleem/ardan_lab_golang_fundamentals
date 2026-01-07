package nlp

import (
	"regexp"
	"strings"

	"github.com/ardanlabs/nlp/stemmer"
)

// "Who's on first?" -> [Who s on first]
var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

// Tokenize returns tokens (lowercase) found in text.
func Tokenize(text string) []string {
	if text == "" {
		return []string{}
	}
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		token = stemmer.Stem(token)
		if token != "" {
			tokens = append(tokens, token)
		}
	}
	return tokens
}
