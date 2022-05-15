package rovermodule

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
)

type RoverModules interface {
	InitialDP() entities.CommandConf
	Rotate(entities.CommandConf, string) entities.CurrentDP
	Move(entities.CommandConf, string, int) entities.CurrentDP
}

// Hardcode here first to can be testable and waiting for implement the interactor
var Scope = 10

func InitialDP() entities.CurrentDP {
	initDP := entities.CurrentDP{
		Degree:     90,
		Position_X: 0,
		Position_Y: 0,
	}
	return initDP
}

func Rotate(c entities.CurrentDP, d string) entities.CurrentDP {

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

func Move(c entities.CurrentDP, m string, u int) entities.CurrentDP {

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
