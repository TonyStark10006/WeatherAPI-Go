package main

import (
	"net/http"
	"weatherAPI/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getWeatherMsg", controllers.GetWeatherMsg)
	mux.HandleFunc("/process", controllers.Process2)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
