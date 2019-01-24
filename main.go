package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Spot : location that can be studied at
type Spot struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Tags     []string `json:"tags"`
	// Longitude float64 `json:"longitude"`
	// Latitude  float64 `json:"latitude"`
	// Category string `json:"category"`
}

var spots []Spot

// term: bubble tea, library, coffee
// Params: open_now: true; sort_by: rating; location: user location
func getSpots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := "https://api.yelp.com/v3/businesses/search?term=restaurants&location=vancouver"
	bearer := "Bearer " + os.Getenv("YELP_TOKEN")

	// NANI
	req, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", bearer)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	json.NewEncoder(w).Encode(string(data))
	return
}

// Get more information on a spot you've selected
func getSpot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spots)
	fmt.Println("Operation hours, category, tags, photos, location")
}

// Users can tag spots with characteristics from selection (eg. yummy, cute, low noise, parking, outlets)
func createRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range spots {
		if item.ID == params["id"] {
			fmt.Println(params["id"])
			// NANI
			fmt.Println(params["tags"])
			spots[i].Tags = append(spots[i].Tags, params["tags"])
			json.NewEncoder(w).Encode(spots)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func main() {

	// Sample Data
	spots = append(spots, Spot{ID: "1", Name: "Green Leaf Sushi", Location: "on broadway", Tags: []string{}})
	spots = append(spots, Spot{ID: "2", Name: "Little Sheep Hot Pot", Location: "also on broadway", Tags: []string{}})

	fmt.Println(spots)

	router := mux.NewRouter()
	router.HandleFunc("/flow", getSpots).Methods("GET")
	router.HandleFunc("/flow/{id}", getSpot).Methods("GET")
	router.HandleFunc("/flow/spot/{id}", createRating).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
