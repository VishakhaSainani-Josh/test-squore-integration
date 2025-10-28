package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Feed struct {
	Section     string `json:"Section"`
	Name        string `json:"Name"`
	FeedMatches string `json:"FeedMatches"`
	ValidFrom   string `json:"ValidFrom"`
	ValidTo     string `json:"ValidTo"`
}

type Match struct {
	A      string `json:"A"`
	B      string `json:"B"`
	ID     string `json:"id"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Court  string `json:"court"`
	Result string `json:"result,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/feeds", func(w http.ResponseWriter, r *http.Request) {
		feeds := []Feed{
			{
				Section:     "SRFI 2025 Season",
				Name:        "Men’s Division A",
				FeedMatches: "/feedmatches/1",
				ValidFrom:   "20250101",
				ValidTo:     "20251231",
			},
			{
				Section:     "SRFI 2025 Season",
				Name:        "Women’s Division A",
				FeedMatches: "/feedmatches/2",
				ValidFrom:   "20250101",
				ValidTo:     "20251231",
			},
		}
		writeJSON(w, feeds)
	})

	mux.HandleFunc("/feedmatches/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/feedmatches/")
		var matches []Match

		if id == "1" {
			matches = []Match{
				{A: "Dish", B: "Mad", ID: "101", Date: "2025-10-28", Time: "12:30", Court: "C4"},
				{A: "Ravi", B: "Arjun", ID: "102", Date: "2025-10-29", Time: "10:00", Court: "C2"},
			}
		} else if id == "2" {
			matches = []Match{
				{A: "Neha", B: "Aditi", ID: "201", Date: "2025-10-28", Time: "14:00", Court: "C3"},
			}
		}

		response := map[string]interface{}{
			"config": map[string]interface{}{
				"name":                     "Sample League " + id,
				"numberOfGamesToWinMatch":  3,
				"numberOfPointsToWinGame":  11,
				"clubName":                 "Chennai Squash Club",
				"clubLogo":                 "https://mysite.com/assets/chennai-logo.png",
			},
			"matches": matches,
		}
		writeJSON(w, response)
	})

	mux.HandleFunc("/match/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/match/")
		match := map[string]interface{}{
			"id":        id,
			"players":   map[string]string{"A": "Dish", "B": "Mad"},
			"court":     "ABSC - C4",
			"event":     map[string]string{"name": "Yadupati Singhania Memorial Squash Tournament", "division": "BU15"},
			"result":    "3-1",
			"gamescores": "11-4, 9-11, 11-6, 11-8",
			"timing": []map[string]interface{}{
				{"start": "2025-10-28T12:30:00+05:30", "end": "2025-10-28T13:10:00+05:30", "offsets": []int{0, 20, 45, 70}},
			},
		}
		writeJSON(w, match)
	})

	log.Println("✅ Squore Feed API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
