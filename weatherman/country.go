package weatherman

import (
	"fmt"
	"log"
	"os"

	"github.com/biter777/countries"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

const (
	countryConfKey = "country"
	defaultCountry = "RS"
)

func ListISOCodes() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Country name", "Country code"})

	allCountries := countries.All()
	for _, c := range allCountries {
		table.Append([]string{fmt.Sprintf("%v", c), c.Alpha2()})
	}

	table.SetRowLine(true)
	table.Render()
}

func SetCountry(c string) {
	if !countryIsValid(c) {
		log.Fatal("Invalid country, please specify country by a valid two-letter ISO 3166 country code>, eg. \"FR\".\n./weatherman country list-iso-codes can be of help.")
	}

	if !configDirExists() {
		createConfigDir()
	}

	fmt.Printf("Setting country to %s\n", countries.ByName(c))
	viper.Set(countryConfKey, c)
	viper.WriteConfig()
	fmt.Println("Done!")
}

func ShowCountry() {
	c := getCountry()
	fmt.Printf("%s (%s)\n", countries.ByName(c), c)
}

func countryIsSet() bool {
	if !viper.IsSet(countryConfKey) {
		return false
	} else {
		return true
	}
}

func countryIsValid(c string) bool {
	if len(c) != 2 || countries.ByName(c) == countries.Unknown {
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
