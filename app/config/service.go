package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	AppConfig  AppConfig  `json:"app"`
	UserConfig UserConfig `json:"user"`
}
type UserConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AppConfig struct {
	PTCClientSecret string `json:"ptc_client_secret"`
	AndroidId       string `json:"android_id"`
	Service         string `json:"service"`
	App             string `json:"app"`
	ClientSIG       string `json:"client_sig"`
	ClientId        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
}

func ReadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
