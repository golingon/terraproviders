// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementauthorizationserver "github.com/golingon/terraproviders/azurerm/3.64.0/apimanagementauthorizationserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementAuthorizationServer creates a new instance of [ApiManagementAuthorizationServer].
func NewApiManagementAuthorizationServer(name string, args ApiManagementAuthorizationServerArgs) *ApiManagementAuthorizationServer {
	return &ApiManagementAuthorizationServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementAuthorizationServer)(nil)

// ApiManagementAuthorizationServer represents the Terraform resource azurerm_api_management_authorization_server.
type ApiManagementAuthorizationServer struct {
	Name      string
	Args      ApiManagementAuthorizationServerArgs
	state     *apiManagementAuthorizationServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementAuthorizationServer].
func (amas *ApiManagementAuthorizationServer) Type() string {
	return "azurerm_api_management_authorization_server"
}

// LocalName returns the local name for [ApiManagementAuthorizationServer].
func (amas *ApiManagementAuthorizationServer) LocalName() string {
	return amas.Name
}

// Configuration returns the configuration (args) for [ApiManagementAuthorizationServer].
func (amas *ApiManagementAuthorizationServer) Configuration() interface{} {
	return amas.Args
}

// DependOn is used for other resources to depend on [ApiManagementAuthorizationServer].
func (amas *ApiManagementAuthorizationServer) DependOn() terra.Reference {
	return terra.ReferenceResource(amas)
}

// Dependencies returns the list of resources [ApiManagementAuthorizationServer] depends_on.
func (amas *ApiManagementAuthorizationServer) Dependencies() terra.Dependencies {
	return amas.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementAuthorizationServer].
func (amas *ApiManagementAuthorizationServer) LifecycleManagement() *terra.Lifecycle {
	return amas.Lifecycle
}

// Attributes returns the attributes for [ApiManagementAuthorizationServer].
func (amas *ApiManagementAuthorizationServer) Attributes() apiManagementAuthorizationServerAttributes {
	return apiManagementAuthorizationServerAttributes{ref: terra.ReferenceResource(amas)}
}

// ImportState imports the given attribute values into [ApiManagementAuthorizationServer]'s state.
func (amas *ApiManagementAuthorizationServer) ImportState(av io.Reader) error {
	amas.state = &apiManagementAuthorizationServerState{}
	if err := json.NewDecoder(av).Decode(amas.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amas.Type(), amas.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementAuthorizationServer] has state.
func (amas *ApiManagementAuthorizationServer) State() (*apiManagementAuthorizationServerState, bool) {
	return amas.state, amas.state != nil
}

// StateMust returns the state for [ApiManagementAuthorizationServer]. Panics if the state is nil.
func (amas *ApiManagementAuthorizationServer) StateMust() *apiManagementAuthorizationServerState {
	if amas.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amas.Type(), amas.LocalName()))
	}
	return amas.state
}

// ApiManagementAuthorizationServerArgs contains the configurations for azurerm_api_management_authorization_server.
type ApiManagementAuthorizationServerArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// AuthorizationEndpoint: string, required
	AuthorizationEndpoint terra.StringValue `hcl:"authorization_endpoint,attr" validate:"required"`
	// AuthorizationMethods: set of string, required
	AuthorizationMethods terra.SetValue[terra.StringValue] `hcl:"authorization_methods,attr" validate:"required"`
	// BearerTokenSendingMethods: set of string, optional
	BearerTokenSendingMethods terra.SetValue[terra.StringValue] `hcl:"bearer_token_sending_methods,attr"`
	// ClientAuthenticationMethod: set of string, optional
	ClientAuthenticationMethod terra.SetValue[terra.StringValue] `hcl:"client_authentication_method,attr"`
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientRegistrationEndpoint: string, required
	ClientRegistrationEndpoint terra.StringValue `hcl:"client_registration_endpoint,attr" validate:"required"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// DefaultScope: string, optional
	DefaultScope terra.StringValue `hcl:"default_scope,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// GrantTypes: set of string, required
	GrantTypes terra.SetValue[terra.StringValue] `hcl:"grant_types,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ResourceOwnerPassword: string, optional
	ResourceOwnerPassword terra.StringValue `hcl:"resource_owner_password,attr"`
	// ResourceOwnerUsername: string, optional
	ResourceOwnerUsername terra.StringValue `hcl:"resource_owner_username,attr"`
	// SupportState: bool, optional
	SupportState terra.BoolValue `hcl:"support_state,attr"`
	// TokenEndpoint: string, optional
	TokenEndpoint terra.StringValue `hcl:"token_endpoint,attr"`
	// Timeouts: optional
	Timeouts *apimanagementauthorizationserver.Timeouts `hcl:"timeouts,block"`
	// TokenBodyParameter: min=0
	TokenBodyParameter []apimanagementauthorizationserver.TokenBodyParameter `hcl:"token_body_parameter,block" validate:"min=0"`
}
type apiManagementAuthorizationServerAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("api_management_name"))
}

// AuthorizationEndpoint returns a reference to field authorization_endpoint of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) AuthorizationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("authorization_endpoint"))
}

// AuthorizationMethods returns a reference to field authorization_methods of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) AuthorizationMethods() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amas.ref.Append("authorization_methods"))
}

// BearerTokenSendingMethods returns a reference to field bearer_token_sending_methods of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) BearerTokenSendingMethods() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amas.ref.Append("bearer_token_sending_methods"))
}

// ClientAuthenticationMethod returns a reference to field client_authentication_method of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ClientAuthenticationMethod() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amas.ref.Append("client_authentication_method"))
}

// ClientId returns a reference to field client_id of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("client_id"))
}

// ClientRegistrationEndpoint returns a reference to field client_registration_endpoint of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ClientRegistrationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("client_registration_endpoint"))
}

// ClientSecret returns a reference to field client_secret of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("client_secret"))
}

// DefaultScope returns a reference to field default_scope of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) DefaultScope() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("default_scope"))
}

// Description returns a reference to field description of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("display_name"))
}

// GrantTypes returns a reference to field grant_types of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) GrantTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amas.ref.Append("grant_types"))
}

// Id returns a reference to field id of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("resource_group_name"))
}

// ResourceOwnerPassword returns a reference to field resource_owner_password of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ResourceOwnerPassword() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("resource_owner_password"))
}

// ResourceOwnerUsername returns a reference to field resource_owner_username of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) ResourceOwnerUsername() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("resource_owner_username"))
}

// SupportState returns a reference to field support_state of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) SupportState() terra.BoolValue {
	return terra.ReferenceAsBool(amas.ref.Append("support_state"))
}

// TokenEndpoint returns a reference to field token_endpoint of azurerm_api_management_authorization_server.
func (amas apiManagementAuthorizationServerAttributes) TokenEndpoint() terra.StringValue {
	return terra.ReferenceAsString(amas.ref.Append("token_endpoint"))
}

func (amas apiManagementAuthorizationServerAttributes) Timeouts() apimanagementauthorizationserver.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementauthorizationserver.TimeoutsAttributes](amas.ref.Append("timeouts"))
}

func (amas apiManagementAuthorizationServerAttributes) TokenBodyParameter() terra.ListValue[apimanagementauthorizationserver.TokenBodyParameterAttributes] {
	return terra.ReferenceAsList[apimanagementauthorizationserver.TokenBodyParameterAttributes](amas.ref.Append("token_body_parameter"))
}

type apiManagementAuthorizationServerState struct {
	ApiManagementName          string                                                     `json:"api_management_name"`
	AuthorizationEndpoint      string                                                     `json:"authorization_endpoint"`
	AuthorizationMethods       []string                                                   `json:"authorization_methods"`
	BearerTokenSendingMethods  []string                                                   `json:"bearer_token_sending_methods"`
	ClientAuthenticationMethod []string                                                   `json:"client_authentication_method"`
	ClientId                   string                                                     `json:"client_id"`
	ClientRegistrationEndpoint string                                                     `json:"client_registration_endpoint"`
	ClientSecret               string                                                     `json:"client_secret"`
	DefaultScope               string                                                     `json:"default_scope"`
	Description                string                                                     `json:"description"`
	DisplayName                string                                                     `json:"display_name"`
	GrantTypes                 []string                                                   `json:"grant_types"`
	Id                         string                                                     `json:"id"`
	Name                       string                                                     `json:"name"`
	ResourceGroupName          string                                                     `json:"resource_group_name"`
	ResourceOwnerPassword      string                                                     `json:"resource_owner_password"`
	ResourceOwnerUsername      string                                                     `json:"resource_owner_username"`
	SupportState               bool                                                       `json:"support_state"`
	TokenEndpoint              string                                                     `json:"token_endpoint"`
	Timeouts                   *apimanagementauthorizationserver.TimeoutsState            `json:"timeouts"`
	TokenBodyParameter         []apimanagementauthorizationserver.TokenBodyParameterState `json:"token_body_parameter"`
}