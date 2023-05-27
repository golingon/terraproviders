// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	directoryrole "github.com/golingon/terraproviders/azuread/2.39.0/directoryrole"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryRole creates a new instance of [DirectoryRole].
func NewDirectoryRole(name string, args DirectoryRoleArgs) *DirectoryRole {
	return &DirectoryRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryRole)(nil)

// DirectoryRole represents the Terraform resource azuread_directory_role.
type DirectoryRole struct {
	Name      string
	Args      DirectoryRoleArgs
	state     *directoryRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryRole].
func (dr *DirectoryRole) Type() string {
	return "azuread_directory_role"
}

// LocalName returns the local name for [DirectoryRole].
func (dr *DirectoryRole) LocalName() string {
	return dr.Name
}

// Configuration returns the configuration (args) for [DirectoryRole].
func (dr *DirectoryRole) Configuration() interface{} {
	return dr.Args
}

// DependOn is used for other resources to depend on [DirectoryRole].
func (dr *DirectoryRole) DependOn() terra.Reference {
	return terra.ReferenceResource(dr)
}

// Dependencies returns the list of resources [DirectoryRole] depends_on.
func (dr *DirectoryRole) Dependencies() terra.Dependencies {
	return dr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryRole].
func (dr *DirectoryRole) LifecycleManagement() *terra.Lifecycle {
	return dr.Lifecycle
}

// Attributes returns the attributes for [DirectoryRole].
func (dr *DirectoryRole) Attributes() directoryRoleAttributes {
	return directoryRoleAttributes{ref: terra.ReferenceResource(dr)}
}

// ImportState imports the given attribute values into [DirectoryRole]'s state.
func (dr *DirectoryRole) ImportState(av io.Reader) error {
	dr.state = &directoryRoleState{}
	if err := json.NewDecoder(av).Decode(dr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dr.Type(), dr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryRole] has state.
func (dr *DirectoryRole) State() (*directoryRoleState, bool) {
	return dr.state, dr.state != nil
}

// StateMust returns the state for [DirectoryRole]. Panics if the state is nil.
func (dr *DirectoryRole) StateMust() *directoryRoleState {
	if dr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dr.Type(), dr.LocalName()))
	}
	return dr.state
}

// DirectoryRoleArgs contains the configurations for azuread_directory_role.
type DirectoryRoleArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TemplateId: string, optional
	TemplateId terra.StringValue `hcl:"template_id,attr"`
	// Timeouts: optional
	Timeouts *directoryrole.Timeouts `hcl:"timeouts,block"`
}
type directoryRoleAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azuread_directory_role.
func (dr directoryRoleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_directory_role.
func (dr directoryRoleAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("display_name"))
}

// Id returns a reference to field id of azuread_directory_role.
func (dr directoryRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azuread_directory_role.
func (dr directoryRoleAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("object_id"))
}

// TemplateId returns a reference to field template_id of azuread_directory_role.
func (dr directoryRoleAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("template_id"))
}

func (dr directoryRoleAttributes) Timeouts() directoryrole.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryrole.TimeoutsAttributes](dr.ref.Append("timeouts"))
}

type directoryRoleState struct {
	Description string                       `json:"description"`
	DisplayName string                       `json:"display_name"`
	Id          string                       `json:"id"`
	ObjectId    string                       `json:"object_id"`
	TemplateId  string                       `json:"template_id"`
	Timeouts    *directoryrole.TimeoutsState `json:"timeouts"`
}
