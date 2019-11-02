package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Env       bool   `yaml:"env"`
	MigrateDB bool   `yaml:"migrate_db"`
	Version   string `yaml:"version"`
	App       struct {
		Port string `yaml:"port"`
	} `yaml:"app"`

	Mysql struct {
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int64  `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"mysql"`

	Log struct {
		Path       string `yaml:"path"`
		AccessName string `yaml:"access_name"`
		ErrorName  string `yaml:"error_name"`
		ErrLevel   string `yaml:"error_level"`
	} `yaml:"log"`

	Redis struct {
		Host         string `yaml:"host"`
		Port         int64  `yaml:"port"`
		Password     string `yaml:"password"`
		Db           int    `yaml:"db"`
		DialTimeOut  int    `yaml:"dialTimeOut"`
		ReadTimeOut  int    `yaml:"readTimeOut"`
		WriteTimeOut int    `yaml:"writeTimeOut"`
		PoolTimeOut  int    `yaml:"poolTimeOut"`
		PoolSize     int    `yaml:"poolSize"`
	} `yaml:"redis"`
}

func loadConfig(file string) *Config {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
