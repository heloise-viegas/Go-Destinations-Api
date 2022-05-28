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
	json.NewDecoder(r.Body).Decode(&dest) //decode values from request and store at dest
	destinations = append(destinations, dest)
	json.NewEncoder(w).Encode(destinations)

}
func updateDest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get id from request

	var dest Destination
	json.NewDecoder(r.Body).Decode(&dest) //get dest vales to upadte and store at &dest

	for index, item := range destinations {
		if item.ID == params["id"] { ///if id matches request id then delete that dest from splice
			destinations = append(destinations[:index], destinations[index+1:]...)
			dest.ID = params["id"] //for new dest values reset the id to original id that was sent in params
			destinations = append(destinations, dest)
			json.NewEncoder(w).Encode(destinations)

		}
	}

}
func deleteDest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range destinations {
		if item.ID == params["id"] {
			destinations = append(destinations[:index], destinations[index+1:]...)
			json.NewEncoder(w).Encode(params["id"])
		}
	}

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
