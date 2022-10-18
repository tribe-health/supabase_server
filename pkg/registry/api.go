package registry

import (
	"supabase_server/pkg/controller"
	"supabase_server/pkg/repository"
	"supabase_server/pkg/usecase/usecase"
)

func (r *registry) NewApiController() controller.Api {
	repo := repository.NewProjectRepository(r.client)
	u := usecase.NewProjectUseCase(repo)

	return controller.NewApiController(u)
}
