package rovermodules

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
)

type roverModules struct {
}

type RoverModules interface {
	InitialDP() entities.CurrentDP
	Rotate(c entities.CurrentDP, d string) entities.CurrentDP
	Move(c entities.CurrentDP, m string, u int) entities.CurrentDP
}

func NewRoverModules() RoverModules {
	return &roverModules{}
}

// Hardcode here first to can be testable and waiting for implement the interactor
var Scope = 10

func (r *roverModules) InitialDP() entities.CurrentDP {
	initDP := entities.CurrentDP{
		Degree:     90,
		Position_X: 0,
		Position_Y: 0,
	}
	return initDP
}

func (r *roverModules) Rotate(c entities.CurrentDP, d string) entities.CurrentDP {

	// Set to revert the circle
	if d == "R" && c.Degree == 0 {
		c.Degree = 360
	}

	// Calculate the next degree follow by the rotate instruction rules
	nextDegree := c.Degree + entities.Instruction.RotateInstruction[d]

	// Reset degree if completed the circle
	if d == "L" && nextDegree == 360 {
		nextDegree = 0
	}

	c.Degree = nextDegree

	return c
}

func (r *roverModules) Move(c entities.CurrentDP, m string, u int) entities.CurrentDP {

	// Calculate block unit when backward moving
	if m == "B" {
		u *= -1
	}

	// If the rover goes to out of scope(including a negative area) will maintain the route
	canMove, newDP := ValidRoute(c, m, u)

	if canMove {
		c = newDP
	}

	return c
}
