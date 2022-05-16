package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCommand(t *testing.T) {
	command, err := NewInput().ReadCommand()

	if err != "" {
		want := "Error: Not found the command file."
		assert.Equal(t, want, err)
	}

	if command.SizeMap < 1 {
		t.Error("Size map can not be zero or negative number")
	}

	for _, v := range command.Command {
		for k := range v {
			if k != "L" && k != "R" && k != "F" && k != "B" {
				t.Error("Some command didn't correct")
			}
		}

	}
}
