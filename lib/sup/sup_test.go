package sup

import (
	"bytes"
	"testing"
)

func TestNewSup(t *testing.T) {
	var b bytes.Buffer
	NewSup(&b).Network("local").Target("ping").Exec()
}