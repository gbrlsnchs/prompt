package prompt

import (
	"bufio"
	"io"
	"os"
	"strings"
)
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
func (p *Prompt) Confirm(inputs map[string]bool) bool {
	return p.ConfirmStatus(inputs) == StatusAccept
}

// ConfirmStatus prompts a message and returns a status depending on input.
func (p *Prompt) ConfirmStatus(inputs map[string]bool) Status {
	var input string
	p.sc.Scan()
	input = strings.TrimSpace(p.sc.Text())
	input = strings.ToLower(input)
	confirm, ok := inputs[input]
	if !ok {
		return StatusNone
	}
	if confirm {
		return StatusAccept
	}
	return StatusDecline
}
