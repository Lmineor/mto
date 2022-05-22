package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	configPath := "../config.yaml"
	cfg, err := ReadConfig(configPath)
	if err != nil{
		t.Logf("err from read config %s", err)
		t.Fail()
	}
	fmt.Println(cfg.System.Addr)

}
