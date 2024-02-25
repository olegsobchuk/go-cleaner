package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const confFileName = "cleaner_config.yml"

var config Config

type Config struct {
	StartPath string   `yaml:"path"`
	RealClean bool     `yaml:"real"`
	IsReady   bool     `yaml:"ready"`
	Exts      []string `yaml:"extensions"`
	FileNames []string `yaml:"files"`
}

func SetConfiguration() {
	if isConfExist() {
		readConfFile()
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
	setDefaultConf()
	data, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("YAML Marshaling err:", err)
	}
	confFile.Write(data)
}

func setDefaultConf() {
	config = Config{
		StartPath: ".",
		RealClean: false,
		IsReady:   false,
		FileNames: []string{"~.ini"},
		Exts:      []string{"lnk"},
	}
}

func isConfExist() bool {
	_, err := os.Stat(confFileName)
	return !errors.Is(err, os.ErrNotExist)
}

func readConfFile() {
	data, err := os.ReadFile(confFileName)
	if err != nil {
		fmt.Println("Read conf file err:", err)
	}
	yaml.Unmarshal(data, &config)
}
