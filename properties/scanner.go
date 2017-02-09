package properties

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"text/scanner"
)

// Internal scanner of properties files. Used when Parse is called, has a cute little state machine
// (think regexp state-machine) to go through all possible states of a properties file.
type propertiesLexer struct {
	// Input and storage of current parse line.
	input scanner.Scanner
	key   bytes.Buffer
	value bytes.Buffer

	// All the parsed properties
	properties []Pair
	err        error
}

func newPropertiesLexer(r io.Reader) *propertiesLexer {
	var scan scanner.Scanner
	scan.Init(r)
	// Use idents to capture larger areas of identifiers. Makes scanning a bit more reasonable.
	scan.Mode = scanner.ScanComments | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanIdents
	scan.Whitespace = 0

	return &propertiesLexer{
		input:      scan,
		properties: []Pair{},
	}

}

// Called throughout the scan process to dump the key, value buffers into a Property Pair.
func (p *propertiesLexer) emitProperty() {
	pair := Pair{
		Key:   p.key.String(),
		Value: p.value.String(),
	}
	p.properties = append(p.properties, pair)
	p.key.Reset()
	p.value.Reset()
}

func (p *propertiesLexer) Run() []Pair {
	var state StateFunc

	for state = startLineState; state != nil; {
		state = state(p)
	}

	return p.properties
}
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

type StateFunc func(*propertiesLexer) StateFunc

func startLineState(p *propertiesLexer) StateFunc {

	for {
		tok := p.input.Scan()
		switch tok {
		case '#':
			return commentState
		case '!':
			return commentState
		case ' ':
			continue
		case '\n':
			continue
		case '\t':
			// Let's just ignore beginning whitespace:
			continue
		case '\r':
			futureTok := p.input.Peek()
			// Only valid use case.
			if futureTok != '\n' {
				return errorState("Invalid whitespace \r character", p)
			} else {
				continue
			}
		case '\\':
			// TODO - support a feature if it is allowed in the spec.
			return errorState("We do not support escaping at the beginning of the line.", p)

		case scanner.EOF:
			// Yay, we are done!
			return nil
		// Wow, we allow most characters. TODO - may need more validation on other chars..
		default:
			p.key.Write([]byte(p.input.TokenText()))
			// There may be more to write of the key. We aren't sure!
			return keyState

		}

	}
}

// Any time you enter the error state, always have the
// TokenText from the last Scanner.Scan(), to help with output.
func errorState(message string, p *propertiesLexer) StateFunc {
	p.err = fmt.Errorf("Parsing failed at %d, %d. %s\n", p.input.Line, p.input.Column, message)
	return nil
}

// When the scanner is reading in a "key". You can get into sticky situations here around escaped characters,
// which we can add support for later.
func keyState(p *propertiesLexer) StateFunc {
	for {
		tok := p.input.Scan()
		switch tok {
		// Weird case since pretty sure most files contain a newline at the end, but eh.
		case scanner.EOF:
			errorState("No value present", p)
			return nil
		case '\n':
			return errorState("No value present", p)
		case '\\':
			// Accept anything after this.
			p.key.Write([]byte(p.input.TokenText()))
			escapedTok := p.input.Scan()
			// TODO - I wonder if \r\n gets escaped properly.. May have to check in the future
			if escapedTok == scanner.EOF {
				// Wow, someone is an asshole.
				return nil
			} else {

				p.key.Write([]byte(p.input.TokenText()))
				continue
			}
		case ' ':
			// Try to consume more
			nextTok := p.input.Peek()
			if nextTok == '=' {
				// Consume =
				p.input.Scan()
				nextTok = p.input.Scan()
				if nextTok != ' ' {
					return errorState("key-value separator ' = ' should have equal spaces around it.", p)
				}
			}
			// TODO - should do a more rigorous test of acceptable characters here.
			return valueState

		case '=':
			return valueState
		case ':':
			return valueState
		default:
			p.key.Write([]byte(p.input.TokenText()))
		}
	}
}

// Accept the universe until a new line. Only time to ignore a new line is escape characters!
func valueState(p *propertiesLexer) StateFunc {
	for {
		tok := p.input.Scan()
		switch tok {
		// Weird case since pretty sure most files contain a newline at the end, but eh.
		case scanner.EOF:
			p.emitProperty()
			return nil
		case '\n':
			p.emitProperty()
			return startLineState
		case '\\':
			// Accept anything after this.
			p.value.Write([]byte(p.input.TokenText()))
			escapedTok := p.input.Scan()
			// TODO - I wonder if \r\n gets escaped properly.. May have to check in the future
			if escapedTok == scanner.EOF {
				// Wow, someone is mean.
				p.emitProperty()
				return nil
			} else {

				p.value.Write([]byte(p.input.TokenText()))
				continue
			}
		default:
			p.value.Write([]byte(p.input.TokenText()))
		}
	}
}

// Accept the universe until a new line. No exceptions here.
func commentState(p *propertiesLexer) StateFunc {
	for {
		tok := p.input.Scan()
		switch tok {
		// Weird case since pretty sure most files contain a newline at the end, but eh.
		case scanner.EOF:
			return nil
		case '\n':
			return startLineState
		}
	}
}
