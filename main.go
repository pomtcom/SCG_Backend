package main

import (
	"./findsequence"
	"./googlemaps"
	"./linenoti"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"
	"log"
	"net/http"
)

type HealthCheck struct {
	Response string `json:"response"`
}

type Answer struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GetAppHealthCheck(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if (*r).Method == "GET" {

	} else if (*r).Method == "POST" {
		fmt.Println("NOT ALLOW for this method")
	}

	fmt.Println(("GET method for healthcheck endpoint is working"))
	var healthCheck HealthCheck
	healthCheck.Response = "OK"
	json.NewEncoder(w).Encode(healthCheck)

}

func findXYZvalue(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	uuid, _ := uuid.NewV4()
	linenoti.LineNotification("findXYZvalue web-service is executing with session id " + uuid.String())

	if (*r).Method == "GET" {

	} else if (*r).Method == "POST" {
		fmt.Println("NOT ALLOW for this method")
	}

	fmt.Println(("findXYZvalue method is executing"))

	var answer Answer
	answer.X = findsequence.FindSequnceDigit(5)
	answer.Y = findsequence.FindSequnceDigit(6)
	answer.Z = findsequence.FindSequnceDigit(7)
	json.NewEncoder(w).Encode(answer)

}

func findRestaurantsBangSue(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	uuid, _ := uuid.NewV4()
	linenoti.LineNotification("findRestaurantsBangSue web-service is executing with session id " + uuid.String())

	if (*r).Method == "GET" {

	} else if (*r).Method == "POST" {
		fmt.Println("NOT ALLOW for this method")
	}

	fmt.Println(("findRestaurantsBangSue method is executing"))

	mapsResponseData := googlemaps.GoogleMapsAPI()

	//fmt.Println(string(responseData))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(mapsResponseData))

}

func main() {
	fmt.Println("SCG webservice backend api is executing")
	router := mux.NewRouter()
	router.HandleFunc("/App-HealthCheck", GetAppHealthCheck).Methods("GET")
	router.HandleFunc("/findXYZ", findXYZvalue).Methods("GET")
	router.HandleFunc("/findRestaurantsBangSue", findRestaurantsBangSue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

