package main

import (
	"encoding/json"
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
	router := mux.NewRouter()
	addData()
	buildRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func addData() {
	vehicles = append(vehicles, Vehicle{ID: 1, Price: 438227.12, Name: "Tesla", numWheels: 4, isManual: false})
	vehicles = append(vehicles, Vehicle{ID: 2, Price: 54555.13, Name: "Toyota", numWheels: 4, isManual: true})
}

func buildRoutes(router *mux.Router) {
	router.HandleFunc("/vehicles", getVehicles).Methods("GET")
	router.HandleFunc("/vehicles/{id}", getVehicle).Methods("GET")
	// router.HandleFunc("/vehicles", createVehicle).Methods("POST")
	// router.HandleFunc("/vehicles/{id}", updateVehicle).Methods("PUT")
	// router.HandleFunc("/vehicles/{id}", deleteVehicle).Methods("DELETE")
}

func getVehicles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	json.NewEncoder(writer).Encode(vehicles)
}

func getVehicle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	params := mux.Vars(request)
	for _, vehicle := range vehicles {
		log.Printf("This is the current vehicle's id: %d and here is the param id: %s", vehicle.ID, params["id"])
		if string(vehicle.ID) == params["id"] {
			json.NewEncoder(writer).Encode(vehicle)
			return
		}

	}
	log.Printf("Hi this is a meme")
	json.NewEncoder(writer).Encode(&Vehicle{})
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
