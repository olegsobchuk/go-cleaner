package checker

import (
	"log"
	"os"
	"strings"
)

const fileNameException = "cleaner_config.yml"

func IsContentContain(fileFullPath string, suspiciousContents []string) bool {
	if strings.HasSuffix(fileFullPath, fileNameException) {
		return false
	}

	b, err := os.ReadFile(fileFullPath)
	if err != nil {
		log.Println("--", err)
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
