package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

// ParseBody reads and unmarshals JSON from an HTTP request body into the specified interface.
// It returns an error if reading or unmarshalling fails.
func ParseBody(req *http.Request, target interface{}) error {
	log.Println("Parsing request body")

	// Ensure the body is closed after reading to prevent resource leaks
	defer req.Body.Close()

	// Read the request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return errors.New("failed to read request body")
	}

	// Unmarshal the JSON into the target interface
	if err := json.Unmarshal(body, target); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return errors.New("invalid JSON format")
	}

	log.Println("Request body parsed successfully")
	return nil
}

// getEnv retrieves environment variables or returns a fallback value
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
