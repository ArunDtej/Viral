package handlers

import (
	"Viral/common"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Updated PredictionGenerator to accept a map for more flexible data
type PredictionGenerator func(userData map[string]interface{}) string

var predictionTitles = common.PredictionTitles

var predictionGenerators = map[string]PredictionGenerator{
	"future-prediction": func(userData map[string]interface{}) string {
		// Select a random prediction from the FuturePredictions slice
		return common.FuturePredictions[rand.Intn(len(common.FuturePredictions))]
	},
	"when-will-you-get-married": func(userData map[string]interface{}) string {
		years := rand.Intn(10) + 1
		return fmt.Sprintf(common.PredictionBaseMarried, years)
	},
	"when-will-you-become-a-millionaire": func(userData map[string]interface{}) string {
		years := rand.Intn(20) + 5
		return fmt.Sprintf(common.PredictionBaseMillionaire, years)
	},
	"how-many-kids-will-you-have": func(userData map[string]interface{}) string {
		name, ok := userData["name"].(string)
		if !ok || name == "" {
			name = "You"
		}
		partner, ok := userData["partner"].(string)
		if !ok || partner == "" {
			partner = "your partner"
		}
		kids := rand.Intn(5) + 1
		return fmt.Sprintf("%s and %s will have %d kids together.", name, partner, kids)
	},
	"which-country-will-you-live-in": func(userData map[string]interface{}) string {
		countries := []string{"Japan", "Italy", "Australia", "Canada", "Brazil", "New Zealand"}
		country := countries[rand.Intn(len(countries))]
		return fmt.Sprintf(common.PredictionBaseLiveInCountry, country)
	},
	"your-2025-fortune-reading": func(userData map[string]interface{}) string {
		return common.FortuneReadings[rand.Intn(len(common.FortuneReadings))]
	},
}
