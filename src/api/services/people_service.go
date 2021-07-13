package services

import "FromGithub/first-step-go/src/api/clients/db_client"

type PeopleService interface {
	AddPerson()
}

type peopleService struct {
	db db_client.DB
}

func NewPeopleService(db db_client.DB) PeopleService {
	return &peopleService{
		db: db,
	}
}

func (pe *peopleService) AddPerson() {
	stms := db_client.NewStatement("addPerson", "Insert into people (`doc_number`, `name`, `last_name`) VALUES ('12345', 'Emi', 'Bent');", nil )
	_, err := pe.db.Query(&stms)
	if err != nil {
		println("something went wrong")
	}
}