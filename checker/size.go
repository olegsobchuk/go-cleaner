package checker

import (
	"log"
	"os"
)

func IsSizeOver(filePath string, limit int64) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		log.Println("Size error -", filePath, err)
	}
	return info.Size() > limit
}
