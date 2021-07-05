package domains

import "FromGithub/first-step-go/src/api/enums"

type Construction struct {
	ID                int                          `json:"id"`
	TypeOfContruction enums.TypeOfConstructionEnum `json:"type_of_contruction"`
	Address           Address                      `json:"address"`
}
