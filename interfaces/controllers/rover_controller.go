package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanakornwry/mars-exploration-project/infrastructure/input"
	"github.com/tanakornwry/mars-exploration-project/interfaces/presenters"
	"github.com/tanakornwry/mars-exploration-project/services/interactors"
)

type roverController struct {
	Input               input.Input
	RoverInteractors    interactors.RoverInteractors
	InterfacePresenters presenters.InterfacePresenters
}

type RoverController interface {
	Explore(c *gin.Context)
	Greeting(c *gin.Context)
}

func NewRoverController(i input.Input, ri interactors.RoverInteractors, ip presenters.InterfacePresenters) RoverController {
	return &roverController{i, ri, ip}
}

func (rc *roverController) Explore(c *gin.Context) {
	// Fetch the command
	commandConf, err := rc.Input.ReadCommand()
	if err != "" {
		c.JSON(http.StatusNotAcceptable, map[string]string{
			"system": err,
		})
		return
	}

	// Start the rover mission
	roverPath := rc.RoverInteractors.StartRover(commandConf)

	// Adjust the response to suit with a human
	response := rc.InterfacePresenters.Response(roverPath)

	c.JSON(http.StatusOK, response)
}

// Just for fun
func (rc *roverController) Greeting(c *gin.Context) {
	greeting := map[string]string{
		"system": "Hello Mars, we're here!!",
	}
	c.JSON(http.StatusOK, greeting)
}
