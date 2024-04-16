// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_identity_provider_aad

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

// Resource represents the Terraform resource azurerm_api_management_identity_provider_aad.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementIdentityProviderAadState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aamipa *Resource) Type() string {
	return "azurerm_api_management_identity_provider_aad"
}

// LocalName returns the local name for [Resource].
func (aamipa *Resource) LocalName() string {
	return aamipa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aamipa *Resource) Configuration() interface{} {
	return aamipa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aamipa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aamipa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aamipa *Resource) Dependencies() terra.Dependencies {
	return aamipa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aamipa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aamipa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aamipa *Resource) Attributes() azurermApiManagementIdentityProviderAadAttributes {
	return azurermApiManagementIdentityProviderAadAttributes{ref: terra.ReferenceResource(aamipa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aamipa *Resource) ImportState(state io.Reader) error {
	aamipa.state = &azurermApiManagementIdentityProviderAadState{}
	if err := json.NewDecoder(state).Decode(aamipa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aamipa.Type(), aamipa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aamipa *Resource) State() (*azurermApiManagementIdentityProviderAadState, bool) {
	return aamipa.state, aamipa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aamipa *Resource) StateMust() *azurermApiManagementIdentityProviderAadState {
	if aamipa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aamipa.Type(), aamipa.LocalName()))
	}
	return aamipa.state
}

// Args contains the configurations for azurerm_api_management_identity_provider_aad.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementIdentityProviderAadAttributes struct {
	ref terra.Reference
}

// AllowedTenants returns a reference to field allowed_tenants of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) AllowedTenants() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aamipa.ref.Append("allowed_tenants"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aamipa.ref.Append("api_management_name"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(aamipa.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(aamipa.ref.Append("client_secret"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aamipa.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aamipa.ref.Append("resource_group_name"))
}

// SigninTenant returns a reference to field signin_tenant of azurerm_api_management_identity_provider_aad.
func (aamipa azurermApiManagementIdentityProviderAadAttributes) SigninTenant() terra.StringValue {
	return terra.ReferenceAsString(aamipa.ref.Append("signin_tenant"))
}

func (aamipa azurermApiManagementIdentityProviderAadAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aamipa.ref.Append("timeouts"))
}

type azurermApiManagementIdentityProviderAadState struct {
	AllowedTenants    []string       `json:"allowed_tenants"`
	ApiManagementName string         `json:"api_management_name"`
	ClientId          string         `json:"client_id"`
	ClientSecret      string         `json:"client_secret"`
	Id                string         `json:"id"`
	ResourceGroupName string         `json:"resource_group_name"`
	SigninTenant      string         `json:"signin_tenant"`
	Timeouts          *TimeoutsState `json:"timeouts"`
}
