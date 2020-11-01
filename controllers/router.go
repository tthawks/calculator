package controllers

import "net/http"

// RegisterControllers registers controllers for routes
func RegisterControllers() {
	cc := newCalculatorController()

	http.Handle("/", *cc)
}
