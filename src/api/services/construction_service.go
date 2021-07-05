package services

import (
	"FromGithub/first-step-go/src/api/domains"
	"FromGithub/first-step-go/src/api/enums"
)

type ConstructionService interface {
	GetAllConstructions() ([]domains.Construction, error)
}

//here comes the data base or kvs or dao
type constructionService struct {
}

func NewConstructionService() ConstructionService{
	return &constructionService{
	}
}

func (cs *constructionService) GetAllConstructions() ([]domains.Construction, error){
	 const1 := domains.Construction{
	 	Address: domains.Address{
	 		Number: 1234, Street: "california Street"},
		 ID:                1,
		 TypeOfContruction: enums.House,
	 }

	 const2 := domains.Construction{
		 Address: domains.Address{
		 	Number: 9876, Street: "chicago Street"},
	 	ID:                2,
	 	TypeOfContruction: enums.Flat,
	 }

	 listConstructions := []domains.Construction{
		 const1, const2,
	 }

	 return listConstructions, nil
}