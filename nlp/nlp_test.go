package nlp

import (
	"os"
	"testing"

	// require stops the code if a test fails
	// Require shows you a diff if a test errors, to diagnose bugs
	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	_ = require.New
	// setup: call a function
	// teardown: defer/t.Cleanup
	text := "Who's on first?"
	tokens := Tokenize(text)
	expected := []string{"who", "s", "on", "first"}
	require.Equal(t, expected, tokens)

	/* Code beofre Testify
	if !slices.Equal(expected, tokens) {
		t.Fatalf("expected %#v, got %#v", expected, tokens)
	}
	*/
}

func TestTokenizeTable(t *testing.T) {
	cases := []struct {
		text   string
		tokens []string
	}{
		{"who's on first", []string{"who", "s", "on", "first"}},
		{"what's on second", []string{"what", "s", "on", "second"}},
		{"", nil},
	}
	for _, tc := range cases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
			/* Code beofre Testify
			if !slices.Equal(tc.tokens, tokens) {
				t.Fatalf("expected %#v, got %#v", tc.tokens, tokens)
			}
			*/
		})
	}
}

/*
Selecting test (can skip test of have them under specific circumstances e.g. a certain environment of if a certain set up is being ran - e.g. all microservices are running)
- "-run" flag: regexp
- build tags - e.g. "//go:build ui" at the top of the file, it will only run test if the build tag is activated
- environment variables,
*/

// In Jenkins use BUILD_NUMBER,
var inCI = os.Getenv("CI") != ""

func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("not in CI")
	}
}
