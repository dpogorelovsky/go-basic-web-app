package handlers

import (
	"net/http"

	"github.com/dpogorelovsky/go-basic-web-app/app/models"
	R "github.com/dpogorelovsky/go-basic-web-app/app/utils/responder"
)

// City - cities handler
type City struct{}

// List - respond with list of cities
func (h *City) List(w http.ResponseWriter, r *http.Request) {
	resp := []models.City{
		models.City{
			ID:   1,
			Name: "New York",
		},
		models.City{
			ID:   2,
			Name: "Dublin",
		},
	}

	R.JSON200(w, resp)
}
