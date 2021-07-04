package app

import "github.com/gin-gonic/gin"

func RunServer(port string) error {
	application := newApplication()
	router := gin.Default()
	urlMapping(router, application)
	return router.Run(":" + port)
}
