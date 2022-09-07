package weatherman

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

type pollution struct {
	List []struct {
		Main struct {
			Aqi float64 `json:"aqi"`
		} `json:"main"`
		Components struct {
			Co   float64 `json:"co"`
			No   float64 `json:"no"`
			No2  float64 `json:"no2"`
			O3   float64 `json:"o3"`
			So2  float64 `json:"so2"`
			Pm25 float64 `json:"pm2_5"`
			Pm10 float64 `json:"pm10"`
			Nh3  float64 `json:"nh3"`
		} `json:"components"`
	} `json:"list"`
}

func GetPollution(args []string) {
	apiKey := getAPIKey()

	lat, lon := cityToLatLon(args[0], apiKey)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/air_pollution?lat=%s&lon=%s&APPID=%s", strconv.FormatFloat(lat, 'f', 4, 64), strconv.FormatFloat(lon, 'f', 4, 64), apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected OpenWeather API response status: ", resp.Status)
	}

	var p pollution

	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAir Quality Index (1/Good -> 5/Very poor): %g (%s)\n", math.Round(p.List[0].Main.Aqi), getAQIDesc(p.List[0].Main.Aqi))
	fmt.Printf("Concentration of Carbon monoxide (CO): %g\n", p.List[0].Components.Co)
	fmt.Printf("Concentration of Nitrogen monoxide (NO): %g\n", p.List[0].Components.No)
	fmt.Printf("Concentration of Nitrogen dioxide (NO2): %g\n", p.List[0].Components.No2)
	fmt.Printf("Concentration of Ozone (O3): %g\n", p.List[0].Components.O3)
	fmt.Printf("Concentration of Sulphur dioxide (SO2): %g\n", p.List[0].Components.So2)
	fmt.Printf("Concentration of Fine particles matter (PM2.5): %g\n", p.List[0].Components.Pm25)
	fmt.Printf("Concentration of Coarse particulate matter (PM10): %g\n", p.List[0].Components.Pm10)
	fmt.Printf("Concentration of Ammonia (NH3): %g\n", p.List[0].Components.Nh3)
}

func getAQIDesc(aqi float64) string {
	switch aqi {
	case 1.0:
		return "Good"
	case 2.0:
		return "Fair"
	case 3.0:
		return "Moderate"
	case 4.0:
		return "Poor"
	case 5.0:
		return "Very Poor"
	default:
		return "?"
	}
}
