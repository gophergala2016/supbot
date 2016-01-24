package sup

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	stackup "github.com/gophergala2016/supbot/Godeps/_workspace/src/github.com/pressly/sup"
)

type Sup struct {
	// stackup app
	*stackup.Stackup
	// loaded Supfile config
	config *stackup.Supfile

	network string
	target  string
	wd      string
	writer  io.Writer
}

var (
	ErrInvalidTarget  = errors.New("invalid target")
	ErrInvalidNetwork = errors.New("invalid network")
)

func stripColor(msg string) string {
	msg = strings.TrimPrefix(msg, stackup.ResetColor)
	for _, c := range stackup.Colors {
		msg = strings.TrimPrefix(msg, c)
	}
	return msg
}

func (s *Sup) Network(n string) *Sup {
	s.network = n
	return s
}

func (s *Sup) Target(t string) *Sup {
	s.target = t
	return s
}

func (s *Sup) Exec() error {
	//TODO: stderr capture
	network, ok := s.config.Networks[s.network]
	if !ok {
		return ErrInvalidNetwork
	}
	command, isCommand := s.config.Commands[s.target]
	if !isCommand {
		return ErrInvalidTarget
	}
	cmds := []*stackup.Command{&command}

	// Do some piping magic here
	old := os.Stdout
	read, write, _ := os.Pipe()

	os.Stdout = write
	err := s.Run(&network, cmds...)
	write.Close()

	scanner := bufio.NewScanner(read)
	var out string
	for scanner.Scan() {
		out = fmt.Sprintf("%s %s\n", out, stripColor(scanner.Text()))
	}
	os.Stdout = old // reset stdout

	_, err = s.writer.Write([]byte(out))
	return err
}

// TODO: Pass in a command directly
// func (s *sup) Cmd() {
// err2 := sup.NewSup(io.Writer).Cmd("Some sup command")
// }

func NewSup(w io.Writer, wdir string) (*Sup, error) {
	// load the supfile
	conf, err := stackup.NewSupfile(fmt.Sprintf("%s/Supfile", wdir))
	if err != nil {
		return nil, err
	}

	// create new Stackup app.
	app, err := stackup.New(conf)
	if err != nil {
		return nil, err
	}

	return &Sup{Stackup: app, writer: w, config: conf}, nil
}
