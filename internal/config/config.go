package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"Addr" env-required:"true"`
}

type Config struct {
	//`yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`  <--- These are struct tags
	Env        string `yaml:"env" env:"ENV" env-required:"true" env-default:"dev"`
	Storage    string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		//we can try getting it from args like go run main.go --config_path=xyz
		flags := flag.String("config", "", "path to the configuration path")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config Path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read the config file: %s", err.Error())
	}
	return &cfg

}
