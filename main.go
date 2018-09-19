package main

import (
	"net/http"
	"testAPI/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getWeatherMsg", controllers.GetWeatherMsg)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
