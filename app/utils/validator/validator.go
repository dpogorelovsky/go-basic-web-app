package validator

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/dpogorelovsky/go-basic-web-app/app/models"
)

// UserCreate - takes json, validates it and returns user
func UserCreate(data []byte) (*models.User, error) {
	u := models.User{}

	err := json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}

	// we just check for '@' in email since we are not email address provider
	if !strings.Contains(u.Email, "@") || u.CityID == 0 {
		err = errors.New("Malformed data")
		return nil, err
	}

	return &u, nil
}

// CityCreate - takes json, validates it and returns city instance
func CityCreate(data []byte) (*models.City, error) {
	c := models.City{}

	err := json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	if c.Name == "" {
		err = errors.New("Malformed data")
		return nil, err
	}

	return &c, nil
}
