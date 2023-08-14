// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementidentityprovideraadb2c "github.com/golingon/terraproviders/azurerm/3.69.0/apimanagementidentityprovideraadb2c"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementIdentityProviderAadb2C creates a new instance of [ApiManagementIdentityProviderAadb2C].
func NewApiManagementIdentityProviderAadb2C(name string, args ApiManagementIdentityProviderAadb2CArgs) *ApiManagementIdentityProviderAadb2C {
	return &ApiManagementIdentityProviderAadb2C{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementIdentityProviderAadb2C)(nil)

// ApiManagementIdentityProviderAadb2C represents the Terraform resource azurerm_api_management_identity_provider_aadb2c.
type ApiManagementIdentityProviderAadb2C struct {
	Name      string
	Args      ApiManagementIdentityProviderAadb2CArgs
	state     *apiManagementIdentityProviderAadb2CState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementIdentityProviderAadb2C].
func (amipa *ApiManagementIdentityProviderAadb2C) Type() string {
	return "azurerm_api_management_identity_provider_aadb2c"
}

// LocalName returns the local name for [ApiManagementIdentityProviderAadb2C].
func (amipa *ApiManagementIdentityProviderAadb2C) LocalName() string {
	return amipa.Name
}

// Configuration returns the configuration (args) for [ApiManagementIdentityProviderAadb2C].
func (amipa *ApiManagementIdentityProviderAadb2C) Configuration() interface{} {
	return amipa.Args
}

// DependOn is used for other resources to depend on [ApiManagementIdentityProviderAadb2C].
func (amipa *ApiManagementIdentityProviderAadb2C) DependOn() terra.Reference {
	return terra.ReferenceResource(amipa)
}

// Dependencies returns the list of resources [ApiManagementIdentityProviderAadb2C] depends_on.
func (amipa *ApiManagementIdentityProviderAadb2C) Dependencies() terra.Dependencies {
	return amipa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementIdentityProviderAadb2C].
func (amipa *ApiManagementIdentityProviderAadb2C) LifecycleManagement() *terra.Lifecycle {
	return amipa.Lifecycle
}

// Attributes returns the attributes for [ApiManagementIdentityProviderAadb2C].
func (amipa *ApiManagementIdentityProviderAadb2C) Attributes() apiManagementIdentityProviderAadb2CAttributes {
	return apiManagementIdentityProviderAadb2CAttributes{ref: terra.ReferenceResource(amipa)}
}

// ImportState imports the given attribute values into [ApiManagementIdentityProviderAadb2C]'s state.
func (amipa *ApiManagementIdentityProviderAadb2C) ImportState(av io.Reader) error {
	amipa.state = &apiManagementIdentityProviderAadb2CState{}
	if err := json.NewDecoder(av).Decode(amipa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amipa.Type(), amipa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementIdentityProviderAadb2C] has state.
func (amipa *ApiManagementIdentityProviderAadb2C) State() (*apiManagementIdentityProviderAadb2CState, bool) {
	return amipa.state, amipa.state != nil
}

// StateMust returns the state for [ApiManagementIdentityProviderAadb2C]. Panics if the state is nil.
func (amipa *ApiManagementIdentityProviderAadb2C) StateMust() *apiManagementIdentityProviderAadb2CState {
	if amipa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amipa.Type(), amipa.LocalName()))
	}
	return amipa.state
}

// ApiManagementIdentityProviderAadb2CArgs contains the configurations for azurerm_api_management_identity_provider_aadb2c.
type ApiManagementIdentityProviderAadb2CArgs struct {
	// AllowedTenant: string, required
	AllowedTenant terra.StringValue `hcl:"allowed_tenant,attr" validate:"required"`
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Authority: string, required
	Authority terra.StringValue `hcl:"authority,attr" validate:"required"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PasswordResetPolicy: string, optional
	PasswordResetPolicy terra.StringValue `hcl:"password_reset_policy,attr"`
	// ProfileEditingPolicy: string, optional
	ProfileEditingPolicy terra.StringValue `hcl:"profile_editing_policy,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SigninPolicy: string, required
	SigninPolicy terra.StringValue `hcl:"signin_policy,attr" validate:"required"`
	// SigninTenant: string, required
	SigninTenant terra.StringValue `hcl:"signin_tenant,attr" validate:"required"`
	// SignupPolicy: string, required
	SignupPolicy terra.StringValue `hcl:"signup_policy,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementidentityprovideraadb2c.Timeouts `hcl:"timeouts,block"`
}
type apiManagementIdentityProviderAadb2CAttributes struct {
	ref terra.Reference
}

// AllowedTenant returns a reference to field allowed_tenant of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) AllowedTenant() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("allowed_tenant"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("api_management_name"))
}

// Authority returns a reference to field authority of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) Authority() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("authority"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("id"))
}

// PasswordResetPolicy returns a reference to field password_reset_policy of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) PasswordResetPolicy() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("password_reset_policy"))
}

// ProfileEditingPolicy returns a reference to field profile_editing_policy of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) ProfileEditingPolicy() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("profile_editing_policy"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("resource_group_name"))
}

// SigninPolicy returns a reference to field signin_policy of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) SigninPolicy() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("signin_policy"))
}

// SigninTenant returns a reference to field signin_tenant of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) SigninTenant() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("signin_tenant"))
}

// SignupPolicy returns a reference to field signup_policy of azurerm_api_management_identity_provider_aadb2c.
func (amipa apiManagementIdentityProviderAadb2CAttributes) SignupPolicy() terra.StringValue {
	return terra.ReferenceAsString(amipa.ref.Append("signup_policy"))
}

func (amipa apiManagementIdentityProviderAadb2CAttributes) Timeouts() apimanagementidentityprovideraadb2c.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementidentityprovideraadb2c.TimeoutsAttributes](amipa.ref.Append("timeouts"))
}

type apiManagementIdentityProviderAadb2CState struct {
	AllowedTenant        string                                             `json:"allowed_tenant"`
	ApiManagementName    string                                             `json:"api_management_name"`
	Authority            string                                             `json:"authority"`
	ClientId             string                                             `json:"client_id"`
	ClientSecret         string                                             `json:"client_secret"`
	Id                   string                                             `json:"id"`
	PasswordResetPolicy  string                                             `json:"password_reset_policy"`
	ProfileEditingPolicy string                                             `json:"profile_editing_policy"`
	ResourceGroupName    string                                             `json:"resource_group_name"`
	SigninPolicy         string                                             `json:"signin_policy"`
	SigninTenant         string                                             `json:"signin_tenant"`
	SignupPolicy         string                                             `json:"signup_policy"`
	Timeouts             *apimanagementidentityprovideraadb2c.TimeoutsState `json:"timeouts"`
}