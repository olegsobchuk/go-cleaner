package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsContentContain(t *testing.T) {
	path := "some/path" + fileNameException
	res := IsContentContain(path, []string{})
	assert.Equal(t, res, false, "not ignored config file")
}
