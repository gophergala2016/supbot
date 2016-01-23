package main

import (
	"io/ioutil"
	"os"

	"github.com/gophergala2016/supbot/Godeps/_workspace/src/github.com/goware/prefixer"
)

func main() {
	// Prefixer accepts anything that implements io.Reader interface
	prefixReader := prefixer.New(os.Stdin, "> ")

	// Read all prefixed lines from STDIN into a buffer
	buffer, _ := ioutil.ReadAll(prefixReader)

	// Write buffer to STDOUT
	os.Stdout.Write(buffer)
}
