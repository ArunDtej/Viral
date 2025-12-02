package handlers

import (
	"Viral/common" // Import the new common package
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
	var req struct {
		Name   string `json:"name"`
		Dob    string `json:"dob"`
		Gender string `json:"gender"`
		Slug   string `json:"slug"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrCannotParseJSON,
		})
	}

	generator, ok := predictionGenerators[req.Slug]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": common.ErrInvalidSlug,
		})
	}

	prediction := generator(req.Name, req.Dob)
	uid := uuid.New().String() // Use uid for PostgreSQL
	log.Printf("Saving prediction with UID: %s", uid)

	title, ok := common.PredictionTitles[req.Slug]
	if !ok {
		title = common.DefaultPredictionTitle // Default title if not found
	}

	// Prepare user_data for JSONB storage
	userDataJSON, err := json.Marshal(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	// Prepare predictions data for JSONB storage
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

	// Insert into PostgreSQL
	_, err = db.Pdb.Exec(`INSERT INTO predictions (uid, page_type, user_data, predictions) VALUES ($1, $2, $3, $4)`,
		uid, req.Slug, userDataJSON, predictionsJSON)
	if err != nil {
		log.Printf("Error saving prediction to PostgreSQL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToSavePrediction,
		})
	}

	return c.JSON(fiber.Map{
		"id":         uid, // Return uid as id
		"page_type":  req.Slug,
		"prediction": prediction,
		"title":      title,
	})
}

func handleViral(c *fiber.Ctx, data interface{}) error {
	predictions := []fiber.Map{
		{"slug": "when-will-you-get-married", "title": common.PredictionTitles["when-will-you-get-married"]},
		{"slug": "when-will-you-become-a-millionaire", "title": common.PredictionTitles["when-will-you-become-a-millionaire"]},
		{"slug": "your-future-childs-face", "title": common.PredictionTitles["your-future-childs-face"]},
		{"slug": "which-country-will-you-live-in", "title": common.PredictionTitles["which-country-will-you-live-in"]},
		{"slug": "your-2025-fortune-reading", "title": common.PredictionTitles["your-2025-fortune-reading"]},
	}

	return c.JSON(fiber.Map{
		"predictions": predictions,
	})
}

func handleFuturePrediction(c *fiber.Ctx, data interface{}) error {
	uid := uuid.New().String() // Use uid for PostgreSQL
	prediction := common.FuturePredictions[rand.Intn(len(common.FuturePredictions))]

	// Prepare user_data for JSONB storage
	userDataJSON, err := json.Marshal(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	// Prepare predictions data for JSONB storage
	predictionsData := map[string]interface{}{
		"prediction_text": prediction,
	}
	predictionsJSON, err := json.Marshal(predictionsData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	// Insert into PostgreSQL
	_, err = db.Pdb.Exec(`INSERT INTO predictions (uid, page_type, user_data, predictions) VALUES ($1, $2, $3, $4)`,
		uid, common.PageTypeFuturePrediction, userDataJSON, predictionsJSON)
	if err != nil {
		log.Printf("Error saving future prediction to PostgreSQL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToSavePrediction,
		})
	}

	return c.JSON(fiber.Map{
		"prediction_id": uid, // Return uid as prediction_id
		"prediction":    prediction,
	})
}

func HandleGetPrediction(c *fiber.Ctx) error {
	uid := c.Params("id") // id is now uid
	log.Printf("Attempting to retrieve prediction with UID: %s", uid)

	var pageType string
	var userDataJSON, predictionsJSON []byte
	var createdAt time.Time

	row := db.Pdb.QueryRow(`SELECT page_type, user_data, predictions, created_at FROM predictions WHERE uid = $1`, uid)
	err := row.Scan(&pageType, &userDataJSON, &predictionsJSON, &createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": common.ErrPredictionNotFound,
			})
		}
		log.Printf("Error retrieving prediction from PostgreSQL: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve prediction", // More generic error for internal issues
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

	// Construct the response similar to how it was returned by KVrocks
	// This might need adjustment based on frontend expectations
	response := fiber.Map{
		"id":         uid,
		"page_type":  pageType,
		"user_data":  userData,
		"predictions": predictions, // This will contain prediction_text and title
		"created_at": createdAt,
	}

	// If it's a "viral" prediction, extract prediction_text and title
	if pageType != common.PageTypeFuturePrediction {
		if predText, ok := predictions["prediction_text"]; ok {
			response["prediction"] = predText
		}
		if title, ok := predictions["title"]; ok {
			response["title"] = title
		}
		delete(response, "predictions") // Remove the raw predictions map if specific fields are extracted
	} else {
		// For future predictions, the 'predictions' JSONB might just contain "prediction_text"
		if predText, ok := predictions["prediction_text"]; ok {
			response["prediction"] = predText
		}
		delete(response, "predictions")
	}


	return c.JSON(response)
}
