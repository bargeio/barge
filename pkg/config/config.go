package config

import (
	"os"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

const configFileName = ".barge.yml"

type Config struct {
	Assets       map[string]interface{} `yaml:"assets"`
	Destinations map[string]interface{} `yaml:"destinations"`
}

type FilesAssetConfig struct {
	AssetType string   `mapstructure:"type"`
	Files     []string `mapstructure:"files"`
}

type DatabaseAssetConfig struct {
	AssetType        string `mapstructure:"type"`
	DatabaseHost     string `mapstructure:"host"`
	DatabasePort     int    `mapstructure:"port"`
	DatabaseUser     string `mapstructure:"user"`
	DatabasePassword string `mapstructure:"password"`
}

type S3DestinationConfig struct {
	DestinationType string `mapstructure:"type"`
	S3AccessKey     string `mapstructure:"access-key"`
	S3Secret        string `mapstructure:"secret"`
	S3Endpoint      string `mapstructure:"endpoint"`
	S3Bucket        string `mapstructure:"bucket"`
	PathTemplate    string `mapstructure:"path-template"`
}

func GetConfigFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return cwd + "/" + configFileName, nil
}

func GetConfig() (*Config, error) {
	// TODO Read config from multiple files and merge them into one
	// Read in the order of .barge.yml.dist and then .barge.yml
	// For now, this just reads the .barge.yml
	configPath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c, err := ParseConfig(f)
	return &c, err
}

func ParseConfig(file *os.File) (Config, error) {
	c := Config{}
	decoder := yaml.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		return c, err
	}

	for name, value := range c.Assets {
		assetConfig := value.(map[interface{}]interface{})
		switch assetConfig["type"] {
		case "database":
			var ad DatabaseAssetConfig
			mapstructure.Decode(assetConfig, &ad)
			c.Assets[name] = ad
		case "files":
			var af FilesAssetConfig
			mapstructure.Decode(assetConfig, &af)
			c.Assets[name] = af
		}
	}

	// Loop over destinations and parse them into structs
	for name, dest := range c.Destinations {
		if name == "default" {
			continue
		}
		destConfig := dest.(map[interface{}]interface{})
		switch destConfig["type"] {
		case "s3":
			var ds3 S3DestinationConfig
			mapstructure.Decode(destConfig, &ds3)
			c.Destinations[name] = ds3
		}
	}

	return c, err
}
