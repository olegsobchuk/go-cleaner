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

	t.Run("read file", func(t *testing.T) {
		filePath := "path.txt"
		suspiciousList := []string{"suspicious"}

		t.Run("file contains suspicious text", func(t *testing.T) {
			fileText := "some suspicious string"
			readFileToString = func(_ string) (string, error) {
				return fileText, nil
			}
			res := IsContentContain(filePath, suspiciousList)

			assert.True(t, res, "can't find text")
		})

		t.Run("file does not contain suspicious text", func(t *testing.T) {
			fileText := "some normal string"
			readFileToString = func(_ string) (string, error) {
				return fileText, nil
			}
			res := IsContentContain(filePath, suspiciousList)

			assert.False(t, res, "can find text")
		})
	})
}
