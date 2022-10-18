// Code generated by ent, DO NOT EDIT.

package ent

import (
	"supabase_server/ent/project"
	"supabase_server/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	projectMixin := schema.Project{}.Mixin()
	projectMixinFields0 := projectMixin[0].Fields()
	_ = projectMixinFields0
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectMixinFields0[0].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectMixinFields0[1].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() uuid.UUID)
}
