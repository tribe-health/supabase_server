package controller

import (
	"context"
	"supabase_server/pkg/entity/model"
	"supabase_server/pkg/usecase/usecase"
)

type project struct {
	projectUsecase usecase.Project
}

func (p project) Get(ctx context.Context, id *model.ID) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p project) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput) (*model.ProjectConnection, error) {
	//TODO implement me
	panic("implement me")
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

func NewProjectController(uc usecase.Project) Project {
	return &project{projectUsecase: uc}
}
