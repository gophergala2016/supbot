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
	_ = s.Stackup // NOTE: godeps..
)

type Sup struct {
	network string
	target  string
	wd      string
	writer  io.Writer
}

func (s *Sup) Network(n string) *Sup {
	s.network = n
	return s
}

func (s *Sup) Setwd(wdir string) *Sup {
	s.wd = wdir
	return s
}

func (s *Sup) Target(t string) *Sup {
	s.target = t
	return s
}

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

// TODO: Pass in a command directly
// func (s *sup) Cmd() {
// err2 := sup.NewSup(io.Writer).Cmd("Some sup command")
// }

func NewSup(w io.Writer) *Sup {
	return &Sup{writer: w}
}
