package api

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// Config API type
type Config struct {
	ServiceAccountJSON []byte
	ProjectID          string
}

var (
	client *datastore.Client
)

// Init api
func Init(cfg Config) error {
	gconf, err := google.JWTConfigFromJSON(cfg.ServiceAccountJSON, datastore.ScopeDatastore)
	if err != nil {
		return err
	}

	ctx := context.Background()

	client, err = datastore.NewClient(ctx, cfg.ProjectID, option.WithTokenSource(gconf.TokenSource(ctx)))
	return err
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
