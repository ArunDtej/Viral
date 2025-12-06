package handlers

import (
	"Viral/common"
	"Viral/db"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type PageData struct {
	PageType string      `json:"page_type"`
	Data     interface{} `json:"data"`
}

func HandlePage(c *fiber.Ctx) error {
	var req PageData
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrCannotParseJSON,
		})
	}

	switch req.PageType {
	case common.PageTypeViral:
		return handleViral(c, req.Data)
	case common.PageTypeFuturePrediction:
		return handleFuturePrediction(c, req.Data)
	case common.PageTypeCats:
		return handleCatsPage(c, req.Data)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrInvalidPageType,
		})
	}
}

var catData = map[string]string{
	"cat1.jpg": "Whiskers",
	"cat2.jpg": "Mittens",
	"cat3.jpg": "Shadow",
}

func handleCatsPage(c *fiber.Ctx, data interface{}) error {
	// get a random cat from the map
	var catNames []string
	for _, name := range catData {
		catNames = append(catNames, name)
	}
	randomCatName := catNames[rand.Intn(len(catNames))]

	// find the filename for the random cat name
	var fileName string
	for f, n := range catData {
		if n == randomCatName {
			fileName = f
			break
		}
	}

	return c.JSON(fiber.Map{
		"file_name": fileName,
		"cat_name":  randomCatName,
	})
}

func HandlePredict(c *fiber.Ctx) error {
	var pageType string
	var userData map[string]interface{}

	// Universal payload structure to capture all possible fields
	var req struct {
		PageType string                 `json:"page_type"`
		UserData map[string]interface{} `json:"user_data"`
		Name     string                 `json:"name"`
		Dob      string                 `json:"dob"`
		Gender   string                 `json:"gender"`
		Partner  string                 `json:"partner"`
		Slug     string                 `json:"slug"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrCannotParseJSON,
		})
	}

	// Determine pageType from either page_type or slug
	// Trim whitespace to handle any accidental spaces
	if strings.TrimSpace(req.PageType) != "" {
		pageType = strings.TrimSpace(req.PageType)
	} else if strings.TrimSpace(req.Slug) != "" {
		pageType = strings.TrimSpace(req.Slug)
	}

	// Log the received values for debugging
	log.Printf("HandlePredict - Received PageType: '%s', Slug: '%s', Final pageType: '%s'", req.PageType, req.Slug, pageType)

	// Consolidate userData
	if req.UserData != nil {
		userData = req.UserData
	} else {
		userData = map[string]interface{}{
			"name":    req.Name,
			"dob":     req.Dob,
			"gender":  req.Gender,
			"partner": req.Partner,
		}
	}

	// Validate pageType is not empty
	if pageType == "" {
		log.Printf("HandlePredict - Empty pageType, PageType: '%s', Slug: '%s'", req.PageType, req.Slug)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrInvalidSlug + ": page_type or slug is required",
		})
	}

	generator, ok := predictionGenerators[pageType]
	if !ok {
		log.Printf("HandlePredict - Invalid slug '%s', available slugs: %v", pageType, getAvailableSlugs())
		// Return both error and message for backward compatibility with old frontend code
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   common.ErrInvalidSlug + ": '" + pageType + "'",
			"message": common.ErrInvalidSlug,
		})
	}

	prediction := generator(userData)
	uid := uuid.New().String()
	log.Printf("Saving prediction with UID: %s", uid)

	title, ok := common.PredictionTitles[pageType]
	if !ok {
		title = common.DefaultPredictionTitle
	}

	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	predictionsData := map[string]interface{}{
		"prediction_text": prediction,
		"title":           title,
	}
	predictionsJSON, err := json.Marshal(predictionsData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	_, err = db.Pdb.Exec(`INSERT INTO predictions (uid, page_type, user_data, predictions) VALUES ($1, $2, $3, $4)`,
		uid, pageType, userDataJSON, predictionsJSON)
	if err != nil {
		log.Printf("Error saving prediction to PostgreSQL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToSavePrediction,
		})
	}

	return c.JSON(fiber.Map{
		"id":         uid,
		"page_type":  pageType,
		"prediction": prediction,
		"title":      title,
	})
}

func handleViral(c *fiber.Ctx, data interface{}) error {
	predictions := []fiber.Map{
		{"slug": "when-will-you-get-married", "title": common.PredictionTitles["when-will-you-get-married"]},
		{"slug": "when-will-you-become-a-millionaire", "title": common.PredictionTitles["when-will-you-become-a-millionaire"]},
		{"slug": "how-many-kids-will-you-have", "title": common.PredictionTitles["how-many-kids-will-you-have"]},
		{"slug": "which-country-will-you-live-in", "title": common.PredictionTitles["which-country-will-you-live-in"]},
		{"slug": "your-2025-fortune-reading", "title": common.PredictionTitles["your-2025-fortune-reading"]},
	}

	return c.JSON(fiber.Map{
		"predictions": predictions,
	})
}

func handleFuturePrediction(c *fiber.Ctx, data interface{}) error {
	uid := uuid.New().String()
	prediction := common.FuturePredictions[rand.Intn(len(common.FuturePredictions))]

	userDataJSON, err := json.Marshal(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	predictionsData := map[string]interface{}{
		"prediction_text": prediction,
	}
	predictionsJSON, err := json.Marshal(predictionsData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}
	_, err = db.Pdb.Exec(`INSERT INTO predictions (uid, page_type, user_data, predictions) VALUES ($1, $2, $3, $4)`,
		uid, common.PageTypeFuturePrediction, userDataJSON, predictionsJSON)
	if err != nil {
		log.Printf("Error saving future prediction to PostgreSQL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToSavePrediction,
		})
	}

	return c.JSON(fiber.Map{
		"prediction_id": uid,
		"prediction":    prediction,
	})
}

func HandleGetPrediction(c *fiber.Ctx) error {
	uid := c.Params("id")
	log.Printf("Attempting to retrieve prediction with UID: %s", uid)

	var id int
	var pageType string
	var userDataJSON, predictionsJSON []byte
	var createdAt time.Time

	row := db.Pdb.QueryRow(`SELECT id, page_type, user_data, predictions, created_at FROM predictions WHERE uid = $1`, uid)
	err := row.Scan(&id, &pageType, &userDataJSON, &predictionsJSON, &createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": common.ErrPredictionNotFound,
			})
		}
		log.Printf("Error retrieving prediction from PostgreSQL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve prediction",
		})
	}

	var userData map[string]interface{}
	err = json.Unmarshal(userDataJSON, &userData)
	if err != nil {
		log.Printf("Error unmarshaling user_data: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToParsePrediction,
		})
	}

	var predictions map[string]interface{}
	err = json.Unmarshal(predictionsJSON, &predictions)
	if err != nil {
		log.Printf("Error unmarshaling predictions: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToParsePrediction,
		})
	}

	// Get previous and next predictions
	prevPredictions, err := getRelatedPredictions(id, "previous")
	if err != nil {
		log.Printf("Error getting previous predictions: %v", err)
		// Not a fatal error, so we can continue
	}

	nextPredictions, err := getRelatedPredictions(id, "next")
	if err != nil {
		log.Printf("Error getting next predictions: %v", err)
		// Not a fatal error, so we can continue
	}

	response := fiber.Map{
		"id":                   uid,
		"page_type":            pageType,
		"user_data":            userData,
		"created_at":           createdAt,
		"previous_predictions": prevPredictions,
		"next_predictions":     nextPredictions,
	}

	if predText, ok := predictions["prediction_text"]; ok {
		response["prediction"] = predText
	}
	if title, ok := predictions["title"]; ok {
		response["title"] = title
	}

	return c.JSON(response)
}

func HandleGetRecentPredictions(c *fiber.Ctx) error {
	rows, err := db.Pdb.Query(`SELECT uid, page_type, user_data, predictions, created_at FROM predictions ORDER BY created_at DESC LIMIT 5`)
	if err != nil {
		log.Printf("Error retrieving recent predictions: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve recent predictions",
		})
	}
	defer rows.Close()

	var recentPredictions []fiber.Map
	for rows.Next() {
		var uid, pageType string
		var userDataJSON, predictionsJSON []byte
		var createdAt time.Time

		if err := rows.Scan(&uid, &pageType, &userDataJSON, &predictionsJSON, &createdAt); err != nil {
			log.Printf("Error scanning recent prediction: %v", err)
			continue
		}

		var userData map[string]interface{}
		if err := json.Unmarshal(userDataJSON, &userData); err != nil {
			log.Printf("Error unmarshaling user_data: %v", err)
			continue
		}

		var predictions map[string]interface{}
		if err := json.Unmarshal(predictionsJSON, &predictions); err != nil {
			log.Printf("Error unmarshaling predictions: %v", err)
			continue
		}

		predictionMap := fiber.Map{
			"uid":        uid,
			"page_type":  pageType,
			"user_data":  userData,
			"created_at": createdAt,
		}

		if predText, ok := predictions["prediction_text"]; ok {
			predictionMap["prediction"] = predText
		}
		if title, ok := predictions["title"]; ok {
			predictionMap["title"] = title
		}

		recentPredictions = append(recentPredictions, predictionMap)
	}

	return c.JSON(fiber.Map{
		"recent_predictions": recentPredictions,
	})
}

type RelatedPrediction struct {
	UID      string `json:"uid"`
	PageType string `json:"page_type"`
}

func getRelatedPredictions(id int, direction string) ([]RelatedPrediction, error) {
	var rows *sql.Rows
	var err error

	if direction == "previous" {
		rows, err = db.Pdb.Query(`SELECT uid, page_type FROM predictions WHERE id < $1 ORDER BY id DESC LIMIT 5`, id)
	} else {
		rows, err = db.Pdb.Query(`SELECT uid, page_type FROM predictions WHERE id > $1 ORDER BY id ASC LIMIT 5`, id)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var relatedPredictions []RelatedPrediction
	for rows.Next() {
		var p RelatedPrediction
		if err := rows.Scan(&p.UID, &p.PageType); err != nil {
			return nil, err
		}
		relatedPredictions = append(relatedPredictions, p)
	}

	// Reverse the order of previous predictions to be ascending
	if direction == "previous" {
		for i, j := 0, len(relatedPredictions)-1; i < j; i, j = i+1, j-1 {
			relatedPredictions[i], relatedPredictions[j] = relatedPredictions[j], relatedPredictions[i]
		}
	}

	return relatedPredictions, nil
}

// getAvailableSlugs returns a list of available prediction slugs for debugging
func getAvailableSlugs() []string {
	slugs := make([]string, 0, len(predictionGenerators))
	for slug := range predictionGenerators {
		slugs = append(slugs, slug)
	}
	return slugs
}
