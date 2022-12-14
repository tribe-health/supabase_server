package repository

import (
	"context"
	"supabase_server/pkg/entity/model"
)

type Project interface {
	Get(ctx context.Context, id *model.ID) (*model.Project, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int,
		where *model.ProjectWhereInput) (*model.ProjectConnection, error)
	Create(ctx context.Context, input model.CreateProjectInput) (*model.Project, error)
	Update(ctx context.Context, input model.UpdateProjectInput) (*model.Project, error)
}
