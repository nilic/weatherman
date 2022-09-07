package weatherman

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	configDir  = ".weatherman"
	configFile = "config"
	configType = "yaml"
)

func getConfigFilePath() string {
	return filepath.Join(getUserHomeDir(), configDir, configFile)
}

func getConfigFileDir() string {
	return filepath.Join(getUserHomeDir(), configDir)
}

func configDirExists() bool {
	_, err := os.Stat(getConfigFileDir())

	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return true
}

func getUserHomeDir() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return homeDir
}

func InitConfig() {
	viper.SetConfigFile(getConfigFilePath())
	viper.SetConfigType(configType)
}
