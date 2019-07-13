package prompt // import "gsr.dev/prompt"

import (
	"bufio"
	"io"
)

// Option is a functional option that enables
// setting private properties to a Prompt pointer.
type Option func(*Prompt)

// Stdin sets a new input buffer.
func Stdin(r io.Reader) Option {
	return func(p *Prompt) {
		p.sc = bufio.NewScanner(r)
	}
}
