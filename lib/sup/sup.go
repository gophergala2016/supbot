/*
Package sup is a library that allows you to set up and execute
a sup CLI command.

The Sup type contained within has methods that can be chained together:

	cmd := sup.New(ioWriter).SetWd("cwd").SetNetwork("local").SetTarget("deploy")
	cmd.Exec()

The working directory for the Sup should be set to a directory that
contains a Supfile.
*/
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

// Sup represents a CLI sup command. It has methods to set it up and
// execute it.
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

// New creates a new Sup. The io.Writer provided to this function will
// receive the output of the sup command.
func New(w io.Writer, wdir string) (*Sup, error) {
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

// SetNetwork sets the sup command to use one of the networks described
// in the Supfile.
func (s *Sup) SetNetwork(n string) *Sup {
	s.network = n
	return s
}

// SetWd sets the working directory where the sup command will be executed.
// This directory should contain a Supfile. If this is not set, this will
// default to the default Dir of an os.exec.Cmd.
func (s *Sup) SetWd(wdir string) *Sup {
	s.wd = wdir
	return s
}

// SetTarget sets the sup command to use one of the targets described in
// the Supfile.
func (s *Sup) SetTarget(t string) *Sup {
	s.target = t
	return s
}

// Exec executes the sup command. Any output will be transmitted to the
// io.Writer that was provided to New().
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
