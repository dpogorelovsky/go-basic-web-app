package models

// User - users model
type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	CityID int    `json:"cityId"`
}
