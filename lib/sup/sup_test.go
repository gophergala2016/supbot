package sup

import (
	"bytes"
	"log"
	"testing"
)

func TestNewSup(t *testing.T) {
	var b bytes.Buffer
	if err := NewSup(&b).Network("local").Target("ping").Setwd("../..").Exec(); err != nil {
		log.Fatalln("Testing error:", err)
	}

	log.Printf("%s", string(b.Bytes()))
}
