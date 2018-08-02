package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Region struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Station struct {
	Name   string  `json:"name"`
	Id     string  `json:"id"`
	Region *Region `json:"region"`
}

var stations []Station

func main() {
	// manually adding a region
	r1 := Region{
		Name: "Shannon",
		Id:   2,
	}

	// manually adding stations as we don't have a db now
	s1 := Station{
		Name:   "Limerick",
		Id:     "0000001043",
		Region: &r1,
	}

	s2 := Station{
		Name:   "Boolick",
		Id:     "0000001011",
		Region: &r1,
	}

	stations = []Station{s1, s2}

	router := mux.NewRouter()
	router.HandleFunc("/stations", GetStations).Methods("GET")
	router.HandleFunc("/stations/{id}", GetStation).Methods("GET")
	router.HandleFunc("/stations", CreateStation).Methods("POST")
	router.HandleFunc("/stations/{id}", DeleteStation).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetStations(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stations)
}

func GetStation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, station := range stations {
		if station.Id == params["id"] {
			json.NewEncoder(w).Encode(station)
		}
	}
}

func CreateStation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var station Station
	_ = json.NewDecoder(r.Body).Decode(&station)
	station.Id = params["id"]
	stations = append(stations, station)
	json.NewEncoder(w).Encode(stations)
}

func DeleteStation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, station := range stations {
		if station.Id == params["id"] {
			stations = append(stations[:i], stations[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(stations)
}
