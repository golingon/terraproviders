// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import (
	"encoding/json"
	"fmt"
	role "github.com/golingon/terraproviders/googleworkspace/0.7.0/role"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRole creates a new instance of [Role].
func NewRole(name string, args RoleArgs) *Role {
	return &Role{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Role)(nil)

// Role represents the Terraform resource googleworkspace_role.
type Role struct {
	Name      string
	Args      RoleArgs
	state     *roleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Role].
func (r *Role) Type() string {
	return "googleworkspace_role"
}

// LocalName returns the local name for [Role].
func (r *Role) LocalName() string {
	return r.Name
}

// Configuration returns the configuration (args) for [Role].
func (r *Role) Configuration() interface{} {
	return r.Args
}

// DependOn is used for other resources to depend on [Role].
func (r *Role) DependOn() terra.Reference {
	return terra.ReferenceResource(r)
}

// Dependencies returns the list of resources [Role] depends_on.
func (r *Role) Dependencies() terra.Dependencies {
	return r.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Role].
func (r *Role) LifecycleManagement() *terra.Lifecycle {
	return r.Lifecycle
}

// Attributes returns the attributes for [Role].
func (r *Role) Attributes() roleAttributes {
	return roleAttributes{ref: terra.ReferenceResource(r)}
}

// ImportState imports the given attribute values into [Role]'s state.
func (r *Role) ImportState(av io.Reader) error {
	r.state = &roleState{}
	if err := json.NewDecoder(av).Decode(r.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", r.Type(), r.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Role] has state.
func (r *Role) State() (*roleState, bool) {
	return r.state, r.state != nil
}

// StateMust returns the state for [Role]. Panics if the state is nil.
func (r *Role) StateMust() *roleState {
	if r.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", r.Type(), r.LocalName()))
	}
	return r.state
}

// RoleArgs contains the configurations for googleworkspace_role.
type RoleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Privileges: min=1
	Privileges []role.Privileges `hcl:"privileges,block" validate:"min=1"`
}
type roleAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of googleworkspace_role.
func (r roleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

// Etag returns a reference to field etag of googleworkspace_role.
func (r roleAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_role.
func (r roleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

// IsSuperAdminRole returns a reference to field is_super_admin_role of googleworkspace_role.
func (r roleAttributes) IsSuperAdminRole() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("is_super_admin_role"))
}

// IsSystemRole returns a reference to field is_system_role of googleworkspace_role.
func (r roleAttributes) IsSystemRole() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("is_system_role"))
}

// Name returns a reference to field name of googleworkspace_role.
func (r roleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r roleAttributes) Privileges() terra.SetValue[role.PrivilegesAttributes] {
	return terra.ReferenceAsSet[role.PrivilegesAttributes](r.ref.Append("privileges"))
}

type roleState struct {
	Description      string                 `json:"description"`
	Etag             string                 `json:"etag"`
	Id               string                 `json:"id"`
	IsSuperAdminRole bool                   `json:"is_super_admin_role"`
	IsSystemRole     bool                   `json:"is_system_role"`
	Name             string                 `json:"name"`
	Privileges       []role.PrivilegesState `json:"privileges"`
}