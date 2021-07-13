package mocks

import (
	"FromGithub/first-step-go/src/api/domains"
)

type constructionServiceMock struct {
	mockGetAllConstructions	func() ([]domains.Construction, error)
}

func (mock *constructionServiceMock) GetAllConstructions() ([]domains.Construction, error) {
	return mock.mockGetAllConstructions()
}