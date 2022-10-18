package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"supabase_server/ent"
	"supabase_server/graph/generated"

	"github.com/google/uuid"
)

// OrganizationID is the resolver for the organizationID field.
func (r *projectResolver) OrganizationID(ctx context.Context, obj *ent.Project) (string, error) {
	panic(fmt.Errorf("not implemented: OrganizationID - organizationID"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ProjectWhereInput) (*ent.ProjectConnection, error) {
	return r.controller.Project.List(ctx, after, first, before, last, where)
}

// OrganizationID is the resolver for the organizationID field.
func (r *createProjectInputResolver) OrganizationID(ctx context.Context, obj *ent.CreateProjectInput, data string) error {
	orgId, err := uuid.Parse(data)
	if err == nil {
		obj.OrganizationID = orgId
	}
	return err
}

// OrganizationID is the resolver for the organizationID field.
func (r *projectWhereInputResolver) OrganizationID(ctx context.Context, obj *ent.ProjectWhereInput, data *string) error {
	if data != nil {
		orgId, err := uuid.Parse(*data)
		if err == nil {
			obj.OrganizationID = &orgId
		}
		return err
	}
	return nil
}

// OrganizationIDNeq is the resolver for the organizationIDNEQ field.
func (r *projectWhereInputResolver) OrganizationIDNeq(ctx context.Context, obj *ent.ProjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDNeq - organizationIDNEQ"))
}

// OrganizationIDIn is the resolver for the organizationIDIn field.
func (r *projectWhereInputResolver) OrganizationIDIn(ctx context.Context, obj *ent.ProjectWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDIn - organizationIDIn"))
}

// OrganizationIDNotIn is the resolver for the organizationIDNotIn field.
func (r *projectWhereInputResolver) OrganizationIDNotIn(ctx context.Context, obj *ent.ProjectWhereInput, data []string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDNotIn - organizationIDNotIn"))
}

// OrganizationIDGt is the resolver for the organizationIDGT field.
func (r *projectWhereInputResolver) OrganizationIDGt(ctx context.Context, obj *ent.ProjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDGt - organizationIDGT"))
}

// OrganizationIDGte is the resolver for the organizationIDGTE field.
func (r *projectWhereInputResolver) OrganizationIDGte(ctx context.Context, obj *ent.ProjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDGte - organizationIDGTE"))
}

// OrganizationIDLt is the resolver for the organizationIDLT field.
func (r *projectWhereInputResolver) OrganizationIDLt(ctx context.Context, obj *ent.ProjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDLt - organizationIDLT"))
}

// OrganizationIDLte is the resolver for the organizationIDLTE field.
func (r *projectWhereInputResolver) OrganizationIDLte(ctx context.Context, obj *ent.ProjectWhereInput, data *string) error {
	panic(fmt.Errorf("not implemented: OrganizationIDLte - organizationIDLTE"))
}

// OrganizationID is the resolver for the organizationID field.
func (r *updateProjectInputResolver) OrganizationID(ctx context.Context, obj *ent.UpdateProjectInput, data *string) error {
	panic(fmt.Errorf("not implemented: OrganizationID - organizationID"))
}

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// CreateProjectInput returns generated.CreateProjectInputResolver implementation.
func (r *Resolver) CreateProjectInput() generated.CreateProjectInputResolver {
	return &createProjectInputResolver{r}
}

// ProjectWhereInput returns generated.ProjectWhereInputResolver implementation.
func (r *Resolver) ProjectWhereInput() generated.ProjectWhereInputResolver {
	return &projectWhereInputResolver{r}
}

// UpdateProjectInput returns generated.UpdateProjectInputResolver implementation.
func (r *Resolver) UpdateProjectInput() generated.UpdateProjectInputResolver {
	return &updateProjectInputResolver{r}
}

type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createProjectInputResolver struct{ *Resolver }
type projectWhereInputResolver struct{ *Resolver }
type updateProjectInputResolver struct{ *Resolver }
