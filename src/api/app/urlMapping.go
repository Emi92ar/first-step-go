package app

import "github.com/gin-gonic/gin"

func urlMapping(r *gin.Engine, app *application) {
	r.GET("/ping", app.healthController.Ping)
	r.GET("/all_constructions", app.constructionController.GetAllConstructions)
	r.POST("/add_person", app.peopleController.AddPersonToDB)
}