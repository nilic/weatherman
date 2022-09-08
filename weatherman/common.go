package weatherman

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func locationToLatLon(l string, apiKey string) (float64, float64) {
	type geo []struct {
		Name    string  `json:"name"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
		Country string  `json:"country"`
	}

	location := makeLocation(l)

	urlGeo := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&APPID=%s&limit=1", location, apiKey)

	resp, err := http.Get(urlGeo)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected OpenWeather API response status: ", resp.Status)
	}

	var g geo

	if err := json.NewDecoder(resp.Body).Decode(&g); err != nil {
		log.Fatal(err)
	} else if len(g) == 0 {
		fmt.Printf("Given location %s not found in the OpenWeather database\n", location)
		log.Fatal("Exiting..")
	}

	return g[0].Lat, g[0].Lon
}

func makeLocation(c string) string {
	var city, country string

	s := strings.Split(c, ",")
	if len(s) == 1 {
		country = getCountry()
	} else if len(s) == 2 && s[1] != "" {
		country = strings.TrimSpace(s[1])
	} else {
		log.Fatal("Invalid location argument, please specify location as \"<city name>,<ISO 3166 country code>\" or by setting default city/country.")
	}
	city = strings.ReplaceAll(strings.TrimSpace(s[0]), " ", "+")

	return city + "," + country
}
