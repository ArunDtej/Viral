package handlers

import (
	"Viral/db"
	"encoding/json"
	"math/rand"
	"time"
	"log" // Added log import

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
			"error": "Cannot parse JSON",
		})
	}

	switch req.PageType {
	case "viral":
		return handleViral(c, req.Data)
	case "future_prediction":
		return handleFuturePrediction(c, req.Data)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page_type",
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
			"error": "Cannot parse JSON",
		})
	}

	generator, ok := predictionGenerators[req.Slug]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid slug",
		})
	}

	prediction := generator(req.Name, req.Dob)
	predictionID := uuid.New().String()
	log.Printf("Saving prediction with ID: %s", predictionID)

	title, ok := predictionTitles[req.Slug]
	if !ok {
		title = "Viral Prediction" // Default title if not found
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
			"error": "Failed to create prediction data",
		})
	}

	rdb := db.GetKVConnection()
	err = rdb.Set(c.Context(), predictionID, jsonData, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save prediction",
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
		{"slug": "when-will-you-get-married", "title": "When will you get married?"},
		{"slug": "when-will-you-become-a-millionaire", "title": "When will you become a millionaire?"},
		{"slug": "your-future-childs-face", "title": "Your future child’s face."},
		{"slug": "which-country-will-you-live-in", "title": "Which country will you live in?"},
		{"slug": "your-2025-fortune-reading", "title": "Your 2025 fortune reading."},
	}

	return c.JSON(fiber.Map{
		"predictions": predictions,
	})
}

var futurePredictions = []string{
	"You will embark on a great adventure in the coming year.",
	"A surprise visitor will bring you unexpected joy.",
	"You will discover a hidden talent you never knew you had.",
	"An old friend will reconnect with you with exciting news.",
	"You will find a solution to a problem that has been bothering you.",
	"A financial windfall is coming your way.",
	"You will travel to a place you've always dreamed of visiting.",
	"You will make a new friend who will change your life for the better.",
	"An opportunity to learn a new skill will present itself.",
	"You will receive a promotion or recognition at work.",
	"You will find a lost item of great sentimental value.",
	"A secret admirer will reveal themselves.",
	"You will overcome a fear that has been holding you back.",
	"You will start a new hobby that brings you immense pleasure.",
	"You will receive an unexpected gift.",
	"You will have a chance to make a positive impact on your community.",
	"You will find a new favorite book or movie.",
	"You will have a memorable meal at a new restaurant.",
	"You will have a chance to spend quality time with loved ones.",
	"You will experience a moment of pure bliss.",
	"You will learn a valuable lesson from a mistake.",
	"You will find a new source of inspiration.",
	"You will have a dream that provides you with a brilliant idea.",
	"You will receive a heartfelt compliment from someone you admire.",
	"You will find a new favorite song that you can't stop listening to.",
	"You will have a chance to help someone in need.",
	"You will discover a new favorite place in your own city.",
	"You will have a day where everything goes perfectly.",
	"You will receive an invitation to an exciting event.",
	"You will find a new reason to be grateful.",
}

func handleFuturePrediction(c *fiber.Ctx, data interface{}) error {
	predictionID := uuid.New().String()
	prediction := futurePredictions[rand.Intn(len(futurePredictions))]

	predictionData := map[string]interface{}{
		"page_type":  "future_prediction",
		"user_data":  data,
		"prediction": prediction,
	}

	jsonData, err := json.Marshal(predictionData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create prediction data",
		})
	}

	rdb := db.GetKVConnection()
	err = rdb.Set(c.Context(), predictionID, jsonData, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save prediction",
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
			"error": "Prediction not found",
		})
	}

	var predictionData map[string]interface{}
	err = json.Unmarshal([]byte(val), &predictionData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse prediction data",
		})
	}

	return c.JSON(predictionData)
}
