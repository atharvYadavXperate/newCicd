package db

import (
	"context"

	"cloud.google.com/go/firestore"
	custiomeerror "github.com/atharvYadavXperate/newCicd/kapad-app/domain/errors"
)

func NewDatabase(ctx context.Context, projectId string) (*Database, error) {
	once.Do(func() {
		c, err := firestore.NewClient(ctx, projectId)
		if err != nil {
			initErr = custiomeerror.ErrDatabaseConnectionFailed
			return
		}
		instance = &Database{
			Client: c,
		}
	})
	if initErr != nil {
		return nil, initErr
	}
	return instance, nil
}

func (db *Database) Close() error {
	if db.Close() != nil {
		return db.Client.Close()
	}
	return nil
}
