package app

import "FromGithub/first-step-go/src/api/controllers"

type application struct {
	healthController controllers.HealthController
}

// connect db and all dependencies
func newApplication() *application {
	return &application{
		healthController: controllers.NewHealthController(),
	}
}