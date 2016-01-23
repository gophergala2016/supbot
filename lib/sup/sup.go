package sup

import (
	"fmt"
	"log"
	"io"
	"os/exec"
)

type sup struct {
	writer io.Writer
	network string
	target string
}

func (s *sup) Network(s string) {
	s.network = s
	return s
}

func (s *sup) Target(t string) {
	s.target = t
	return s
}

func (s *sup) Exec() {
	cmd := fmt.SprintF("sup %v %v", s.network, s.target)
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := s.writer.Write(out); err != nil {
		log.Fatal(err)
	}
	return s
}

// TODO: Pass in a command directly
// func (s *sup) Cmd() {
// err2 := sup.NewSup(io.Writer).Cmd("Some sup command")	
// }


func NewSup(w io.Writer) {
	return &sup{writer: w}	
}

