package models

import "errors"

// Add function of a calcultaor
func Add(a, b float64) float64 {
	return a + b
}

// Subtract function of a calcultaor
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply function of a calcultaor
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide function of a calcultaor
func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}

	return a / b, nil
}
