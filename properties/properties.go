package properties

// Package properties deals with parsing .properties files are mostly used in java~ish applications.
// Currently there is only support for "=" spaces, and escaped characters are still there :(

import (
	"io"
)

// Pair is a key-value pair pulled from a properties file.
type Pair struct {
	Key   string
	Value string
}

// Parse can parse and read property files into maps.
func Parse(input io.Reader) ([]Pair, error) {
	lexer := newPropertiesLexer(input)
	props := lexer.Run()
	return props, lexer.err
}

// ParseIntoMap is a convenience method that disregards ordering in favor of getting the values you need.
// useful for just "getting things"
func ParseIntoMap(input io.Reader) (map[string]string, error) {
	output, err := Parse(input)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for _, pair := range output {
		m[pair.Key] = pair.Value
	}

	return m, nil
}
