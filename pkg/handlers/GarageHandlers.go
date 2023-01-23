package handlers

import (
	"net/http"
)

func Actuate(w http.ResponseWriter, r *http.Request)        {}
func GetState(w http.ResponseWriter, r *http.Request)       {}
func GetTemperature(w http.ResponseWriter, r *http.Request) {}
func GetHumidity(w http.ResponseWriter, r *http.Request)    {}
func GetHeatIndex(w http.ResponseWriter, r *http.Request)   {}
