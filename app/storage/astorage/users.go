package astorage

import "github.com/dpogorelovsky/go-basic-web-app/app/models"

// UserStorage - abstract storage for users to describe functionality
type UserStorage interface {
	GetUserList(limit, offset int) ([]*models.User, error)
}
