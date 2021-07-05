package domains

const (
	Flat = iota
	House
	Garage
)

type Construction struct {
	//Address Address
	ID                int
	TypeOfContruction string
}

type typeOfContruction struct {
	typeOfContruction string
}
