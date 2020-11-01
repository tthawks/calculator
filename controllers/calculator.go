package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tthawks/calculator/models"
)

type calculatorController struct{}

var result models.CacheEntry

func (cc calculatorController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte("Please use GET method"))
		return
	}

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s?%s", r.URL.Path, r.URL.RawQuery)))
	key := fmt.Sprintf("%x\n", h.Sum(nil))

	data, cached := models.CacheGet(key)

	if cached {
		enc := json.NewEncoder(w)
		enc.Encode(data)
		return
	}

	data.Key = key
	result = data

	switch r.URL.Path {
	case "/add":
		cc.add(w, r)
	case "/subtract":
		cc.sub(w, r)
	case "/multiply":
		cc.mul(w, r)
	case "/divide":
		cc.div(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Canot find path that you are looking for: \"%s\"", r.URL.Path)))
	}
}

func (cc *calculatorController) add(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractVariables(r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	result.Action = "add"
	result.X = x
	result.Y = y
	result.Answer = models.Add(x, y)

	encodeResponseAsJSON(result, w)

	result.Cached = true
	models.CacheSet(result.Key, result)
	return
}

func (cc *calculatorController) sub(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractVariables(r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	result.Action = "subtract"
	result.X = x
	result.Y = y
	result.Answer = models.Subtract(x, y)

	encodeResponseAsJSON(result, w)

	result.Cached = true
	models.CacheSet(result.Key, result)
	return
}

func (cc *calculatorController) mul(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractVariables(r.URL.Query())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	result.Action = "multiply"
	result.X = x
	result.Y = y
	result.Answer = models.Multiply(x, y)

	encodeResponseAsJSON(result, w)

	result.Cached = true
	models.CacheSet(result.Key, result)
	return
}

func (cc *calculatorController) div(w http.ResponseWriter, r *http.Request) {
	x, y, err := extractVariables(r.URL.Query())

	var answer float64
	if err == nil {
		answer, err = models.Divide(x, y)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	result.Action = "divide"
	result.X = x
	result.Y = y
	result.Answer = answer

	encodeResponseAsJSON(result, w)

	result.Cached = true
	models.CacheSet(result.Key, result)
	return

}

func extractVariables(p map[string][]string) (float64, float64, error) {
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

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func newCalculatorController() *calculatorController {
	return &calculatorController{}
}
