// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	apimanagementidentityprovidermicrosoft "github.com/golingon/terraproviders/azurerm/3.98.0/apimanagementidentityprovidermicrosoft"
	"io"
)

// NewApiManagementIdentityProviderMicrosoft creates a new instance of [ApiManagementIdentityProviderMicrosoft].
func NewApiManagementIdentityProviderMicrosoft(name string, args ApiManagementIdentityProviderMicrosoftArgs) *ApiManagementIdentityProviderMicrosoft {
	return &ApiManagementIdentityProviderMicrosoft{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementIdentityProviderMicrosoft)(nil)

// ApiManagementIdentityProviderMicrosoft represents the Terraform resource azurerm_api_management_identity_provider_microsoft.
type ApiManagementIdentityProviderMicrosoft struct {
	Name      string
	Args      ApiManagementIdentityProviderMicrosoftArgs
	state     *apiManagementIdentityProviderMicrosoftState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementIdentityProviderMicrosoft].
func (amipm *ApiManagementIdentityProviderMicrosoft) Type() string {
	return "azurerm_api_management_identity_provider_microsoft"
}

// LocalName returns the local name for [ApiManagementIdentityProviderMicrosoft].
func (amipm *ApiManagementIdentityProviderMicrosoft) LocalName() string {
	return amipm.Name
}

// Configuration returns the configuration (args) for [ApiManagementIdentityProviderMicrosoft].
func (amipm *ApiManagementIdentityProviderMicrosoft) Configuration() interface{} {
	return amipm.Args
}

// DependOn is used for other resources to depend on [ApiManagementIdentityProviderMicrosoft].
func (amipm *ApiManagementIdentityProviderMicrosoft) DependOn() terra.Reference {
	return terra.ReferenceResource(amipm)
}

// Dependencies returns the list of resources [ApiManagementIdentityProviderMicrosoft] depends_on.
func (amipm *ApiManagementIdentityProviderMicrosoft) Dependencies() terra.Dependencies {
	return amipm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementIdentityProviderMicrosoft].
func (amipm *ApiManagementIdentityProviderMicrosoft) LifecycleManagement() *terra.Lifecycle {
	return amipm.Lifecycle
}

// Attributes returns the attributes for [ApiManagementIdentityProviderMicrosoft].
func (amipm *ApiManagementIdentityProviderMicrosoft) Attributes() apiManagementIdentityProviderMicrosoftAttributes {
	return apiManagementIdentityProviderMicrosoftAttributes{ref: terra.ReferenceResource(amipm)}
}

// ImportState imports the given attribute values into [ApiManagementIdentityProviderMicrosoft]'s state.
func (amipm *ApiManagementIdentityProviderMicrosoft) ImportState(av io.Reader) error {
	amipm.state = &apiManagementIdentityProviderMicrosoftState{}
	if err := json.NewDecoder(av).Decode(amipm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amipm.Type(), amipm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementIdentityProviderMicrosoft] has state.
func (amipm *ApiManagementIdentityProviderMicrosoft) State() (*apiManagementIdentityProviderMicrosoftState, bool) {
	return amipm.state, amipm.state != nil
}

// StateMust returns the state for [ApiManagementIdentityProviderMicrosoft]. Panics if the state is nil.
func (amipm *ApiManagementIdentityProviderMicrosoft) StateMust() *apiManagementIdentityProviderMicrosoftState {
	if amipm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amipm.Type(), amipm.LocalName()))
	}
	return amipm.state
}

// ApiManagementIdentityProviderMicrosoftArgs contains the configurations for azurerm_api_management_identity_provider_microsoft.
type ApiManagementIdentityProviderMicrosoftArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementidentityprovidermicrosoft.Timeouts `hcl:"timeouts,block"`
}
type apiManagementIdentityProviderMicrosoftAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_microsoft.
func (amipm apiManagementIdentityProviderMicrosoftAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amipm.ref.Append("api_management_name"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_identity_provider_microsoft.
func (amipm apiManagementIdentityProviderMicrosoftAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(amipm.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_identity_provider_microsoft.
func (amipm apiManagementIdentityProviderMicrosoftAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(amipm.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_microsoft.
func (amipm apiManagementIdentityProviderMicrosoftAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amipm.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_microsoft.
func (amipm apiManagementIdentityProviderMicrosoftAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amipm.ref.Append("resource_group_name"))
}

func (amipm apiManagementIdentityProviderMicrosoftAttributes) Timeouts() apimanagementidentityprovidermicrosoft.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementidentityprovidermicrosoft.TimeoutsAttributes](amipm.ref.Append("timeouts"))
}

type apiManagementIdentityProviderMicrosoftState struct {
	ApiManagementName string                                                `json:"api_management_name"`
	ClientId          string                                                `json:"client_id"`
	ClientSecret      string                                                `json:"client_secret"`
	Id                string                                                `json:"id"`
	ResourceGroupName string                                                `json:"resource_group_name"`
	Timeouts          *apimanagementidentityprovidermicrosoft.TimeoutsState `json:"timeouts"`
}
