package hal

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendCommand(t *testing.T) {

	testCommands := map[string]struct{ err bool }{
		`help`: {
			err: false,
		}, // should show help.
		/*
			`deploy`: {
				err: true,
			}, // should fail because we didn't specify a network.
			`prod deploy`: {
				err: true,
			}, // should fail because we don't have a repo.
			`set-repo git@github.com:xiam/go-playground.git`: {
				err: false,
			}, // should set the repo
			`prod deploy`: {
				err: false,
			}, // should execute production restart.
			`git@github.com:xiam/go-playground.git prod deploy`: {
				err: false,
			}, // should set the repo and exec prod deploy
			`git@github.com:xiam/go-playground.git prod deploy`: {
				err: false,
			}, // should set the repo and exec prod deploy
			`git@github.com:xiam/vanity.git/issue-1 prod deploy`: {
				err: false,
			}, // should set the repo and exec prod deploy
			`prod deploy`: {
				err: false,
			}, // should execute production deploy.
			`git@github.com:xiam/go-playground.git prod deploy`: {
				err: false,
			}, // should set the repo and exec prod deploy
			`prod deploy`: {
				err: false,
			}, // should execute production reploy.
		*/
	}

	buf := bytes.NewBuffer(nil)

	hal := NewHal(buf)

	for c, v := range testCommands {
		_, err := hal.Write([]byte(c))
		if err == nil {
			assert.False(t, v.err)
		} else {
			assert.True(t, v.err)
		}
	}

}
