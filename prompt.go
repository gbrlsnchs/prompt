package prompt // import "gsr.dev/prompt"

import (
	"bufio"
	"os"
	"strings"
)

// Answer is an input answer.
type Answer int

const (
	// True means input is truthy.
	True Answer = iota
	// False means input is falsy.
	False
	// None means input didn't match anything.
	None
)

// Inputs is a dictionary that stores
// whether an input should be accepted.
type Inputs map[string]bool

// Prompt is a text prompter.
type Prompt struct {
	sc *bufio.Scanner
	// Options.
	transform func(string) string
}

// New initializes a new prompt using r. If r is nil, os.Stdin is used.
func New(opts ...Option) *Prompt {
	p := Prompt{
		transform: func(s string) string {
			return strings.TrimSpace(s)
		},
	}
	for _, opt := range opts {
		opt(&p)
	}
	if p.sc == nil {
		p.sc = bufio.NewScanner(os.Stdin)
	}
	return &p
}

// Answer prompts a message and returns an answer depending on input.
// The input is trimmed by default.
func (p *Prompt) Answer(in Inputs) Answer {
	s := p.Text()
	if confirm, ok := in[p.transform(s)]; ok {
		if confirm {
			return True
		}
		return False
	}
	return None
}

// Confirm prompts a message and check whether the input is truthy.
// The input is trimmed by default.
func (p *Prompt) Confirm(in Inputs) bool {
	return p.Answer(in) == True
}

// Text scans an input and returns it as is.
func (p *Prompt) Text() string {
	p.sc.Scan()
	return p.sc.Text()
}
