package access

import "time"

type credential struct {
	id int64
	username string
	passwordHash string
	email string
	lastUpdatedTime time.Time
	createdTime time.Time
}

type user struct {
	userID int64
	firstName string
	lastName string
	dateOfBirth time.Time
	lastUpdatedTime time.Time
	createdTime time.Time
}

type credentialToUserLookup struct {
	credentialID int64
	userID int64
}