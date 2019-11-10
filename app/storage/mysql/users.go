package mysql

import (
	"database/sql"

	"github.com/dpogorelovsky/go-basic-web-app/app/models"
)

// GetUserList - list of users with city as its name, not city_id
func (s *storage) GetUserList(limit, offset int) ([]*models.User, error) {
	resp := []*models.User{}
	err := s.driver.Select(&resp, "SELECT u.id, u.name, u.email, u.age, c.name as city FROM users u LEFT JOIN cities c on u.city_id = c.id LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return resp, nil
}

// GetUserListRaw - pure list of users, no possible relations
func (s *storage) GetUserListRaw(limit, offset int) ([]*models.User, error) {
	resp := []*models.User{}
	err := s.driver.Select(&resp, "SELECT * FROM users LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}
	return resp, nil
}

// CreateUser - dumb user insertion, no validation.
func (s *storage) CreateUser(u *models.User) (err error) {
	_, err = s.driver.Exec("INSERT into users (email, name, age, city_id) VALUES (?, ?, ?, ?)",
		u.Email, u.Name, u.Age, u.CityID)

	return
}

// UpdateUser - override whole user entry
func (s *storage) UpdateUser(u *models.User) (err error) {
	_, err = s.driver.Exec("UPDATE users set email = ?, name = ?, age = ?, city_id = ? WHERE id = ?",
		u.Email, u.Name, u.Age, u.CityID, u.ID)

	return
}

// DeleteUser - removes user entry
func (s *storage) DeleteUser(ID int) (err error) {
	_, err = s.driver.Exec("DELETE FROM users WHERE id = ?", ID)

	return
}

// GetUser - fetching user by ID
func (s *storage) GetUser(ID int) (*models.User, error) {
	resp := []*models.User{}
	// since id is primary key - no need to set limit=1
	err := s.driver.Select(&resp, "SELECT * FROM users WHERE id = ?", ID)
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
