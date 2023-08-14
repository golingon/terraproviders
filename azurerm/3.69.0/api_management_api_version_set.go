// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapiversionset "github.com/golingon/terraproviders/azurerm/3.69.0/apimanagementapiversionset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementApiVersionSet creates a new instance of [ApiManagementApiVersionSet].
func NewApiManagementApiVersionSet(name string, args ApiManagementApiVersionSetArgs) *ApiManagementApiVersionSet {
	return &ApiManagementApiVersionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiVersionSet)(nil)

// ApiManagementApiVersionSet represents the Terraform resource azurerm_api_management_api_version_set.
type ApiManagementApiVersionSet struct {
	Name      string
	Args      ApiManagementApiVersionSetArgs
	state     *apiManagementApiVersionSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApiVersionSet].
func (amavs *ApiManagementApiVersionSet) Type() string {
	return "azurerm_api_management_api_version_set"
}

// LocalName returns the local name for [ApiManagementApiVersionSet].
func (amavs *ApiManagementApiVersionSet) LocalName() string {
	return amavs.Name
}

// Configuration returns the configuration (args) for [ApiManagementApiVersionSet].
func (amavs *ApiManagementApiVersionSet) Configuration() interface{} {
	return amavs.Args
}

// DependOn is used for other resources to depend on [ApiManagementApiVersionSet].
func (amavs *ApiManagementApiVersionSet) DependOn() terra.Reference {
	return terra.ReferenceResource(amavs)
}

// Dependencies returns the list of resources [ApiManagementApiVersionSet] depends_on.
func (amavs *ApiManagementApiVersionSet) Dependencies() terra.Dependencies {
	return amavs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApiVersionSet].
func (amavs *ApiManagementApiVersionSet) LifecycleManagement() *terra.Lifecycle {
	return amavs.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApiVersionSet].
func (amavs *ApiManagementApiVersionSet) Attributes() apiManagementApiVersionSetAttributes {
	return apiManagementApiVersionSetAttributes{ref: terra.ReferenceResource(amavs)}
}

// ImportState imports the given attribute values into [ApiManagementApiVersionSet]'s state.
func (amavs *ApiManagementApiVersionSet) ImportState(av io.Reader) error {
	amavs.state = &apiManagementApiVersionSetState{}
	if err := json.NewDecoder(av).Decode(amavs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amavs.Type(), amavs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApiVersionSet] has state.
func (amavs *ApiManagementApiVersionSet) State() (*apiManagementApiVersionSetState, bool) {
	return amavs.state, amavs.state != nil
}

// StateMust returns the state for [ApiManagementApiVersionSet]. Panics if the state is nil.
func (amavs *ApiManagementApiVersionSet) StateMust() *apiManagementApiVersionSetState {
	if amavs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amavs.Type(), amavs.LocalName()))
	}
	return amavs.state
}

// ApiManagementApiVersionSetArgs contains the configurations for azurerm_api_management_api_version_set.
type ApiManagementApiVersionSetArgs struct {
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
	Timeouts *apimanagementapiversionset.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiVersionSetAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("resource_group_name"))
}

// VersionHeaderName returns a reference to field version_header_name of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) VersionHeaderName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("version_header_name"))
}

// VersionQueryName returns a reference to field version_query_name of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) VersionQueryName() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("version_query_name"))
}

// VersioningScheme returns a reference to field versioning_scheme of azurerm_api_management_api_version_set.
func (amavs apiManagementApiVersionSetAttributes) VersioningScheme() terra.StringValue {
	return terra.ReferenceAsString(amavs.ref.Append("versioning_scheme"))
}

func (amavs apiManagementApiVersionSetAttributes) Timeouts() apimanagementapiversionset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapiversionset.TimeoutsAttributes](amavs.ref.Append("timeouts"))
}

type apiManagementApiVersionSetState struct {
	ApiManagementName string                                    `json:"api_management_name"`
	Description       string                                    `json:"description"`
	DisplayName       string                                    `json:"display_name"`
	Id                string                                    `json:"id"`
	Name              string                                    `json:"name"`
	ResourceGroupName string                                    `json:"resource_group_name"`
	VersionHeaderName string                                    `json:"version_header_name"`
	VersionQueryName  string                                    `json:"version_query_name"`
	VersioningScheme  string                                    `json:"versioning_scheme"`
	Timeouts          *apimanagementapiversionset.TimeoutsState `json:"timeouts"`
}
