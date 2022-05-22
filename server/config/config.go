package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Server struct {
	System System `json:"system" yaml:"system"`
	Mysql Mysql `json:"mysql" yaml:"mysql"`
}

func ReadConfig(configPath string) (*Server, error) {
	if configPath == "" {
		configPath = "./config.yaml"
	}
	var config Server
	buff, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(buff, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}