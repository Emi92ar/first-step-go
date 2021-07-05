package enums

import (
	"bytes"
	"encoding/json"
)

type TypeOfConstructionEnum int

const (
	Flat = iota
	House
	Garage
)

var toString = map[TypeOfConstructionEnum]string{
	Flat:   "Flat",
	House:  "House",
	Garage: "Garage",
}

var toID = map[string]TypeOfConstructionEnum{
	"Flat":   Flat,
	"House":  House,
	"Garage": Garage,
}

func (g TypeOfConstructionEnum) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[g])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (g *TypeOfConstructionEnum) UnmarshalJSON(b []byte) error {
	var j string
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}
	*g = toID[j]
	return nil
}