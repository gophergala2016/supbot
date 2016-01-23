package sup

import (
	"io"
	"log"
	"os/exec"
	"strings"
)

var supCommand string

func init () {
	out, err := exec.Command("/usr/bin/which", "sup").Output()
	if err != nil {
		log.Fatalln("sup init issue:", err)
		return
	}
	supCommand = strings.TrimSpace(string(out))
}

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
	// log.Printf("Command: %v %v %v\n", supCommand, s.network, s.target)
	cmd := exec.Command(supCommand, s.network, s.target)
	cmd.Dir = s.wd
	// log.Printf("Working Directory: %v", cmd.Dir)
	
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	if _, err := s.writer.Write(out); err != nil {
		return err
	}
	return nil
}

// TODO: Pass in a command directly
// func (s *sup) Cmd() {
// err2 := sup.NewSup(io.Writer).Cmd("Some sup command")
// }

func NewSup(w io.Writer) *Sup {
	return &Sup{writer: w}
}
