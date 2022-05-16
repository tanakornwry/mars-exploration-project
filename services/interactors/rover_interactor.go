package interactors

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
	rovermodules "github.com/tanakornwry/mars-exploration-project/services/modules"
	roverpresenters "github.com/tanakornwry/mars-exploration-project/services/presenters"
)

type roverInteractors struct {
	RoverModules    rovermodules.RoverModules
	RoverPresenters roverpresenters.RoverPresenters
}

type RoverInteractors interface {
	StartRover(c entities.CommandConf) []entities.CurrentDP
}

func NewRoverInteractors(m rovermodules.RoverModules, p roverpresenters.RoverPresenters) RoverInteractors {
	return &roverInteractors{m, p}
}

func (ui *roverInteractors) StartRover(c entities.CommandConf) []entities.CurrentDP {

	var roverModules = rovermodules.NewRoverModules()

	// Set the rover to start direction and position (90(N):0,0)
	var currentDP = roverModules.InitialDP()

	// Keep every steps the rover has moved
	var roverPath = []entities.CurrentDP{
		currentDP,
	}

	for _, v := range c.Command {
		for act, unit := range v {
			var rAct entities.CurrentDP

			// Call the rover module to execute command
			switch {
			case act == "L" || act == "HL" || act == "R" || act == "HR":
				rAct = roverModules.Rotate(currentDP, act)
			case act == "F" || act == "B":
				rAct = roverModules.Move(c.SizeMap, currentDP, act, unit)
			}

			// Set the new rover direction and position
			currentDP = rAct

			// Keep the rover path
			roverPath = append(roverPath, rAct)
		}
	}

	return ui.RoverPresenters.RoverResponse(roverPath)
}
