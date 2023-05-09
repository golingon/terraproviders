// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	directoryroleassignment "github.com/golingon/terraproviders/azuread/2.38.0/directoryroleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryRoleAssignment creates a new instance of [DirectoryRoleAssignment].
func NewDirectoryRoleAssignment(name string, args DirectoryRoleAssignmentArgs) *DirectoryRoleAssignment {
	return &DirectoryRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryRoleAssignment)(nil)

// DirectoryRoleAssignment represents the Terraform resource azuread_directory_role_assignment.
type DirectoryRoleAssignment struct {
	Name      string
	Args      DirectoryRoleAssignmentArgs
	state     *directoryRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryRoleAssignment].
func (dra *DirectoryRoleAssignment) Type() string {
	return "azuread_directory_role_assignment"
}

// LocalName returns the local name for [DirectoryRoleAssignment].
func (dra *DirectoryRoleAssignment) LocalName() string {
	return dra.Name
}

// Configuration returns the configuration (args) for [DirectoryRoleAssignment].
func (dra *DirectoryRoleAssignment) Configuration() interface{} {
	return dra.Args
}

// DependOn is used for other resources to depend on [DirectoryRoleAssignment].
func (dra *DirectoryRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(dra)
}

// Dependencies returns the list of resources [DirectoryRoleAssignment] depends_on.
func (dra *DirectoryRoleAssignment) Dependencies() terra.Dependencies {
	return dra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryRoleAssignment].
func (dra *DirectoryRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return dra.Lifecycle
}

// Attributes returns the attributes for [DirectoryRoleAssignment].
func (dra *DirectoryRoleAssignment) Attributes() directoryRoleAssignmentAttributes {
	return directoryRoleAssignmentAttributes{ref: terra.ReferenceResource(dra)}
}

// ImportState imports the given attribute values into [DirectoryRoleAssignment]'s state.
func (dra *DirectoryRoleAssignment) ImportState(av io.Reader) error {
	dra.state = &directoryRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(dra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dra.Type(), dra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryRoleAssignment] has state.
func (dra *DirectoryRoleAssignment) State() (*directoryRoleAssignmentState, bool) {
	return dra.state, dra.state != nil
}

// StateMust returns the state for [DirectoryRoleAssignment]. Panics if the state is nil.
func (dra *DirectoryRoleAssignment) StateMust() *directoryRoleAssignmentState {
	if dra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dra.Type(), dra.LocalName()))
	}
	return dra.state
}

// DirectoryRoleAssignmentArgs contains the configurations for azuread_directory_role_assignment.
type DirectoryRoleAssignmentArgs struct {
	// AppScopeId: string, optional
	AppScopeId terra.StringValue `hcl:"app_scope_id,attr"`
	// AppScopeObjectId: string, optional
	AppScopeObjectId terra.StringValue `hcl:"app_scope_object_id,attr"`
	// DirectoryScopeId: string, optional
	DirectoryScopeId terra.StringValue `hcl:"directory_scope_id,attr"`
	// DirectoryScopeObjectId: string, optional
	DirectoryScopeObjectId terra.StringValue `hcl:"directory_scope_object_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrincipalObjectId: string, required
	PrincipalObjectId terra.StringValue `hcl:"principal_object_id,attr" validate:"required"`
	// RoleId: string, required
	RoleId terra.StringValue `hcl:"role_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *directoryroleassignment.Timeouts `hcl:"timeouts,block"`
}
type directoryRoleAssignmentAttributes struct {
	ref terra.Reference
}

// AppScopeId returns a reference to field app_scope_id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) AppScopeId() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("app_scope_id"))
}

// AppScopeObjectId returns a reference to field app_scope_object_id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) AppScopeObjectId() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("app_scope_object_id"))
}

// DirectoryScopeId returns a reference to field directory_scope_id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) DirectoryScopeId() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("directory_scope_id"))
}

// DirectoryScopeObjectId returns a reference to field directory_scope_object_id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) DirectoryScopeObjectId() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("directory_scope_object_id"))
}

// Id returns a reference to field id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("id"))
}

// PrincipalObjectId returns a reference to field principal_object_id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) PrincipalObjectId() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("principal_object_id"))
}

// RoleId returns a reference to field role_id of azuread_directory_role_assignment.
func (dra directoryRoleAssignmentAttributes) RoleId() terra.StringValue {
	return terra.ReferenceAsString(dra.ref.Append("role_id"))
}

func (dra directoryRoleAssignmentAttributes) Timeouts() directoryroleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryroleassignment.TimeoutsAttributes](dra.ref.Append("timeouts"))
}

type directoryRoleAssignmentState struct {
	AppScopeId             string                                 `json:"app_scope_id"`
	AppScopeObjectId       string                                 `json:"app_scope_object_id"`
	DirectoryScopeId       string                                 `json:"directory_scope_id"`
	DirectoryScopeObjectId string                                 `json:"directory_scope_object_id"`
	Id                     string                                 `json:"id"`
	PrincipalObjectId      string                                 `json:"principal_object_id"`
	RoleId                 string                                 `json:"role_id"`
	Timeouts               *directoryroleassignment.TimeoutsState `json:"timeouts"`
}
