package handlers

import (
	"Viral/common"
	"Viral/db"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
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
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrInvalidPageType,
		})
	}
}

func HandlePredict(c *fiber.Ctx) error {
	var pageType string
	var userData map[string]interface{}

	// Attempt to parse the new payload structure
	var newReq struct {
		PageType string                 `json:"page_type"`
		UserData map[string]interface{} `json:"user_data"`
	}
	if err := c.BodyParser(&newReq); err == nil && newReq.PageType != "" {
		pageType = newReq.PageType
		userData = newReq.UserData
	} else {
		// If new payload parsing fails or is incomplete, attempt to parse the old payload structure
		var oldReq struct {
			Name   string `json:"name"`
			Dob    string `json:"dob"`
			Gender string `json:"gender"`
			Slug   string `json:"slug"`
		}
		if err := c.BodyParser(&oldReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": common.ErrCannotParseJSON,
			})
		}
		pageType = oldReq.Slug
		userData = map[string]interface{}{
			"name":   oldReq.Name,
			"dob":    oldReq.Dob,
			"gender": oldReq.Gender,
		}
	}

	generator, ok := predictionGenerators[pageType]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrInvalidSlug,
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
