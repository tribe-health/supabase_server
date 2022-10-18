package usecase

import (
	"context"
	"supabase_server/pkg/entity/model"
	"supabase_server/pkg/usecase/repository"
)

type project struct {
	projectRepository repository.Project
}

func (p project) Get(ctx context.Context, id *model.ID) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p project) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput) (*model.ProjectConnection, error) {
	return p.projectRepository.List(ctx, after, first, before, last, where)
}

func (p project) Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p project) Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

type Project interface {
	Get(ctx context.Context, id *model.ID) (*model.Project, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int,
		where *model.ProjectWhereInput) (*model.ProjectConnection, error)
	Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error)
	Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error)
}

func NewProjectUseCase(r repository.Project) Project {
	return &project{projectRepository: r}
}
