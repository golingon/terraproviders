// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	serviceprincipal "github.com/golingon/terraproviders/azuread/2.38.0/serviceprincipal"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicePrincipal creates a new instance of [ServicePrincipal].
func NewServicePrincipal(name string, args ServicePrincipalArgs) *ServicePrincipal {
	return &ServicePrincipal{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicePrincipal)(nil)

// ServicePrincipal represents the Terraform resource azuread_service_principal.
type ServicePrincipal struct {
	Name      string
	Args      ServicePrincipalArgs
	state     *servicePrincipalState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicePrincipal].
func (sp *ServicePrincipal) Type() string {
	return "azuread_service_principal"
}

// LocalName returns the local name for [ServicePrincipal].
func (sp *ServicePrincipal) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [ServicePrincipal].
func (sp *ServicePrincipal) Configuration() interface{} {
	return sp.Args
}

// DependOn is used for other resources to depend on [ServicePrincipal].
func (sp *ServicePrincipal) DependOn() terra.Reference {
	return terra.ReferenceResource(sp)
}

// Dependencies returns the list of resources [ServicePrincipal] depends_on.
func (sp *ServicePrincipal) Dependencies() terra.Dependencies {
	return sp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicePrincipal].
func (sp *ServicePrincipal) LifecycleManagement() *terra.Lifecycle {
	return sp.Lifecycle
}

// Attributes returns the attributes for [ServicePrincipal].
func (sp *ServicePrincipal) Attributes() servicePrincipalAttributes {
	return servicePrincipalAttributes{ref: terra.ReferenceResource(sp)}
}

// ImportState imports the given attribute values into [ServicePrincipal]'s state.
func (sp *ServicePrincipal) ImportState(av io.Reader) error {
	sp.state = &servicePrincipalState{}
	if err := json.NewDecoder(av).Decode(sp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sp.Type(), sp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicePrincipal] has state.
func (sp *ServicePrincipal) State() (*servicePrincipalState, bool) {
	return sp.state, sp.state != nil
}

// StateMust returns the state for [ServicePrincipal]. Panics if the state is nil.
func (sp *ServicePrincipal) StateMust() *servicePrincipalState {
	if sp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sp.Type(), sp.LocalName()))
	}
	return sp.state
}

// ServicePrincipalArgs contains the configurations for azuread_service_principal.
type ServicePrincipalArgs struct {
	// AccountEnabled: bool, optional
	AccountEnabled terra.BoolValue `hcl:"account_enabled,attr"`
	// AlternativeNames: set of string, optional
	AlternativeNames terra.SetValue[terra.StringValue] `hcl:"alternative_names,attr"`
	// AppRoleAssignmentRequired: bool, optional
	AppRoleAssignmentRequired terra.BoolValue `hcl:"app_role_assignment_required,attr"`
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoginUrl: string, optional
	LoginUrl terra.StringValue `hcl:"login_url,attr"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// NotificationEmailAddresses: set of string, optional
	NotificationEmailAddresses terra.SetValue[terra.StringValue] `hcl:"notification_email_addresses,attr"`
	// Owners: set of string, optional
	Owners terra.SetValue[terra.StringValue] `hcl:"owners,attr"`
	// PreferredSingleSignOnMode: string, optional
	PreferredSingleSignOnMode terra.StringValue `hcl:"preferred_single_sign_on_mode,attr"`
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// UseExisting: bool, optional
	UseExisting terra.BoolValue `hcl:"use_existing,attr"`
	// AppRoles: min=0
	AppRoles []serviceprincipal.AppRoles `hcl:"app_roles,block" validate:"min=0"`
	// Oauth2PermissionScopes: min=0
	Oauth2PermissionScopes []serviceprincipal.Oauth2PermissionScopes `hcl:"oauth2_permission_scopes,block" validate:"min=0"`
	// FeatureTags: min=0
	FeatureTags []serviceprincipal.FeatureTags `hcl:"feature_tags,block" validate:"min=0"`
	// Features: min=0
	Features []serviceprincipal.Features `hcl:"features,block" validate:"min=0"`
	// SamlSingleSignOn: optional
	SamlSingleSignOn *serviceprincipal.SamlSingleSignOn `hcl:"saml_single_sign_on,block"`
	// Timeouts: optional
	Timeouts *serviceprincipal.Timeouts `hcl:"timeouts,block"`
}
type servicePrincipalAttributes struct {
	ref terra.Reference
}

// AccountEnabled returns a reference to field account_enabled of azuread_service_principal.
func (sp servicePrincipalAttributes) AccountEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("account_enabled"))
}

// AlternativeNames returns a reference to field alternative_names of azuread_service_principal.
func (sp servicePrincipalAttributes) AlternativeNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sp.ref.Append("alternative_names"))
}

// AppRoleAssignmentRequired returns a reference to field app_role_assignment_required of azuread_service_principal.
func (sp servicePrincipalAttributes) AppRoleAssignmentRequired() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("app_role_assignment_required"))
}

// AppRoleIds returns a reference to field app_role_ids of azuread_service_principal.
func (sp servicePrincipalAttributes) AppRoleIds() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("app_role_ids"))
}

// ApplicationId returns a reference to field application_id of azuread_service_principal.
func (sp servicePrincipalAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("application_id"))
}

// ApplicationTenantId returns a reference to field application_tenant_id of azuread_service_principal.
func (sp servicePrincipalAttributes) ApplicationTenantId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("application_tenant_id"))
}

// Description returns a reference to field description of azuread_service_principal.
func (sp servicePrincipalAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_service_principal.
func (sp servicePrincipalAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("display_name"))
}

// HomepageUrl returns a reference to field homepage_url of azuread_service_principal.
func (sp servicePrincipalAttributes) HomepageUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("homepage_url"))
}

// Id returns a reference to field id of azuread_service_principal.
func (sp servicePrincipalAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// LoginUrl returns a reference to field login_url of azuread_service_principal.
func (sp servicePrincipalAttributes) LoginUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("login_url"))
}

// LogoutUrl returns a reference to field logout_url of azuread_service_principal.
func (sp servicePrincipalAttributes) LogoutUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("logout_url"))
}

// Notes returns a reference to field notes of azuread_service_principal.
func (sp servicePrincipalAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("notes"))
}

// NotificationEmailAddresses returns a reference to field notification_email_addresses of azuread_service_principal.
func (sp servicePrincipalAttributes) NotificationEmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sp.ref.Append("notification_email_addresses"))
}

// Oauth2PermissionScopeIds returns a reference to field oauth2_permission_scope_ids of azuread_service_principal.
func (sp servicePrincipalAttributes) Oauth2PermissionScopeIds() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("oauth2_permission_scope_ids"))
}

// ObjectId returns a reference to field object_id of azuread_service_principal.
func (sp servicePrincipalAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("object_id"))
}

// Owners returns a reference to field owners of azuread_service_principal.
func (sp servicePrincipalAttributes) Owners() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sp.ref.Append("owners"))
}

// PreferredSingleSignOnMode returns a reference to field preferred_single_sign_on_mode of azuread_service_principal.
func (sp servicePrincipalAttributes) PreferredSingleSignOnMode() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("preferred_single_sign_on_mode"))
}

// RedirectUris returns a reference to field redirect_uris of azuread_service_principal.
func (sp servicePrincipalAttributes) RedirectUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("redirect_uris"))
}

// SamlMetadataUrl returns a reference to field saml_metadata_url of azuread_service_principal.
func (sp servicePrincipalAttributes) SamlMetadataUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("saml_metadata_url"))
}

// ServicePrincipalNames returns a reference to field service_principal_names of azuread_service_principal.
func (sp servicePrincipalAttributes) ServicePrincipalNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("service_principal_names"))
}

// SignInAudience returns a reference to field sign_in_audience of azuread_service_principal.
func (sp servicePrincipalAttributes) SignInAudience() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("sign_in_audience"))
}

// Tags returns a reference to field tags of azuread_service_principal.
func (sp servicePrincipalAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sp.ref.Append("tags"))
}

// Type returns a reference to field type of azuread_service_principal.
func (sp servicePrincipalAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("type"))
}

// UseExisting returns a reference to field use_existing of azuread_service_principal.
func (sp servicePrincipalAttributes) UseExisting() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("use_existing"))
}

func (sp servicePrincipalAttributes) AppRoles() terra.ListValue[serviceprincipal.AppRolesAttributes] {
	return terra.ReferenceAsList[serviceprincipal.AppRolesAttributes](sp.ref.Append("app_roles"))
}

func (sp servicePrincipalAttributes) Oauth2PermissionScopes() terra.ListValue[serviceprincipal.Oauth2PermissionScopesAttributes] {
	return terra.ReferenceAsList[serviceprincipal.Oauth2PermissionScopesAttributes](sp.ref.Append("oauth2_permission_scopes"))
}

func (sp servicePrincipalAttributes) FeatureTags() terra.ListValue[serviceprincipal.FeatureTagsAttributes] {
	return terra.ReferenceAsList[serviceprincipal.FeatureTagsAttributes](sp.ref.Append("feature_tags"))
}

func (sp servicePrincipalAttributes) Features() terra.ListValue[serviceprincipal.FeaturesAttributes] {
	return terra.ReferenceAsList[serviceprincipal.FeaturesAttributes](sp.ref.Append("features"))
}

func (sp servicePrincipalAttributes) SamlSingleSignOn() terra.ListValue[serviceprincipal.SamlSingleSignOnAttributes] {
	return terra.ReferenceAsList[serviceprincipal.SamlSingleSignOnAttributes](sp.ref.Append("saml_single_sign_on"))
}

func (sp servicePrincipalAttributes) Timeouts() serviceprincipal.TimeoutsAttributes {
	return terra.ReferenceAsSingle[serviceprincipal.TimeoutsAttributes](sp.ref.Append("timeouts"))
}

type servicePrincipalState struct {
	AccountEnabled             bool                                           `json:"account_enabled"`
	AlternativeNames           []string                                       `json:"alternative_names"`
	AppRoleAssignmentRequired  bool                                           `json:"app_role_assignment_required"`
	AppRoleIds                 map[string]string                              `json:"app_role_ids"`
	ApplicationId              string                                         `json:"application_id"`
	ApplicationTenantId        string                                         `json:"application_tenant_id"`
	Description                string                                         `json:"description"`
	DisplayName                string                                         `json:"display_name"`
	HomepageUrl                string                                         `json:"homepage_url"`
	Id                         string                                         `json:"id"`
	LoginUrl                   string                                         `json:"login_url"`
	LogoutUrl                  string                                         `json:"logout_url"`
	Notes                      string                                         `json:"notes"`
	NotificationEmailAddresses []string                                       `json:"notification_email_addresses"`
	Oauth2PermissionScopeIds   map[string]string                              `json:"oauth2_permission_scope_ids"`
	ObjectId                   string                                         `json:"object_id"`
	Owners                     []string                                       `json:"owners"`
	PreferredSingleSignOnMode  string                                         `json:"preferred_single_sign_on_mode"`
	RedirectUris               []string                                       `json:"redirect_uris"`
	SamlMetadataUrl            string                                         `json:"saml_metadata_url"`
	ServicePrincipalNames      []string                                       `json:"service_principal_names"`
	SignInAudience             string                                         `json:"sign_in_audience"`
	Tags                       []string                                       `json:"tags"`
	Type                       string                                         `json:"type"`
	UseExisting                bool                                           `json:"use_existing"`
	AppRoles                   []serviceprincipal.AppRolesState               `json:"app_roles"`
	Oauth2PermissionScopes     []serviceprincipal.Oauth2PermissionScopesState `json:"oauth2_permission_scopes"`
	FeatureTags                []serviceprincipal.FeatureTagsState            `json:"feature_tags"`
	Features                   []serviceprincipal.FeaturesState               `json:"features"`
	SamlSingleSignOn           []serviceprincipal.SamlSingleSignOnState       `json:"saml_single_sign_on"`
	Timeouts                   *serviceprincipal.TimeoutsState                `json:"timeouts"`
}
