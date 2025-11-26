package middlewares

import (
	"encoding/json"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var dangerousWords = []string{
	"select", "insert", "update", "delete", "drop", "union", "--",
	"rm", "mv", "cp", "sh", "bash", "cmd", "powershell",
}

var dangerousSymbols = []string{
	";", "`", "|", "&", "$", ">", "<", "(", ")", "\\",
}

var skipKeys = map[string]bool{
	"access_token":  true,
	"refresh_token": true,
	"token":         true,
	"jwt":           true,
	"target":        true,
}

// Detect email-like strings
func isEmail(s string) bool {
	// simple email detection (doesn't need to be RFC strict)
	return strings.Contains(s, "@") && strings.Contains(s, ".")
}

// Detect URL-like strings
func isURL(s string) bool {
	return strings.HasPrefix(strings.ToLower(s), "http://") ||
		strings.HasPrefix(strings.ToLower(s), "https://")
}

// Match dangerous word only as a whole word, not part of another word
func matchWholeWord(text, word string) bool {
	re := regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
	return re.MatchString(text)
}

func isDangerous(value string) bool {
	decoded := value
	if d, err := url.QueryUnescape(value); err == nil {
		decoded = d
	}
	normalized := strings.ToLower(decoded)

	// Skip clearly safe patterns
	if isEmail(normalized) || isURL(normalized) {
		return false
	}

	// Check dangerous words (whole word only)
	for _, word := range dangerousWords {
		if matchWholeWord(normalized, word) {
			log.Println("Blocked: matched word:", word)
			return true
		}
	}

	// Check dangerous symbols (anywhere)
	for _, symbol := range dangerousSymbols {
		if strings.Contains(decoded, symbol) {
			log.Println("Blocked: matched symbol:", symbol)
			return true
		}
	}

	return false
}

func SanitizeInputMiddleware(c *fiber.Ctx) error {

	if !strings.HasPrefix(c.Path(), "/app/") {
		return c.Next()
	}

	// Query parameters
	for key, val := range c.Queries() {
		if skipKeys[strings.ToLower(key)] {
			continue
		}
		if isDangerous(val) {
			log.Printf("Blocked query param [%s]: %s", key, val)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "error",
				"message": "Blocked suspicious input in query parameters",
			})
		}
	}

	// Form inputs
	if c.Is("application/x-www-form-urlencoded") || c.Is("multipart/form-data") {
		if form, err := c.MultipartForm(); err == nil {
			for key, vals := range form.Value {
				if skipKeys[strings.ToLower(key)] {
					continue
				}
				for _, val := range vals {
					if isDangerous(val) {
						log.Printf("Blocked form param [%s]: %s", key, val)
						return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
							"status":  "error",
							"message": "Blocked suspicious input in form parameters",
						})
					}
				}
			}
		}
	}

	// JSON body (regardless of content-type)
	bodyBytes := c.Body()
	if json.Valid(bodyBytes) {
		var body map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &body); err == nil {
			for key, val := range body {
				if skipKeys[strings.ToLower(key)] {
					continue
				}
				strVal, ok := val.(string)
				if !ok {
					continue
				}
				if isDangerous(strVal) {
					log.Printf("Blocked JSON param [%s]: %s", key, strVal)
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"status":  "error",
						"message": "Blocked suspicious input in JSON body",
					})
				}
			}
		}
	}

	return c.Next()
}
