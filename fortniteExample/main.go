package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func GetWeapons(w http.ResponseWriter, r *http.Request) {
	weapons := LoadWeaponsCsv()
	json.NewEncoder(w).Encode(weapons)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/weapons", GetWeapons).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
