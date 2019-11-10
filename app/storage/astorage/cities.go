package astorage

import "github.com/dpogorelovsky/go-basic-web-app/app/models"

// CityStorage - abstract city storage
type CityStorage interface {
	GetCityList(limit, offset int) ([]*models.City, error)
	CreateCity(*models.City) error
	GetCity(ID int) (*models.City, error)
	UpdateCity(*models.City) error
	DeleteCity(ID int) error
}
