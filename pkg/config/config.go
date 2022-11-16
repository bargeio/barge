package config

import "os"

const configFileName = ".barge.yml"

type Config struct {
	assets       map[string]interface{}
	destinations map[string]interface{}
}

func GetConfigFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return cwd + "/" + configFileName, nil
}

func GetConfig() (Config, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	f, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	c := Config{}
	return c, nil
}
