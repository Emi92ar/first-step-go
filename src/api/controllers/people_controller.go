package controllers

import (
	"FromGithub/first-step-go/src/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PeopleController interface {
	AddPersonToDB(c *gin.Context)
}

type peopleController struct {
	PeopleService services.PeopleService
}

func NewPeopleController(peopleService services.PeopleService) PeopleController {
	return &peopleController{
		PeopleService: peopleService,

	}
}

func (cc *peopleController) AddPersonToDB(c *gin.Context) {
	//ctx := c.Request.Context()

	cc.PeopleService.AddPerson()

	c.JSON(http.StatusOK, "Person added")
}
