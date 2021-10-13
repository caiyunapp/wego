package backends

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ringsaturn/wego/iface"
)

const (
	CAIYUNAPI       = "http://api.caiyunapp.com/v2.6/%s/%s/weather?lang=%s&dailysteps=%s&alert=true&unit=metric:v2"
	CAIYUNDATE_TMPL = "2006-01-02T15:04-07:00"
)

type CaiyunConfig struct {
	apiKey string
	lang   string
	// debug  bool
	tz *time.Location
}

func (c *CaiyunConfig) Setup() {
	flag.StringVar(&c.apiKey, "caiyun-api-key", "", "forecast backend: the api `KEY` to use")
	flag.StringVar(&c.lang, "caiyun-lang", "en", "forecast backend: the `LANGUAGE` to request from caiyunapp.com/")
	// flag.BoolVar(&c.debug, "forecast-debug", false, "forecast backend: print raw requests and responses")
}

var SkyconToIfaceCode map[string]iface.WeatherCode

func init() {
	//SkyconToIfaceCode["CLEAR_DAY"] = iface.CodeSunny
	SkyconToIfaceCode = map[string]iface.WeatherCode{
		"CLEAR_DAY":           iface.CodeSunny,
		"CLEAR_NIGHT":         iface.CodeSunny,
		"PARTLY_CLOUDY_DAY":   iface.CodePartlyCloudy,
		"PARTLY_CLOUDY_NIGHT": iface.CodePartlyCloudy,
		"CLOUDY":              iface.CodeCloudy,
		"LIGHT_HAZE":          iface.CodeUnknown,
		"MODERATE_HAZE":       iface.CodeUnknown,
		"HEAVY_HAZE":          iface.CodeUnknown,
		"LIGHT_RAIN":          iface.CodeLightRain,
		"MODERATE_RAIN":       iface.CodeLightRain,
		"HEAVY_RAIN":          iface.CodeHeavyRain,
		"STORM_RAIN":          iface.CodeHeavyRain,
		"FOG":                 iface.CodeFog,
		"LIGHT_SNOW":          iface.CodeLightSnow,
		"MODERATE_SNOW":       iface.CodeLightSnow,
		"HEAVY_SNOW":          iface.CodeHeavySnow,
		"STORM_SNOW":          iface.CodeHeavySnow,
		"DUST":                iface.CodeUnknown,
		"SAND":                iface.CodeUnknown,
		"WIND":                iface.CodeUnknown,
	}
}

func ParseCoordinates(latlng string) (float64, float64, error) {
	s := strings.Split(latlng, ",")
	if len(s) != 2 {
		return 0, 0, fmt.Errorf("input %v split to %v parts", latlng, len(s))
	}

	lat, err := strconv.ParseFloat(s[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parse Coodinates failed input %v get parts %v", latlng, s[0])
	}

	lng, err := strconv.ParseFloat(s[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parse Coodinates failed input %v get parts %v", latlng, s[1])
	}
	return lng, lat, nil
}

func (c *CaiyunConfig) Fetch(location string, numdays int) iface.Data {
	res := iface.Data{}
	lat, lng, err := ParseCoordinates(location)
	if err != nil {
		panic(err)
	}
	cyLocation := fmt.Sprintf("%v,%v", lng, lat)
	url := fmt.Sprintf(CAIYUNAPI, c.apiKey, cyLocation, c.lang, strconv.FormatInt(int64(numdays), 10))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	weatherData := &CaiyunWeather{}
	if err := json.Unmarshal(body, weatherData); err != nil {
		panic(err)
	}
	res.Current.Desc = weatherData.Result.ForecastKeypoint
	res.Current.TempC = func() *float32 {
		x := float32(weatherData.Result.Realtime.Temperature)
		return &x
	}()
	if code, ok := SkyconToIfaceCode[weatherData.Result.Realtime.Skycon]; ok {
		res.Current.Code = code
	} else {
		res.Current.Code = iface.CodeUnknown
	}
	if adcodes := weatherData.Result.Alert.Adcodes; len(adcodes) != 0 {
		if len(adcodes) == 3 {
			res.Location = adcodes[1].Name + adcodes[2].Name
		}
		if len(adcodes) == 2 {
			res.Location = adcodes[0].Name + adcodes[1].Name
		}
	} else {
		res.Location = "第三红岸基地"
	}
	res.Current.WinddirDegree = func() *int {
		x := int(weatherData.Result.Realtime.Wind.Direction)
		return &x
	}()
	res.Current.WindspeedKmph = func() *float32 {
		x := float32(weatherData.Result.Realtime.Wind.Speed)
		return &x
	}()
	res.Current.PrecipM = func() *float32 {
		x := float32(weatherData.Result.Realtime.Precipitation.Local.Intensity)
		return &x
	}()
	res.Current.FeelsLikeC = func() *float32 {
		x := float32(weatherData.Result.Realtime.ApparentTemperature)
		return &x
	}()
	res.Current.Humidity = func() *int {
		x := int(weatherData.Result.Realtime.Humidity * 100)
		return &x
	}()
	res.Current.ChanceOfRainPercent = func() *int {
		x := int(weatherData.Result.Minutely.Probability[0] * 100)
		return &x
	}()
	res.Current.VisibleDistM = func() *float32 {
		x := float32(weatherData.Result.Realtime.Visibility)
		return &x
	}()
	dailyDataSlice := []iface.Day{}
	for i := 0; i < numdays; i++ {
		weatherDailyData := weatherData.Result.Daily

		dailyData := iface.Day{
			Date: func() time.Time {
				x, err := time.Parse(CAIYUNDATE_TMPL, weatherDailyData.Temperature[i].Date)
				if err != nil {
					panic(err)
				}
				return x
			}(),
			Slots: []iface.Cond{},
		}

		dailyData.Astronomy = iface.Astro{
			Sunrise: func() time.Time {
				s := strings.Split(weatherDailyData.Astro[i].Sunset.Time, ":")
				hourStr := s[0]
				minuteStr := s[1]
				hour, err := strconv.Atoi(hourStr)
				if err != nil {
					panic(err)
				}
				minute, err := strconv.Atoi(minuteStr)
				if err != nil {
					panic(err)
				}
				x := time.Date(dailyData.Date.Year(), dailyData.Date.Month(), dailyData.Date.Day(), hour, minute, 0, 0, dailyData.Date.Location())
				return x
			}(),
			Sunset: func() time.Time {
				s := strings.Split(weatherDailyData.Astro[i].Sunset.Time, ":")
				hourStr := s[0]
				minuteStr := s[1]
				hour, err := strconv.Atoi(hourStr)
				if err != nil {
					panic(err)
				}
				minute, err := strconv.Atoi(minuteStr)
				if err != nil {
					panic(err)
				}
				x := time.Date(dailyData.Date.Year(), dailyData.Date.Month(), dailyData.Date.Day(), hour, minute, 0, 0, dailyData.Date.Location())
				return x
			}(),
		}

		dateStr := weatherDailyData.Temperature[i].Date[0:10]

		weatherHourlyData := weatherData.Result.Hourly

		for index, houryTmp := range weatherData.Result.Hourly.Temperature {
			if !strings.Contains(houryTmp.Datetime, dateStr) {
				continue
			}
			dailyData.Slots = append(dailyData.Slots, iface.Cond{
				TempC: func() *float32 {
					x := float32(weatherData.Result.Hourly.Temperature[index].Value)
					return &x
				}(),
				VisibleDistM: func() *float32 {
					x := float32(weatherHourlyData.Visibility[index].Value)
					return &x
				}(),
				Humidity: func() *int {
					x := int(weatherHourlyData.Humidity[index].Value)
					return &x
				}(),
				WindspeedKmph: func() *float32 {
					x := float32(weatherHourlyData.Wind[index].Speed)
					return &x
				}(),
				WinddirDegree: func() *int {
					x := int(weatherHourlyData.Wind[index].Direction)
					return &x
				}(),
				Time: func() time.Time {
					x, err := time.Parse(CAIYUNDATE_TMPL, houryTmp.Datetime)
					if err != nil {
						panic(err)
					}
					return x
				}(),
				Code: func() iface.WeatherCode {
					if code, ok := SkyconToIfaceCode[weatherHourlyData.Skycon[index].Value]; ok {
						return code
					} else {
						return iface.CodeUnknown
					}
				}(),
				PrecipM: func() *float32 {
					x := float32(weatherHourlyData.Precipitation[index].Value)
					return &x
				}(),
				FeelsLikeC: func() *float32 {
					x := float32(weatherData.Result.Hourly.ApparentTemperature[index].Value)
					return &x
				}(),
			})
		}

		dailyDataSlice = append(dailyDataSlice, dailyData)
	}
	res.Forecast = dailyDataSlice

	res.GeoLoc = &iface.LatLon{
		Latitude:  float32(weatherData.Location[0]),
		Longitude: float32(weatherData.Location[1]),
	}
	return res
}

func init() {
	iface.AllBackends["caiyunapp.com"] = &CaiyunConfig{}
}
