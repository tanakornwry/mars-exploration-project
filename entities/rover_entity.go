package entities

type CurrentDP struct {
	Direction  string
	Position_X int
	Position_Y int
}

type CommandConf struct {
	SizeMap int
	Command []string
}
