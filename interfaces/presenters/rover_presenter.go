package presenters

import (
	"fmt"

	"github.com/tanakornwry/mars-exploration-project/entities"
)

type interfacePresenters struct {
}

type InterfacePresenters interface {
	Response(rr []entities.CurrentDP) []string
}

func NewInterfaceRes() InterfacePresenters {
	return &interfacePresenters{}
}

func (ip *interfacePresenters) Response(rr []entities.CurrentDP) []string {
	var resp []string
	for _, v := range rr {
		resp = append(resp, fmt.Sprintf("%s:%d,%d", convertDegree(v.Degree), v.Position_X, v.Position_Y))
	}

	return resp
}

// Used for converting degree to direction
func convertDegree(dg int) string {
	mapDirection := map[int]string{
		0:   "E",
		45:  "NE",
		90:  "N",
		135: "NW",
		180: "W",
		225: "SW",
		270: "S",
		315: "SE",
	}

	return mapDirection[dg]
}
