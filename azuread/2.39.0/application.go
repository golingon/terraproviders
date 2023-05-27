// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	application "github.com/golingon/terraproviders/azuread/2.39.0/application"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplication creates a new instance of [Application].
func NewApplication(name string, args ApplicationArgs) *Application {
	return &Application{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Application)(nil)

// Application represents the Terraform resource azuread_application.
type Application struct {
	Name      string
	Args      ApplicationArgs
	state     *applicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Application].
func (a *Application) Type() string {
	return "azuread_application"
}

// LocalName returns the local name for [Application].
func (a *Application) LocalName() string {
	return a.Name
}

// Configuration returns the configuration (args) for [Application].
func (a *Application) Configuration() interface{} {
	return a.Args
}

// DependOn is used for other resources to depend on [Application].
func (a *Application) DependOn() terra.Reference {
	return terra.ReferenceResource(a)
}

// Dependencies returns the list of resources [Application] depends_on.
func (a *Application) Dependencies() terra.Dependencies {
	return a.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Application].
func (a *Application) LifecycleManagement() *terra.Lifecycle {
	return a.Lifecycle
}

// Attributes returns the attributes for [Application].
func (a *Application) Attributes() applicationAttributes {
	return applicationAttributes{ref: terra.ReferenceResource(a)}
}

// ImportState imports the given attribute values into [Application]'s state.
func (a *Application) ImportState(av io.Reader) error {
	a.state = &applicationState{}
	if err := json.NewDecoder(av).Decode(a.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", a.Type(), a.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Application] has state.
func (a *Application) State() (*applicationState, bool) {
	return a.state, a.state != nil
}

// StateMust returns the state for [Application]. Panics if the state is nil.
func (a *Application) StateMust() *applicationState {
	if a.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", a.Type(), a.LocalName()))
	}
	return a.state
}

// ApplicationArgs contains the configurations for azuread_application.
type ApplicationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DeviceOnlyAuthEnabled: bool, optional
	DeviceOnlyAuthEnabled terra.BoolValue `hcl:"device_only_auth_enabled,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// FallbackPublicClientEnabled: bool, optional
	FallbackPublicClientEnabled terra.BoolValue `hcl:"fallback_public_client_enabled,attr"`
	// GroupMembershipClaims: set of string, optional
	GroupMembershipClaims terra.SetValue[terra.StringValue] `hcl:"group_membership_claims,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentifierUris: set of string, optional
	IdentifierUris terra.SetValue[terra.StringValue] `hcl:"identifier_uris,attr"`
	// LogoImage: string, optional
	LogoImage terra.StringValue `hcl:"logo_image,attr"`
	// MarketingUrl: string, optional
	MarketingUrl terra.StringValue `hcl:"marketing_url,attr"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// Oauth2PostResponseRequired: bool, optional
	Oauth2PostResponseRequired terra.BoolValue `hcl:"oauth2_post_response_required,attr"`
	// Owners: set of string, optional
	Owners terra.SetValue[terra.StringValue] `hcl:"owners,attr"`
	// PreventDuplicateNames: bool, optional
	PreventDuplicateNames terra.BoolValue `hcl:"prevent_duplicate_names,attr"`
	// PrivacyStatementUrl: string, optional
	PrivacyStatementUrl terra.StringValue `hcl:"privacy_statement_url,attr"`
	// ServiceManagementReference: string, optional
	ServiceManagementReference terra.StringValue `hcl:"service_management_reference,attr"`
	// SignInAudience: string, optional
	SignInAudience terra.StringValue `hcl:"sign_in_audience,attr"`
	// SupportUrl: string, optional
	SupportUrl terra.StringValue `hcl:"support_url,attr"`
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// TemplateId: string, optional
	TemplateId terra.StringValue `hcl:"template_id,attr"`
	// TermsOfServiceUrl: string, optional
	TermsOfServiceUrl terra.StringValue `hcl:"terms_of_service_url,attr"`
	// Api: optional
	Api *application.Api `hcl:"api,block"`
	// AppRole: min=0
	AppRole []application.AppRole `hcl:"app_role,block" validate:"min=0"`
	// FeatureTags: min=0
	FeatureTags []application.FeatureTags `hcl:"feature_tags,block" validate:"min=0"`
	// OptionalClaims: optional
	OptionalClaims *application.OptionalClaims `hcl:"optional_claims,block"`
	// PublicClient: optional
	PublicClient *application.PublicClient `hcl:"public_client,block"`
	// RequiredResourceAccess: min=0
	RequiredResourceAccess []application.RequiredResourceAccess `hcl:"required_resource_access,block" validate:"min=0"`
	// SinglePageApplication: optional
	SinglePageApplication *application.SinglePageApplication `hcl:"single_page_application,block"`
	// Timeouts: optional
	Timeouts *application.Timeouts `hcl:"timeouts,block"`
	// Web: optional
	Web *application.Web `hcl:"web,block"`
}
type applicationAttributes struct {
	ref terra.Reference
}

// AppRoleIds returns a reference to field app_role_ids of azuread_application.
func (a applicationAttributes) AppRoleIds() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("app_role_ids"))
}

// ApplicationId returns a reference to field application_id of azuread_application.
func (a applicationAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("application_id"))
}

// Description returns a reference to field description of azuread_application.
func (a applicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("description"))
}

// DeviceOnlyAuthEnabled returns a reference to field device_only_auth_enabled of azuread_application.
func (a applicationAttributes) DeviceOnlyAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("device_only_auth_enabled"))
}

// DisabledByMicrosoft returns a reference to field disabled_by_microsoft of azuread_application.
func (a applicationAttributes) DisabledByMicrosoft() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("disabled_by_microsoft"))
}

// DisplayName returns a reference to field display_name of azuread_application.
func (a applicationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("display_name"))
}

// FallbackPublicClientEnabled returns a reference to field fallback_public_client_enabled of azuread_application.
func (a applicationAttributes) FallbackPublicClientEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("fallback_public_client_enabled"))
}

// GroupMembershipClaims returns a reference to field group_membership_claims of azuread_application.
func (a applicationAttributes) GroupMembershipClaims() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("group_membership_claims"))
}

// Id returns a reference to field id of azuread_application.
func (a applicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

// IdentifierUris returns a reference to field identifier_uris of azuread_application.
func (a applicationAttributes) IdentifierUris() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("identifier_uris"))
}

// LogoImage returns a reference to field logo_image of azuread_application.
func (a applicationAttributes) LogoImage() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("logo_image"))
}

// LogoUrl returns a reference to field logo_url of azuread_application.
func (a applicationAttributes) LogoUrl() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("logo_url"))
}

// MarketingUrl returns a reference to field marketing_url of azuread_application.
func (a applicationAttributes) MarketingUrl() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("marketing_url"))
}

// Notes returns a reference to field notes of azuread_application.
func (a applicationAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("notes"))
}

// Oauth2PermissionScopeIds returns a reference to field oauth2_permission_scope_ids of azuread_application.
func (a applicationAttributes) Oauth2PermissionScopeIds() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("oauth2_permission_scope_ids"))
}

// Oauth2PostResponseRequired returns a reference to field oauth2_post_response_required of azuread_application.
func (a applicationAttributes) Oauth2PostResponseRequired() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("oauth2_post_response_required"))
}

// ObjectId returns a reference to field object_id of azuread_application.
func (a applicationAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("object_id"))
}

// Owners returns a reference to field owners of azuread_application.
func (a applicationAttributes) Owners() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("owners"))
}

// PreventDuplicateNames returns a reference to field prevent_duplicate_names of azuread_application.
func (a applicationAttributes) PreventDuplicateNames() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("prevent_duplicate_names"))
}

// PrivacyStatementUrl returns a reference to field privacy_statement_url of azuread_application.
func (a applicationAttributes) PrivacyStatementUrl() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("privacy_statement_url"))
}

// PublisherDomain returns a reference to field publisher_domain of azuread_application.
func (a applicationAttributes) PublisherDomain() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("publisher_domain"))
}

// ServiceManagementReference returns a reference to field service_management_reference of azuread_application.
func (a applicationAttributes) ServiceManagementReference() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("service_management_reference"))
}

// SignInAudience returns a reference to field sign_in_audience of azuread_application.
func (a applicationAttributes) SignInAudience() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("sign_in_audience"))
}

// SupportUrl returns a reference to field support_url of azuread_application.
func (a applicationAttributes) SupportUrl() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("support_url"))
}

// Tags returns a reference to field tags of azuread_application.
func (a applicationAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("tags"))
}

// TemplateId returns a reference to field template_id of azuread_application.
func (a applicationAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("template_id"))
}

// TermsOfServiceUrl returns a reference to field terms_of_service_url of azuread_application.
func (a applicationAttributes) TermsOfServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("terms_of_service_url"))
}

func (a applicationAttributes) Api() terra.ListValue[application.ApiAttributes] {
	return terra.ReferenceAsList[application.ApiAttributes](a.ref.Append("api"))
}

func (a applicationAttributes) AppRole() terra.SetValue[application.AppRoleAttributes] {
	return terra.ReferenceAsSet[application.AppRoleAttributes](a.ref.Append("app_role"))
}

func (a applicationAttributes) FeatureTags() terra.ListValue[application.FeatureTagsAttributes] {
	return terra.ReferenceAsList[application.FeatureTagsAttributes](a.ref.Append("feature_tags"))
}

func (a applicationAttributes) OptionalClaims() terra.ListValue[application.OptionalClaimsAttributes] {
	return terra.ReferenceAsList[application.OptionalClaimsAttributes](a.ref.Append("optional_claims"))
}

func (a applicationAttributes) PublicClient() terra.ListValue[application.PublicClientAttributes] {
	return terra.ReferenceAsList[application.PublicClientAttributes](a.ref.Append("public_client"))
}

func (a applicationAttributes) RequiredResourceAccess() terra.SetValue[application.RequiredResourceAccessAttributes] {
	return terra.ReferenceAsSet[application.RequiredResourceAccessAttributes](a.ref.Append("required_resource_access"))
}

func (a applicationAttributes) SinglePageApplication() terra.ListValue[application.SinglePageApplicationAttributes] {
	return terra.ReferenceAsList[application.SinglePageApplicationAttributes](a.ref.Append("single_page_application"))
}

func (a applicationAttributes) Timeouts() application.TimeoutsAttributes {
	return terra.ReferenceAsSingle[application.TimeoutsAttributes](a.ref.Append("timeouts"))
}

func (a applicationAttributes) Web() terra.ListValue[application.WebAttributes] {
	return terra.ReferenceAsList[application.WebAttributes](a.ref.Append("web"))
}

type applicationState struct {
	AppRoleIds                  map[string]string                         `json:"app_role_ids"`
	ApplicationId               string                                    `json:"application_id"`
	Description                 string                                    `json:"description"`
	DeviceOnlyAuthEnabled       bool                                      `json:"device_only_auth_enabled"`
	DisabledByMicrosoft         string                                    `json:"disabled_by_microsoft"`
	DisplayName                 string                                    `json:"display_name"`
	FallbackPublicClientEnabled bool                                      `json:"fallback_public_client_enabled"`
	GroupMembershipClaims       []string                                  `json:"group_membership_claims"`
	Id                          string                                    `json:"id"`
	IdentifierUris              []string                                  `json:"identifier_uris"`
	LogoImage                   string                                    `json:"logo_image"`
	LogoUrl                     string                                    `json:"logo_url"`
	MarketingUrl                string                                    `json:"marketing_url"`
	Notes                       string                                    `json:"notes"`
	Oauth2PermissionScopeIds    map[string]string                         `json:"oauth2_permission_scope_ids"`
	Oauth2PostResponseRequired  bool                                      `json:"oauth2_post_response_required"`
	ObjectId                    string                                    `json:"object_id"`
	Owners                      []string                                  `json:"owners"`
	PreventDuplicateNames       bool                                      `json:"prevent_duplicate_names"`
	PrivacyStatementUrl         string                                    `json:"privacy_statement_url"`
	PublisherDomain             string                                    `json:"publisher_domain"`
	ServiceManagementReference  string                                    `json:"service_management_reference"`
	SignInAudience              string                                    `json:"sign_in_audience"`
	SupportUrl                  string                                    `json:"support_url"`
	Tags                        []string                                  `json:"tags"`
	TemplateId                  string                                    `json:"template_id"`
	TermsOfServiceUrl           string                                    `json:"terms_of_service_url"`
	Api                         []application.ApiState                    `json:"api"`
	AppRole                     []application.AppRoleState                `json:"app_role"`
	FeatureTags                 []application.FeatureTagsState            `json:"feature_tags"`
	OptionalClaims              []application.OptionalClaimsState         `json:"optional_claims"`
	PublicClient                []application.PublicClientState           `json:"public_client"`
	RequiredResourceAccess      []application.RequiredResourceAccessState `json:"required_resource_access"`
	SinglePageApplication       []application.SinglePageApplicationState  `json:"single_page_application"`
	Timeouts                    *application.TimeoutsState                `json:"timeouts"`
	Web                         []application.WebState                    `json:"web"`
}
