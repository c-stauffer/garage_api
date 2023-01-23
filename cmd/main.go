package main

import (
	"log"
	"net/http"

	"code.crogge.rs/chris/garage_api/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/actuate", handlers.Actuate).Methods(http.MethodGet)
	router.HandleFunc("/state", handlers.GetState).Methods(http.MethodGet)
	router.HandleFunc("/temperature", handlers.GetTemperature).Methods(http.MethodGet)
	router.HandleFunc("/humidity", handlers.GetHumidity).Methods(http.MethodGet)
	router.HandleFunc("/heatindex", handlers.GetHeatIndex).Methods(http.MethodGet)

	c := cors.Default()
	handler := c.Handler(router)

	log.Println("API is running!")
	http.ListenAndServe(":4001", handler)
}
