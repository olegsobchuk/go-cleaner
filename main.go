package main

import (
	"log"
	"os"

	"go-cleaner/configurator"
)

const dumpFilePath = "./dump_file.txt"

var (
	dumpFile *os.File
	config   = &configurator.Config
)

func main() {
	configurator.Init()

	if !config.IsReady {
		log.Println("Not ready, check and adjust 'cleaner_config.yml' file to ger ready")
		return
	}

	createDumpFile()
	defer dumpFile.Close()

	_, err := dumpFile.WriteString(representation())
	if err != nil {
		log.Println(err)
	}
	log.Printf("Real mode %t \n", config.RealClean)

	log.Println("--->> Go...")
	checkAndRemove(config.StartPath)

	log.Println("<<--- Finished")
	printStats(config.RealClean)
	log.Println(representation())
}
