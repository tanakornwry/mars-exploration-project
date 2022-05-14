package rovermodule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanakornwry/mars-exploration-project/entities"
)

func TestInitialDP(t *testing.T) {
	c := entities.CurrentDP{Direction: "N", Position_X: 0, Position_Y: 0}

	assert.Equal(t, c, InitialDP())
}

func TestRotate(t *testing.T) {
	current := InitialDP()
	L, R := "L", "R"

	var when string
	var expected entities.CurrentDP

	t.Run("Turning right checking", func(t *testing.T) {
		when = R

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "W", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "S", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "E", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "N", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})

	t.Run("Turning right checking", func(t *testing.T) {
		when = L

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "E", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "S", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "W", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		current = Rotate(current, when)
		expected = entities.CurrentDP{Direction: "N", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})
}

func TestMove(t *testing.T) {
	current := InitialDP()
	F := "F"

	var when = F
	var unit = 1
	var expected entities.CurrentDP

	t.Run("Move to positive area", func(t *testing.T) {
		// Move to quadrant Y
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "N", Position_X: 0, Position_Y: 1}
		assert.Equal(t, expected, current)

		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "N", Position_X: 0, Position_Y: 2}
		assert.Equal(t, expected, current)

		current.Direction = "S"
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "S", Position_X: 0, Position_Y: 1}
		assert.Equal(t, expected, current)

		// Move to quadrant Y
		current.Direction = "W"
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "W", Position_X: 1, Position_Y: 1}
		assert.Equal(t, expected, current)

		current.Direction = "W"
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "W", Position_X: 2, Position_Y: 1}
		assert.Equal(t, expected, current)

		current.Direction = "E"
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "E", Position_X: 1, Position_Y: 1}
		assert.Equal(t, expected, current)
	})

	t.Run("Move to negative area", func(t *testing.T) {
		// Move to quadrant Y
		current = InitialDP()

		current.Direction = "S"
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "S", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)

		// Move to quadrant X
		current.Direction = "E"
		current = Move(current, when, unit)
		expected = entities.CurrentDP{Direction: "E", Position_X: 0, Position_Y: 0}
		assert.Equal(t, expected, current)
	})
}
