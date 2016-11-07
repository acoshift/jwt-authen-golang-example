package model

import "time"

// Token model
type Token struct {
	Base
	Token        string    `json:"-"`
	UserID       int64     `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
	LastAccessAt time.Time `json:"lastAccessAt"`
}

// Stamp current time to token
func (x *Token) Stamp() {
	x.LastAccessAt = time.Now()
	if x.CreatedAt.IsZero() {
		x.CreatedAt = x.LastAccessAt
	}
}
