package prompt_test

import (
	"fmt"
	"os"
	"testing"

	"gsr.dev/prompt"
)

var testIn = prompt.Inputs{
	"y":   true,
	"yes": true,
	"n":   false,
	"no":  false,
}

func TestPrompt(t *testing.T) {
	t.Run("ConfirmStatus", func(t *testing.T) {
		testCases := []struct {
			name string
			in   prompt.Inputs
			want prompt.Answer
		}{
			{
				name: "y",
				in:   testIn,
				want: prompt.True,
			},
			{
				name: "n",
				in:   testIn,
				want: prompt.False,
			},
			{
				name: "Y",
				in:   testIn,
				want: prompt.Undefined,
			},
			{
				name: "N",
				in:   testIn,
				want: prompt.Undefined,
			},
			{
				name: "yes",
				in:   testIn,
				want: prompt.True,
			},
			{
				name: "no",
				in:   testIn,
				want: prompt.False,
			},
			{
				name: "YES",
				in:   testIn,
				want: prompt.Undefined,
			},
			{
				name: "NO",
				in:   testIn,
				want: prompt.Undefined,
			},
			{
				name: "__yes__",
				in:   testIn,
				want: prompt.True,
			},
			{
				name: "__no__",
				in:   testIn,
				want: prompt.False,
			},
			{
				name: "__YES__",
				in:   testIn,
				want: prompt.Undefined,
			},
			{
				name: "__NO__",
				in:   testIn,
				want: prompt.Undefined,
			},
			{
				name: "foo",
				in:   testIn,
				want: prompt.Undefined,
			},
		}
		for _, tc := range testCases {
			t.Run("", func(t *testing.T) {
				f, err := os.Open(fmt.Sprintf("testdata/inputs/%s.txt", tc.name))
				if err != nil {
					t.Fatal(err)
				}
				os.Stdin = f
				p := prompt.New()
				if want, got := tc.want, p.Answer(tc.in); want != got {
					t.Errorf("want %d, got %d", want, got)
				}
			})
		}
	})

	t.Run("Confirm", func(t *testing.T) {
		testCases := []struct {
			name string
			in   prompt.Inputs
			want bool
		}{
			{
				name: "y",
				in:   testIn,
				want: true,
			},
			{
				name: "n",
				in:   testIn,
				want: false,
			},
			{
				name: "Y",
				in:   testIn,
				want: false,
			},
			{
				name: "N",
				in:   testIn,
				want: false,
			},
			{
				name: "yes",
				in:   testIn,
				want: true,
			},
			{
				name: "no",
				in:   testIn,
				want: false,
			},
			{
				name: "YES",
				in:   testIn,
				want: false,
			},
			{
				name: "NO",
				in:   testIn,
				want: false,
			},
			{
				name: "__yes__",
				in:   testIn,
				want: true,
			},
			{
				name: "__no__",
				in:   testIn,
				want: false,
			},
			{
				name: "__YES__",
				in:   testIn,
				want: false,
			},
			{
				name: "__NO__",
				in:   testIn,
				want: false,
			},
			{
				name: "foo",
				in:   testIn,
				want: false,
			},
		}
		for _, tc := range testCases {
			f, err := os.Open(fmt.Sprintf("testdata/inputs/%s.txt", tc.name))
			if err != nil {
				t.Fatal(err)
			}
			os.Stdin = f
			t.Run(tc.name, func(t *testing.T) {
				p := prompt.New()
				if want, got := tc.want, p.Confirm(tc.in); want != got {
					t.Errorf("want %t, got %t", want, got)
				}
			})
		}
	})

	t.Run("Text", func(t *testing.T) {
		testCases := []struct {
			name string
			want string
		}{
			{
				name: "y",
				want: "y",
			},
			{
				name: "n",
				want: "n",
			},
			{
				name: "Y",
				want: "Y",
			},
			{
				name: "N",
				want: "N",
			},
			{
				name: "yes",
				want: "yes",
			},
			{
				name: "no",
				want: "no",
			},
			{
				name: "YES",
				want: "YES",
			},
			{
				name: "NO",
				want: "NO",
			},
			{
				name: "__yes__",
				want: "  yes  ",
			},
			{
				name: "__no__",
				want: "  no  ",
			},
			{
				name: "__YES__",
				want: "  YES  ",
			},
			{
				name: "__NO__",
				want: "  NO  ",
			},
			{
				name: "foo",
				want: "foo",
			},
		}
		for _, tc := range testCases {
			t.Run(tc.want, func(t *testing.T) {
				f, err := os.Open(fmt.Sprintf("testdata/inputs/%s.txt", tc.name))
				if err != nil {
					t.Fatal(err)
				}
				os.Stdin = f
				p := prompt.New()
				if want, got := tc.want, p.Text(); want != got {
					t.Errorf("want %q, got %q", want, got)
				}
			})
		}
	})
}
