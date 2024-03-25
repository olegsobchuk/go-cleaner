package main

import (
	"fmt"
	"go-cleaner/checker"
	"log"
	"os"
	"path"
	"sort"
)

func checkAndRemove(dirPath string) error {
	printPath := true
	stats.FolderChecked++
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Println(err)
		return err
	}

	// check dir-s in the end
	sort.Slice(entries, func(i, j int) bool { return entries[j].IsDir() })
	for _, entry := range entries {
		fullFilePath := path.Join(dirPath, entry.Name())
		if entry.IsDir() {
			checkAndRemove(fullFilePath)
			// we don't need to check error here
			// if internal folder is absent it means that it's been deleted or renamed
			// so we just skip this check
			continue
		}

		stats.FileChecked++

		// Ignore file by WhiteList extension or Size limit
		beIgored := checker.IsExtMatch(entry, config.Exts.WhiteList) ||
			checker.IsSizeOver(fullFilePath, config.SizeConfig.Threshold)

		if beIgored {
			continue
		}

		isMatch := checker.IsExtMatch(entry, config.Exts.BlackList) ||
			checker.IsNameMatch(entry.Name(), config.Files.BlackList) ||
			checker.IsZero(fullFilePath, config.SizeConfig.CatchZero) ||
			checker.IsContentContain(fullFilePath, config.Contents)

		if isMatch {
			catchFile(fullFilePath)

			if printPath {
				dumpFile.WriteString(fmt.Sprintln(dirPath))
				log.Println(dirPath)
				printPath = false
			}
			dumpFile.WriteString(fmt.Sprintf("   X: %s \n", entry.Name()))
			log.Printf("   X: %s \n", entry.Name())
		}
	}
	dumpFile.Sync()
	return nil
}

func catchFile(filePath string) {
	if config.RealClean {
		err := os.Remove(filePath)
		if err != nil {
			log.Println(err)
			return
		}
		stats.RemovedCount++
	} else {
		stats.FoundCount++
	}
}
