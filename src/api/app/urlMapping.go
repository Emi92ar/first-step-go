package app

import "github.com/gin-gonic/gin"

func urlMapping(r *gin.Engine, app *application) {
	r.GET("/ping", app.healthController.Ping)
}