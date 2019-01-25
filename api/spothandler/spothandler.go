package spothandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Spot : location that can be studied at
type Spot struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location string   `json:"location"`
	Tags     []string `json:"tags"`
	// Category string `json:"category"`
}

var spots []Spot

// GetSpots : (term-  bubble tea, library, coffee)
func GetSpots(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query
	category := query()["term"]
	loc := query()["location"]
	// url := "https://api.yelp.com/v3/businesses/search?limit=10&sort_by=rating&term=coffee&location=vancouver"
	url := "https://api.yelp.com/v3/businesses/search?limit=3&open_now=true&sort_by=rating&term=" + category[0] + "&location=" + loc[0]
	bearer := "Bearer JedlKdM0T8RoPKWK5BW2O1-mLE4vUe6JNH2S-78CrlmEMErDdd2DXRuFfBWe4sl3eg-ckkt3aNhdwWh0-OUheboqbFbH2NvjsILnbavJwTSQ59B4Ef6FtrSUjrJHW3Yx"
	// os.Getenv("YELP_API")

	// New Client
	client := &http.Client{}

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", bearer)
	res, err := client.Do(r)
	if err != nil {
		fmt.Printf("HTTP request failed with error %s\n", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	jsonData, err := json.Marshal(string(data))
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("spots.txt", jsonData, 0644)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	return
}

// GetSpot : Get more information on a spot you've selected
func GetSpot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spots)
	fmt.Println("Operation hours, category, tags, photos, location")
}

// CreateRating : Users can tag spots with characteristics from selection (eg. yummy, cute, low noise, parking, outlets)
func CreateRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range spots {
		if item.ID == params["id"] {
			fmt.Println(params["tags"])
			spots[i].Tags = append(spots[i].Tags, params["tags"])
			json.NewEncoder(w).Encode(spots)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
