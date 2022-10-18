package controller

import (
	"context"
	ht "github.com/ogen-go/ogen/http"
	oas "supabase_server/api"
	"supabase_server/pkg/entity/model"
	"supabase_server/pkg/usecase/usecase"
	"time"
)

type ApiController interface {
	oas.Handler

	SecurityHandler() oas.SecurityHandler
}

type api struct {
	ApiController

	projectUsecase usecase.Project
}

type security struct {
	oas.SecurityHandler
}

// HandleApiKeyHeader handles apiKeyHeader security.
func (security) HandleApiKeyHeader(ctx context.Context, operationName string, t oas.ApiKeyHeader) (context.Context, error) {
	return ctx, nil
}

// HandleApiKeyParam handles apiKeyParam security.
func (security) HandleApiKeyParam(ctx context.Context, operationName string, t oas.ApiKeyParam) (context.Context, error) {
	return ctx, nil
}

// HandleBearer handles bearer security.
func (security) HandleBearer(ctx context.Context, operationName string, t oas.Bearer) (context.Context, error) {
	return ctx, nil
}

func (api) SecurityHandler() oas.SecurityHandler {
	return &security{}
}
func (api) CreateFunction(ctx context.Context, req oas.CreateFunctionBody,
	params oas.CreateFunctionParams) (r oas.CreateFunctionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateOrganization implements createOrganization operation.
//
// Create an organization.
//
// POST /v1/organizations
func (api) CreateOrganization(ctx context.Context, req oas.CreateOrganizationBody) (r oas.CreateOrganizationRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateProject implements createProject operation.
//
// Create a project.
//
// POST /v1/projects
func (api) CreateProject(ctx context.Context, req oas.CreateProjectBody) (r oas.ProjectResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSecrets implements createSecrets operation.
//
// Creates multiple secrets and adds them to the specified project.
//
// POST /v1/projects/{ref}/secrets
func (api) CreateSecrets(ctx context.Context, req []oas.CreateSecretBody, params oas.CreateSecretsParams) (r oas.CreateSecretsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteFunction implements deleteFunction operation.
//
// Deletes a function with the specified slug from the specified project.
//
// DELETE /v1/projects/{ref}/functions/{function_slug}
func (api) DeleteFunction(ctx context.Context, params oas.DeleteFunctionParams) (r oas.DeleteFunctionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteSecrets implements deleteSecrets operation.
//
// Deletes all secrets with the given names from the specified project.
//
// DELETE /v1/projects/{ref}/secrets
func (api) DeleteSecrets(ctx context.Context, req []string, params oas.DeleteSecretsParams) (r oas.DeleteSecretsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetConfig implements getConfig operation.
//
// Gets project's pgsodium config.
//
// GET /v1/projects/{ref}/pgsodium
func (api) GetConfig(ctx context.Context, params oas.GetConfigParams) (r oas.GetConfigRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFunction implements getFunction operation.
//
// Retrieves a function with the specified slug and project.
//
// GET /v1/projects/{ref}/functions/{function_slug}
func (api) GetFunction(ctx context.Context, params oas.GetFunctionParams) (r oas.GetFunctionRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFunctions implements getFunctions operation.
//
// Returns all functions you've previously added to the specified project.
//
// GET /v1/projects/{ref}/functions
func (api) GetFunctions(ctx context.Context, params oas.GetFunctionsParams) (r oas.GetFunctionsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetOrganizations implements getOrganizations operation.
//
// Returns a list of organizations that you currently belong to.
//
// GET /v1/organizations
func (api) GetOrganizations(ctx context.Context) (r oas.GetOrganizationsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetProjects implements getProjects operation.
//
// Returns a list of all projects you've previously created.
//
// GET /v1/projects
func (a api) GetProjects(ctx context.Context) (r []oas.ProjectResponse, e error) {
	var first = 0
	connection, err := a.projectUsecase.List(ctx, nil, &first, nil, nil, &model.ProjectWhereInput{})
	if err != nil {
		return nil, err
	}

	var responses []oas.ProjectResponse
	for _, v := range connection.Edges {
		responses = append(responses, oas.ProjectResponse{
			ID:             v.Node.ID.String(),
			OrganizationID: v.Node.OrganizationID.String(),
			Name:           v.Node.Name,
			Region:         v.Node.Region,
			CreatedAt:      v.Node.CreatedAt.Format(time.RFC3339),
		})
	}

	return responses, err
}

// GetSecrets implements getSecrets operation.
//
// Returns all secrets you've previously added to the specified project.
//
// GET /v1/projects/{ref}/secrets
func (api) GetSecrets(ctx context.Context, params oas.GetSecretsParams) (r oas.GetSecretsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateConfig implements updateConfig operation.
//
// Updates project's pgsodium config. Updating the root_key can cause all data encrypted with the
// older key to become inaccessible.
//
// PUT /v1/projects/{ref}/pgsodium
func (api) UpdateConfig(ctx context.Context, req oas.UpdatePgsodiumConfigBody, params oas.UpdateConfigParams) (r oas.UpdateConfigRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateFunction implements updateFunction operation.
//
// Updates a function with the specified slug and project.
//
// PATCH /v1/projects/{ref}/functions/{function_slug}
func (api) UpdateFunction(ctx context.Context, req oas.UpdateFunctionBody, params oas.UpdateFunctionParams) (r oas.UpdateFunctionRes, _ error) {
	return r, ht.ErrNotImplemented
}

type Api interface {
	ApiController
}

func NewApiController(ucProject usecase.Project) Api {
	return &api{
		projectUsecase: ucProject,
	}
}
