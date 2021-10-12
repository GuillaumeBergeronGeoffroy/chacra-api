package main

import (
	"fmt"
	"net/http"

	s "github.com/GuillaumeBergeronGeoffroy/chacra-api/service"

	"github.com/spf13/viper"
)

type Config struct {
	Services []s.ServiceConfig
	Gateway  map[string]string
}

func loadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func main() {
	config, err := loadConfig()
	if err != nil || len(config.Services) == 0 {
		println("Error while reading env file:", err)
	}
	for _, service := range config.Services {
		actions, err := s.InitService(service, config.Gateway)
		if err != nil {
			fmt.Println("Error while initializing service: ", err)
			continue
		}
		for route, action := range actions {
			http.HandleFunc(route, action)
		}
	}
	http.ListenAndServe(":3000", nil)
}
