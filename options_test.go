package prompt_test

import (
	"fmt"
	"os"
	"testing"

	"gsr.dev/prompt"
)

func TestOptions(t *testing.T) {
	t.Run("CaseInsensitive", func(t *testing.T) {
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
				want: prompt.None,
			},
			{
				name: "N",
				in:   testIn,
				want: prompt.None,
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
				want: prompt.None,
			},
			{
				name: "NO",
				in:   testIn,
				want: prompt.None,
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
				want: prompt.None,
			},
			{
				name: "__NO__",
				in:   testIn,
				want: prompt.None,
			},
			{
				name: "foo",
				in:   testIn,
				want: prompt.None,
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
				if want, got := tc.want, p.Answer(tc.in); got != want {
					t.Errorf("want %d, got %d", want, got)
				}
			})
		}
	})
}
