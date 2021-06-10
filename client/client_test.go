package client

import (
	"fmt"
	"testing"

	"github.com/jdxj/qweather/config"
)

func initConfig() {
	err := config.ReadConfigure("../config/config.yaml")
	if err != nil {
		panic(err)
	}
}

func TestClient_SearchCity(t *testing.T) {
	initConfig()

	c := NewClient(ModeDebug)

	lp := &LocationParam{
		Location: "beijing",
		Key:      config.QWeather.Key,
		Adm:      "",
		Range:    "",
		Number:   "",
		Gzip:     "",
		Lang:     "",
	}

	locations, err := c.SearchLocation(lp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	for _, v := range locations {
		fmt.Printf("%v\n", *v)
	}
}

func TestLocationParam_Encode(t *testing.T) {
	lp := LocationParam{
		Location: "1",
		Key:      "2",
		Adm:      "3",
		Range:    "",
		Number:   "5",
		Gzip:     "6",
		Lang:     "7",
	}
	fmt.Printf("%s\n", lp.Encode())
}

func TestClient_QueryLiveWeather(t *testing.T) {
	initConfig()

	c := NewClient(ModeDebug)

	lwp := &LiveWeatherParam{
		Location: "101010100",
		Key:      config.QWeather.Key,
		Gzip:     "",
		Lang:     "",
		Unit:     "",
	}
	now, err := c.QueryLiveWeather(lwp)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", *now)
}
