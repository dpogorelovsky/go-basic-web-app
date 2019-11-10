package models

// User - users model
type User struct {
	ID     int    `json:"id" db:"id"`
	Email  string `json:"email" db:"email"`
	Name   string `json:"name" db:"name"`
	Age    int    `json:"age" db:"age"`
	CityID int    `json:"cityId,omitempty" db:"city_id"`
	City   string `json:"city,omitempty" db:"city"` // city relation, to show city name, not ID
}
