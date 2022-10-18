package registry

import (
	"supabase_server/pkg/controller"
	"supabase_server/pkg/repository"
	"supabase_server/pkg/usecase/usecase"
)

func (r *registry) NewProjectController() controller.Project {
	repo := repository.NewProjectRepository(r.client)
	u := usecase.NewProjectUseCase(repo)

	return controller.NewProjectController(u)
}
