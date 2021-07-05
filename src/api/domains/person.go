package domains

const (
	Owner = iota
	Tenant
)

type Person struct {
	Name           string `json:"Name"`
	Age            int    `json:"Age"`
	DocumentNumber int    `json:"document_number"`
}
