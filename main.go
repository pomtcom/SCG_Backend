package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

	if (*r).Method == "GET" {

	} else if (*r).Method == "POST" {
		fmt.Println("NOT ALLOW for this method")
	}

	fmt.Println(("findXYZvalue method is executing"))

	var answer Answer
	answer.X = findSequnceDigit(5)
	answer.Y = findSequnceDigit(6)
	answer.Z = findSequnceDigit(7)
	json.NewEncoder(w).Encode(answer)

}

func main() {
	fmt.Println("SCG webservice backend api is executing")
	//fmt.Print(findXvalue())
	//fmt.Print(findYvalue())
	//fmt.Print(findZvalue())
	router := mux.NewRouter()
	router.HandleFunc("/App-HealthCheck", GetAppHealthCheck).Methods("GET")
	router.HandleFunc("/findXYZ", findXYZvalue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func findSequnceDigit(index int) int{
	n1 := 3
	n2 := 5

	if index == 0 {
		return n1
	}else if index == 1{
		return n2
	}

	nTarget := 0
	for i := 3; i <= index; i++ {
		nTarget = n1 + n2 + 1
		n1 = n2
		n2 = nTarget
	}
	return nTarget
}