package ngorm

import "time"

// Model this struct is provided by default to create an id,createdAt and updatedAt
type Table struct {
	ID        uint `ngorm:"pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
