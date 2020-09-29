package config

import (
	"os"
)

var PORT string = "8080"
var HOST string = ""
var MONGO_PATH string = "mongodb://localhost:27017"
var MONGO_DB string = "api"

func init() {
	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}
	if host := os.Getenv("HOST"); host != "" {
		HOST = host
	}
	if mongopath := os.Getenv("MONGO_PATH"); mongopath != "" {
		MONGO_PATH = mongopath
	}
	if mongodb := os.Getenv("MONGO_DB"); mongodb != "" {
		MONGO_DB = mongodb
	}
}
