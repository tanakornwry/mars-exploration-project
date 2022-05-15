package roverpresenters

import "github.com/tanakornwry/mars-exploration-project/entities"

type roverPresentation struct {
}

type RoverPresenters interface {
	RoverResponse(rr []entities.CurrentDP) []entities.CurrentDP
}

func NewRoverPresenters() RoverPresenters {
	return &roverPresentation{}
}

// If the rover has been change the output to outside, update here!!
func (rp *roverPresentation) RoverResponse(rr []entities.CurrentDP) []entities.CurrentDP {
	return rr
}
