package sup

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

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
	ErrInvalidTarget = errors.New("invalid target")
)

func (s *Sup) Network(n string) *Sup {
	s.network = n
	return s
}

func (s *Sup) Target(t string) *Sup {
	s.target = t
	return s
}

func (s *Sup) Exec() error {
	//cmd := exec.Command("sup", s.network, s.target)
	//cmd.Dir = s.wd
	//log.Println("Command:", strings.Join(cmd.Args, " "))
	//log.Printf("Working Directory: %v", cmd.Dir)

	//var outbuf bytes.Buffer
	//var errbuf bytes.Buffer

	//cmd.Stdout = &outbuf
	//cmd.Stderr = &errbuf

	//err := cmd.Run()
	//if err != nil {
	//s.writer.Write(errbuf.Bytes())
	//return err
	//}
	network, _ := s.config.Networks[s.network]

	command, isCommand := s.config.Commands[s.target]
	if !isCommand {
		return ErrInvalidTarget
	}
	cmds := []*stackup.Command{&command}

	// Do some piping magic here
	old := os.Stdout
	oldErr := os.Stderr
	read, write, _ := os.Pipe()
	errRead, errWrite, _ := os.Pipe()

	os.Stdout = write
	os.Stderr = errWrite

	write.Close()
	errWrite.Close()

	var out string
	var errOut string
	if err := s.Run(&network, cmds...); err == nil {

		//out := fmt.Sprintf("<@%s>: \n", msg.User)
		scanner := bufio.NewScanner(read)
		for scanner.Scan() {
			out = fmt.Sprintf("%s %s\n", out, scanner.Text())
		}

		scanner = bufio.NewScanner(errRead)
		for scanner.Scan() {
			errOut = fmt.Sprintf("%s %s\n", errOut, scanner.Text())
		}

	}
	os.Stdout = old    // reset stdout
	os.Stderr = oldErr // reset stderr

	if err != nil {
		log.Println("got this RUN error %v\n\n", err)
	}

	log.Println("got this stdout: %v", out)
	log.Println("got this stderr: %v", errOut)

	//_, err = s.writer.Write(outbuf.Bytes())
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
