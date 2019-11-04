package handlers

import (
	"net/http"

	"github.com/dpogorelovsky/go-basic-web-app/app/models"
	R "github.com/dpogorelovsky/go-basic-web-app/app/utils/responder"
)

// User - users handler/controller
type User struct{}

// List - returning list of users
func (h *User) List(w http.ResponseWriter, r *http.Request) {

	resp := []models.User{
		models.User{
			ID:     1,
			Email:  "some@email.com",
			Name:   "John",
			Age:    25,
			CityID: 1,
		},
		models.User{
			ID:     2,
			Email:  "some2@email.com",
			Name:   "Eric",
			Age:    17,
			CityID: 2,
		},
	}

	R.JSON200(w, resp)
}
