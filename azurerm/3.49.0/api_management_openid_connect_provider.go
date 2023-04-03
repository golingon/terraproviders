// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementopenidconnectprovider "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementopenidconnectprovider"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementOpenidConnectProvider creates a new instance of [ApiManagementOpenidConnectProvider].
func NewApiManagementOpenidConnectProvider(name string, args ApiManagementOpenidConnectProviderArgs) *ApiManagementOpenidConnectProvider {
	return &ApiManagementOpenidConnectProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementOpenidConnectProvider)(nil)

// ApiManagementOpenidConnectProvider represents the Terraform resource azurerm_api_management_openid_connect_provider.
type ApiManagementOpenidConnectProvider struct {
	Name      string
	Args      ApiManagementOpenidConnectProviderArgs
	state     *apiManagementOpenidConnectProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementOpenidConnectProvider].
func (amocp *ApiManagementOpenidConnectProvider) Type() string {
	return "azurerm_api_management_openid_connect_provider"
}

// LocalName returns the local name for [ApiManagementOpenidConnectProvider].
func (amocp *ApiManagementOpenidConnectProvider) LocalName() string {
	return amocp.Name
}

// Configuration returns the configuration (args) for [ApiManagementOpenidConnectProvider].
func (amocp *ApiManagementOpenidConnectProvider) Configuration() interface{} {
	return amocp.Args
}

// DependOn is used for other resources to depend on [ApiManagementOpenidConnectProvider].
func (amocp *ApiManagementOpenidConnectProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(amocp)
}

// Dependencies returns the list of resources [ApiManagementOpenidConnectProvider] depends_on.
func (amocp *ApiManagementOpenidConnectProvider) Dependencies() terra.Dependencies {
	return amocp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementOpenidConnectProvider].
func (amocp *ApiManagementOpenidConnectProvider) LifecycleManagement() *terra.Lifecycle {
	return amocp.Lifecycle
}

// Attributes returns the attributes for [ApiManagementOpenidConnectProvider].
func (amocp *ApiManagementOpenidConnectProvider) Attributes() apiManagementOpenidConnectProviderAttributes {
	return apiManagementOpenidConnectProviderAttributes{ref: terra.ReferenceResource(amocp)}
}

// ImportState imports the given attribute values into [ApiManagementOpenidConnectProvider]'s state.
func (amocp *ApiManagementOpenidConnectProvider) ImportState(av io.Reader) error {
	amocp.state = &apiManagementOpenidConnectProviderState{}
	if err := json.NewDecoder(av).Decode(amocp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amocp.Type(), amocp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementOpenidConnectProvider] has state.
func (amocp *ApiManagementOpenidConnectProvider) State() (*apiManagementOpenidConnectProviderState, bool) {
	return amocp.state, amocp.state != nil
}

// StateMust returns the state for [ApiManagementOpenidConnectProvider]. Panics if the state is nil.
func (amocp *ApiManagementOpenidConnectProvider) StateMust() *apiManagementOpenidConnectProviderState {
	if amocp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amocp.Type(), amocp.LocalName()))
	}
	return amocp.state
}

// ApiManagementOpenidConnectProviderArgs contains the configurations for azurerm_api_management_openid_connect_provider.
type ApiManagementOpenidConnectProviderArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MetadataEndpoint: string, required
	MetadataEndpoint terra.StringValue `hcl:"metadata_endpoint,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementopenidconnectprovider.Timeouts `hcl:"timeouts,block"`
}
type apiManagementOpenidConnectProviderAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("api_management_name"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("client_secret"))
}

// Description returns a reference to field description of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("id"))
}

// MetadataEndpoint returns a reference to field metadata_endpoint of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) MetadataEndpoint() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("metadata_endpoint"))
}

// Name returns a reference to field name of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_openid_connect_provider.
func (amocp apiManagementOpenidConnectProviderAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amocp.ref.Append("resource_group_name"))
}

func (amocp apiManagementOpenidConnectProviderAttributes) Timeouts() apimanagementopenidconnectprovider.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementopenidconnectprovider.TimeoutsAttributes](amocp.ref.Append("timeouts"))
}

type apiManagementOpenidConnectProviderState struct {
	ApiManagementName string                                            `json:"api_management_name"`
	ClientId          string                                            `json:"client_id"`
	ClientSecret      string                                            `json:"client_secret"`
	Description       string                                            `json:"description"`
	DisplayName       string                                            `json:"display_name"`
	Id                string                                            `json:"id"`
	MetadataEndpoint  string                                            `json:"metadata_endpoint"`
	Name              string                                            `json:"name"`
	ResourceGroupName string                                            `json:"resource_group_name"`
	Timeouts          *apimanagementopenidconnectprovider.TimeoutsState `json:"timeouts"`
}
