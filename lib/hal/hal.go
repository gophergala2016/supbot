package hal

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/gophergala2016/supbot/lib/git"
	"github.com/gophergala2016/supbot/lib/sup"
)

var space = []byte(` `)

// Making sure this is a reader.
var _ = io.Writer(&Hal{})

var (
	msgMissingCommand = `Say what?!`
)

type Hal struct {
	out io.Writer
	cwd string // current working directory.
}

var (
	errMissingCommand    = errors.New(`Missing command.`)
	errIncompleteCommand = errors.New(`Incomplete command.`)
)

func NewHal(out io.Writer) *Hal {
	return &Hal{
		out: out,
	}
}

func (h *Hal) Write(cmd []byte) (n int, err error) {
	l := len(cmd)

	chunks := bytes.Split(cmd, space)

	if len(chunks) < 1 {
		h.out.Write([]byte(msgMissingCommand))
		return l, errMissingCommand
	}

	s := string(chunks[0])

	switch s {
	case "help":
		h.out.Write([]byte(`[repository]/[branch] [network] [target]`))
		return l, nil
	default:
		if len(chunks) > 0 {
			switch string(chunks[0]) {
			case "set-repo":
				if len(chunks) > 1 {
					cwd := string(chunks[1])
					if cwd != "" {
						// TODO: check this is an actual repo.
						h.cwd = cwd
						h.out.Write([]byte(fmt.Sprintf("You current repo is %q", h.cwd)))
						return l, nil
					}
				}
				h.out.Write([]byte(fmt.Sprintf("Try `set-repo [repo-url]`")))
				return l, errMissingCommand
			}
			if h.cwd != "" {

				// TODO: grab branch name from URL, if any.

				repo, err := git.Clone(h.cwd)
				if err != nil {
					return l, err
				}

				if err := repo.Checkout("master"); err != nil {
					return l, err
				}

				// TODO: insert sup magic here.
				var outbuf bytes.Buffer
				cmd := sup.NewSup(&outbuf).Setwd(repo.Dir())

				if len(chunks) > 0 {
					cmd.Network(string(chunks[0]))
				}
				if len(chunks) > 1 {
					cmd.Target(string(chunks[1]))
				}

				err = cmd.Exec()

				h.out.Write(outbuf.Bytes())
				return l, err
			} else {
				h.out.Write([]byte(fmt.Sprintf("Missing repo, try `set-repo [repo-url]`")))
			}
			return l, errMissingCommand
		}
	}

	return l, nil
}
