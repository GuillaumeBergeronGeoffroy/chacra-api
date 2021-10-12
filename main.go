package main

import (
	"fmt"
	"net/http"

	s "github.com/GuillaumeBergeronGeoffroy/chacra-api/service"

	"github.com/spf13/viper"
)

type Config struct {
	services []s.ServiceConfig
	gateway  map[string]string
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	config := Config{}
	viper.Unmarshal(&config)
	for _, service := range config.services {
		actions, err := s.InitService(service, config.gateway)
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
