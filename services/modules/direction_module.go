package rovermodules

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
)

func ValidRoute(c entities.CurrentDP, m string, u int) (bool, entities.CurrentDP) {

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

	// Extended skill to the Rover that it can half rotate
	switch {
	case c.Degree == 45: // NE
		c.Position_X += u
		c.Position_Y += u
	case c.Degree == 135: // NW
		c.Position_X -= u
		c.Position_Y += u
	case c.Degree == 225: // SW
		c.Position_X -= u
		c.Position_Y -= u
	case c.Degree == 315: // SE
		c.Position_X += u
		c.Position_Y -= u
	}

	var canMove bool
	switch {
	case c.Position_X >= 0 && c.Position_X <= Scope && c.Position_Y >= 0 && c.Position_Y <= Scope:
		canMove = true
	default:
		canMove = false
	}

	return canMove, c
}
