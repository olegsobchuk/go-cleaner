package main

import (
	"fmt"
	"log"
)

const version = "2.0.0"

var stats Stats

type Stats struct {
	FolderChecked uint16
	FileChecked   uint16
	FoundCount    uint16
	RemovedCount  uint16
}

func representation() string {
	text := "Version: %s * Developped by Oleh Sobchuk tel: 0730240643\n"
	return fmt.Sprintf(text, version)
}

func printStats(realClean bool) {
	log.Printf("\nChecked: %d folder(s), %d file(s)\n", stats.FileChecked, stats.FileChecked)
	if realClean {
		log.Printf("\nRemoved: %d file(s)\n\n", stats.RemovedCount)
	} else {
		log.Printf("\nFound: %d file(s)\n\n", stats.FoundCount)
	}
}
