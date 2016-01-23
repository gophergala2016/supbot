package hal

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type testCommand struct {
	command string
	err     bool
}

func TestSendCommand(t *testing.T) {

	testCommands := []testCommand{
		{
			command: "help",
			err:     false,
		},
		{
			command: "prod deploy",
			err:     true,
		},
		{
			command: "local ping",
			err:     true,
		},
		{
			command: `set-repo https://github.com/gophergala2016/supbot.git`,
			err:     false,
		},
		{
			command: "local ping",
			err:     false,
		},
	}

	buf := bytes.NewBuffer(nil)

	hal := NewHal(buf)

	for _, c := range testCommands {
		_, err := hal.Write([]byte(c.command))
		log.Printf("command: %q", c.command)
		if err == nil {
			assert.False(t, c.err)
		} else {
			assert.True(t, c.err)
		}
		log.Printf("out: %s, %q", buf.String(), err)
		buf.Reset()
	}

}
