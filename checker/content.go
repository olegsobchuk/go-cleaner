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
	fileString = strings.ReplaceAll(fileString, "\x00", "")
	for _, str := range suspiciousContents {
		lowerStr := strings.ToLower(str)

		if strings.Contains(fileString, lowerStr) {
			return true
		}
	}
	return false
}
