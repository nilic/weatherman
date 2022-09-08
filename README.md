## ☀️ weatherman - OpenWeather CLI ☀️
weatherman is a CLI tool that can retrieve [OpenWeather](https://openweathermap.org/) data such as current weather, forecast and pollution information for a given location and display it in a human-readable format.

## Getting started

1. Download weatherman from the [Releases](https://github.com/nilic/weatherman/releases) page
2. Register on the [OpenWeather](https://openweathermap.org/) site and get a free API key
3. Save the API key to weatherman configuration by running `weatherman key save <your_API_key>`

Now you can start using weatherman! 

## Usage

```
Usage:
  weatherman [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  country     Commands for managing default country configuration
  current     Retrieve current weather conditions for a given location
  forecast    Retrieve 5-day weather forecast for a given location
  help        Help about any command
  key         Commands for managing OpenWeather API key
  location    Commands for managing default location configuration
  pollution   Retrieve pollution information for a given location
  version     Returns the current version of weatherman

Flags:
  -h, --help   help for weatherman
```

### Get current weather

```
$ weatherman current Paris,FR

Location: Palais-Royal (FR)
Current temperature: 17°C
Feels like: 17°C
Weather conditions: scattered clouds
Wind speed: 3 m/s
Pressure: 1009 mbar
Humidity: 84%
Sunrise: 07:17
Sunset: 20:20
```

### Get 5-day weather forecast

```
$ weatherman forecast "New York,US"

Location: New York (US)
+------------------+-------------+------------------+------------+-------------------------+
|    LOCAL TIME    | TEMPERATURE |    CONDITIONS    | WIND SPEED | CHANCE OF PRECIPITATION |
+------------------+-------------+------------------+------------+-------------------------+
| 08.09.2022 17:00 | 25°C        | scattered clouds | 4 m/s      | 0%                      |
| 08.09.2022 20:00 | 25°C        | few clouds       | 2 m/s      | 0%                      |
| 08.09.2022 23:00 | 24°C        | clear sky        | 1 m/s      | 0%                      |
| 09.09.2022 02:00 | 22°C        | clear sky        | 1 m/s      | 0%                      |
| 09.09.2022 05:00 | 20°C        | clear sky        | 3 m/s      | 0%                      |
| 09.09.2022 08:00 | 20°C        | clear sky        | 4 m/s      | 0%                      |
| 09.09.2022 11:00 | 24°C        | clear sky        | 3 m/s      | 0%                      |
| 09.09.2022 14:00 | 26°C        | clear sky        | 3 m/s      | 0%                      |
| 09.09.2022 17:00 | 26°C        | clear sky        | 3 m/s      | 0%                      |
| 09.09.2022 20:00 | 23°C        | clear sky        | 4 m/s      | 0%                      |
| 09.09.2022 23:00 | 22°C        | clear sky        | 3 m/s      | 0%                      |
| 10.09.2022 02:00 | 21°C        | clear sky        | 2 m/s      | 0%                      |
| 10.09.2022 05:00 | 20°C        | clear sky        | 2 m/s      | 0%                      |
| 10.09.2022 08:00 | 21°C        | clear sky        | 1 m/s      | 0%                      |
| 10.09.2022 11:00 | 26°C        | clear sky        | 2 m/s      | 0%                      |
| 10.09.2022 14:00 | 27°C        | few clouds       | 4 m/s      | 0%                      |
| 10.09.2022 17:00 | 27°C        | overcast clouds  | 4 m/s      | 0%                      |
| 10.09.2022 20:00 | 25°C        | overcast clouds  | 4 m/s      | 0%                      |
| 10.09.2022 23:00 | 23°C        | overcast clouds  | 4 m/s      | 0%                      |
| 11.09.2022 02:00 | 22°C        | overcast clouds  | 3 m/s      | 0%                      |
| 11.09.2022 05:00 | 22°C        | overcast clouds  | 3 m/s      | 0%                      |
| 11.09.2022 08:00 | 22°C        | overcast clouds  | 2 m/s      | 0%                      |
| 11.09.2022 11:00 | 23°C        | overcast clouds  | 3 m/s      | 0%                      |
| 11.09.2022 14:00 | 25°C        | overcast clouds  | 2 m/s      | 0%                      |
| 11.09.2022 17:00 | 24°C        | overcast clouds  | 4 m/s      | 3%                      |
| 11.09.2022 20:00 | 24°C        | light rain       | 2 m/s      | 24%                     |
| 11.09.2022 23:00 | 23°C        | overcast clouds  | 3 m/s      | 26%                     |
| 12.09.2022 02:00 | 22°C        | light rain       | 4 m/s      | 63%                     |
| 12.09.2022 05:00 | 21°C        | moderate rain    | 4 m/s      | 100%                    |
| 12.09.2022 08:00 | 22°C        | light rain       | 4 m/s      | 100%                    |
| 12.09.2022 11:00 | 24°C        | light rain       | 4 m/s      | 46%                     |
| 12.09.2022 14:00 | 27°C        | overcast clouds  | 2 m/s      | 42%                     |
| 12.09.2022 17:00 | 27°C        | overcast clouds  | 0 m/s      | 22%                     |
| 12.09.2022 20:00 | 26°C        | overcast clouds  | 1 m/s      | 22%                     |
| 12.09.2022 23:00 | 24°C        | scattered clouds | 1 m/s      | 26%                     |
| 13.09.2022 02:00 | 23°C        | light rain       | 0 m/s      | 27%                     |
| 13.09.2022 05:00 | 23°C        | light rain       | 2 m/s      | 27%                     |
| 13.09.2022 08:00 | 23°C        | light rain       | 1 m/s      | 74%                     |
| 13.09.2022 11:00 | 25°C        | light rain       | 1 m/s      | 57%                     |
| 13.09.2022 14:00 | 29°C        | light rain       | 1 m/s      | 54%                     |
+------------------+-------------+------------------+------------+-------------------------+
```

### Get pollution information

```
$ weatherman pollution Belgrade,RS

Air Quality Index (1/Good -> 5/Very poor): 2 (Fair)
Concentration of Carbon monoxide (CO): 233.65
Concentration of Nitrogen monoxide (NO): 0
Concentration of Nitrogen dioxide (NO2): 13.54
Concentration of Ozone (O3): 39.34
Concentration of Sulphur dioxide (SO2): 1.4
Concentration of Fine particles matter (PM2.5): 11.02
Concentration of Coarse particulate matter (PM10): 16.24
Concentration of Ammonia (NH3): 4.43
```

## Specifying location

When running `current`, `forecast` and `pollution` commands, a geographical location needs to be specified in format:

`<city name>,<two-letter ISO 3166 country code>`

eg. `Paris,FR` or `"New York,US"` (note the necessary quotes when city name includes one or more space characters). 

Command `weatherman country list-iso-codes` can be used to display a list of all known countries along with their ISO 3166 codes.

### Setting a default location

`weatherman` allows for configuring a default location which will be used if no location is specified:

```
$ weatherman location set Chichester,GB
Setting location to Chichester,GB
Done!

$ weatherman current

Location: Chichester (GB)
Current temperature: 17°C
Feels like: 17°C
Weather conditions: broken clouds
Wind speed: 3 m/s
Pressure: 1002 mbar
Humidity: 93%
Sunrise: 06:27
Sunset: 19:34
```

If default location is not set, information for `Belgrade,RS` will be displayed when no location is supplied for the `current`, `forecast` or `pollution` commands. 

### Setting a default country

Similarly to default location, a default country can be set so that the country information can be omitted:

```
$ weatherman country set DE
Setting country to Germany
Done!

$ weatherman current Berlin

Location: Alt-Kölln (DE)
Current temperature: 16°C
Feels like: 16°C
Weather conditions: few clouds
Wind speed: 4 m/s
Pressure: 1006 mbar
Humidity: 94%
Sunrise: 06:28
Sunset: 19:40

$ weatherman current Hamburg

Location: Hamburg (DE)
Current temperature: 14°C
Feels like: 14°C
Weather conditions: fog
Wind speed: 2 m/s
Pressure: 1007 mbar
Humidity: 98%
Sunrise: 06:41
Sunset: 19:54
```
If default country is not set, Serbia (`RS`) is presumed.