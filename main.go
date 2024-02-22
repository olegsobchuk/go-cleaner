package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

const dumpFilePath = "./dump_file.txt"

var (
	dumpFile     *os.File
	folderCount  uint16
	fileCount    uint16
	removedCount uint16
	foundCount   uint16
)

func main() {
	SetConfiguration()

	if !config.IsReady {
		fmt.Println("Not ready, check config file")
		return
	}

	if !config.RealClean {
		var err error
		dumpFile, err = os.Create(dumpFilePath)
		if err != nil {
			log.Fatalln(err)
		}
		defer dumpFile.Close()
		defer dumpFile.Sync()
	}

	fmt.Printf("Real mode %t \n", config.RealClean)

	fmt.Println("--->> Go...")
	checkAndRemove(config.StartPath)
	fmt.Println("<<--- Finished")
	printStats()
	presentation()
}

func supportedExtensions() []string {
	return []string{"lnk"}
}

func checkAndRemove(dirPath string) {
	folderCount++
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	printPath := true

	for _, entry := range entries {
		newPath := path.Join(dirPath, entry.Name())
		if entry.IsDir() {
			checkAndRemove(newPath)
			continue
		}

		fileCount++
		ext := extension(entry.Name())

		if slices.Contains(supportedExtensions(), ext) {
			if config.RealClean {
				err := os.Remove(newPath)
				if err != nil {
					fmt.Println(err)
					continue
				}
				removedCount++
			} else {
				foundCount++
				_, err = dumpFile.WriteString(fmt.Sprintf("%s\n", newPath))
				if err != nil {
					log.Println("Write to file error:", err)
				}
			}
			if printPath {
				fmt.Println(dirPath)
				printPath = false
			}
			fmt.Printf("   XXX: %s \n", entry.Name())
		}
	}
}

func extension(fileName string) string {
	ext := filepath.Ext(fileName)
	ext = strings.TrimPrefix(ext, ".")
	return strings.ToLower(ext)
}
