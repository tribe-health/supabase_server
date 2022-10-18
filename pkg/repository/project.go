package repository

import (
	"context"
	"supabase_server/ent"
	"supabase_server/pkg/entity/model"
	"supabase_server/pkg/usecase/repository"
)

type projectRepository struct {
	client *ent.Client
}

func (p projectRepository) Get(ctx context.Context, id *model.ID) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.ProjectWhereInput) (*model.ProjectConnection, error) {
	results, err := p.client.Project.Query().Paginate(ctx, after, first, before, last, ent.WithProjectFilter(where.Filter))
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (p projectRepository) Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p projectRepository) Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error) {
	//TODO implement me
	panic("implement me")
}

func NewProjectRepository(client *ent.Client) repository.Project {
	return &projectRepository{
		client: client,
	}
}
