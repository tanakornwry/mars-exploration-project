package input

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCommand(t *testing.T) {
	command, err := NewInput().ReadCommand()
	fmt.Println(command)

	if err != "" {
		want := "Error: Not found the command file."
		assert.Equal(t, want, err)
	}

	if command.SizeMap < 1 {
		t.Error("Size map can not be zero or negative number")
	}

	for _, v := range command.Command {
		for k := range v {
			if k != "L" && k != "HL" && k != "R" && k != "HR" && k != "F" && k != "B" {
				t.Error("Some command didn't correct")
			}
		}

	}

	fmt.Println(command)
	if len(command.Command) != 26 {
		t.Error("Refactor command was incorrect!!")
	}
}
