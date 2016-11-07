package model

// User model
type User struct {
	Base
	HasPassword
	HasTimestamp
	Username string `json:"username"`
}
