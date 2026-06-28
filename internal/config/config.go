package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env:"HTTP_SERVER_ADDRESS" env-default:"localhost:8080"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-default:"production"`
	StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-default:"storage/storage.db"`
	HttpServer  HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPathFlag := flag.String("config", "", "Path to config file")
		flag.Parse()
		configPath = *configPathFlag
	}
	if configPath == "" {
		panic("config path is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("Config file not exist")
	}
	var cfg Config
	e := cleanenv.ReadConfig(configPath, &cfg)
	if e != nil {
		log.Fatal("Couldn't get data")
	}
	return &cfg
}
