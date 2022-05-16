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

func TestHalfRotate(t *testing.T) {
	var roverModules = NewRoverModules()

	current := roverModules.InitialDP()
	var expected entities.CurrentDP

	// Half right
	current = roverModules.Rotate(current, "HR")
	expected = entities.CurrentDP{Degree: 45, Position_X: 0, Position_Y: 0}
	assert.Equal(t, expected, current)

	// Half right and convert full circle
	current = roverModules.Rotate(current, "R")
	expected = entities.CurrentDP{Degree: 315, Position_X: 0, Position_Y: 0}
	assert.Equal(t, expected, current)

	current = roverModules.InitialDP()
	// Half Left
	current = roverModules.Rotate(current, "HL")
	expected = entities.CurrentDP{Degree: 135, Position_X: 0, Position_Y: 0}
	assert.Equal(t, expected, current)

	// Half right and convert full circle
	current.Degree = 270
	current = roverModules.Rotate(current, "L")
	expected = entities.CurrentDP{Degree: 0, Position_X: 0, Position_Y: 0}
	assert.Equal(t, expected, current)

	current.Degree = 315
	current = roverModules.Rotate(current, "L")
	expected = entities.CurrentDP{Degree: 45, Position_X: 0, Position_Y: 0}
	assert.Equal(t, expected, current)
}

func TestMove(t *testing.T) {
	var roverModules = NewRoverModules()

	current := roverModules.InitialDP()
	scope := 10
	F := "F"

	var when = F
	var unit = 1
	var expected entities.CurrentDP

	t.Run("Move to positive area", func(t *testing.T) {
		// Move to quadrant Y
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 1}
		assert.Equal(t, expected, current)

		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 90, Position_X: 0, Position_Y: 2}
		assert.Equal(t, expected, current)

		current.Degree = 270
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 270, Position_X: 0, Position_Y: 1}
		assert.Equal(t, expected, current)

		// Move to quadrant X
		current.Degree = 0
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 0, Position_X: 1, Position_Y: 1}
		assert.Equal(t, expected, current)

		current.Degree = 0
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 0, Position_X: 2, Position_Y: 1}
		assert.Equal(t, expected, current)

		current.Degree = 180
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 180, Position_X: 1, Position_Y: 1}
		assert.Equal(t, expected, current)
	})

	t.Run("Move to negative area", func(t *testing.T) {
		// Move to quadrant Y
		current = roverModules.InitialDP()

		current.Degree = 270
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 270, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// Move to quadrant X
		current.Degree = 180
		current = roverModules.Move(scope, current, when, unit)
		expected = entities.CurrentDP{Degree: 180, Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})
}

func TestBackwardMoving(t *testing.T) {
	var roverModules = NewRoverModules()

	current := roverModules.InitialDP()
	scope := 10
	F := "F"
	B := "B"

	var when = F
	var unit = 1
	current.Degree = 90
	current.Position_X = 1
	current.Position_Y = 1

	when = B
	current = roverModules.Move(scope, current, when, unit)
	expected := entities.CurrentDP{Degree: 90, Position_X: 1, Position_Y: 0}
	assert.Equal(t, expected, current)
}
