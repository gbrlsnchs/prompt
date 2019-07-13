# prompt
[![CircleCI](https://circleci.com/gh/gbrlsnchs/prompt.svg?style=shield)](https://circleci.com/gh/gbrlsnchs/prompt)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/prompt?status.svg)](https://godoc.org/github.com/gbrlsnchs/prompt)

Agnostic prompt for Go.

## Usage
### Example
#### Confirming input from stdin
```go
import "gsr.dev/prompt"

func main() {
	p := prompt.New()
	confirm := p.Confirm(prompt.Inputs{
		"yes": true,
		"no":  false,
	})
}
```

#### Using options
##### Case insensitive prompt
```go
import "gsr.dev/prompt"

func main() {
	p := prompt.New(prompt.CaseInsensitive)
	confirm := p.Confirm(prompt.Inputs{
		"yes": true,
		"no":  false,
	})
}
```

##### Reading from another source
```go
import (
	"strings"

	"gsr.dev/prompt"
)

func main() {
	r := strings.NewReader("yes\n")
	p := prompt.New(prompt.ReadFrom(r))
	confirm := p.Confirm(prompt.Inputs{
		"yes": true,
		"no":  false,
	})
}
```
