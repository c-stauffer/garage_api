package main

import (
	"log"
	"net/http"

	"code.crogge.rs/chris/garage_api/pkg/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// func BasicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		user, pass, ok := r.BasicAuth()

// 		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
// 			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		handler(w, r)
// 	}
// }

func main() {
	router := mux.NewRouter()
	// user := "username"
	// pass := "password"
	// realm := "Enter username and password."

	router.HandleFunc("/actuate", handlers.Actuate).Methods(http.MethodGet)
	router.HandleFunc("/ping", handlers.Actuate).Methods(http.MethodGet)
	router.HandleFunc("/doorstate", handlers.GetState).Methods(http.MethodGet)
	router.HandleFunc("/temperature", handlers.GetTemperature).Methods(http.MethodGet)
	router.HandleFunc("/humidity", handlers.GetHumidity).Methods(http.MethodGet)
	router.HandleFunc("/heatindex", handlers.GetHeatIndex).Methods(http.MethodGet)

	// router.HandleFunc("/actuate", BasicAuth(handlers.Actuate, user, pass, realm)).Methods(http.MethodGet)
	// router.HandleFunc("/ping", BasicAuth(handlers.Ping, user, pass, realm)).Methods(http.MethodGet)
	// router.HandleFunc("/doorstate", BasicAuth(handlers.GetState, user, pass, realm)).Methods(http.MethodGet)
	// router.HandleFunc("/temperature", BasicAuth(handlers.GetTemperature, user, pass, realm)).Methods(http.MethodGet)
	// router.HandleFunc("/humidity", BasicAuth(handlers.GetHumidity, user, pass, realm)).Methods(http.MethodGet)
	// router.HandleFunc("/heatindex", BasicAuth(handlers.GetHeatIndex, user, pass, realm)).Methods(http.MethodGet)

	c := cors.Default()
	handler := c.Handler(router)

	log.Println("API is running!")
	http.ListenAndServe(":4001", handler)
}
