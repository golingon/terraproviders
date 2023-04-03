// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementidentityprovidertwitter "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementidentityprovidertwitter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementIdentityProviderTwitter creates a new instance of [ApiManagementIdentityProviderTwitter].
func NewApiManagementIdentityProviderTwitter(name string, args ApiManagementIdentityProviderTwitterArgs) *ApiManagementIdentityProviderTwitter {
	return &ApiManagementIdentityProviderTwitter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementIdentityProviderTwitter)(nil)

// ApiManagementIdentityProviderTwitter represents the Terraform resource azurerm_api_management_identity_provider_twitter.
type ApiManagementIdentityProviderTwitter struct {
	Name      string
	Args      ApiManagementIdentityProviderTwitterArgs
	state     *apiManagementIdentityProviderTwitterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementIdentityProviderTwitter].
func (amipt *ApiManagementIdentityProviderTwitter) Type() string {
	return "azurerm_api_management_identity_provider_twitter"
}

// LocalName returns the local name for [ApiManagementIdentityProviderTwitter].
func (amipt *ApiManagementIdentityProviderTwitter) LocalName() string {
	return amipt.Name
}

// Configuration returns the configuration (args) for [ApiManagementIdentityProviderTwitter].
func (amipt *ApiManagementIdentityProviderTwitter) Configuration() interface{} {
	return amipt.Args
}

// DependOn is used for other resources to depend on [ApiManagementIdentityProviderTwitter].
func (amipt *ApiManagementIdentityProviderTwitter) DependOn() terra.Reference {
	return terra.ReferenceResource(amipt)
}

// Dependencies returns the list of resources [ApiManagementIdentityProviderTwitter] depends_on.
func (amipt *ApiManagementIdentityProviderTwitter) Dependencies() terra.Dependencies {
	return amipt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementIdentityProviderTwitter].
func (amipt *ApiManagementIdentityProviderTwitter) LifecycleManagement() *terra.Lifecycle {
	return amipt.Lifecycle
}

// Attributes returns the attributes for [ApiManagementIdentityProviderTwitter].
func (amipt *ApiManagementIdentityProviderTwitter) Attributes() apiManagementIdentityProviderTwitterAttributes {
	return apiManagementIdentityProviderTwitterAttributes{ref: terra.ReferenceResource(amipt)}
}

// ImportState imports the given attribute values into [ApiManagementIdentityProviderTwitter]'s state.
func (amipt *ApiManagementIdentityProviderTwitter) ImportState(av io.Reader) error {
	amipt.state = &apiManagementIdentityProviderTwitterState{}
	if err := json.NewDecoder(av).Decode(amipt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amipt.Type(), amipt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementIdentityProviderTwitter] has state.
func (amipt *ApiManagementIdentityProviderTwitter) State() (*apiManagementIdentityProviderTwitterState, bool) {
	return amipt.state, amipt.state != nil
}

// StateMust returns the state for [ApiManagementIdentityProviderTwitter]. Panics if the state is nil.
func (amipt *ApiManagementIdentityProviderTwitter) StateMust() *apiManagementIdentityProviderTwitterState {
	if amipt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amipt.Type(), amipt.LocalName()))
	}
	return amipt.state
}

// ApiManagementIdentityProviderTwitterArgs contains the configurations for azurerm_api_management_identity_provider_twitter.
type ApiManagementIdentityProviderTwitterArgs struct {
	// ApiKey: string, required
	ApiKey terra.StringValue `hcl:"api_key,attr" validate:"required"`
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiSecretKey: string, required
	ApiSecretKey terra.StringValue `hcl:"api_secret_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementidentityprovidertwitter.Timeouts `hcl:"timeouts,block"`
}
type apiManagementIdentityProviderTwitterAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of azurerm_api_management_identity_provider_twitter.
func (amipt apiManagementIdentityProviderTwitterAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(amipt.ref.Append("api_key"))
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_twitter.
func (amipt apiManagementIdentityProviderTwitterAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amipt.ref.Append("api_management_name"))
}

// ApiSecretKey returns a reference to field api_secret_key of azurerm_api_management_identity_provider_twitter.
func (amipt apiManagementIdentityProviderTwitterAttributes) ApiSecretKey() terra.StringValue {
	return terra.ReferenceAsString(amipt.ref.Append("api_secret_key"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_twitter.
func (amipt apiManagementIdentityProviderTwitterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amipt.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_twitter.
func (amipt apiManagementIdentityProviderTwitterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amipt.ref.Append("resource_group_name"))
}

func (amipt apiManagementIdentityProviderTwitterAttributes) Timeouts() apimanagementidentityprovidertwitter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementidentityprovidertwitter.TimeoutsAttributes](amipt.ref.Append("timeouts"))
}

type apiManagementIdentityProviderTwitterState struct {
	ApiKey            string                                              `json:"api_key"`
	ApiManagementName string                                              `json:"api_management_name"`
	ApiSecretKey      string                                              `json:"api_secret_key"`
	Id                string                                              `json:"id"`
	ResourceGroupName string                                              `json:"resource_group_name"`
	Timeouts          *apimanagementidentityprovidertwitter.TimeoutsState `json:"timeouts"`
}
