// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementidentityproviderfacebook "github.com/golingon/terraproviders/azurerm/3.66.0/apimanagementidentityproviderfacebook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementIdentityProviderFacebook creates a new instance of [ApiManagementIdentityProviderFacebook].
func NewApiManagementIdentityProviderFacebook(name string, args ApiManagementIdentityProviderFacebookArgs) *ApiManagementIdentityProviderFacebook {
	return &ApiManagementIdentityProviderFacebook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementIdentityProviderFacebook)(nil)

// ApiManagementIdentityProviderFacebook represents the Terraform resource azurerm_api_management_identity_provider_facebook.
type ApiManagementIdentityProviderFacebook struct {
	Name      string
	Args      ApiManagementIdentityProviderFacebookArgs
	state     *apiManagementIdentityProviderFacebookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementIdentityProviderFacebook].
func (amipf *ApiManagementIdentityProviderFacebook) Type() string {
	return "azurerm_api_management_identity_provider_facebook"
}

// LocalName returns the local name for [ApiManagementIdentityProviderFacebook].
func (amipf *ApiManagementIdentityProviderFacebook) LocalName() string {
	return amipf.Name
}

// Configuration returns the configuration (args) for [ApiManagementIdentityProviderFacebook].
func (amipf *ApiManagementIdentityProviderFacebook) Configuration() interface{} {
	return amipf.Args
}

// DependOn is used for other resources to depend on [ApiManagementIdentityProviderFacebook].
func (amipf *ApiManagementIdentityProviderFacebook) DependOn() terra.Reference {
	return terra.ReferenceResource(amipf)
}

// Dependencies returns the list of resources [ApiManagementIdentityProviderFacebook] depends_on.
func (amipf *ApiManagementIdentityProviderFacebook) Dependencies() terra.Dependencies {
	return amipf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementIdentityProviderFacebook].
func (amipf *ApiManagementIdentityProviderFacebook) LifecycleManagement() *terra.Lifecycle {
	return amipf.Lifecycle
}

// Attributes returns the attributes for [ApiManagementIdentityProviderFacebook].
func (amipf *ApiManagementIdentityProviderFacebook) Attributes() apiManagementIdentityProviderFacebookAttributes {
	return apiManagementIdentityProviderFacebookAttributes{ref: terra.ReferenceResource(amipf)}
}

// ImportState imports the given attribute values into [ApiManagementIdentityProviderFacebook]'s state.
func (amipf *ApiManagementIdentityProviderFacebook) ImportState(av io.Reader) error {
	amipf.state = &apiManagementIdentityProviderFacebookState{}
	if err := json.NewDecoder(av).Decode(amipf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amipf.Type(), amipf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementIdentityProviderFacebook] has state.
func (amipf *ApiManagementIdentityProviderFacebook) State() (*apiManagementIdentityProviderFacebookState, bool) {
	return amipf.state, amipf.state != nil
}

// StateMust returns the state for [ApiManagementIdentityProviderFacebook]. Panics if the state is nil.
func (amipf *ApiManagementIdentityProviderFacebook) StateMust() *apiManagementIdentityProviderFacebookState {
	if amipf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amipf.Type(), amipf.LocalName()))
	}
	return amipf.state
}

// ApiManagementIdentityProviderFacebookArgs contains the configurations for azurerm_api_management_identity_provider_facebook.
type ApiManagementIdentityProviderFacebookArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// AppSecret: string, required
	AppSecret terra.StringValue `hcl:"app_secret,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementidentityproviderfacebook.Timeouts `hcl:"timeouts,block"`
}
type apiManagementIdentityProviderFacebookAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_identity_provider_facebook.
func (amipf apiManagementIdentityProviderFacebookAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amipf.ref.Append("api_management_name"))
}

// AppId returns a reference to field app_id of azurerm_api_management_identity_provider_facebook.
func (amipf apiManagementIdentityProviderFacebookAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(amipf.ref.Append("app_id"))
}

// AppSecret returns a reference to field app_secret of azurerm_api_management_identity_provider_facebook.
func (amipf apiManagementIdentityProviderFacebookAttributes) AppSecret() terra.StringValue {
	return terra.ReferenceAsString(amipf.ref.Append("app_secret"))
}

// Id returns a reference to field id of azurerm_api_management_identity_provider_facebook.
func (amipf apiManagementIdentityProviderFacebookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amipf.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_identity_provider_facebook.
func (amipf apiManagementIdentityProviderFacebookAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amipf.ref.Append("resource_group_name"))
}

func (amipf apiManagementIdentityProviderFacebookAttributes) Timeouts() apimanagementidentityproviderfacebook.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementidentityproviderfacebook.TimeoutsAttributes](amipf.ref.Append("timeouts"))
}

type apiManagementIdentityProviderFacebookState struct {
	ApiManagementName string                                               `json:"api_management_name"`
	AppId             string                                               `json:"app_id"`
	AppSecret         string                                               `json:"app_secret"`
	Id                string                                               `json:"id"`
	ResourceGroupName string                                               `json:"resource_group_name"`
	Timeouts          *apimanagementidentityproviderfacebook.TimeoutsState `json:"timeouts"`
}
