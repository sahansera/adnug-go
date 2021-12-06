package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Forecast struct {
	Date        time.Time `json:"date"`
	Temperature int       `json:"temperature"`
	Summary     string    `json:"summary"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	var summaries = []string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

	var forecasts = []Forecast{}
	for i := 0; i < 5; i++ {
		forecast := Forecast{
			Date:        time.Now().AddDate(0, 0, i),
			Temperature: rand.Intn(20),
			Summary:     summaries[rand.Intn(len(summaries))],
		}
		forecasts = append(forecasts, forecast)
	}
	reponse, _ := json.Marshal(&forecasts)
	fmt.Fprintf(w, string(reponse))
}

func main() {
	http.HandleFunc("/weatherforecast", hello)
	http.ListenAndServe(":8080", nil)
}
