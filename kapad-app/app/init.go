package app

import "os"

func (a *Application) Init() {
	a.ProjectId = os.Getenv("PROJECT_ID")
}

func (a *Application) GetProjectId() string {
	return a.ProjectId
}
