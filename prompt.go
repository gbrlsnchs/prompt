package prompt // import "gsr.dev/prompt"

import (
	"bufio"
	"os"
	"strings"
)

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

// Confirm prompts a message and check whether the input is acceptable.
// The input is transformed using the Transform function.
func (p *Prompt) Confirm(inputs map[string]bool) bool {
	return p.ConfirmStatus(inputs) == StatusAccept
}

// ConfirmStatus prompts a message and returns a status depending on input.
// The input is transformed using the Transform function.
func (p *Prompt) ConfirmStatus(inputs map[string]bool) Status {
	in := p.Response()
	if confirm, ok := inputs[Transform(in)]; ok {
		if confirm {
			return StatusAccept
		}
		return StatusDecline
	}
	return StatusNone
}

// Response scans an input and returns it as it is.
func (p *Prompt) Response() string {
	p.sc.Scan()
	return p.sc.Text()
}
