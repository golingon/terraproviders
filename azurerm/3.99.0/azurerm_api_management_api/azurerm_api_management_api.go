// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_api

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

// Resource represents the Terraform resource azurerm_api_management_api.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApiManagementApiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aama *Resource) Type() string {
	return "azurerm_api_management_api"
}

// LocalName returns the local name for [Resource].
func (aama *Resource) LocalName() string {
	return aama.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aama *Resource) Configuration() interface{} {
	return aama.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aama *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aama)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aama *Resource) Dependencies() terra.Dependencies {
	return aama.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aama *Resource) LifecycleManagement() *terra.Lifecycle {
	return aama.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aama *Resource) Attributes() azurermApiManagementApiAttributes {
	return azurermApiManagementApiAttributes{ref: terra.ReferenceResource(aama)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aama *Resource) ImportState(state io.Reader) error {
	aama.state = &azurermApiManagementApiState{}
	if err := json.NewDecoder(state).Decode(aama.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aama.Type(), aama.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aama *Resource) State() (*azurermApiManagementApiState, bool) {
	return aama.state, aama.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aama *Resource) StateMust() *azurermApiManagementApiState {
	if aama.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aama.Type(), aama.LocalName()))
	}
	return aama.state
}

// Args contains the configurations for azurerm_api_management_api.
type Args struct {
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
	Contact *Contact `hcl:"contact,block"`
	// Import: optional
	Import *Import `hcl:"import,block"`
	// License: optional
	License *License `hcl:"license,block"`
	// Oauth2Authorization: optional
	Oauth2Authorization *Oauth2Authorization `hcl:"oauth2_authorization,block"`
	// OpenidAuthentication: optional
	OpenidAuthentication *OpenidAuthentication `hcl:"openid_authentication,block"`
	// SubscriptionKeyParameterNames: optional
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNames `hcl:"subscription_key_parameter_names,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApiManagementApiAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("api_management_name"))
}

// ApiType returns a reference to field api_type of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) ApiType() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("api_type"))
}

// Description returns a reference to field description of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("id"))
}

// IsCurrent returns a reference to field is_current of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) IsCurrent() terra.BoolValue {
	return terra.ReferenceAsBool(aama.ref.Append("is_current"))
}

// IsOnline returns a reference to field is_online of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) IsOnline() terra.BoolValue {
	return terra.ReferenceAsBool(aama.ref.Append("is_online"))
}

// Name returns a reference to field name of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("name"))
}

// Path returns a reference to field path of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("path"))
}

// Protocols returns a reference to field protocols of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Protocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aama.ref.Append("protocols"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("resource_group_name"))
}

// Revision returns a reference to field revision of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Revision() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("revision"))
}

// RevisionDescription returns a reference to field revision_description of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) RevisionDescription() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("revision_description"))
}

// ServiceUrl returns a reference to field service_url of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) ServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("service_url"))
}

// SoapPassThrough returns a reference to field soap_pass_through of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) SoapPassThrough() terra.BoolValue {
	return terra.ReferenceAsBool(aama.ref.Append("soap_pass_through"))
}

// SourceApiId returns a reference to field source_api_id of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) SourceApiId() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("source_api_id"))
}

// SubscriptionRequired returns a reference to field subscription_required of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(aama.ref.Append("subscription_required"))
}

// TermsOfServiceUrl returns a reference to field terms_of_service_url of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) TermsOfServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("terms_of_service_url"))
}

// Version returns a reference to field version of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("version"))
}

// VersionDescription returns a reference to field version_description of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) VersionDescription() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("version_description"))
}

// VersionSetId returns a reference to field version_set_id of azurerm_api_management_api.
func (aama azurermApiManagementApiAttributes) VersionSetId() terra.StringValue {
	return terra.ReferenceAsString(aama.ref.Append("version_set_id"))
}

func (aama azurermApiManagementApiAttributes) Contact() terra.ListValue[ContactAttributes] {
	return terra.ReferenceAsList[ContactAttributes](aama.ref.Append("contact"))
}

func (aama azurermApiManagementApiAttributes) Import() terra.ListValue[ImportAttributes] {
	return terra.ReferenceAsList[ImportAttributes](aama.ref.Append("import"))
}

func (aama azurermApiManagementApiAttributes) License() terra.ListValue[LicenseAttributes] {
	return terra.ReferenceAsList[LicenseAttributes](aama.ref.Append("license"))
}

func (aama azurermApiManagementApiAttributes) Oauth2Authorization() terra.ListValue[Oauth2AuthorizationAttributes] {
	return terra.ReferenceAsList[Oauth2AuthorizationAttributes](aama.ref.Append("oauth2_authorization"))
}

func (aama azurermApiManagementApiAttributes) OpenidAuthentication() terra.ListValue[OpenidAuthenticationAttributes] {
	return terra.ReferenceAsList[OpenidAuthenticationAttributes](aama.ref.Append("openid_authentication"))
}

func (aama azurermApiManagementApiAttributes) SubscriptionKeyParameterNames() terra.ListValue[SubscriptionKeyParameterNamesAttributes] {
	return terra.ReferenceAsList[SubscriptionKeyParameterNamesAttributes](aama.ref.Append("subscription_key_parameter_names"))
}

func (aama azurermApiManagementApiAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aama.ref.Append("timeouts"))
}

type azurermApiManagementApiState struct {
	ApiManagementName             string                               `json:"api_management_name"`
	ApiType                       string                               `json:"api_type"`
	Description                   string                               `json:"description"`
	DisplayName                   string                               `json:"display_name"`
	Id                            string                               `json:"id"`
	IsCurrent                     bool                                 `json:"is_current"`
	IsOnline                      bool                                 `json:"is_online"`
	Name                          string                               `json:"name"`
	Path                          string                               `json:"path"`
	Protocols                     []string                             `json:"protocols"`
	ResourceGroupName             string                               `json:"resource_group_name"`
	Revision                      string                               `json:"revision"`
	RevisionDescription           string                               `json:"revision_description"`
	ServiceUrl                    string                               `json:"service_url"`
	SoapPassThrough               bool                                 `json:"soap_pass_through"`
	SourceApiId                   string                               `json:"source_api_id"`
	SubscriptionRequired          bool                                 `json:"subscription_required"`
	TermsOfServiceUrl             string                               `json:"terms_of_service_url"`
	Version                       string                               `json:"version"`
	VersionDescription            string                               `json:"version_description"`
	VersionSetId                  string                               `json:"version_set_id"`
	Contact                       []ContactState                       `json:"contact"`
	Import                        []ImportState                        `json:"import"`
	License                       []LicenseState                       `json:"license"`
	Oauth2Authorization           []Oauth2AuthorizationState           `json:"oauth2_authorization"`
	OpenidAuthentication          []OpenidAuthenticationState          `json:"openid_authentication"`
	SubscriptionKeyParameterNames []SubscriptionKeyParameterNamesState `json:"subscription_key_parameter_names"`
	Timeouts                      *TimeoutsState                       `json:"timeouts"`
}
