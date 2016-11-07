package model

import "time"

// HasTimestamp base model
type HasTimestamp struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Stamp current time to model
func (x *HasTimestamp) Stamp() {
	x.UpdatedAt = time.Now()
	if x.CreatedAt.IsZero() {
		x.CreatedAt = x.UpdatedAt
	}
}
