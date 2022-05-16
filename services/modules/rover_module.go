package rovermodules

import (
	"math"

	"github.com/tanakornwry/mars-exploration-project/entities"
)

type roverModules struct {
}

type RoverModules interface {
	InitialDP() entities.CurrentDP
	Rotate(c entities.CurrentDP, d string) entities.CurrentDP
	Move(s int, c entities.CurrentDP, m string, u int) entities.CurrentDP
}

func NewRoverModules() RoverModules {
	return &roverModules{}
}

func (r *roverModules) InitialDP() entities.CurrentDP {
	initDP := entities.CurrentDP{
		Degree:     90,
		Position_X: 0,
		Position_Y: 0,
	}
	return initDP
}

func (r *roverModules) Rotate(c entities.CurrentDP, d string) entities.CurrentDP {

	// Calculate the next degree follow by the rotate instruction rules
	nextDegree := c.Degree + entities.Instruction.RotateInstruction[d]

	// Reset degree if completed the circle
	switch {
	case (d == "R" || d == "HR") && nextDegree < 0:
		nextDegree = 360 + nextDegree
	case (d == "L" || d == "HL") && nextDegree == 360:
		nextDegree = 0
	case (d == "L" || d == "HL") && nextDegree > 360:
		nextDegree = 360 - nextDegree
		nextDegree = int(math.Abs(float64(nextDegree)))
	}

	c.Degree = nextDegree

	return c
}

var Scope int

func (r *roverModules) Move(s int, c entities.CurrentDP, m string, u int) entities.CurrentDP {
	Scope = s

	// If the rover goes to out of scope(including a negative area) will maintain the route
	if canMove, newDP := ValidRoute(c, m, u); canMove {
		c = newDP
	}

	return c
}
