package main

import "fmt"

func presentation() {
	text := `Version: 1.0.1 * Developped by Oleh Sobchuk tel: 0730240643`
	fmt.Println(text)
}

func printStats() {
	fmt.Printf("\nChecked: %d folder(s), %d file(s)\n", folderCount, fileCount)
	if config.RealClean {
		fmt.Printf("\nRemoved: %d file(s)\n\n", removedCount)
	} else {
		fmt.Printf("\nFound: %d file(s)\n\n", foundCount)
	}
}
