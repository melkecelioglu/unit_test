package server

import (
	"encoding/json"
	"io"
	"net/http"
)

type Weather struct {
	City     string `json:"city"`
	Forecast string `json:"forecast"`
}

func GetWeather(url string) (*Weather, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	var weather Weather
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}
