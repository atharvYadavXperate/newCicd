package app

import db "github.com/atharvYadavXperate/newCicd/kapad-app/firestore"

type App struct {
	ProjectId string
	Database  *db.Database
}
