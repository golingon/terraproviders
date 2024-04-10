// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewProjectIamCustomRole creates a new instance of [ProjectIamCustomRole].
func NewProjectIamCustomRole(name string, args ProjectIamCustomRoleArgs) *ProjectIamCustomRole {
	return &ProjectIamCustomRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectIamCustomRole)(nil)

// ProjectIamCustomRole represents the Terraform resource google_project_iam_custom_role.
type ProjectIamCustomRole struct {
	Name      string
	Args      ProjectIamCustomRoleArgs
	state     *projectIamCustomRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectIamCustomRole].
func (picr *ProjectIamCustomRole) Type() string {
	return "google_project_iam_custom_role"
}

// LocalName returns the local name for [ProjectIamCustomRole].
func (picr *ProjectIamCustomRole) LocalName() string {
	return picr.Name
}

// Configuration returns the configuration (args) for [ProjectIamCustomRole].
func (picr *ProjectIamCustomRole) Configuration() interface{} {
	return picr.Args
}

// DependOn is used for other resources to depend on [ProjectIamCustomRole].
func (picr *ProjectIamCustomRole) DependOn() terra.Reference {
	return terra.ReferenceResource(picr)
}

// Dependencies returns the list of resources [ProjectIamCustomRole] depends_on.
func (picr *ProjectIamCustomRole) Dependencies() terra.Dependencies {
	return picr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectIamCustomRole].
func (picr *ProjectIamCustomRole) LifecycleManagement() *terra.Lifecycle {
	return picr.Lifecycle
}

// Attributes returns the attributes for [ProjectIamCustomRole].
func (picr *ProjectIamCustomRole) Attributes() projectIamCustomRoleAttributes {
	return projectIamCustomRoleAttributes{ref: terra.ReferenceResource(picr)}
}

// ImportState imports the given attribute values into [ProjectIamCustomRole]'s state.
func (picr *ProjectIamCustomRole) ImportState(av io.Reader) error {
	picr.state = &projectIamCustomRoleState{}
	if err := json.NewDecoder(av).Decode(picr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", picr.Type(), picr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectIamCustomRole] has state.
func (picr *ProjectIamCustomRole) State() (*projectIamCustomRoleState, bool) {
	return picr.state, picr.state != nil
}

// StateMust returns the state for [ProjectIamCustomRole]. Panics if the state is nil.
func (picr *ProjectIamCustomRole) StateMust() *projectIamCustomRoleState {
	if picr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", picr.Type(), picr.LocalName()))
	}
	return picr.state
}

// ProjectIamCustomRoleArgs contains the configurations for google_project_iam_custom_role.
type ProjectIamCustomRoleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Permissions: set of string, required
	Permissions terra.SetValue[terra.StringValue] `hcl:"permissions,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RoleId: string, required
	RoleId terra.StringValue `hcl:"role_id,attr" validate:"required"`
	// Stage: string, optional
	Stage terra.StringValue `hcl:"stage,attr"`
	// Title: string, required
	Title terra.StringValue `hcl:"title,attr" validate:"required"`
}
type projectIamCustomRoleAttributes struct {
	ref terra.Reference
}

// Deleted returns a reference to field deleted of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Deleted() terra.BoolValue {
	return terra.ReferenceAsBool(picr.ref.Append("deleted"))
}

// Description returns a reference to field description of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("description"))
}

// Id returns a reference to field id of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("id"))
}

// Name returns a reference to field name of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("name"))
}

// Permissions returns a reference to field permissions of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Permissions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](picr.ref.Append("permissions"))
}

// Project returns a reference to field project of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("project"))
}

// RoleId returns a reference to field role_id of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) RoleId() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("role_id"))
}

// Stage returns a reference to field stage of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("stage"))
}

// Title returns a reference to field title of google_project_iam_custom_role.
func (picr projectIamCustomRoleAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(picr.ref.Append("title"))
}

type projectIamCustomRoleState struct {
	Deleted     bool     `json:"deleted"`
	Description string   `json:"description"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Project     string   `json:"project"`
	RoleId      string   `json:"role_id"`
	Stage       string   `json:"stage"`
	Title       string   `json:"title"`
}
