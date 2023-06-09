// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementidentityprovideraad "github.com/golingon/terraproviders/azurerm/3.64.0/apimanagementidentityprovideraad"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementIdentityProviderAad creates a new instance of [ApiManagementIdentityProviderAad].
func NewApiManagementIdentityProviderAad(name string, args ApiManagementIdentityProviderAadArgs) *ApiManagementIdentityProviderAad {
	return &ApiManagementIdentityProviderAad{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementIdentityProviderAad)(nil)

// ApiManagementIdentityProviderAad represents the Terraform resource azurerm_api_management_identity_provider_aad.
type ApiManagementIdentityProviderAad struct {
	Name      string
	Args      ApiManagementIdentityProviderAadArgs
	state     *apiManagementIdentityProviderAadState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementIdentityProviderAad].
func (amipa *ApiManagementIdentityProviderAad) Type() string {
	return "azurerm_api_management_identity_provider_aad"
}

// LocalName returns the local name for [ApiManagementIdentityProviderAad].
func (amipa *ApiManagementIdentityProviderAad) LocalName() string {
	return amipa.Name
}

// Configuration returns the configuration (args) for [ApiManagementIdentityProviderAad].
func (amipa *ApiManagementIdentityProviderAad) Configuration() interface{} {
	return amipa.Args
}

// DependOn is used for other resources to depend on [ApiManagementIdentityProviderAad].
func (amipa *ApiManagementIdentityProviderAad) DependOn() terra.Reference {
	return terra.ReferenceResource(amipa)
}

// Dependencies returns the list of resources [ApiManagementIdentityProviderAad] depends_on.
func (amipa *ApiManagementIdentityProviderAad) Dependencies() terra.Dependencies {
	return amipa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementIdentityProviderAad].
func (amipa *ApiManagementIdentityProviderAad) LifecycleManagement() *terra.Lifecycle {
	return amipa.Lifecycle
}

// Attributes returns the attributes for [ApiManagementIdentityProviderAad].
func (amipa *ApiManagementIdentityProviderAad) Attributes() apiManagementIdentityProviderAadAttributes {
	return apiManagementIdentityProviderAadAttributes{ref: terra.ReferenceResource(amipa)}
}

// ImportState imports the given attribute values into [ApiManagementIdentityProviderAad]'s state.
func (amipa *ApiManagementIdentityProviderAad) ImportState(av io.Reader) error {
	amipa.state = &apiManagementIdentityProviderAadState{}
	if err := json.NewDecoder(av).Decode(amipa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amipa.Type(), amipa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementIdentityProviderAad] has state.
func (amipa *ApiManagementIdentityProviderAad) State() (*apiManagementIdentityProviderAadState, bool) {
	return amipa.state, amipa.state != nil
}

// StateMust returns the state for [ApiManagementIdentityProviderAad]. Panics if the state is nil.
func (amipa *ApiManagementIdentityProviderAad) StateMust() *apiManagementIdentityProviderAadState {
	if amipa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amipa.Type(), amipa.LocalName()))
	}
	return amipa.state
}

// ApiManagementIdentityProviderAadArgs contains the configurations for azurerm_api_management_identity_provider_aad.
type ApiManagementIdentityProviderAadArgs struct {
	// AllowedTenants: list of string, required
	AllowedTenants terra.ListValue[terra.StringValue] `hcl:"allowed_tenants,attr" validate:"required"`
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
	// SigninTenant: string, optional
	SigninTenant terra.StringValue `hcl:"signin_tenant,attr"`
	// Timeouts: optional
	Timeouts *apimanagementidentityprovideraad.Timeouts `hcl:"timeouts,block"`
}
type apiManagementIdentityProviderAadAttributes struct {
	ref terra.Reference
}

// AllowedTenants returns a reference to field allowed_tenants of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) AllowedTenants() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](amipa.ref.Append("allowed_tenants"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("api_management_name"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("resource_group_name"))
}

// SigninTenant returns a reference to field signin_tenant of azurerm_api_management_identity_provider_aad.
func (amipa apiManagementIdentityProviderAadAttributes) SigninTenant() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("signin_tenant"))
}

func (amipa apiManagementIdentityProviderAadAttributes) Timeouts() apimanagementidentityprovideraad.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementidentityprovideraad.TimeoutsAttributes](amipa.ref.Append("timeouts"))
}

type apiManagementIdentityProviderAadState struct {
	AllowedTenants    []string                                        `json:"allowed_tenants"`
	ApiManagementName string                                          `json:"api_management_name"`
	ClientId          string                                          `json:"client_id"`
	ClientSecret      string                                          `json:"client_secret"`
	Id                string                                          `json:"id"`
	ResourceGroupName string                                          `json:"resource_group_name"`
	SigninTenant      string                                          `json:"signin_tenant"`
	Timeouts          *apimanagementidentityprovideraad.TimeoutsState `json:"timeouts"`
}
