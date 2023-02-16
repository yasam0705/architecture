package entity

import "time"

type File struct {
	Guid      string
	UserId    string
	FileName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
