package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}

// env-default:"production
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"  env-default:"production"` //abstract tax
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")

		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exit: %s", configPath)

	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config files: %s", err.Error())
	}
	return &cfg

}
