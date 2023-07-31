// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	customdirectoryrole "github.com/golingon/terraproviders/azuread/2.41.0/customdirectoryrole"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCustomDirectoryRole creates a new instance of [CustomDirectoryRole].
func NewCustomDirectoryRole(name string, args CustomDirectoryRoleArgs) *CustomDirectoryRole {
	return &CustomDirectoryRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CustomDirectoryRole)(nil)

// CustomDirectoryRole represents the Terraform resource azuread_custom_directory_role.
type CustomDirectoryRole struct {
	Name      string
	Args      CustomDirectoryRoleArgs
	state     *customDirectoryRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CustomDirectoryRole].
func (cdr *CustomDirectoryRole) Type() string {
	return "azuread_custom_directory_role"
}

// LocalName returns the local name for [CustomDirectoryRole].
func (cdr *CustomDirectoryRole) LocalName() string {
	return cdr.Name
}

// Configuration returns the configuration (args) for [CustomDirectoryRole].
func (cdr *CustomDirectoryRole) Configuration() interface{} {
	return cdr.Args
}

// DependOn is used for other resources to depend on [CustomDirectoryRole].
func (cdr *CustomDirectoryRole) DependOn() terra.Reference {
	return terra.ReferenceResource(cdr)
}

// Dependencies returns the list of resources [CustomDirectoryRole] depends_on.
func (cdr *CustomDirectoryRole) Dependencies() terra.Dependencies {
	return cdr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CustomDirectoryRole].
func (cdr *CustomDirectoryRole) LifecycleManagement() *terra.Lifecycle {
	return cdr.Lifecycle
}

// Attributes returns the attributes for [CustomDirectoryRole].
func (cdr *CustomDirectoryRole) Attributes() customDirectoryRoleAttributes {
	return customDirectoryRoleAttributes{ref: terra.ReferenceResource(cdr)}
}

// ImportState imports the given attribute values into [CustomDirectoryRole]'s state.
func (cdr *CustomDirectoryRole) ImportState(av io.Reader) error {
	cdr.state = &customDirectoryRoleState{}
	if err := json.NewDecoder(av).Decode(cdr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdr.Type(), cdr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CustomDirectoryRole] has state.
func (cdr *CustomDirectoryRole) State() (*customDirectoryRoleState, bool) {
	return cdr.state, cdr.state != nil
}

// StateMust returns the state for [CustomDirectoryRole]. Panics if the state is nil.
func (cdr *CustomDirectoryRole) StateMust() *customDirectoryRoleState {
	if cdr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdr.Type(), cdr.LocalName()))
	}
	return cdr.state
}

// CustomDirectoryRoleArgs contains the configurations for azuread_custom_directory_role.
type CustomDirectoryRoleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TemplateId: string, optional
	TemplateId terra.StringValue `hcl:"template_id,attr"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Permissions: min=1
	Permissions []customdirectoryrole.Permissions `hcl:"permissions,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *customdirectoryrole.Timeouts `hcl:"timeouts,block"`
}
type customDirectoryRoleAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cdr.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cdr.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cdr.ref.Append("enabled"))
}

// Id returns a reference to field id of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdr.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(cdr.ref.Append("object_id"))
}

// TemplateId returns a reference to field template_id of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(cdr.ref.Append("template_id"))
}

// Version returns a reference to field version of azuread_custom_directory_role.
func (cdr customDirectoryRoleAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(cdr.ref.Append("version"))
}

func (cdr customDirectoryRoleAttributes) Permissions() terra.SetValue[customdirectoryrole.PermissionsAttributes] {
	return terra.ReferenceAsSet[customdirectoryrole.PermissionsAttributes](cdr.ref.Append("permissions"))
}

func (cdr customDirectoryRoleAttributes) Timeouts() customdirectoryrole.TimeoutsAttributes {
	return terra.ReferenceAsSingle[customdirectoryrole.TimeoutsAttributes](cdr.ref.Append("timeouts"))
}

type customDirectoryRoleState struct {
	Description string                                 `json:"description"`
	DisplayName string                                 `json:"display_name"`
	Enabled     bool                                   `json:"enabled"`
	Id          string                                 `json:"id"`
	ObjectId    string                                 `json:"object_id"`
	TemplateId  string                                 `json:"template_id"`
	Version     string                                 `json:"version"`
	Permissions []customdirectoryrole.PermissionsState `json:"permissions"`
	Timeouts    *customdirectoryrole.TimeoutsState     `json:"timeouts"`
}
