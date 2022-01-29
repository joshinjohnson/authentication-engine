package models

import "time"

type UserCredential struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
}