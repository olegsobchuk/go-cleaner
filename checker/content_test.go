package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsContentContain(t *testing.T) {
	t.Run("ignores file with config file name", func(t *testing.T) {
		path := "some/path/" + fileNameException
		res := IsContentContain(path, []string{})
		assert.False(t, res, "not ignored config file")
	})

}
