package controllers

import (
	"encoding/json"
	"net/http"
)

func GetWeatherMsg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	good := WeatherMsg{
		StatusCode: 200,
		UpdateAt:   "2018/06/19",
		Data:       []string{"good", "haha"},
	}
	// good := []string{"good", "haha"}
	output, _ := json.MarshalIndent(&good, "", "\t\t")
	w.Write(output)
	// fmt.Println("哈哈")
}

type WeatherMsg struct {
	StatusCode int      `json:"status_code"`
	UpdateAt   string   `json:"update_at"`
	Data       []string `json:"data"`
}
