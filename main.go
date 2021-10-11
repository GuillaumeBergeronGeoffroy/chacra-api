package main

import (
	"fmt"
	"net/http"

	s "./service"
	"github.com/spf13/viper"
)

type Config struct {
	ServicesConfig []s.ServiceConfig
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	config := Config{}
	viper.Unmarshal(&config)
	for _, serviceConfig := range config.ServicesConfig {
		serviceActions, err := s.InitService(serviceConfig)
		if err != nil {
			fmt.Println("Error while initializing service: ", err)
			continue
		}
		for route, action := range serviceActions {
			http.HandleFunc(route, action)
		}
	}
	http.ListenAndServe(":3000", nil)
}
