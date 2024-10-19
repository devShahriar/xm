package config

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

var CommandLineConfig *CmdConfig

func GetCmdConfig() *CmdConfig {
	if CommandLineConfig == nil {
		CommandLineConfig = &CmdConfig{}
	}
	return CommandLineConfig
}

type CmdConfig struct {
	ConfigPath string
}

func (i *CmdConfig) ReadConfig() {
	yamlFile, err := os.ReadFile(i.ConfigPath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
	AppConfig = &config
	spew.Dump(AppConfig)
}
