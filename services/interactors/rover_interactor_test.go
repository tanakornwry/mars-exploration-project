package interactors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanakornwry/mars-exploration-project/entities"
	rovermodules "github.com/tanakornwry/mars-exploration-project/services/modules"
)

func TestStartRover(t *testing.T) {
	roverModules := rovermodules.NewRoverModules()
	roverInteractors := NewRoverInteractors(roverModules)

	// Mock Data and Result base on sample Input/Output
	var c entities.CommandConf
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

	have := roverInteractors.StartRover(c)
	want := []entities.CurrentDP{
		{Degree: 90, Position_X: 0, Position_Y: 0},
		{Degree: 0, Position_X: 0, Position_Y: 0},
		{Degree: 0, Position_X: 1, Position_Y: 0},
		{Degree: 90, Position_X: 1, Position_Y: 0},
		{Degree: 90, Position_X: 1, Position_Y: 1},
		{Degree: 180, Position_X: 1, Position_Y: 1},
		{Degree: 270, Position_X: 1, Position_Y: 1},
		{Degree: 270, Position_X: 1, Position_Y: 0},
		{Degree: 180, Position_X: 1, Position_Y: 0},
	}

	assert.Equal(t, want, have)
}
