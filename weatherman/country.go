package weatherman

import (
	"fmt"
	"log"

	"github.com/biter777/countries"
	"github.com/spf13/viper"
)

const (
	countryConfKey = "country"
	defaultCountry = "RS"
)

func countryIsSet() bool {
	if !viper.IsSet(countryConfKey) {
		return false
	} else {
		return true
	}
}

func getCountry() string {
	if countryIsSet() {
		return viper.GetString(countryConfKey)
	} else {
		return defaultCountry
	}
}

func SetCountry(c string) {
	if !countryIsValid(c) {
		log.Fatal("Invalid country, please specify country by a valid two-letter ISO 3166 country code>, eg. \"FR\".")
	}

	if !configDirExists() {
		createConfigDir()
	}

	fmt.Printf("Setting country to %s\n", c)
	viper.Set(countryConfKey, c)
	viper.WriteConfig()
	fmt.Println("Done!")
}

func ShowCountry() {
	fmt.Print(getCountry())
}

func countryIsValid(c string) bool {
	if len(c) != 2 || countries.ByName(c) == countries.Unknown {
		return false
	} else {
		return true
	}
}
