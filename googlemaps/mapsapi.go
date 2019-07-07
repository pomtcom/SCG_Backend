package googlemaps

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GoogleMapsAPI() string {
	apiKey := "GOOLGE API TOKEN"
	response, err := http.Get("https://maps.googleapis.com/maps/api/place/textsearch/json?query=restaurants+in+Bang+Sue&key=" + apiKey)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}
