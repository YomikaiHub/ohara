package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": app.config.version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, req, err)
	}
}
