// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementidentityprovidergoogle "github.com/golingon/terraproviders/azurerm/3.52.0/apimanagementidentityprovidergoogle"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementIdentityProviderGoogle creates a new instance of [ApiManagementIdentityProviderGoogle].
func NewApiManagementIdentityProviderGoogle(name string, args ApiManagementIdentityProviderGoogleArgs) *ApiManagementIdentityProviderGoogle {
	return &ApiManagementIdentityProviderGoogle{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementIdentityProviderGoogle)(nil)

// ApiManagementIdentityProviderGoogle represents the Terraform resource azurerm_api_management_identity_provider_google.
type ApiManagementIdentityProviderGoogle struct {
	Name      string
	Args      ApiManagementIdentityProviderGoogleArgs
	state     *apiManagementIdentityProviderGoogleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementIdentityProviderGoogle].
func (amipg *ApiManagementIdentityProviderGoogle) Type() string {
	return "azurerm_api_management_identity_provider_google"
}

// LocalName returns the local name for [ApiManagementIdentityProviderGoogle].
func (amipg *ApiManagementIdentityProviderGoogle) LocalName() string {
	return amipg.Name
}

// Configuration returns the configuration (args) for [ApiManagementIdentityProviderGoogle].
func (amipg *ApiManagementIdentityProviderGoogle) Configuration() interface{} {
	return amipg.Args
}

// DependOn is used for other resources to depend on [ApiManagementIdentityProviderGoogle].
func (amipg *ApiManagementIdentityProviderGoogle) DependOn() terra.Reference {
	return terra.ReferenceResource(amipg)
}

// Dependencies returns the list of resources [ApiManagementIdentityProviderGoogle] depends_on.
func (amipg *ApiManagementIdentityProviderGoogle) Dependencies() terra.Dependencies {
	return amipg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementIdentityProviderGoogle].
func (amipg *ApiManagementIdentityProviderGoogle) LifecycleManagement() *terra.Lifecycle {
	return amipg.Lifecycle
}

// Attributes returns the attributes for [ApiManagementIdentityProviderGoogle].
func (amipg *ApiManagementIdentityProviderGoogle) Attributes() apiManagementIdentityProviderGoogleAttributes {
	return apiManagementIdentityProviderGoogleAttributes{ref: terra.ReferenceResource(amipg)}
}

// ImportState imports the given attribute values into [ApiManagementIdentityProviderGoogle]'s state.
func (amipg *ApiManagementIdentityProviderGoogle) ImportState(av io.Reader) error {
	amipg.state = &apiManagementIdentityProviderGoogleState{}
	if err := json.NewDecoder(av).Decode(amipg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amipg.Type(), amipg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementIdentityProviderGoogle] has state.
func (amipg *ApiManagementIdentityProviderGoogle) State() (*apiManagementIdentityProviderGoogleState, bool) {
	return amipg.state, amipg.state != nil
}

// StateMust returns the state for [ApiManagementIdentityProviderGoogle]. Panics if the state is nil.
func (amipg *ApiManagementIdentityProviderGoogle) StateMust() *apiManagementIdentityProviderGoogleState {
	if amipg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amipg.Type(), amipg.LocalName()))
	}
	return amipg.state
}

// ApiManagementIdentityProviderGoogleArgs contains the configurations for azurerm_api_management_identity_provider_google.
type ApiManagementIdentityProviderGoogleArgs struct {
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
	Timeouts *apimanagementidentityprovidergoogle.Timeouts `hcl:"timeouts,block"`
}
type apiManagementIdentityProviderGoogleAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_google.
func (amipg apiManagementIdentityProviderGoogleAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amipg.ref.Append("api_management_name"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_identity_provider_google.
func (amipg apiManagementIdentityProviderGoogleAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(amipg.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_identity_provider_google.
func (amipg apiManagementIdentityProviderGoogleAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(amipg.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_google.
func (amipg apiManagementIdentityProviderGoogleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amipg.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_google.
func (amipg apiManagementIdentityProviderGoogleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amipg.ref.Append("resource_group_name"))
}

func (amipg apiManagementIdentityProviderGoogleAttributes) Timeouts() apimanagementidentityprovidergoogle.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementidentityprovidergoogle.TimeoutsAttributes](amipg.ref.Append("timeouts"))
}

type apiManagementIdentityProviderGoogleState struct {
	ApiManagementName string                                             `json:"api_management_name"`
	ClientId          string                                             `json:"client_id"`
	ClientSecret      string                                             `json:"client_secret"`
	Id                string                                             `json:"id"`
	ResourceGroupName string                                             `json:"resource_group_name"`
	Timeouts          *apimanagementidentityprovidergoogle.TimeoutsState `json:"timeouts"`
}
