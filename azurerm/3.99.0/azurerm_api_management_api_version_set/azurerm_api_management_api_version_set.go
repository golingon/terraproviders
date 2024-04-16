// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_api_version_set

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_api_management_api_version_set.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementApiVersionSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aamavs *Resource) Type() string {
	return "azurerm_api_management_api_version_set"
}

// LocalName returns the local name for [Resource].
func (aamavs *Resource) LocalName() string {
	return aamavs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aamavs *Resource) Configuration() interface{} {
	return aamavs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aamavs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aamavs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aamavs *Resource) Dependencies() terra.Dependencies {
	return aamavs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aamavs *Resource) LifecycleManagement() *terra.Lifecycle {
	return aamavs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aamavs *Resource) Attributes() azurermApiManagementApiVersionSetAttributes {
	return azurermApiManagementApiVersionSetAttributes{ref: terra.ReferenceResource(aamavs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aamavs *Resource) ImportState(state io.Reader) error {
	aamavs.state = &azurermApiManagementApiVersionSetState{}
	if err := json.NewDecoder(state).Decode(aamavs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aamavs.Type(), aamavs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aamavs *Resource) State() (*azurermApiManagementApiVersionSetState, bool) {
	return aamavs.state, aamavs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aamavs *Resource) StateMust() *azurermApiManagementApiVersionSetState {
	if aamavs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aamavs.Type(), aamavs.LocalName()))
	}
	return aamavs.state
}

// Args contains the configurations for azurerm_api_management_api_version_set.
type Args struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VersionHeaderName: string, optional
	VersionHeaderName terra.StringValue `hcl:"version_header_name,attr"`
	// VersionQueryName: string, optional
	VersionQueryName terra.StringValue `hcl:"version_query_name,attr"`
	// VersioningScheme: string, required
	VersioningScheme terra.StringValue `hcl:"versioning_scheme,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementApiVersionSetAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("resource_group_name"))
}

// VersionHeaderName returns a reference to field version_header_name of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) VersionHeaderName() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("version_header_name"))
}

// VersionQueryName returns a reference to field version_query_name of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) VersionQueryName() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("version_query_name"))
}

// VersioningScheme returns a reference to field versioning_scheme of azurerm_api_management_api_version_set.
func (aamavs azurermApiManagementApiVersionSetAttributes) VersioningScheme() terra.StringValue {
	return terra.ReferenceAsString(aamavs.ref.Append("versioning_scheme"))
}

func (aamavs azurermApiManagementApiVersionSetAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aamavs.ref.Append("timeouts"))
}

type azurermApiManagementApiVersionSetState struct {
	ApiManagementName string         `json:"api_management_name"`
	Description       string         `json:"description"`
	DisplayName       string         `json:"display_name"`
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	ResourceGroupName string         `json:"resource_group_name"`
	VersionHeaderName string         `json:"version_header_name"`
	VersionQueryName  string         `json:"version_query_name"`
	VersioningScheme  string         `json:"versioning_scheme"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
