package entity

import "time"

type User struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
	FirstName string    `json:"first_name" db:"first_name"  binding:"required"`
	LastName  string    `json:"last_name" db:"last_name"  binding:"required"`
	City      string    `json:"city" db:"city"  binding:"required"`
	BirthDate time.Time `json:"birth_date" db:"birth_date"  binding:"required"`
}

type CreateUserRequest struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Password  string    `json:"password"  db:"password"  binding:"required"`
	FirstName string    `json:"first_name" db:"first_name"  binding:"required"`
	LastName  string    `json:"last_name" db:"last_name"  binding:"required"`
	City      string    `json:"city" db:"city"  binding:"required"`
	BirthDate time.Time `json:"birth_date" db:"birth_date"  binding:"required"`
}
