package sup

import (
	"bytes"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	var b bytes.Buffer
	if err := New(&b).SetNetwork("local").SetTarget("ping").SetWd("../../cmd/supbot").Exec(); err != nil {
		log.Fatalln("Testing error:", err)
	}

	log.Printf("%s", string(b.Bytes()))
}
