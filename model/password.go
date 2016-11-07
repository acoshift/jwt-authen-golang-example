package model

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 13

// HasPassword base type
type HasPassword struct {
	Password string `datastore:",noindex" json:"-"`
}

// SetPassword hash password then set to model
func (x *HasPassword) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}
	x.Password = string(hashed)
	return nil
}

// ComparePassword verify password and hashed
func (x *HasPassword) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(x.Password), []byte(password))
	return err == nil
}
