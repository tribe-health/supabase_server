package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	mixin "supabase_server/ent/schema/mixins"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

/*
type CreateProjectBody struct {
	DbPass         string                  `json:"db_pass"`
	Name           string                  `json:"name"`
	OrganizationID string                  `json:"organization_id"`
	Plan           CreateProjectBodyPlan   `json:"plan"`
	Region         CreateProjectBodyRegion `json:"region"`
	KpsEnabled     OptBool                 `json:"kps_enabled"`
}
*/

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.UUID("organization_id", uuid.UUID{}),
		field.String("region"),
		field.Bool("kps_enabled"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}

func (Project) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Project) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}
