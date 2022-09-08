package weatherman

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const apiKeyConfKey = "apiKey"

func apiKeyIsSet() bool {
	if !viper.IsSet(apiKeyConfKey) || viper.GetString(apiKeyConfKey) == "" {
		return false
	} else {
		return true
	}
}

func SaveAPIKey(apiKey string) {
	if !configDirExists() {
		createConfigDir()
	}

	viper.Set(apiKeyConfKey, apiKey)

	fmt.Println("Writing API key to configuration file..")
	viper.WriteConfig()
	fmt.Println("Done!")
}

func ShowAPIKey() {
	fmt.Print(getAPIKey())
}

func DeleteAPIKey() {
	if apiKeyIsSet() {
		fmt.Println("Removing API key from the configuration file..")
		viper.Set(apiKeyConfKey, "")
		viper.WriteConfig()
		fmt.Println("Done!")
	} else {
		fmt.Println("API key not found in the configuration file. Nothing to delete.")
	}
}

func getAPIKey() string {
	if !apiKeyIsSet() {
		log.Fatal("OpenWeather API key not found, please run ./weatherman key save to add one")
	}

	return viper.GetString(apiKeyConfKey)
}
