package interactors

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
	rovermodules "github.com/tanakornwry/mars-exploration-project/services/modules"
)

type roverInteractors struct {
	RoverModules rovermodules.RoverModules
}

type RoverInteractors interface {
	StartRover(c entities.CommandConf) []entities.CurrentDP
}

func NewRoverInteractors(m rovermodules.RoverModules) RoverInteractors {
	return &roverInteractors{m}
}

func (ui *roverInteractors) StartRover(c entities.CommandConf) []entities.CurrentDP {

	// Mock command
	c.SizeMap = 20
	c.Command = []map[string]int{
		{"R": 1},
		{"F": 1},
		{"L": 1},
		{"F": 1},
		{"L": 1},
		{"L": 1},
		{"F": 1},
		{"R": 1},
	}

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
			case act == "L" || act == "R":
				rAct = roverModules.Rotate(currentDP, act)
			case act == "F":
				rAct = roverModules.Move(currentDP, act, unit)
			}

			// Set the new rover direction and position
			currentDP = rAct

			// Keep the rover path
			roverPath = append(roverPath, rAct)
		}
	}

	return roverPath
}
