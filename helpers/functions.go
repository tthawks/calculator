package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ExtractVariables method extracts needed variables for caluclations
func ExtractVariables(p map[string][]string) (float64, float64, error) {
	_, exists := p["x"]
	if !exists || len(p["x"][0]) < 1 {
		return 0, 0, errors.New("No X value provided")
	}

	_, exists = p["y"]
	if !exists || len(p["y"][0]) < 1 {
		return 0, 0, errors.New("No Y value provided")
	}

	x, err := strconv.ParseFloat(p["x"][0], 64)
	if err != nil {
		return 0, 0, errors.New("Invalid value provided for X (expected number)")
	}

	y, err := strconv.ParseFloat(p["y"][0], 64)
	if err != nil {
		return 0, 0, errors.New("Invalid value provided for Y (expected number)")
	}

	return x, y, nil
}

// EncodeResponseAsJSON method encodes response as JSON object
func EncodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
