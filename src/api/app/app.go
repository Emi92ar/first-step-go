package app

import (
	"FromGithub/first-step-go/src/api/clients/db_client"
	"FromGithub/first-step-go/src/api/controllers"
	"FromGithub/first-step-go/src/api/datasource"
	"FromGithub/first-step-go/src/api/services"
)

type application struct {
	healthController       controllers.HealthController
	constructionController controllers.ConstructionController
	peopleController controllers.PeopleController
}

// connect db and all dependencies
func newApplication() *application {
	dbConnection := datasource.GetDBConnection()
	dbClient := db_client.NewDBClient(dbConnection)
	constructionService := services.NewConstructionService()
	peopleService := services.NewPeopleService(dbClient)

	return &application{
		healthController:       controllers.NewHealthController(),
		constructionController: controllers.NewConstructionController(constructionService),
		peopleController : controllers.NewPeopleController(peopleService),

	}
}
