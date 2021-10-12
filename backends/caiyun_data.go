package backends

type CaiyunWeather struct {
	Status     string    `json:"status"`
	APIVersion string    `json:"api_version"`
	APIStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    int       `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Alert struct {
			Status  string `json:"status"`
			Content []struct {
				Province      string    `json:"province"`
				Status        string    `json:"status"`
				Code          string    `json:"code"`
				Description   string    `json:"description"`
				RegionID      string    `json:"regionId"`
				County        string    `json:"county"`
				Pubtimestamp  int       `json:"pubtimestamp"`
				Latlon        []float64 `json:"latlon"`
				City          string    `json:"city"`
				AlertID       string    `json:"alertId"`
				Title         string    `json:"title"`
				Adcode        string    `json:"adcode"`
				Source        string    `json:"source"`
				Location      string    `json:"location"`
				RequestStatus string    `json:"request_status"`
			} `json:"content"`
			Adcodes []struct {
				Adcode int    `json:"adcode"`
				Name   string `json:"name"`
			} `json:"adcodes"`
		} `json:"alert"`
		Realtime struct {
			Status      string  `json:"status"`
			Temperature float64 `json:"temperature"`
			Humidity    float64 `json:"humidity"`
			Cloudrate   float64 `json:"cloudrate"`
			Skycon      string  `json:"skycon"`
			Visibility  float64 `json:"visibility"`
			Dswrf       float64 `json:"dswrf"`
			Wind        struct {
				Speed     float64 `json:"speed"`
				Direction float64 `json:"direction"`
			} `json:"wind"`
			Pressure            float64 `json:"pressure"`
			ApparentTemperature float64 `json:"apparent_temperature"`
			Precipitation       struct {
				Local struct {
					Status     string  `json:"status"`
					Datasource string  `json:"datasource"`
					Intensity  float64 `json:"intensity"`
				} `json:"local"`
				Nearest struct {
					Status    string  `json:"status"`
					Distance  float64 `json:"distance"`
					Intensity float64 `json:"intensity"`
				} `json:"nearest"`
			} `json:"precipitation"`
			AirQuality struct {
				Pm25 int     `json:"pm25"`
				Pm10 int     `json:"pm10"`
				O3   int     `json:"o3"`
				So2  int     `json:"so2"`
				No2  int     `json:"no2"`
				Co   float64 `json:"co"`
				Aqi  struct {
					Chn int `json:"chn"`
					Usa int `json:"usa"`
				} `json:"aqi"`
				Description struct {
					Chn string `json:"chn"`
					Usa string `json:"usa"`
				} `json:"description"`
			} `json:"air_quality"`
			LifeIndex struct {
				Ultraviolet struct {
					Index float64 `json:"index"`
					Desc  string  `json:"desc"`
				} `json:"ultraviolet"`
				Comfort struct {
					Index int    `json:"index"`
					Desc  string `json:"desc"`
				} `json:"comfort"`
			} `json:"life_index"`
		} `json:"realtime"`
		Minutely struct {
			Status          string    `json:"status"`
			Datasource      string    `json:"datasource"`
			Precipitation2H []float64 `json:"precipitation_2h"`
			Precipitation   []float64 `json:"precipitation"`
			Probability     []float64 `json:"probability"`
			Description     string    `json:"description"`
		} `json:"minutely"`
		Hourly struct {
			Status        string `json:"status"`
			Description   string `json:"description"`
			Precipitation []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"precipitation"`
			Temperature []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"temperature"`
			ApparentTemperature []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"apparent_temperature"`
			Wind []struct {
				Datetime  string  `json:"datetime"`
				Speed     float64 `json:"speed"`
				Direction float64 `json:"direction"`
			} `json:"wind"`
			Humidity []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"humidity"`
			Cloudrate []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"cloudrate"`
			Skycon []struct {
				Datetime string `json:"datetime"`
				Value    string `json:"value"`
			} `json:"skycon"`
			Pressure []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"pressure"`
			Visibility []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"visibility"`
			Dswrf []struct {
				Datetime string  `json:"datetime"`
				Value    float64 `json:"value"`
			} `json:"dswrf"`
			AirQuality struct {
				Aqi []struct {
					Datetime string `json:"datetime"`
					Value    struct {
						Chn int `json:"chn"`
						Usa int `json:"usa"`
					} `json:"value"`
				} `json:"aqi"`
				Pm25 []struct {
					Datetime string `json:"datetime"`
					Value    int    `json:"value"`
				} `json:"pm25"`
			} `json:"air_quality"`
		} `json:"hourly"`
		Daily struct {
			Status string `json:"status"`
			Astro  []struct {
				Date    string `json:"date"`
				Sunrise struct {
					Time string `json:"time"`
				} `json:"sunrise"`
				Sunset struct {
					Time string `json:"time"`
				} `json:"sunset"`
			} `json:"astro"`
			Precipitation []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"precipitation"`
			Temperature []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature"`
			Temperature08H20H []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature_08h_20h"`
			Temperature20H32H []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"temperature_20h_32h"`
			Wind []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind"`
			Wind08H20H []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind_08h_20h"`
			Wind20H32H []struct {
				Date string `json:"date"`
				Max  struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"max"`
				Min struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"min"`
				Avg struct {
					Speed     float64 `json:"speed"`
					Direction float64 `json:"direction"`
				} `json:"avg"`
			} `json:"wind_20h_32h"`
			Humidity []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"humidity"`
			Cloudrate []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"cloudrate"`
			Pressure []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"pressure"`
			Visibility []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"visibility"`
			Dswrf []struct {
				Date string  `json:"date"`
				Max  float64 `json:"max"`
				Min  float64 `json:"min"`
				Avg  float64 `json:"avg"`
			} `json:"dswrf"`
			AirQuality struct {
				Aqi []struct {
					Date string `json:"date"`
					Max  struct {
						Chn int `json:"chn"`
						Usa int `json:"usa"`
					} `json:"max"`
					Avg struct {
						Chn float64 `json:"chn"`
						Usa float64 `json:"usa"`
					} `json:"avg"`
					Min struct {
						Chn int `json:"chn"`
						Usa int `json:"usa"`
					} `json:"min"`
				} `json:"aqi"`
				Pm25 []struct {
					Date string  `json:"date"`
					Max  int     `json:"max"`
					Avg  float64 `json:"avg"`
					Min  int     `json:"min"`
				} `json:"pm25"`
			} `json:"air_quality"`
			Skycon []struct {
				Date  string `json:"date"`
				Value string `json:"value"`
			} `json:"skycon"`
			Skycon08H20H []struct {
				Date  string `json:"date"`
				Value string `json:"value"`
			} `json:"skycon_08h_20h"`
			Skycon20H32H []struct {
				Date  string `json:"date"`
				Value string `json:"value"`
			} `json:"skycon_20h_32h"`
			LifeIndex struct {
				Ultraviolet []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"`
				} `json:"ultraviolet"`
				CarWashing []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"`
				} `json:"carWashing"`
				Dressing []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"`
				} `json:"dressing"`
				Comfort []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"`
				} `json:"comfort"`
				ColdRisk []struct {
					Date  string `json:"date"`
					Index string `json:"index"`
					Desc  string `json:"desc"`
				} `json:"coldRisk"`
			} `json:"life_index"`
		} `json:"daily"`
		Primary          int    `json:"primary"`
		ForecastKeypoint string `json:"forecast_keypoint"`
	} `json:"result"`
}
