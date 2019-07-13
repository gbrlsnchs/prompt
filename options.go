package prompt // import "gsr.dev/prompt"

import (
	"bufio"
	"io"
	"strings"
)

// Option is a functional option that enables
// setting private properties to a Prompt pointer.
type Option func(*Prompt)

// CaseInsensitive transforms answers to lowercase.
func CaseInsensitive(p *Prompt) {
	prev := p.transform
	p.transform = func(s string) string {
		return strings.ToLower(prev(s))
	}
}

// ReadFrom sets a new input buffer.
func ReadFrom(r io.Reader) Option {
	return func(p *Prompt) {
		p.sc = bufio.NewScanner(r)
	}
}
