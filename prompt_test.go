package prompt_test

import (
	"io"
	"strings"
	"testing"

	"github.com/gbrlsnchs/prompt"
)

func TestPromptConfirm(t *testing.T) {
	testCases := []struct {
		r       io.Reader
		accept  []string
		decline []string

		msg  string
		want bool
	}{
		// default "y"
		{
			r:       strings.NewReader("y\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Confirm? [y(es)/n(o)] ",
			want:    true,
		},
		// default "n"
		{
			r:       strings.NewReader("n\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Are you sure? [y(es)/n(o)] ",
			want:    false,
		},
		// capital "y"
		{
			r:       strings.NewReader("Y\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Is that right? [y(es)/n(o)] ",
			want:    true,
		},
		// capital "n"
		{
			r:       strings.NewReader("N\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Accept all? [y(es)/n(o)] ",
			want:    false,
		},
		// default "yes"
		{
			r:       strings.NewReader("yes\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Confirm? [y(es)/n(o)] ",
			want:    true,
		},
		// default "no"
		{
			r:       strings.NewReader("no\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Are you sure? [y(es)/n(o)] ",
			want:    false,
		},
		// capital "yes"
		{
			r:       strings.NewReader("YES\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Is that right? [y(es)/n(o)] ",
			want:    true,
		},
		// capital "no"
		{
			r:       strings.NewReader("NO\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Accept all? [y(es)/n(o)] ",
			want:    false,
		},
		// default whitespaced "yes"
		{
			r:       strings.NewReader("  yes  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Confirm? [y(es)/n(o)] ",
			want:    true,
		},
		// default whitespaced "no"
		{
			r:       strings.NewReader("  no  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Are you sure? [y(es)/n(o)] ",
			want:    false,
		},
		// capital whitespaced "yes"
		{
			r:       strings.NewReader("  YES  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Is that right? [y(es)/n(o)] ",
			want:    true,
		},
		// capital whitespaced "no"
		{
			r:       strings.NewReader("  NO  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Accept all? [y(es)/n(o)] ",
			want:    false,
		},
		// none
		{
			r:       strings.NewReader("foo\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Please answer yes or no: ",
			want:    false,
		},
	}
	var stdin strings.Builder
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			stdin.Reset()
			p := prompt.New(tc.r, &stdin)
			p.SetAccept(tc.accept...)
			p.SetDecline(tc.decline...)
			if want, got := tc.want, p.Confirm(tc.msg); want != got {
				t.Errorf("want %t, got %t", want, got)
			}
			if want, got := tc.msg, stdin.String(); want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}

func TestPromptConfirmStatus(t *testing.T) {
	testCases := []struct {
		r       io.Reader
		accept  []string
		decline []string

		msg  string
		want prompt.Status
	}{
		// default "y"
		{
			r:       strings.NewReader("y\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Confirm? [y(es)/n(o)] ",
			want:    prompt.StatusAccept,
		},
		// default "n"
		{
			r:       strings.NewReader("n\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Are you sure? [y(es)/n(o)] ",
			want:    prompt.StatusDecline,
		},
		// capital "y"
		{
			r:       strings.NewReader("Y\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Is that right? [y(es)/n(o)] ",
			want:    prompt.StatusAccept,
		},
		// capital "n"
		{
			r:       strings.NewReader("N\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Accept all? [y(es)/n(o)] ",
			want:    prompt.StatusDecline,
		},
		// default "yes"
		{
			r:       strings.NewReader("yes\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Confirm? [y(es)/n(o)] ",
			want:    prompt.StatusAccept,
		},
		// default "no"
		{
			r:       strings.NewReader("no\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Are you sure? [y(es)/n(o)] ",
			want:    prompt.StatusDecline,
		},
		// capital "yes"
		{
			r:       strings.NewReader("YES\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Is that right? [y(es)/n(o)] ",
			want:    prompt.StatusAccept,
		},
		// capital "no"
		{
			r:       strings.NewReader("NO\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Accept all? [y(es)/n(o)] ",
			want:    prompt.StatusDecline,
		},
		// default whitespaced "yes"
		{
			r:       strings.NewReader("  yes  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Confirm? [y(es)/n(o)] ",
			want:    prompt.StatusAccept,
		},
		// default whitespaced "no"
		{
			r:       strings.NewReader("  no  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Are you sure? [y(es)/n(o)] ",
			want:    prompt.StatusDecline,
		},
		// capital whitespaced "yes"
		{
			r:       strings.NewReader("  YES  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Is that right? [y(es)/n(o)] ",
			want:    prompt.StatusAccept,
		},
		// capital whitespaced "no"
		{
			r:       strings.NewReader("  NO  \n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Accept all? [y(es)/n(o)] ",
			want:    prompt.StatusDecline,
		},
		// none
		{
			r:       strings.NewReader("foo\n"),
			accept:  []string{"y", "yes"},
			decline: []string{"n", "no"},
			msg:     "Please answer yes or no: ",
			want:    prompt.StatusNone,
		},
	}
	var stdin strings.Builder
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			stdin.Reset()
			p := prompt.New(tc.r, &stdin)
			p.SetAccept(tc.accept...)
			p.SetDecline(tc.decline...)
			if want, got := tc.want, p.ConfirmStatus(tc.msg); want != got {
				t.Errorf("want %d, got %d", want, got)
			}
			if want, got := tc.msg, stdin.String(); want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}
