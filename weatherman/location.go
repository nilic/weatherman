package weatherman

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const (
	locationConfKey = "location"
	defaultLocation = "Belgrade,RS"
)

func locationIsSet() bool {
	if !viper.IsSet(locationConfKey) {
		return false
	} else {
		return true
	}
}

func getLocation() string {
	if locationIsSet() {
		return viper.GetString(locationConfKey)
	} else {
		return defaultLocation
	}
}

func SetLocation(l string) {
	if !locationIsValid(l) {
		log.Fatal("Invalid location format, please specify location as \"<city name>,<ISO 3166 country code>\", eg \"Paris,FR\".")
	}

	if !configDirExists() {
		createConfigDir()
	}

	fmt.Printf("Setting location to %s\n", l)
	viper.Set(locationConfKey, l)
	viper.WriteConfig()
	fmt.Println("Done!")
}

func ShowLocation() {
	fmt.Print(getLocation())
}

func locationIsValid(l string) bool {
	if s := strings.Split(l, ","); len(s) != 2 || len(s[0]) == 0 || len(s[1]) != 2 {
		return false
	} else {
		return true
	}
}
