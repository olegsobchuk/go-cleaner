package configurator

import (
	"errors"
	"log"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

const confFileName = "cleaner_config.yml"

var (
	Config Configuration
)

type ExtList []string
type FileList []string

type Exts struct {
	WhiteList ExtList `yaml:"whitelist"`
	BlackList ExtList `yaml:"blacklist"`
}

type Files struct {
	WhiteList FileList `yaml:"whitelist"`
	BlackList FileList `yaml:"blacklist"`
}

type SizeConfig struct {
	Threshold int64 `yaml:"ignore_more"`
	CatchZero bool  `yaml:"catch_zero"`
}

type Configuration struct {
	StartPath  string     `yaml:"path"`
	RealClean  bool       `yaml:"real"`
	IsReady    bool       `yaml:"ready"`
	SizeConfig SizeConfig `yaml:"size"`
	Exts       Exts       `yaml:"extensions"`
	Files      Files      `yaml:"files"`
	Contents   []string   `yaml:"content"`
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
	var defaultStartPath = "."
	var defaultBlackList = ExtList{"lnk", "ini2", "bin", "tmp"}
	var defaultWhiteListDocs = ExtList{"doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf"}
	var defaultWhiteListImgs = ExtList{"png", "jpg", "jpeg", "raw"}
	var defaultWhiteListExts = slices.Concat(defaultWhiteListDocs, defaultWhiteListImgs)
	var defaultBlackListFiles = FileList{"~.ini2"}
	var defaultSizeLimit int64 = 5_000_000
	var defaultContentBlacklist = []string{"powershell"}

	Config = Configuration{
		StartPath: defaultStartPath,
		RealClean: false,
		IsReady:   false,
		SizeConfig: SizeConfig{
			Threshold: defaultSizeLimit,
			CatchZero: true,
		},
		Files: Files{
			BlackList: defaultBlackListFiles,
		},
		Exts: Exts{
			WhiteList: defaultWhiteListExts,
			BlackList: defaultBlackList,
		},
		Contents: defaultContentBlacklist,
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
