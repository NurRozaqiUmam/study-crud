package config

import "github.com/tkanos/gonfig"

// configuration, struct yang menyimpan konfigurasi dari aplikasi
type Configuration struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBName     string `json:"db_name"`
	SignKey    string `json:"sign_key"`
}

var conf Configuration

func GetConfig() Configuration {
	err := gonfig.GetConf("config/config.json", &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
