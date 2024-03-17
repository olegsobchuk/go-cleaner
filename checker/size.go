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
	log.Println("l", limit, "f", filePath, info.Size())
	return info.Size() > limit
}
