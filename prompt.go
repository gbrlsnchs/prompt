package prompt

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Status is a response status.
type Status int

const (
	// StatusAccept means input matches accepting answers.
	StatusAccept Status = iota
	// StatusDecline means input matches declining answers.
	StatusDecline
	// StatusNone means input matches neither accepting nor declining answers.
	StatusNone
)

// Prompt is a text prompter.
type Prompt struct {
	want map[string]struct{}
	deny map[string]struct{}

	sc *bufio.Scanner
	w  io.Writer
}

// New initializes a new prompt with proper writer and reader.
func New(r io.Reader, w io.Writer) *Prompt {
	return &Prompt{
		want: make(map[string]struct{}),
		deny: make(map[string]struct{}),
		sc:   bufio.NewScanner(r),
		w:    w,
	}
}

// SetAccept sets all accepted answers.
func (p *Prompt) SetAccept(a ...string) {
	for _, aa := range a {
		p.want[aa] = struct{}{}
	}
}

// SetDecline sets all declining answers.
func (p *Prompt) SetDecline(a ...string) {
	for _, aa := range a {
		p.deny[aa] = struct{}{}
	}
}

// Confirm prompts a message and check whether the input is acceptable.
func (p *Prompt) Confirm(msg string) bool {
	return p.ConfirmStatus(msg) == StatusAccept
}

// ConfirmStatus prompts a message and returns a status depending on input.
func (p *Prompt) ConfirmStatus(msg string) Status {
	fmt.Fprint(p.w, msg)
	var input string
	p.sc.Scan()
	input = strings.TrimSpace(p.sc.Text())
	input = strings.ToLower(input)
	if _, ok := p.want[input]; ok {
		return StatusAccept
	}
	if _, ok := p.deny[input]; ok {
		return StatusDecline
	}
	return StatusNone
}
