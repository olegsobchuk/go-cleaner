package main

import (
	"fmt"
	"log"
	"os"
)

const version = "2.0.4"

var stats Stats

type Stats struct {
	FolderChecked uint16
	FileChecked   uint16
	FoundCount    uint16
	RemovedCount  uint16
}

func createDumpFile() {
	var err error
	dumpFile, err = os.Create(dumpFilePath)
	if err != nil {
		log.Fatalln(err)
	}
}

func representation() string {
	text := "Version: %s * Developped by Oleh Sobchuk tel: 0730240643\n\n"
	return fmt.Sprintf(text, version)
}

func printStats(realClean bool) {
	totalStatInfo := fmt.Sprintf("\nChecked: %d folder(s), %d file(s)\n", stats.FolderChecked, stats.FileChecked)
	var processedFilesStat string

	dumpFile.WriteString(totalStatInfo)
	log.Println(totalStatInfo)
	if realClean {
		processedFilesStat = fmt.Sprintf("\nRemoved: %d file(s)\n\n", stats.RemovedCount)
	} else {
		processedFilesStat = fmt.Sprintf("\nFound: %d file(s)\n\n", stats.FoundCount)
	}
	dumpFile.WriteString(processedFilesStat)
	log.Println(processedFilesStat)
}
