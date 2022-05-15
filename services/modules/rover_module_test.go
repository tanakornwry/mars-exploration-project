package rovermodules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanakornwry/mars-exploration-project/entities"
)

func TestInitialDP(t *testing.T) {
	var roverModules = NewRoverModules()

	c := entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 0}

	assert.Equal(t, c, roverModules.InitialDP())
}

func TestRotate(t *testing.T) {
	var roverModules = NewRoverModules()

	current := roverModules.InitialDP()
	L, R := "L", "R"

	var when string
	var expected entities.CurrentDP

	t.Run("Turning right checking", func(t *testing.T) {
		when = R

		// N
		current = roverModules.Rotate(current, when)
		// E
		expected = entities.CurrentDP{Degree: 0, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// E
		current = roverModules.Rotate(current, when)
		// S
		expected = entities.CurrentDP{Degree: 270, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// S
		current = roverModules.Rotate(current, when)
		// W
		expected = entities.CurrentDP{Degree: 180, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// W
		current = roverModules.Rotate(current, when)
		// N
		expected = entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})

	t.Run("Turning right checking", func(t *testing.T) {
		current = roverModules.InitialDP()
		when = L

		// N
		current = roverModules.Rotate(current, when)
		// W
		expected = entities.CurrentDP{Degree: 180, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// W
		current = roverModules.Rotate(current, when)
		// S
		expected = entities.CurrentDP{Degree: 270, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// S
		current = roverModules.Rotate(current, when)
		// E
		expected = entities.CurrentDP{Degree: 0, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// E
		current = roverModules.Rotate(current, when)
		// N
		expected = entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})
}

func TestMove(t *testing.T) {
	var roverModules = NewRoverModules()

	current := roverModules.InitialDP()
	F := "F"

	var when = F
	var unit = 1
	var expected entities.CurrentDP

	t.Run("Move to positive area", func(t *testing.T) {
		// Move to quadrant Y
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 1}
		assert.Equal(t, expected, current)

		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 2}
		assert.Equal(t, expected, current)

		current.Degree = 270
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 270, Position_X: 0, Position_Y: 1}
		assert.Equal(t, expected, current)

		// Move to quadrant X
		current.Degree = 0
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 0, Position_X: 1, Position_Y: 1}
		assert.Equal(t, expected, current)

		current.Degree = 0
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 0, Position_X: 2, Position_Y: 1}
		assert.Equal(t, expected, current)

		current.Degree = 180
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 180, Position_X: 1, Position_Y: 1}
		assert.Equal(t, expected, current)
	})

	t.Run("Move to negative area", func(t *testing.T) {
		// Move to quadrant Y
		current = roverModules.InitialDP()

		current.Degree = 270
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 270, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// Move to quadrant X
		current.Degree = 180
		current = roverModules.Move(current, when, unit)
		expected = entities.CurrentDP{Degree: 180, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})
}
