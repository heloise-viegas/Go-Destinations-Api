package main

import (
	"net/http"

	"github.com/gorilla/mux"
)
	
type Destination struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
}

var destinations []Destination

func getDestinations(w http.ResponseWriter,r *http.Request){

}
func getDestById(w http.ResponseWriter,r *http.Request){
	
}
func createDest(w http.ResponseWriter,r *http.Request){
	
}
func updateDest(w http.ResponseWriter,r *http.Request){
	
}
func deleteDest(w http.ResponseWriter,r *http.Request){
	
}

func main(){
router := mux.NewRouter()

router.HandleFunc("/destinations",getDestinations).Methods("GET")
router.HandleFunc("/destination/{id}",getDestById).Methods("GET")
router.HandleFunc("/destinations",createDest).Methods("POST")
router.HandleFunc("/destination/{id}",updateDest).Methods("PUT")
router.HandleFunc("/destination/{id}",deleteDest).Methods("Delete")
}


