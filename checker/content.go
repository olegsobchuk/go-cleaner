package checker

import (
	"log"
	"os"
	"strings"
)

const fileNameException = "cleaner_config.yml"

// needed this implementattion for tests
var readFileToString = func(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func IsContentContain(fileFullPath string, suspiciousContents []string) bool {
	if strings.HasSuffix(fileFullPath, fileNameException) {
		return false
	}

	fileString, err := readFileToString(fileFullPath)
	if err != nil {
		log.Println("--", err)
		return false
	}

	fileString = strings.ToLower(fileString)
	fileString = strings.ReplaceAll(fileString, "\x00", "")
	log.Println(fileString)
	for _, str := range suspiciousContents {
		lowerStr := strings.ToLower(str)

		if strings.Contains(fileString, lowerStr) {
			return true
		}
	}
	return false
}
