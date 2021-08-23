package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server struct {
		Token string `yaml:"token"`
	}
	AccountInfo struct {
		Iban  string `yaml:"iban"`
		Swift string `yaml:"swift"`
		Name  string `yaml:"name"`
	}
}

var Cfg Config

func Initialization() {

	readFile(&Cfg)
	fmt.Printf("%+v", Cfg)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config/config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
