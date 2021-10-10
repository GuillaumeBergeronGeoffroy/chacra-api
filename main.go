package main

import (
	"net/http"
)

type Action func(w http.ResponseWriter, r *http.Request)
type Actions map[string]Action
type Services []interface{}

// use viper package to read .env file
// return the value of the key
func viperEnvVariable(key string) (value string, err string) {
	// viper.SetConfigFile(".env")
	// err = viper.ReadInConfig()
	// if err != null {
	// 	return
	// }
	// value, ok := viper.Get(key).(string)
	// if !ok {
	// 	err = "Invalid type assert"
	// } else if !value {
	// 	err = "Invalid value"
	// }
	// return
}

func main() {
	// services := Services{}
	// for _, val := range services {
	// 	value, error := viperEnvVariable(val.key)
	// }
	// // for service / params in .env
	// // execute service settup
	// http.ListenAndServe(":3000", nil)
}
