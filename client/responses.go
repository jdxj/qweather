package client

type GeneralResponse struct {
	Code       string `json:"code"`
	Refer      *Refer `json:"refer"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`

	Locations []*Location `json:"location"`
	Now       *Now        `json:"now"`
}

type Refer struct {
	Sources []string `json:"sources"`
	License []string `json:"license"`
}

// Location 用于城市搜索
type Location struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm1      string `json:"adm1"`
	Adm2      string `json:"adm2"`
	Country   string `json:"country"`
	TZ        string `json:"tz"`
	UTCOffset string `json:"utcOffset"`
	IsDst     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

// Now 实况天气
type Now struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	PrecIP    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}
