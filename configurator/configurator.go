package configurator

import (
	"errors"
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

const (
	ConfigFileName         = "cleaner_config.yml"
	defaultStartPath       = "."
	defaultSizeLimit int64 = 5_000_000
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

type Configurator struct {
	os OsWrapper
}

func NewConfigurator(os OsWrapper) Configurator {
	return Configurator{os}
}

func (c Configurator) GetConfiguration() (config *Configuration, fileExists bool, err error) {
	if c.isFilePresent(ConfigFileName) {
		config, err := c.readConfigurationFromFile(ConfigFileName)
		return config, true, err
	}

	return getDefaultConfiguration(), false, nil
}

func (c Configurator) SaveConfigurationToFile(config *Configuration, filePath string) error {
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to convert default configuration to YAML: %w", err)
	}

	permissions := os.FileMode(0666) // everybody may read and write the file
	return c.os.WriteFile(filePath, yamlBytes, permissions)
}

func getDefaultConfiguration() *Configuration {
	var (
		defaultBlackList        = ExtList{"lnk", "ini2", "bin", "tmp"}
		defaultWhiteListDocs    = ExtList{"doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf"}
		defaultWhiteListImgs    = ExtList{"png", "jpg", "jpeg", "raw"}
		defaultWhiteListExts    = slices.Concat(defaultWhiteListDocs, defaultWhiteListImgs)
		defaultBlackListFiles   = FileList{"~.ini2"}
		defaultContentBlacklist = []string{"powershell"}
	)

	return &Configuration{
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

func (c Configurator) isFilePresent(filePath string) bool {
	_, err := c.os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func (c Configurator) readConfigurationFromFile(filePath string) (*Configuration, error) {
	fileContent, err := c.os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	var config Configuration
	err = yaml.Unmarshal(fileContent, &config)

	return &config, err
}
