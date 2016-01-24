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
	"bytes"
	"io"
	"log"
	"os/exec"
	"strings"

	s "github.com/gophergala2016/supbot/Godeps/_workspace/src/github.com/pressly/sup"
)

var supCommand string

var (
	_ s.Stackup // NOTE: godeps..
)

// Sup represents a CLI sup command. It has methods to set it up and
// execute it.
type Sup struct {
	network string
	target  string
	wd      string
	writer  io.Writer
}

// New creates a new Sup. The io.Writer provided to this function will
// receive the output of the sup command.
func New(w io.Writer) *Sup {
	return &Sup{writer: w}
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
	cmd := exec.Command("sup", s.network, s.target)
	cmd.Dir = s.wd
	log.Println("Command:", strings.Join(cmd.Args, " "))
	log.Printf("Working Directory: %v", cmd.Dir)

	var outbuf bytes.Buffer
	var errbuf bytes.Buffer

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil {
		s.writer.Write(errbuf.Bytes())
		return err
	}

	_, err = s.writer.Write(outbuf.Bytes())
	return err
}
