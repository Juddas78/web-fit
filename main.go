package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// upper: chest, upper back, middle back, shoulders, biceps, triceps
// lower: lower back, glutes, quads, hamstrings, calves, abs
var upper = []string{"chest", "upper back", "middle back", "shoulders", "biceps", "triceps"}
var lower = []string{"lower back", "glutes", "quads", "hamstrings", "calves", "abs"}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET").Schemes("http")
	r.HandleFunc("/workout", HomeHandler)
	http.ListenAndServe(":8000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createWorkout(lower, 5))
	fmt.Println("request recieved")
}
