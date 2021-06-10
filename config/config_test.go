package config

import (
	"fmt"
	"testing"
)

func TestReadConfigure(t *testing.T) {
	err := ReadConfigure("config.yaml")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", QWeather)
}
