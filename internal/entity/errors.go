package entity

import "fmt"

func Codes(code int) string {
	switch code {
	default:
		return ""
	}
}

var (
	ErrNotFound = fmt.Errorf("object not found")
)
