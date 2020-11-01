package model

import "time"

type User struct {
	ID          int
	Name        string
	BirthDate   *time.Time
	BaptismDate *time.Time
	SidiDate    *time.Time
	CreatedAt   time.Time
}
