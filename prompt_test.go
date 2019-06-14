package prompt_test

import (
	"io"
	"strings"
	"testing"

	"gsr.dev/prompt"
)

var testInputs = map[string]bool{
	"y":   true,
	"yes": true,
	"n":   false,
	"no":  false,
}

func TestPromptConfirm(t *testing.T) {
	testCases := []struct {
		r      io.Reader
		inputs map[string]bool
		want   bool
	}{
		// default "y"
		{
			r:      strings.NewReader("y\n"),
			inputs: testInputs,
			want:   true,
		},
		// default "n"
		{
			r:      strings.NewReader("n\n"),
			inputs: testInputs,
			want:   false,
		},
		// capital "y"
		{
			r:      strings.NewReader("Y\n"),
			inputs: testInputs,
			want:   true,
		},
		// capital "n"
		{
			r:      strings.NewReader("N\n"),
			inputs: testInputs,
			want:   false,
		},
		// default "yes"
		{
			r:      strings.NewReader("yes\n"),
			inputs: testInputs,
			want:   true,
		},
		// default "no"
		{
			r:      strings.NewReader("no\n"),
			inputs: testInputs,
			want:   false,
		},
		// capital "yes"
		{
			r:      strings.NewReader("YES\n"),
			inputs: testInputs,
			want:   true,
		},
		// capital "no"
		{
			r:      strings.NewReader("NO\n"),
			inputs: testInputs,
			want:   false,
		},
		// default whitespaced "yes"
		{
			r:      strings.NewReader("  yes  \n"),
			inputs: testInputs,
			want:   true,
		},
		// default whitespaced "no"
		{
			r:      strings.NewReader("  no  \n"),
			inputs: testInputs,
			want:   false,
		},
		// capital whitespaced "yes"
		{
			r:      strings.NewReader("  YES  \n"),
			inputs: testInputs,
			want:   true,
		},
		// capital whitespaced "no"
		{
			r:      strings.NewReader("  NO  \n"),
			inputs: testInputs,
			want:   false,
		},
		// none
		{
			r:      strings.NewReader("foo\n"),
			inputs: testInputs,
			want:   false,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			p := prompt.New(tc.r)
			if want, got := tc.want, p.Confirm(tc.inputs); want != got {
				t.Errorf("want %t, got %t", want, got)
			}
		})
	}
}

func TestPromptConfirmStatus(t *testing.T) {
	testCases := []struct {
		r      io.Reader
		inputs map[string]bool
		want   prompt.Status
	}{
		// default "y"
		{
			r:      strings.NewReader("y\n"),
			inputs: testInputs,
			want:   prompt.StatusAccept,
		},
		// default "n"
		{
			r:      strings.NewReader("n\n"),
			inputs: testInputs,
			want:   prompt.StatusDecline,
		},
		// capital "y"
		{
			r:      strings.NewReader("Y\n"),
			inputs: testInputs,
			want:   prompt.StatusAccept,
		},
		// capital "n"
		{
			r:      strings.NewReader("N\n"),
			inputs: testInputs,
			want:   prompt.StatusDecline,
		},
		// default "yes"
		{
			r:      strings.NewReader("yes\n"),
			inputs: testInputs,
			want:   prompt.StatusAccept,
		},
		// default "no"
		{
			r:      strings.NewReader("no\n"),
			inputs: testInputs,
			want:   prompt.StatusDecline,
		},
		// capital "yes"
		{
			r:      strings.NewReader("YES\n"),
			inputs: testInputs,
			want:   prompt.StatusAccept,
		},
		// capital "no"
		{
			r:      strings.NewReader("NO\n"),
			inputs: testInputs,
			want:   prompt.StatusDecline,
		},
		// default whitespaced "yes"
		{
			r:      strings.NewReader("  yes  \n"),
			inputs: testInputs,
			want:   prompt.StatusAccept,
		},
		// default whitespaced "no"
		{
			r:      strings.NewReader("  no  \n"),
			inputs: testInputs,
			want:   prompt.StatusDecline,
		},
		// capital whitespaced "yes"
		{
			r:      strings.NewReader("  YES  \n"),
			inputs: testInputs,
			want:   prompt.StatusAccept,
		},
		// capital whitespaced "no"
		{
			r:      strings.NewReader("  NO  \n"),
			inputs: testInputs,
			want:   prompt.StatusDecline,
		},
		// none
		{
			r:      strings.NewReader("foo\n"),
			inputs: testInputs,
			want:   prompt.StatusNone,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			p := prompt.New(tc.r)
			if want, got := tc.want, p.ConfirmStatus(tc.inputs); want != got {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}

func TestPromptResponse(t *testing.T) {
	testCases := []struct {
		r    io.Reader
		want string
	}{
		{strings.NewReader("test\n"), "test"},
		{strings.NewReader("FooBar\n"), "FooBar"},
	}
	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			p := prompt.New(tc.r)
			if want, got := tc.want, p.Response(); want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}
