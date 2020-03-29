package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	ID        int     `json:"id"`
	Name      string  `json:"name,omitempty"`
	Price     float32 `json:"price,omitempty"`
	numWheels int     `json:"numWheels,omitempty"`
	isManual  bool    `json:"isManual"`
}

var vehicles []Vehicle

func main() {
	fmt.Println("Hello World!")
	router := mux.NewRouter()
	buildRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func buildRoutes(router *mux.Router) {
	router.HandleFunc("/bar", GetInfo).Methods("GET")
}

func GetInfo(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hi")
	// initHeaders(writer)
	json.NewEncoder(writer).Encode(writer)
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
