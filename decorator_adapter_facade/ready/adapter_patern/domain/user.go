package domain

import "time"

type (
	User struct {
		ID          int
		Name        string
		Email       string
		ProjectId   string
		PhoneNumber string
		Country     string
		BirthDate   time.Time
		DateCreated time.Time
		// Other fields
	}
)
