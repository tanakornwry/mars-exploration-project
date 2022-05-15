package input

import (
	"testing"
)

func TestReadCommand(t *testing.T) {
	command := NewInput().ReadCommand()

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
