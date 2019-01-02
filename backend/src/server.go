package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//JSON struct defining the information TODO: Update the schema after discussion
type sillyPoints struct {
	Name       string  `json:"name"`
	Age        float64 `json:"age"`
	Married    bool    `json:"married"`
	Batter     bool    `json:"batter"`
	Bowler     bool    `json:"bowler"`
	AllRounder bool    `json:"allrounder"`
}

func sillyPointsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("We welcome you to Silly Point")
	//sp := "Welcome to SP"
	sp := sillyPoints{}
	sp.Name = "Bhuvaneshwar Kumar"
	sp.Age = 28
	sp.Batter = false
	sp.Bowler = true
	sp.Married = true
	sp.AllRounder = true
	json.NewEncoder(w).Encode(sp)
}

func server(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/app/v1/sp", sillyPointsHandler).Methods("GET")
	log.Printf("Silly Point is running on: %s", addr)
	http.ListenAndServe(addr, router)
}

func main() {
	server(":9090")
}
