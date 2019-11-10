package mysql

import (
	"database/sql"

	"github.com/dpogorelovsky/go-basic-web-app/app/models"
)

// GetCityList - cities list
func (s *storage) GetCityList(limit, offset int) ([]*models.City, error) {
	resp := []*models.City{}
	err := s.driver.Select(&resp, "SELECT * FROM cities LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return resp, nil
}

// CreateCity - city insertion
func (s *storage) CreateCity(c *models.City) (err error) {
	_, err = s.driver.Exec("INSERT into cities (name) VALUES (?)",
		c.Name)

	return
}

// UpdateCity - override whole city entry
func (s *storage) UpdateCity(c *models.City) (err error) {
	_, err = s.driver.Exec("UPDATE city set name = ? WHERE id = ?",
		c.Name, c.ID)

	return
}

// DeleteCity - removes city entry
func (s *storage) DeleteCity(ID int) (err error) {
	_, err = s.driver.Exec("DELETE FROM cities WHERE id = ?", ID)

	return
}

// GetCity - fetching city by ID
func (s *storage) GetCity(ID int) (*models.City, error) {
	resp := []*models.City{}
	// since id is primary key - no need to set limit=1
	err := s.driver.Select(&resp, "SELECT * FROM cities WHERE id = ?", ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	if len(resp) < 1 {
		return nil, nil
	}

	return resp[0], nil
}
