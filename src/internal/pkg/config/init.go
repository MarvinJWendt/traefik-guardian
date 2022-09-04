package config

import (
	"errors"
	"io/fs"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func Initzialize() error {
	err := createConfigIfNotExists()
	if err != nil {
		return err
	}

	err = readConfig()
	if err != nil {
		return err
	}

	return nil
}

func Reload() error {
	err := readConfig()
	if err != nil {
		return err
	}

	return nil
}

func createConfigIfNotExists() error {
	// Check if config path exists
	_, err := os.Stat(CONFIG_DIR)
	if errors.Is(err, fs.ErrNotExist) {
		// Create config path
		err := os.MkdirAll(CONFIG_DIR, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Check if config file exists
	_, err = os.Stat(CONFIG_DIR + "/" + CONFIG_FILE)
	if errors.Is(err, fs.ErrNotExist) {
		logrus.Debug("config file does not exist, creating new one")
		// Parse InitConfig to YAML
		yamlData, err := yaml.Marshal(InitConfig)
		if err != nil {
			return err
		}

		// Write YAML to config file
		err = os.WriteFile(CONFIG_DIR+"/"+CONFIG_FILE, yamlData, os.ModePerm)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	return nil
}

func readConfig() error {
	// Read config file
	logrus.Debug("reading config file")
	yamlData, err := os.ReadFile(CONFIG_DIR + "/" + CONFIG_FILE)
	if err != nil {
		return err
	}

	// Parse YAML to Config
	err = yaml.Unmarshal(yamlData, &Cfg)
	if err != nil {
		return err
	}

	return nil
}
