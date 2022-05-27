package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Destination struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var destinations []Destination

func getDestinations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(destinations) //encode and return destinations splice
}
func getDestById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range destinations {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
func createDest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//create id using rand func from math library
	var dest Destination
	json.NewDecoder(r.Body).Decode(dest) //decode values from request and store at dest
	destinations = append(destinations, dest)
	json.NewEncoder(w).Encode(destinations)

}
func updateDest(w http.ResponseWriter, r *http.Request) {

}
func deleteDest(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/destinations", getDestinations).Methods("GET")
	router.HandleFunc("/destination/{id}", getDestById).Methods("GET")
	router.HandleFunc("/destinations", createDest).Methods("POST")
	router.HandleFunc("/destination/{id}", updateDest).Methods("PUT")
	router.HandleFunc("/destination/{id}", deleteDest).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", router))
}
