// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapi "github.com/golingon/terraproviders/azurerm/3.58.0/apimanagementapi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementApi creates a new instance of [ApiManagementApi].
func NewApiManagementApi(name string, args ApiManagementApiArgs) *ApiManagementApi {
	return &ApiManagementApi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApi)(nil)

// ApiManagementApi represents the Terraform resource azurerm_api_management_api.
type ApiManagementApi struct {
	Name      string
	Args      ApiManagementApiArgs
	state     *apiManagementApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApi].
func (ama *ApiManagementApi) Type() string {
	return "azurerm_api_management_api"
}

// LocalName returns the local name for [ApiManagementApi].
func (ama *ApiManagementApi) LocalName() string {
	return ama.Name
}

// Configuration returns the configuration (args) for [ApiManagementApi].
func (ama *ApiManagementApi) Configuration() interface{} {
	return ama.Args
}

// DependOn is used for other resources to depend on [ApiManagementApi].
func (ama *ApiManagementApi) DependOn() terra.Reference {
	return terra.ReferenceResource(ama)
}

// Dependencies returns the list of resources [ApiManagementApi] depends_on.
func (ama *ApiManagementApi) Dependencies() terra.Dependencies {
	return ama.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApi].
func (ama *ApiManagementApi) LifecycleManagement() *terra.Lifecycle {
	return ama.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApi].
func (ama *ApiManagementApi) Attributes() apiManagementApiAttributes {
	return apiManagementApiAttributes{ref: terra.ReferenceResource(ama)}
}

// ImportState imports the given attribute values into [ApiManagementApi]'s state.
func (ama *ApiManagementApi) ImportState(av io.Reader) error {
	ama.state = &apiManagementApiState{}
	if err := json.NewDecoder(av).Decode(ama.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ama.Type(), ama.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApi] has state.
func (ama *ApiManagementApi) State() (*apiManagementApiState, bool) {
	return ama.state, ama.state != nil
}

// StateMust returns the state for [ApiManagementApi]. Panics if the state is nil.
func (ama *ApiManagementApi) StateMust() *apiManagementApiState {
	if ama.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ama.Type(), ama.LocalName()))
	}
	return ama.state
}

// ApiManagementApiArgs contains the configurations for azurerm_api_management_api.
type ApiManagementApiArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// ApiType: string, optional
	ApiType terra.StringValue `hcl:"api_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Protocols: set of string, optional
	Protocols terra.SetValue[terra.StringValue] `hcl:"protocols,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Revision: string, required
	Revision terra.StringValue `hcl:"revision,attr" validate:"required"`
	// RevisionDescription: string, optional
	RevisionDescription terra.StringValue `hcl:"revision_description,attr"`
	// ServiceUrl: string, optional
	ServiceUrl terra.StringValue `hcl:"service_url,attr"`
	// SoapPassThrough: bool, optional
	SoapPassThrough terra.BoolValue `hcl:"soap_pass_through,attr"`
	// SourceApiId: string, optional
	SourceApiId terra.StringValue `hcl:"source_api_id,attr"`
	// SubscriptionRequired: bool, optional
	SubscriptionRequired terra.BoolValue `hcl:"subscription_required,attr"`
	// TermsOfServiceUrl: string, optional
	TermsOfServiceUrl terra.StringValue `hcl:"terms_of_service_url,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// VersionDescription: string, optional
	VersionDescription terra.StringValue `hcl:"version_description,attr"`
	// VersionSetId: string, optional
	VersionSetId terra.StringValue `hcl:"version_set_id,attr"`
	// Contact: optional
	Contact *apimanagementapi.Contact `hcl:"contact,block"`
	// Import: optional
	Import *apimanagementapi.Import `hcl:"import,block"`
	// License: optional
	License *apimanagementapi.License `hcl:"license,block"`
	// Oauth2Authorization: optional
	Oauth2Authorization *apimanagementapi.Oauth2Authorization `hcl:"oauth2_authorization,block"`
	// OpenidAuthentication: optional
	OpenidAuthentication *apimanagementapi.OpenidAuthentication `hcl:"openid_authentication,block"`
	// SubscriptionKeyParameterNames: optional
	SubscriptionKeyParameterNames *apimanagementapi.SubscriptionKeyParameterNames `hcl:"subscription_key_parameter_names,block"`
	// Timeouts: optional
	Timeouts *apimanagementapi.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api.
func (ama apiManagementApiAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("api_management_name"))
}

// ApiType returns a reference to field api_type of azurerm_api_management_api.
func (ama apiManagementApiAttributes) ApiType() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("api_type"))
}

// Description returns a reference to field description of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api.
func (ama apiManagementApiAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("id"))
}

// IsCurrent returns a reference to field is_current of azurerm_api_management_api.
func (ama apiManagementApiAttributes) IsCurrent() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("is_current"))
}

// IsOnline returns a reference to field is_online of azurerm_api_management_api.
func (ama apiManagementApiAttributes) IsOnline() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("is_online"))
}

// Name returns a reference to field name of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("name"))
}

// Path returns a reference to field path of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("path"))
}

// Protocols returns a reference to field protocols of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Protocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ama.ref.Append("protocols"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api.
func (ama apiManagementApiAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("resource_group_name"))
}

// Revision returns a reference to field revision of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Revision() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("revision"))
}

// RevisionDescription returns a reference to field revision_description of azurerm_api_management_api.
func (ama apiManagementApiAttributes) RevisionDescription() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("revision_description"))
}

// ServiceUrl returns a reference to field service_url of azurerm_api_management_api.
func (ama apiManagementApiAttributes) ServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("service_url"))
}

// SoapPassThrough returns a reference to field soap_pass_through of azurerm_api_management_api.
func (ama apiManagementApiAttributes) SoapPassThrough() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("soap_pass_through"))
}

// SourceApiId returns a reference to field source_api_id of azurerm_api_management_api.
func (ama apiManagementApiAttributes) SourceApiId() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("source_api_id"))
}

// SubscriptionRequired returns a reference to field subscription_required of azurerm_api_management_api.
func (ama apiManagementApiAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("subscription_required"))
}

// TermsOfServiceUrl returns a reference to field terms_of_service_url of azurerm_api_management_api.
func (ama apiManagementApiAttributes) TermsOfServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("terms_of_service_url"))
}

// Version returns a reference to field version of azurerm_api_management_api.
func (ama apiManagementApiAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("version"))
}

// VersionDescription returns a reference to field version_description of azurerm_api_management_api.
func (ama apiManagementApiAttributes) VersionDescription() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("version_description"))
}

// VersionSetId returns a reference to field version_set_id of azurerm_api_management_api.
func (ama apiManagementApiAttributes) VersionSetId() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("version_set_id"))
}

func (ama apiManagementApiAttributes) Contact() terra.ListValue[apimanagementapi.ContactAttributes] {
	return terra.ReferenceAsList[apimanagementapi.ContactAttributes](ama.ref.Append("contact"))
}

func (ama apiManagementApiAttributes) Import() terra.ListValue[apimanagementapi.ImportAttributes] {
	return terra.ReferenceAsList[apimanagementapi.ImportAttributes](ama.ref.Append("import"))
}

func (ama apiManagementApiAttributes) License() terra.ListValue[apimanagementapi.LicenseAttributes] {
	return terra.ReferenceAsList[apimanagementapi.LicenseAttributes](ama.ref.Append("license"))
}

func (ama apiManagementApiAttributes) Oauth2Authorization() terra.ListValue[apimanagementapi.Oauth2AuthorizationAttributes] {
	return terra.ReferenceAsList[apimanagementapi.Oauth2AuthorizationAttributes](ama.ref.Append("oauth2_authorization"))
}

func (ama apiManagementApiAttributes) OpenidAuthentication() terra.ListValue[apimanagementapi.OpenidAuthenticationAttributes] {
	return terra.ReferenceAsList[apimanagementapi.OpenidAuthenticationAttributes](ama.ref.Append("openid_authentication"))
}

func (ama apiManagementApiAttributes) SubscriptionKeyParameterNames() terra.ListValue[apimanagementapi.SubscriptionKeyParameterNamesAttributes] {
	return terra.ReferenceAsList[apimanagementapi.SubscriptionKeyParameterNamesAttributes](ama.ref.Append("subscription_key_parameter_names"))
}

func (ama apiManagementApiAttributes) Timeouts() apimanagementapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapi.TimeoutsAttributes](ama.ref.Append("timeouts"))
}

type apiManagementApiState struct {
	ApiManagementName             string                                                `json:"api_management_name"`
	ApiType                       string                                                `json:"api_type"`
	Description                   string                                                `json:"description"`
	DisplayName                   string                                                `json:"display_name"`
	Id                            string                                                `json:"id"`
	IsCurrent                     bool                                                  `json:"is_current"`
	IsOnline                      bool                                                  `json:"is_online"`
	Name                          string                                                `json:"name"`
	Path                          string                                                `json:"path"`
	Protocols                     []string                                              `json:"protocols"`
	ResourceGroupName             string                                                `json:"resource_group_name"`
	Revision                      string                                                `json:"revision"`
	RevisionDescription           string                                                `json:"revision_description"`
	ServiceUrl                    string                                                `json:"service_url"`
	SoapPassThrough               bool                                                  `json:"soap_pass_through"`
	SourceApiId                   string                                                `json:"source_api_id"`
	SubscriptionRequired          bool                                                  `json:"subscription_required"`
	TermsOfServiceUrl             string                                                `json:"terms_of_service_url"`
	Version                       string                                                `json:"version"`
	VersionDescription            string                                                `json:"version_description"`
	VersionSetId                  string                                                `json:"version_set_id"`
	Contact                       []apimanagementapi.ContactState                       `json:"contact"`
	Import                        []apimanagementapi.ImportState                        `json:"import"`
	License                       []apimanagementapi.LicenseState                       `json:"license"`
	Oauth2Authorization           []apimanagementapi.Oauth2AuthorizationState           `json:"oauth2_authorization"`
	OpenidAuthentication          []apimanagementapi.OpenidAuthenticationState          `json:"openid_authentication"`
	SubscriptionKeyParameterNames []apimanagementapi.SubscriptionKeyParameterNamesState `json:"subscription_key_parameter_names"`
	Timeouts                      *apimanagementapi.TimeoutsState                       `json:"timeouts"`
}
