package model

import (
	"cloud.google.com/go/datastore"
)

// Base type provides datastore-based model
type Base struct {
	key *datastore.Key
	ID  int64 `datastore:"-" json:"id"`
}

// Key return datastore key or nil
func (x *Base) Key() *datastore.Key {
	return x.key
}

// SetKey sets key and id to new given key
func (x *Base) SetKey(key *datastore.Key) {
	x.key = key
	x.ID = key.ID()
}
