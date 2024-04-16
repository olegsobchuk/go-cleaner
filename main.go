package main

import (
	"log"
	"os"

	"go-cleaner/configurator"
)

const dumpFilePath = "./dump_file.txt"

var (
	dumpFile *os.File
)

func main() {
	config := configurator.NewConfigurator(configurator.BuiltInOs{})
	configuration, fileExists, err := config.GetConfiguration()
	printErrAndExit(err)

	if !fileExists {
		err := config.SaveConfigurationToFile(configuration, configurator.ConfigFileName)
		printErrAndExit(err)

		log.Printf("Configuration file %s with default options was created in current directory. "+
			"Check its content, update it and rerun the command.\n", configurator.ConfigFileName)
	}

	if !configuration.IsReady {
		log.Println("Configuraiton is not ready, adjust 'cleaner_config.yml' file to get it ready.")
		return
	}

	createDumpFile()
	defer dumpFile.Close()

	_, err = dumpFile.WriteString(representation())
	if err != nil {
		log.Println(err)
	}
	log.Printf("Real mode %t \n", configuration.RealClean)

	log.Println("--->> Go...")
	checkAndRemove(configuration.StartPath, configuration)

	log.Println("<<--- Finished")
	printStats(configuration.RealClean)
	log.Println(representation())
}

func printErrAndExit(err error) {
	if err != nil {
		log.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}
