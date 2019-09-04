package main

import (
	confstate "github.com/mamemomonga/go-confstate"
)

type Configs struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Timeout  int    `yaml:"timeout"`
	Debug    bool   `yaml:"debug"`
}

func LoadConfigs(cf string) error {
	confstate.ConfigsFile = cf
	confstate.DefaultConfigsFile = "ap7900.yaml"
	confstate.DefaultBaseDirType = confstate.DBTWork
	confstate.Debug = false

	confstate.Configs = &Configs{
		Username: "apc",
		Password: "apc",
		Address:  "192.168.1.100:23",
		Timeout:  10,
		Debug:    false,
	}
	if err := confstate.LoadConfigs(); err != nil {
		return err
	}
	return nil
}
func Conf() *Configs {
	return confstate.Configs.(*Configs)
}
