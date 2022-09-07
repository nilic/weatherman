package weatherman

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	configDir   = ".weatherman"
	configFile  = "config"
	configType  = "yaml"
	defaultCity = "Belgrade,RS"
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
	_ = viper.ReadInConfig()
}

func createConfigDir() {
	d := getConfigFileDir()
	fmt.Printf("Creating configuration directory %s\n", d)
	err := os.Mkdir(d, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
