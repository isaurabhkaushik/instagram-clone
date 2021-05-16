package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Env string    `yaml:"env" validate:"required"`
	DB  *Database `yaml:"database" validate:"required"`
}

type Database struct {
	Host     string `yaml:"host" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Dbname   string `yaml:"dbname" validate:"required"`
	Port     int    `yaml:"port" validate:"required"`
}

func LoadConfig(filepath string) *Config {
	f, err := os.Open(filepath)
	if err != nil {
		log.Printf("LoadConfig: failed to open config.yml with error: %v", err)
		panic(nil)
	}
	defer f.Close()

	var cfg *Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Printf("LoadConfig: failed to decode config.yml with error: %v", err)
		panic(nil)
	}
	return cfg
}
