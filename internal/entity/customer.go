package entity

import "time"

type Customer struct {
	GUID      string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
