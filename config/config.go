package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	QWeather qWeather
)

func ReadConfigure(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	cfg := new(configure)
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	QWeather = cfg.QWeather
	return nil
}

type configure struct {
	QWeather qWeather `yaml:"qweather"`
}

type qWeather struct {
	Key string `yaml:"key"`
}
