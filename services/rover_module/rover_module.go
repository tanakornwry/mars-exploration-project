package rovermodule

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
)

type RoverModules interface {
	InitialDP() entities.CommandConf
	Rotate(entities.CommandConf, string) entities.CurrentDP
	Move(entities.CommandConf, string, int) entities.CurrentDP
	forwardMoving(entities.CurrentDP)
	backwardMoving(entities.CurrentDP)
}

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

	switch {
	case c.Degree == 0: // Move to E direction on X axis
		c.Position_X += u
	case c.Degree == 180: // Move to W direction on X axis
		c.Position_X -= u
	case c.Degree == 90: // Move to N direction on Y axis
		c.Position_Y += u
	case c.Degree == 270: // Move to S direction on Y axis
		c.Position_Y -= u
	}

	return c
}
