package checker

import (
	"io/fs"
	"os"
	"strings"
)

var suspiciousContents = []string{"t"}

func IsContentContain(file fs.DirEntry) bool {
	b, err := os.ReadFile(file.Name())
	if err != nil {
		return false
	}

	fileString := strings.ToLower(string(b))
	for _, str := range suspiciousContents {
		if strings.Contains(fileString, str) {
			return true
		}
	}
	return false
}
