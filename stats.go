package main

import "fmt"

const version = "1.0.2"

var stats Stats

type Stats struct {
	FolderChecked uint16
	FileChecked   uint16
	FoundCount    uint16
	RemovedCount  uint16
}

func presentation() {
	text := "Version: %s * Developped by Oleh Sobchuk tel: 0730240643\n"
	fmt.Printf(text, version)
}

func printStats() {
	fmt.Printf("\nChecked: %d folder(s), %d file(s)\n", stats.FileChecked, stats.FileChecked)
	if config.RealClean {
		fmt.Printf("\nRemoved: %d file(s)\n\n", stats.RemovedCount)
	} else {
		fmt.Printf("\nFound: %d file(s)\n\n", stats.FoundCount)
	}
}
