package controllers

import (
	"FromGithub/first-step-go/src/api/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHealthController(t *testing.T) {
	t.Run(t.Name(), func(t *testing.T) {
		expected := &healtController{}
		got := NewHealthController()
		assert.EqualValues(t, got, expected)
	})
}

func Test_healtController_Ping(t *testing.T) {

	t.Run(t.Name(), func(t *testing.T) {
		response := httptest.NewRecorder()
		c := mocks.GetMockedContext(http.MethodGet, "/ping", nil, response)
		controller := NewHealthController()
		controller.Ping(c)

		assert.EqualValues(t, http.StatusOK, response.Code)
		assert.EqualValues(t, "{\"message\":\"pong\"}", response.Body.String())
	})

}
