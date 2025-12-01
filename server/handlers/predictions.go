package handlers

import (
	"fmt"
	"math/rand"
	"time"
	"Viral/common" // Import the new common package
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type PredictionGenerator func(name, dob string) string

var predictionTitles = common.PredictionTitles // Use the constant from common package

var predictionGenerators = map[string]PredictionGenerator{
	"future-prediction": func(name, dob string) string {
		days := rand.Intn(365) + 1
		return fmt.Sprintf(common.PredictionBaseAmazingEvent, days, name)
	},
	"when-will-you-get-married": func(name, dob string) string {
		years := rand.Intn(10) + 1
		return fmt.Sprintf(common.PredictionBaseMarried, years)
	},
	"when-will-you-become-a-millionaire": func(name, dob string) string {
		years := rand.Intn(20) + 5
		return fmt.Sprintf(common.PredictionBaseMillionaire, years)
	},
	"your-future-childs-face": func(name, dob string) string {
		return common.PredictionBaseFutureChild
	},
	"which-country-will-you-live-in": func(name, dob string) string {
		countries := []string{"Japan", "Italy", "Australia", "Canada", "Brazil", "New Zealand"}
		country := countries[rand.Intn(len(countries))]
		return fmt.Sprintf(common.PredictionBaseLiveInCountry, country)
	},
	"your-2025-fortune-reading": func(name, dob string) string {
		return common.FortuneReadings[rand.Intn(len(common.FortuneReadings))]
	},
}
