package checker

import (
	"log"
	"os"
)

func IsSizeOver(filePath string, limit int64) bool {
	return size(filePath) > limit
}

func IsZero(filePath string, toCheck bool) bool {
	if toCheck {
		return size(filePath) == 0
	}
	return false
}

func size(filePath string) int64 {
	info, err := os.Stat(filePath)
	if err != nil {
		log.Println("Size error -", filePath, err)
	}
	return info.Size()
}
