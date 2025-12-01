package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type PredictionGenerator func(name, dob string) string

var predictionTitles = map[string]string{
	"future-prediction": "Future Prediction",
	"when-will-you-get-married": "When will you get married?",
	"when-will-you-become-a-millionaire": "When will you become a millionaire?",
	"your-future-childs-face": "Your future child's face",
	"which-country-will-you-live-in": "Which country will you live in?",
	"your-2025-fortune-reading": "Your 2025 fortune reading",
}

var predictionGenerators = map[string]PredictionGenerator{
	"future-prediction": func(name, dob string) string {
		days := rand.Intn(365) + 1
		return fmt.Sprintf("In the next %d days, something amazing will happen to %s.", days, name)
	},
	"when-will-you-get-married": func(name, dob string) string {
		years := rand.Intn(10) + 1
		return fmt.Sprintf("You will get married in %d years.", years)
	},
	"when-will-you-become-a-millionaire": func(name, dob string) string {
		years := rand.Intn(20) + 5
		return fmt.Sprintf("You will become a millionaire in %d years.", years)
	},
	"your-future-childs-face": func(name, dob string) string {
		return "Here is what your future child could look like."
	},
	"which-country-will-you-live-in": func(name, dob string) string {
		countries := []string{"Japan", "Italy", "Australia", "Canada", "Brazil", "New Zealand"}
		country := countries[rand.Intn(len(countries))]
		return fmt.Sprintf("You will live in %s.", country)
	},
	"your-2025-fortune-reading": func(name, dob string) string {
		fortunes := []string{
			"A great opportunity will arise.",
			"You will find unexpected joy.",
			"A long-held wish will come true.",
		}
		return fortunes[rand.Intn(len(fortunes))]
	},
}
