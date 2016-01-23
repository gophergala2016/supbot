package git

import (
	"testing"

	"github.com/gophergala2016/supbot/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestCloneCheckout(t *testing.T) {
	assert := assert.New(t)

	repo, err := Clone("https://github.com/gophergala2016/supbot.git")
	assert.NoError(err)
	assert.NotNil(repo)

	assert.NotEmpty(repo.Dir())

	err = repo.Checkout("master")
	assert.NoError(err)
}
