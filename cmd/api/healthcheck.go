package main

import (
	"net/http"
)

// Healthcheck godoc
//
//	@Summary	Gives a brief healthcheck details about the applications
//	@Tags		healthcheck
//	@Router		/healthcheck [get]
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
