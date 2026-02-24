package db

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	customerror "github.com/atharvYadavXperate/newCicd/kapad-app/domain/errors"
)

func (db *Database) Create(ctx context.Context, collection string, data interface{}) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	if db.Client == nil {
		return nil, nil, customerror.ErrDatabaseConnectionFailed
	}
	return db.Client.Collection(collection).Add(ctx, data)
}

func (db *Database) Update(ctx context.Context, collection string, docId string, fieldsToUpdate map[string]interface{}) (*firestore.WriteResult, error) {
	if db.Client == nil {
		return nil, customerror.ErrDatabaseConnectionFailed
	}
	fieldsToUpdate["updatedAt"] = time.Now()

	var updates []firestore.Update
	for k, v := range fieldsToUpdate {
		updates = append(updates, firestore.Update{Path: k, Value: v})
	}
	return db.Client.Collection(collection).Doc(docId).Update(ctx, updates)
}

func (db *Database) Delete(ctx context.Context, collection string, docId string) (*firestore.WriteResult, error) {
	if db.Client == nil {
		return nil, customerror.ErrDatabaseConnectionFailed
	}
	return db.Client.Collection(collection).Doc(docId).Delete(ctx)
}
