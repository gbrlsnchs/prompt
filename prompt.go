package prompt

import (
	"bufio"
	"io"
	"os"
)

// Prompt is a text prompter.
type Prompt struct {
	sc *bufio.Scanner
}

// New initializes a new prompt with proper writer and reader.
func New(r io.Reader) *Prompt {
	if r == nil {
		r = os.Stdin
	}
	return &Prompt{sc: bufio.NewScanner(r)}
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
	confirm, ok := inputs[Transform(in)]
	if !ok {
		return StatusNone
	}
	if confirm {
		return StatusAccept
	}
	return StatusDecline
}

// Response scans an input and returns it as it is.
func (p *Prompt) Response() string {
	p.sc.Scan()
	return p.sc.Text()
}
