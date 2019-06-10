package prompt

import "strings"

// Transform is the function applied to the text input on prompt.
var Transform = func(s string) string {
	input := strings.TrimSpace(s)
	return strings.ToLower(input)
}
