package main

import (
	"math/rand"
	"time"
)

// upper: chest, upper back, middle back, shoulders, biceps, triceps
// lower: lower back, glutes, quads, hamstrings, calves, abs

var chest = []string{
	"Barbell Bench Press",
	"Incline Barbell Bench Press",
	"Dumbbell Bench Press",
	"Incline Dumbbell Bench Press",
	"Weighted Dips for Chest",
	"Decline Bench Press",
}

var low_back = []string{
	"Barbell Deadlift",
	"Good Mornings",
	"Back Extensions",
}
var mid_back = []string{
	"Pull-ups",
	"Wide-Grip Pull-up",
	"Standing T-Bar Row",
	"Wide-Grip Pulldown",
	"Close-Grip Pulldown",
	"Seated Cable Row",
}
var upper_back = []string{
	"Shrug",
	"Bent-Over Barbell Row",
	"Farmer' Carry",
	"Face Pull",
}

var abs = []string{
	"Cable Crunch",
	"Hanging Dumbbell Knee Raise",
	"Landmine",
	"Captain’s Chair Leg Raise",
	"Ab-Wheel Rollout",
	"Plank",
	"Weighted Decline Sit-up",
	"Bicycle",
}

var shoulder = []string{

	"Standing Barbell Overhead Press",
	"Seated Dumbbell Overhead Press",
	"Arnold Shoulder Press",
	"Bent-Over Reverse Fly",
	"Lateral Raise",
	"Front Dumbbell Raise",
}

var leg = []string{

	"Barbell Squat",
	"Standing Calf Raise",
	"Calf Press",
	"Romanian Deadlift",
	"Leg Press",
	"Hack Squat",
	"Dumbbell Lunge",
}

var biceps = []string{

	"Seated Hammer Curl",
	"Standing Barbell Curl",
	"Inverted Rows",
	"Zottman Curl",
	"Weighted Chin-ups",
	"Incline-Bench Curl",
	"Preacher EZ-Bar Curl",
	"Standing Cable Curl",
}
var triceps = []string{

	"Close-Grip Bench Press",
	"Skullcrusher",
	"Weighted Dips",
	"Cable Push-Down",
	"Dumbbell Kickback",
	"Seated Overhead Dumbbell Extension",
}

var lifts map[string][]string

type Workout struct {
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	Sets []Set  `json:"sets"`
	Name string `json:"name"`
}

type Set struct {
	Reps      int `json:"reps"`
	Rest      int `json:"rest"`
	Intensity int `json:"intensity"`
}

type Groups []string

func createWorkout(groups Groups, intensity int) Workout {
	// to do assumed hypertrophic, add support for others here
	// to do add used exercises list to pass to chooselift
	r := rand.NewSource(time.Now().UnixNano())
	s := rand.New(r)
	exercises := []Exercise{}
	for _, group := range groups {
		exercise := Exercise{Sets: chooseSets(intensity), Name: chooseLift(group, s)}
		// fmt.Printf("Group: %s, Lift: %s\n", group, exercise.name)
		exercises = append(exercises, exercise)
	}
	return Workout{exercises}
}

func chooseSets(intensity int) []Set {
	sets := []Set{}
	switch intensity {
	case 8, 9, 10:
		sets = []Set{{12, 90, intensity - 2}, {12, 90, intensity - 1}, {10, 90, intensity}, {10, 90, intensity}, {8, 90, intensity}}
	case 5, 6, 7:
		sets = []Set{{12, 90, intensity - 2}, {10, 90, intensity - 1}, {10, 90, intensity}, {8, 90, intensity}}
	case 1, 2, 3, 4:
		sets = []Set{{12, 90, intensity - 2}, {10, 90, intensity - 1}, {8, 90, intensity}}
	}
	return sets
}

func chooseLift(group string, s *rand.Rand) string {
	switch group {
	case "chest":
		return chest[s.Intn(len(chest))]
	case "lower back":
		return low_back[s.Intn(len(low_back))]
	case "upper back":
		return upper_back[s.Intn(len(upper_back))]
	case "middle back":
		return mid_back[s.Intn(len(mid_back))]
	case "abs":
		return abs[s.Intn(len(abs))]
	case "triceps":
		return triceps[s.Intn(len(triceps))]
	case "biceps":
		return biceps[s.Intn(len(biceps))]
	case "legs", "glutes", "quads", "hamstrings", "calves":
		return leg[s.Intn(len(leg))]
	case "shoulders":
		return shoulder[s.Intn(len(shoulder))]
	default:
		return "group not found"
	}
}
