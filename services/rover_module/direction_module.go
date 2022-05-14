package rovermodule

import "github.com/tanakornwry/mars-exploration-project/entities"

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

	var canMove bool
	switch {
	case c.Position_X >= 0 && c.Position_X <= Scope && c.Position_Y >= 0 && c.Position_Y <= Scope:
		canMove = true
	default:
		canMove = false
	}

	return canMove, c
}
