// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	approleassignment "github.com/golingon/terraproviders/azuread/2.38.0/approleassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppRoleAssignment creates a new instance of [AppRoleAssignment].
func NewAppRoleAssignment(name string, args AppRoleAssignmentArgs) *AppRoleAssignment {
	return &AppRoleAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppRoleAssignment)(nil)

// AppRoleAssignment represents the Terraform resource azuread_app_role_assignment.
type AppRoleAssignment struct {
	Name      string
	Args      AppRoleAssignmentArgs
	state     *appRoleAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppRoleAssignment].
func (ara *AppRoleAssignment) Type() string {
	return "azuread_app_role_assignment"
}

// LocalName returns the local name for [AppRoleAssignment].
func (ara *AppRoleAssignment) LocalName() string {
	return ara.Name
}

// Configuration returns the configuration (args) for [AppRoleAssignment].
func (ara *AppRoleAssignment) Configuration() interface{} {
	return ara.Args
}

// DependOn is used for other resources to depend on [AppRoleAssignment].
func (ara *AppRoleAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(ara)
}

// Dependencies returns the list of resources [AppRoleAssignment] depends_on.
func (ara *AppRoleAssignment) Dependencies() terra.Dependencies {
	return ara.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppRoleAssignment].
func (ara *AppRoleAssignment) LifecycleManagement() *terra.Lifecycle {
	return ara.Lifecycle
}

// Attributes returns the attributes for [AppRoleAssignment].
func (ara *AppRoleAssignment) Attributes() appRoleAssignmentAttributes {
	return appRoleAssignmentAttributes{ref: terra.ReferenceResource(ara)}
}

// ImportState imports the given attribute values into [AppRoleAssignment]'s state.
func (ara *AppRoleAssignment) ImportState(av io.Reader) error {
	ara.state = &appRoleAssignmentState{}
	if err := json.NewDecoder(av).Decode(ara.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ara.Type(), ara.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppRoleAssignment] has state.
func (ara *AppRoleAssignment) State() (*appRoleAssignmentState, bool) {
	return ara.state, ara.state != nil
}

// StateMust returns the state for [AppRoleAssignment]. Panics if the state is nil.
func (ara *AppRoleAssignment) StateMust() *appRoleAssignmentState {
	if ara.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ara.Type(), ara.LocalName()))
	}
	return ara.state
}

// AppRoleAssignmentArgs contains the configurations for azuread_app_role_assignment.
type AppRoleAssignmentArgs struct {
	// AppRoleId: string, required
	AppRoleId terra.StringValue `hcl:"app_role_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrincipalObjectId: string, required
	PrincipalObjectId terra.StringValue `hcl:"principal_object_id,attr" validate:"required"`
	// ResourceObjectId: string, required
	ResourceObjectId terra.StringValue `hcl:"resource_object_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *approleassignment.Timeouts `hcl:"timeouts,block"`
}
type appRoleAssignmentAttributes struct {
	ref terra.Reference
}

// AppRoleId returns a reference to field app_role_id of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) AppRoleId() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("app_role_id"))
}

// Id returns a reference to field id of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("id"))
}

// PrincipalDisplayName returns a reference to field principal_display_name of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) PrincipalDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("principal_display_name"))
}

// PrincipalObjectId returns a reference to field principal_object_id of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) PrincipalObjectId() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("principal_object_id"))
}

// PrincipalType returns a reference to field principal_type of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) PrincipalType() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("principal_type"))
}

// ResourceDisplayName returns a reference to field resource_display_name of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) ResourceDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("resource_display_name"))
}

// ResourceObjectId returns a reference to field resource_object_id of azuread_app_role_assignment.
func (ara appRoleAssignmentAttributes) ResourceObjectId() terra.StringValue {
	return terra.ReferenceAsString(ara.ref.Append("resource_object_id"))
}

func (ara appRoleAssignmentAttributes) Timeouts() approleassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[approleassignment.TimeoutsAttributes](ara.ref.Append("timeouts"))
}

type appRoleAssignmentState struct {
	AppRoleId            string                           `json:"app_role_id"`
	Id                   string                           `json:"id"`
	PrincipalDisplayName string                           `json:"principal_display_name"`
	PrincipalObjectId    string                           `json:"principal_object_id"`
	PrincipalType        string                           `json:"principal_type"`
	ResourceDisplayName  string                           `json:"resource_display_name"`
	ResourceObjectId     string                           `json:"resource_object_id"`
	Timeouts             *approleassignment.TimeoutsState `json:"timeouts"`
}
