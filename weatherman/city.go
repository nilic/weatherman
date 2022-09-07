package weatherman

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const cityConfKey = "city"

func cityIsSet() bool {
	if !viper.IsSet(cityConfKey) {
		return false
	} else {
		return true
	}
}

func getCity() string {
	if cityIsSet() {
		return viper.GetString(cityConfKey)
	} else {
		return defaultCity
	}
}

func SetCity(c string) {
	if !cityValid(c) {
		log.Fatal("Invalid city format, please specify city as \"<city name>,<ISO 3166 country code>\", eg \"Paris,FR\".")
	}

	if !configDirExists() {
		createConfigDir()
	}

	fmt.Printf("Setting city to %s\n", c)
	viper.Set(cityConfKey, c)
	viper.WriteConfig()
	fmt.Println("Done!")
}

func ShowCity() {
	fmt.Print(getCity())
}

func cityValid(c string) bool {
	if s := strings.Split(c, ","); len(s) != 2 || len(s[0]) == 0 || len(s[1]) != 2 {
		return false
	} else {
		return true
	}
}
