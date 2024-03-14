package configurator

import (
	"errors"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const confFileName = "cleaner_config.yml"

var (
	defaultBlackList     = []string{"lnk", "ini", "bin", "tmp"}
	defaultWhiteListDocs = []string{"doc", "docx", "xls", "xlsx"}
	defaultWhiteListImgs = []string{"png", "jpg"}
	Config               Configuration
)

type ExtList []string

type Exts struct {
	WhiteList ExtList `yaml:"whitelist"`
	BlackList ExtList `yaml:"blacklist"`
}

type Configuration struct {
	StartPath string   `yaml:"path"`
	RealClean bool     `yaml:"real"`
	IsReady   bool     `yaml:"ready"`
	Exts      Exts     `yaml:"extensions"`
	FileNames []string `yaml:"files"`
	Contents  []string `yaml:"content"`
}

func Init() {
	if isConfExist() {
		readConfFile()
	} else {
		addConfigFile()
	}
}

func addConfigFile() {
	confFile, err := os.Create(confFileName)
	if err != nil {
		log.Println("Create config error:", err)
	}
	defer confFile.Close()
	setDefaultConf()
	data, err := yaml.Marshal(Config)
	if err != nil {
		log.Println("YAML Marshaling err:", err)
	}
	confFile.Write(data)
}

func setDefaultConf() {
	Config = Configuration{
		StartPath: ".",
		RealClean: false,
		IsReady:   false,
		FileNames: []string{"~.ini"},
		Exts: Exts{
			WhiteList: append(defaultWhiteListDocs, defaultWhiteListImgs...),
			BlackList: defaultBlackList,
		},
		Contents: []string{"powershell"},
	}
}

func isConfExist() bool {
	_, err := os.Stat(confFileName)
	return !errors.Is(err, os.ErrNotExist)
}

func readConfFile() {
	data, err := os.ReadFile(confFileName)
	if err != nil {
		log.Println("Read conf file err:", err)
	}
	yaml.Unmarshal(data, &Config)
}
