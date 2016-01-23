package hal

import (
	"bytes"
	"fmt"
	"io"
)

var space = []byte(` `)

// Making sure this is a reader.
var _ = io.Writer(&Hal{})

type Hal struct {
	out io.Writer
}

func NewHal(out io.Writer) *Hal {
	return &Hal{
		out: out,
	}
}

func (h *Hal) Write(cmd []byte) (n int, err error) {
	l := len(cmd)

	chunks := bytes.Split(cmd, space)
	if len(chunks) <= 1 {
		s := string(chunks[0])

		switch s {
		case "help":
			h.out.Write([]byte(`[repository]/[branch] [network] [target]`))
			return l, nil
		default:
			h.out.Write([]byte(fmt.Sprintf("Unknown command %q", string(cmd))))
			return l, nil
		}
	}

	return l, nil
}
