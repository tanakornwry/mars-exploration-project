package registry

import (
	"github.com/tanakornwry/mars-exploration-project/interfaces/controllers"
)

type registry struct{}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controllers.AppController {
	return r.NewRoverController()
}
