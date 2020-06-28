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
var userGroup = []string{}

func main() {
	fmt.Println("Starting development server :)")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET").Schemes("http")
	r.HandleFunc("/workout", HomeHandler).Methods("POST", "GET")
	http.ListenAndServe(":8000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var groups []string
	var workout Workout
	method := r.Method
	fmt.Println(r.Body)
	if method != "POST" {
		workout = createWorkout(lower, 5)
	} else {
		_ = json.NewDecoder(r.Body).Decode(groups)
		err := r.ParseForm()
		if err != nil {
			fmt.Errorf("Got this error", err)
		} else {
			for _, group := range groups {
				fmt.Println("%s", group)
				groups = append(userGroup, group)
			}
			workout = createWorkout(userGroup, 5)
		}
	}
	json.NewEncoder(w).Encode(workout)
	fmt.Println("request recieved", r.Method, workout)
}
