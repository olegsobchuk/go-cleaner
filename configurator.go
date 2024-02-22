package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const confFileName = "cleaner_config.yml"

type Config struct {
	StartPath string `yaml:"path"`
	RealClean bool   `yaml:"real"`
}

func SetConfiguration() {
	if isConfExist() {
		// read conf
	} else {
		addConfigFile()
	}
}

func addConfigFile() {
	confFile, err := os.Create(confFileName)
	if err != nil {
		fmt.Println("Create config error:", err)
	}
	defer confFile.Close()
	conf := defaultConf()
	data, err := yaml.Marshal(conf)
	if err != nil {
		fmt.Println("YAML Marshaling err:", err)
	}
	confFile.Write(data)
}

func defaultConf() Config {
	return Config{
		StartPath: ".",
		RealClean: false,
	}
}

func isConfExist() bool {
	_, err := os.Stat(confFileName)
	return !errors.Is(err, os.ErrNotExist)
}
