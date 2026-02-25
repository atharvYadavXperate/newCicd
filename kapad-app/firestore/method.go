package db

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	customerror "github.com/atharvYadavXperate/newCicd/kapad-app/domain/errors"
	"github.com/atharvYadavXperate/newCicd/kapad-app/domain/helpers"
)

func (db *Database) GetById(ctx context.Context, collection, docID string, result interface{}) error {
	if db.Client == nil {
		return customerror.ErrDatabaseConnectionFailed
	}

	docSnap, err := db.Client.Collection(collection).Doc(docID).Get(ctx)
	if err != nil {
		return err
	}

	return docSnap.DataTo(result)
}

func (db *Database) FindOne(ctx context.Context, collection, field string, value any, result interface{}) error {
	iter := db.Client.Collection(collection).Where(field, "==", value).Limit(1).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		return err
	}
	return doc.DataTo(result)
}

func (db *Database) Create(ctx context.Context, collection string, data interface{}) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	if db.Client == nil {
		return nil, nil, customerror.ErrDatabaseConnectionFailed
	}
	return db.Client.Collection(collection).Add(ctx, data)
}

func (db *Database) CreateWithCustomId(ctx context.Context, collection string, docID string, data interface{}) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	if db.Client == nil {
		return nil, nil, customerror.ErrDatabaseConnectionFailed
	}
	docID = helpers.HashValueDeterministic(docID)
	docRef := db.Client.Collection(collection).Doc(docID)
	writerResult, err := docRef.Create(ctx, data)
	return docRef, writerResult, err
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
