package registry

import (
	"majezanu/capstone/internal/contracts/controller"
)

type registry struct {
	csvPath string
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(csvPath string) Registry {
	return &registry{csvPath}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
