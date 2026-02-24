package db

import (
	"sync"

	"cloud.google.com/go/firestore"
)

type Database struct {
	Client *firestore.Client
}

var (
	instance *Database
	once     sync.Once
	initErr  error
)
