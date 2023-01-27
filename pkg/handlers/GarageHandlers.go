package handlers

import (
	"io"
	"log"
	"net/http"
	"time"
)

const microAddress string = "http://192.168.1.45:6262/"
const defaultRequestTimeoutSeconds int = 5

func Ping(w http.ResponseWriter, r *http.Request) {
	sendGetRequest(w, "ping", defaultRequestTimeoutSeconds)
}

func Actuate(w http.ResponseWriter, r *http.Request) {
	sendGetRequest(w, "actuate", defaultRequestTimeoutSeconds)
}

func GetState(w http.ResponseWriter, r *http.Request) {
	sendGetRequest(w, "state", defaultRequestTimeoutSeconds)
}

func GetTemperature(w http.ResponseWriter, r *http.Request) {
	sendGetRequest(w, "temperature", defaultRequestTimeoutSeconds)
}

func GetHumidity(w http.ResponseWriter, r *http.Request) {
	sendGetRequest(w, "humidity", defaultRequestTimeoutSeconds)
}

func GetHeatIndex(w http.ResponseWriter, r *http.Request) {
	sendGetRequest(w, "heatindex", defaultRequestTimeoutSeconds)
}

func sendGetRequest(w http.ResponseWriter, uri string, timeoutSeconds int) {
	client := http.Client{
		Timeout: time.Duration(timeoutSeconds) * time.Second,
	}

	w.Header().Add("Content-Type", "text/plain")

	resp, err := client.Get(microAddress + uri)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
