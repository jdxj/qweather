package client

import (
	"fmt"
	"net/url"
	"reflect"
)

const (
	format = "%s?%s"

	// 城市信息搜索
	APICity = "https://geoapi.qweather.com/v2/city/lookup"
	// 实况天气
	APILiveWeather = "https://devapi.qweather.com/v7/weather/now"
)

func encodeParam(i interface{}) string {
	iType := reflect.TypeOf(i).Elem()
	iValue := reflect.ValueOf(i).Elem()

	param := url.Values{}
	for i := 0; i < iType.NumField(); i++ {
		k := iType.Field(i).Tag.Get("qw")
		v := iValue.Field(i).Interface().(string)
		if v == "" {
			continue
		}
		param.Set(k, v)
	}
	return param.Encode()
}

type LocationParam struct {
	Location string `qw:"location"`
	Key      string `qw:"key"`
	Adm      string `qw:"adm"`
	Range    string `qw:"range"`
	Number   string `qw:"number"`
	Gzip     string `qw:"gzip"`
	Lang     string `qw:"lang"`
}

func (lp *LocationParam) Encode() string {
	return fmt.Sprintf(format, APICity, encodeParam(lp))
}

type LiveWeatherParam struct {
	Location string `qw:"location"` // LocationID/经纬度
	Key      string `qw:"key"`
	Gzip     string `qw:"gzip"` // y/n
	Lang     string `qw:"lang"`
	Unit     string `qw:"unit"` // m/i
}

func (lwp *LiveWeatherParam) Encode() string {
	return fmt.Sprintf(format, APILiveWeather, encodeParam(lwp))
}
