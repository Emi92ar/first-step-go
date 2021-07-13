package controllers

import (
	"FromGithub/first-step-go/src/api/mocks"
	"FromGithub/first-step-go/src/api/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewConstructionController(t *testing.T) {
	t.Run(t.Name(), func(t *testing.T) {
		expected := &constructionController{}

		got := NewConstructionController(nil)
		assert.EqualValues(t, got, expected)
	})
}

func Test_constructionController_GetAllConstructions(t *testing.T) {
	t.Run(t.Name(), func(t *testing.T) {
		response := httptest.NewRecorder()
		c := mocks.GetMockedContext(http.MethodGet, "/all_constructions", nil, response)
		constructionService := services.NewConstructionService() //make it global
		controller := NewConstructionController(constructionService)
		controller.GetAllConstructions(c)


		assert.EqualValues(t, http.StatusOK, response.Code)
		assert.EqualValues(t, "[{\"id\":1,\"type_of_contruction\":\"House\",\"address\":{\"number\":1234,\"street\":\"california Street\"}},{\"id\":2,\"type_of_contruction\":\"Flat\",\"address\":{\"number\":9876,\"street\":\"chicago Street\"}}]", response.Body.String())
	})
}
