package app

import "os"

func (a *App) Init() {
	a.ProjectId = os.Getenv("PROJECT_ID")
}

func (a *App) GetProjectId() string {
	return a.ProjectId
}
