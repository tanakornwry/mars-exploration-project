package rovermodule

import (
	"github.com/tanakornwry/mars-exploration-project/entities"
)

type RoverService interface {
	InitialDP() entities.CommandConf
	Rotate(entities.CommandConf) entities.CurrentDP
	Move(entities.CommandConf) entities.CurrentDP
}

func InitialDP() entities.CurrentDP {
	return entities.CurrentDP{}
}

func Rotate(c entities.CurrentDP, d string) entities.CurrentDP {
	return entities.CurrentDP{}
}

func Move(c entities.CurrentDP, m string, u int) entities.CurrentDP {
	return entities.CurrentDP{}
}
