package weatherman

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type weatherForecast struct {
	Cnt  int `json:"cnt"` // A number of timestamps returned in the API response
	List []struct {
		Dt   int `json:"dt"` // Time of data forecasted, unix, UTC
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
		Pop   float64 `json:"pop"`
		DtTxt string  `json:"dt_txt"` // Time of data forecasted, ISO, UTC
	} `json:"list"`
	City struct {
		Name     string `json:"name"`
		Country  string `json:"country"`
		Timezone int    `json:"timezone"`
	} `json:"city"`
}

func GetWeatherForecast(args []string) {
	apiKey := getAPIKey()

	var c string

	if len(args) > 0 {
		c = args[0]
	} else {
		c = getCity()
	}

	lat, lon := cityToLatLon(c, apiKey)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?lat=%s&lon=%s&APPID=%s&units=metric", strconv.FormatFloat(lat, 'f', 4, 64), strconv.FormatFloat(lon, 'f', 4, 64), apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Unexpected response status: ", resp.Status)
	}

	var f weatherForecast

	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Local time", "Temperature", "Conditions", "Wind speed", "Chance of precipitation"})

	fmt.Printf("\nLocation: %s (%s)\n", f.City.Name, f.City.Country)
	for _, v := range f.List {
		conditions := ""
		for _, condition := range v.Weather {
			conditions = conditions + condition.Description + ", "
		}

		fTime := time.Unix(int64(v.Dt), 0).UTC().Add(time.Second * time.Duration(f.City.Timezone))
		temp := strconv.FormatFloat(math.Round(v.Main.Temp), 'f', 0, 64) + "°C"
		wind := strconv.FormatFloat(math.Round(v.Wind.Speed), 'f', 0, 64) + " m/s"
		pop := strconv.FormatFloat(math.Round(v.Pop*100), 'f', 0, 64) + "%"

		table.Append([]string{fTime.Format("02.01.2006 15:04"), temp, strings.TrimSuffix(conditions, ", "), wind, pop})
	}

	table.Render()
}
