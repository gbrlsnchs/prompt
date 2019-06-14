package prompt_test

import (
	"testing"

	"gsr.dev/prompt"
)

func TestTransform(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"test", "test"},
		{"  TEST  ", "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			if want, got := tc.want, prompt.Transform(tc.s); want != got {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}
