package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config model
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Url struct {
		Business string `yaml:"business"`
	} `yaml:"url"`
}

func Get() *Config {
	f, err := os.Open("configs/config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return &config
}
