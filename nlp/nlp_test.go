package nlp

import (
	"io"
	"os"
	"strings"
	"testing"

	// require stops the code if a test fails
	// Require shows you a diff if a test errors, to diagnose bugs
	"github.com/BurntSushi/toml"
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

// Exercise: Resd test cased from tokenize_cases.toml
// instead of in-memory slice
func TestTokenizeTable(t *testing.T) {
	// cases := []struct {
	// 	text   string
	// 	tokens []string
	// }{
	// 	{"who's on first", []string{"who", "s", "on", "first"}},
	// 	{"what's on second", []string{"what", "s", "on", "second"}},
	// 	{"", nil},
	// }
	type testCases struct {
		Text   string
		Tokens []string
		Name   string
	}
	file, err := os.Open("./testdata/tokenize_cases.toml")
	require.NoError(t, err)
	defer file.Close()

	b, err := io.ReadAll(file)
	require.NoError(t, err)
	var cases struct {
		Case []testCases `toml:"case"`
	}
	err = toml.Unmarshal([]byte(b), &cases)
	require.NoError(t, err)

	// Instructors method
	// cases := toml.NewDecoder(file)
	//_, err = dec.Decode(&cases)

	for _, tc := range cases.Case {
		name := tc.Name
		if name == "" {
			name = tc.Text
		}
		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
			require.Equal(t, tc.Tokens, tokens)
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

// FUZZ Test

func FuzzTokenizer(f *testing.F) {
	f.Add("") // can add own test data - empty string
	fn := func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)
		for _, tok := range tokens {
			require.Contains(t, lText, tok)
		}
	}
	f.Fuzz(fn)
}
