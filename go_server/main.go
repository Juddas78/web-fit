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
	r.HandleFunc("/workout", WorkoutHandler).Methods("POST", "GET")
	http.ListenAndServe(":8000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	workout, _ := createWorkout(upper, 5)
	json.NewEncoder(w).Encode(workout)
	fmt.Println("home request recieved")
}

func WorkoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("workout request recieved\n")
	w.Header().Set("Content-Type", "application/json")
	var groups []string
	var workout Workout
	var err error
	method := r.Method
	if method != "POST" {
		userGroup = upper
	} else {
		err := json.NewDecoder(r.Body).Decode(&groups)
		// err := r.ParseForm()
		if err != nil {
			_ = fmt.Errorf("Got this error %s", err)
		} else {
			for _, group := range groups {
				fmt.Printf("group added %s\n", group)
				userGroup = append(userGroup, group)
			}
		}
	}
	fmt.Printf("usergroup: %s\n\n", userGroup)
	workout, err = createWorkout(userGroup, 5)
	if err == nil {
		_ = fmt.Errorf("Error recieved: %s", err)
	}
	json.NewEncoder(w).Encode(workout)
	userGroup = []string{}
}
