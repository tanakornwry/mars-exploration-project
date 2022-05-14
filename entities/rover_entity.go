package entities

type CurrentDP struct {
	Degree     int
	Position_X int
	Position_Y int
}

type CommandConf struct {
	SizeMap int
	Command [][2]int // 1 is command, 2 is block
}

type InstructionStruct struct {
	RotateInstruction map[string]int
}

var Instruction = InstructionStruct{
	RotateInstruction: map[string]int{
		"R": -90,
		"L": 90,
	},
}
