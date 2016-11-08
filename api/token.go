package api

import (
	"jwt-authen-golang-example/model"
	"time"

	"cloud.google.com/go/datastore"
)

const kindToken = "Token"

// CreateToken save new token to database
func CreateToken(token string, userID int64) error {
	ctx, cancel := getContext()
	defer cancel()

	var err error
	tk := &model.Token{
		Token:  token,
		UserID: userID,
	}
	tk.Stamp()
	key := datastore.NewIncompleteKey(ctx, kindToken, nil)
	key, err = client.Put(ctx, key, tk)
	if err != nil {
		return err
	}
	tk.SetKey(key)
	return nil
}

func getToken(token string) (*model.Token, error) {
	ctx, cancel := getContext()
	defer cancel()

	var tk model.Token
	var err error
	q := datastore.
		NewQuery(kindToken).
		Filter("Token =", token).
		Limit(1)
	key, err := client.Run(ctx, q).Next(&tk)
	if err == datastore.Done {
		// token not found
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	tk.SetKey(key)
	return &tk, nil
}

// DeleteToken delete a token from datastore
func DeleteToken(token string) error {
	tk, err := getToken(token)
	if err != nil {
		return err
	}
	ctx, cancel := getContext()
	defer cancel()
	return client.Delete(ctx, tk.Key())
}

// ValidateToken validate and update token last access timestamp
func ValidateToken(token string, userID int64, expiresInFromLastAccess time.Duration) (bool, error) {
	tk, err := getToken(token)
	if err != nil {
		return false, err
	}
	if tk == nil || tk.UserID != userID {
		return false, nil
	}
	if time.Now().After(tk.LastAccessAt.Add(expiresInFromLastAccess)) {
		// token expired
		// remove expired token from database
		go DeleteToken(token)
		return false, nil
	}
	tk.Stamp()
	go func(tk model.Token) {
		ctx, cancel := getContext()
		defer cancel()
		client.Put(ctx, tk.Key(), &tk)
	}(*tk)
	return true, nil
}
