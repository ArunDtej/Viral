package handlers

import (
	"Viral/db"
	"encoding/json"
	"math/rand"
	"time"
	"log"
	"Viral/common" // Import the new common package

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
	predictionID := uuid.New().String()
	log.Printf("Saving prediction with ID: %s", predictionID)

	title, ok := common.PredictionTitles[req.Slug]
	if !ok {
		title = common.DefaultPredictionTitle // Default title if not found
	}

	predictionData := map[string]interface{}{
		"page_type":  req.Slug,
		"user_data":  req,
		"prediction": prediction,
		"title":      title, // Include the title
	}

	jsonData, err := json.Marshal(predictionData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	rdb := db.GetKVConnection()
	err = rdb.Set(c.Context(), predictionID, jsonData, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToSavePrediction,
		})
	}

	return c.JSON(fiber.Map{
		"id":         predictionID,
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
	predictionID := uuid.New().String()
	prediction := common.FuturePredictions[rand.Intn(len(common.FuturePredictions))]

	predictionData := map[string]interface{}{
		"page_type":  common.PageTypeFuturePrediction,
		"user_data":  data,
		"prediction": prediction,
	}

	jsonData, err := json.Marshal(predictionData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToCreatePrediction,
		})
	}

	rdb := db.GetKVConnection()
	err = rdb.Set(c.Context(), predictionID, jsonData, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToSavePrediction,
		})
	}

	return c.JSON(fiber.Map{
		"prediction_id": predictionID,
		"prediction":    prediction,
	})
}

func HandleGetPrediction(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("Attempting to retrieve prediction with ID: %s", id)
	rdb := db.GetKVConnection()

	val, err := rdb.Get(c.Context(), id).Result()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": common.ErrPredictionNotFound,
		})
	}

	var predictionData map[string]interface{}
	err = json.Unmarshal([]byte(val), &predictionData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": common.ErrFailedToParsePrediction,
		})
	}

	return c.JSON(predictionData)
}
