package app

import (
	"context"
	"os"

	db "github.com/atharvYadavXperate/newCicd/kapad-app/firestore"
)

func (a *App) Init() {
	a.ProjectId = os.Getenv("PROJECT_ID")
}

func (a *App) GetProjectId() string {
	return a.ProjectId
}

func (a *App) InitDatabase() error {
	database, err := db.NewDatabase(context.Background(), a.ProjectId)
	if err != nil {
		return err
	}
	a.Database = database
	return nil
}
