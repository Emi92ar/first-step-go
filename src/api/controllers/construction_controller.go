package controllers

import (
	"FromGithub/first-step-go/src/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConstructionController interface {
	GetAllConstructions(c *gin.Context)
}

type constructionController struct {
	ConstructionService services.ConstructionService
}

func NewConstructionController(constructionService services.ConstructionService) ConstructionController {
	return &constructionController{
		ConstructionService: constructionService,
	}
}

func (cc *constructionController) GetAllConstructions(c *gin.Context) {
	//ctx := c.Request.Context()

	listConstructions, err := cc.ConstructionService.GetAllConstructions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, listConstructions)
}