package services

import "FromGithub/first-step-go/src/api/domains"

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
		 ID:                1,
		 TypeOfContruction: "garage",
	 }

	 const2 := domains.Construction{
	 	ID: 2,
	 	TypeOfContruction: "house",
	 }

	 listConstructions := []domains.Construction{
		 const1, const2,
	 }
	 return listConstructions, nil
}