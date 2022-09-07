package weatherman

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

const apiKeyConfKey = "apiKey"

func APIKeyIsSet() bool {
	_ = viper.ReadInConfig()

	if !viper.IsSet(apiKeyConfKey) || viper.GetString(apiKeyConfKey) == "" {
		return false
	} else {
		return true
	}
}

func getAPIKey() string {
	if !APIKeyIsSet() {
		log.Fatal("OpenWeather API key not found, please run ./weatherman key save to add one")
	}

	return viper.GetString(apiKeyConfKey)
}

func SaveAPIKey(args []string) {
	if !configDirExists() {
		d := getConfigFileDir()
		fmt.Printf("Creating configuration directory %s\n", d)
		err := os.Mkdir(d, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	viper.Set(apiKeyConfKey, args[0])

	f := getConfigFilePath()
	fmt.Printf("Writing API key to configuration file %s\n", f)
	viper.WriteConfig()
	fmt.Println("Done!")
}

func ShowAPIKey() {
	fmt.Print(getAPIKey())
}

func DeleteAPIKey() {
	f := getConfigFilePath()
	if APIKeyIsSet() {
		fmt.Printf("Removing API key from the configuration file %s\n", f)
		viper.Set(apiKeyConfKey, "")
		viper.WriteConfig()
		fmt.Println("Done!")
	} else {
		fmt.Printf("API key not found in the configuration file %s\nNothing to delete.", f)
	}
}
