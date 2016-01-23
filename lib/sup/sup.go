package sup

import (
	"fmt"
	"io"
	"os/exec"
)

type Sup struct {
	writer io.Writer
	network string
	target string
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
	cmd := fmt.Sprintf("sup %v %v", s.network, s.target)
	out, err := exec.Command(cmd).Output()
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

