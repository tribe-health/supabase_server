package registry

import (
	"supabase_server/ent"
	"supabase_server/pkg/controller"
)

type registry struct {
	client *ent.Client
}

// Registry is an interface of registry
type Registry interface {
	NewController() controller.Controller
}

// New registers entire controller with dependencies
func New(client *ent.Client) Registry {
	return &registry{
		client: client,
	}
}

// NewController generates controllers
func (r *registry) NewController() controller.Controller {
	return controller.Controller{
		Project: r.NewProjectController(),
		Api:     r.NewApiController(),
	}
}
