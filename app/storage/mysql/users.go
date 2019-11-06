package mysql

import "github.com/dpogorelovsky/go-basic-web-app/app/models"

func (s *storage) GetUserList(limit, offset int) ([]*models.User, error) {
	resp := []*models.User{
		&models.User{
			ID:     1,
			Email:  "some@email.com",
			Name:   "John",
			Age:    25,
			CityID: 1,
		},
		&models.User{
			ID:     2,
			Email:  "some2@email.com",
			Name:   "Eric",
			Age:    17,
			CityID: 2,
		},
	}

	return resp, nil
}
