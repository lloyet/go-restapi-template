package app

import (
	"api/config"
	"api/routes"
	"fmt"
	"log"
	"net/http"
	"time"
)

var api *http.Server

func Start() {
	fmt.Println("Listening on " + config.HOST + ":" + config.PORT)
	log.Fatal(api.ListenAndServe())
}

func init() {
	api = &http.Server{
		Addr:           config.HOST + ":" + config.PORT,
		Handler:        routes.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
