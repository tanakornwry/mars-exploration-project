package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanakornwry/mars-exploration-project/entities"
)

func TestRespose(t *testing.T) {
	rr := []entities.CurrentDP{
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

	ip := NewInterfaceRes()
	have := ip.Response(rr)
	want := []string{
		"N:0,0",
		"E:0,0",
		"E:1,0",
		"N:1,0",
		"N:1,1",
		"W:1,1",
		"S:1,1",
		"S:1,0",
		"W:1,0",
	}

	assert.Equal(t, have, want)
}
