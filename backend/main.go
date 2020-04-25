package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	ID     int     `json:"id"`
	Name   string  `json:"name,omitempty"`
	Price  float32 `json:"price,omitempty"`
	Wheels int     `json:"numWheels,omitempty"`
	Manual bool    `json:"isManual"`
}

var vehicles []Vehicle

func main() {
	router := mux.NewRouter()
	addData()
	buildRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func addData() {
	vehicles = append(vehicles, Vehicle{ID: 1, Price: 438227.12, Name: "Tesla", Wheels: 4, Manual: false})
	vehicles = append(vehicles, Vehicle{ID: 2, Price: 54555.13, Name: "Toyota", Wheels: 4, Manual: true})
}

func buildRoutes(router *mux.Router) {
	router.HandleFunc("/vehicles", getVehicles).Methods("GET")
	router.HandleFunc("/vehicles/{id}", getVehicle).Methods("GET")
	router.HandleFunc("/vehicles", createVehicle).Methods("POST")
	router.HandleFunc("/vehicles/{id}", updateVehicle).Methods("PUT")
	router.HandleFunc("/vehicles/{id}", deleteVehicle).Methods("DELETE")
}

func getVehicles(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	json.NewEncoder(writer).Encode(vehicles)
}

func getVehicle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	params := mux.Vars(request)
	for _, vehicle := range vehicles {
		if params["id"] == strconv.Itoa(vehicle.ID) {
			json.NewEncoder(writer).Encode(vehicle)
			return
		}

	}
	json.NewEncoder(writer).Encode(&Vehicle{})
}

func createVehicle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	var v Vehicle
	json.NewDecoder(request.Body).Decode(&v)
	v.ID = len(vehicles) + 1
	vehicles = append(vehicles, v)
	json.NewEncoder(writer).Encode(v)
}

func updateVehicle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	var v Vehicle
	json.NewDecoder(request.Body).Decode(&v)
	id, err := strconv.Atoi(strings.TrimPrefix(request.RequestURI, "/vehicles/"))
	if err != nil {
		log.Fatal(err)
	}
	vehicles[id-1] = v
	json.NewEncoder(writer).Encode(v)
}

func deleteVehicle(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	params := mux.Vars(request)
	for index, v := range vehicles {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatal(err)
		}
		if v.ID == id {
			fmt.Printf("I am deleteing vehicles %v", vehicles)
			vehicles = remove(vehicles, index)
			json.NewEncoder(writer).Encode(vehicles)
			return
		}
	}
	json.NewEncoder(writer).Encode(vehicles)
}

func remove(s []Vehicle, i int) []Vehicle {
	s[i] = s[len(s)-1]  // copy last elem to index i
	return s[:len(s)-1] // erase last elem and truncate slice
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
