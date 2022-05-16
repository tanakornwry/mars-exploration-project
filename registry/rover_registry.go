package registry

import (
	"github.com/tanakornwry/mars-exploration-project/infrastructure/input"
	"github.com/tanakornwry/mars-exploration-project/interfaces/controllers"
	"github.com/tanakornwry/mars-exploration-project/interfaces/presenters"
	"github.com/tanakornwry/mars-exploration-project/services/interactors"
	rovermodules "github.com/tanakornwry/mars-exploration-project/services/modules"
	roverpresenters "github.com/tanakornwry/mars-exploration-project/services/presenters"
)

func (r *registry) NewInput() input.Input {
	return input.NewInput()
}

func (r *registry) NewRoverModules() rovermodules.RoverModules {
	return rovermodules.NewRoverModules()
}

func (r *registry) NewRoverPresenters() roverpresenters.RoverPresenters {
	return roverpresenters.NewRoverPresenters()
}

func (r *registry) NewRoverInteractors() interactors.RoverInteractors {
	return interactors.NewRoverInteractors(r.NewRoverModules(), r.NewRoverPresenters())
}

func (r *registry) NewInterfacePresenter() presenters.InterfacePresenters {
	return presenters.NewInterfaceRes()
}
func (r *registry) NewRoverController() controllers.RoverController {
	return controllers.NewRoverController(r.NewInput(), r.NewRoverInteractors(), r.NewInterfacePresenter())
}
