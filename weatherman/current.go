package weatherman

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type weatherCurrent struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Sys struct {
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
		Country string `json:"country"`
	} `json:"sys"`
	Name     string `json:"name"`
	Timezone int    `json:"timezone"`
}

func GetWeatherCurrent(args []string) {
	apiKey := getAPIKey()

	var l string

	if len(args) > 0 {
		l = args[0]
	} else {
		l = getLocation()
	}

	lat, lon := locationToLatLon(l, apiKey)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&APPID=%s&units=metric", strconv.FormatFloat(lat, 'f', 4, 64), strconv.FormatFloat(lon, 'f', 4, 64), apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected OpenWeather API response status: ", resp.Status)
	}

	var w weatherCurrent
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		log.Fatal(err)
	}

	var conditions string
	for _, condition := range w.Weather {
		conditions = conditions + condition.Description + ", "
	}

	fmt.Printf("\nLocation: %s (%s)\n", w.Name, w.Sys.Country)
	fmt.Printf("Current temperature: %g°C\n", math.Round(w.Main.Temp))
	fmt.Printf("Feels like: %g°C\n", math.Round(w.Main.FeelsLike))
	fmt.Printf("Weather conditions: %s\n", strings.TrimSuffix(conditions, ", "))
	fmt.Printf("Wind speed: %g m/s\n", math.Round(w.Wind.Speed))
	fmt.Printf("Pressure: %d mbar\n", w.Main.Pressure)
	fmt.Printf("Humidity: %d%%\n", w.Main.Humidity)
	fmt.Println("Sunrise:", time.Unix(int64(w.Sys.Sunrise), 0).UTC().Add(time.Second*time.Duration(w.Timezone)).Format("15:04"))
	fmt.Println("Sunset:", time.Unix(int64(w.Sys.Sunset), 0).UTC().Add(time.Second*time.Duration(w.Timezone)).Format("15:04"))
	fmt.Println()
}
