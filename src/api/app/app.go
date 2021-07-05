package app

import (
	"FromGithub/first-step-go/src/api/controllers"
	"FromGithub/first-step-go/src/api/services"
)

type application struct {
	healthController       controllers.HealthController
	constructionController controllers.ConstructionController
}

// connect db and all dependencies
func newApplication() *application {
	constructionService := services.NewConstructionService()

	return &application{
		healthController:       controllers.NewHealthController(),
		constructionController: controllers.NewConstructionController(constructionService),
	}
}
