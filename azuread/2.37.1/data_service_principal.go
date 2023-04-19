// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	dataserviceprincipal "github.com/golingon/terraproviders/azuread/2.37.1/dataserviceprincipal"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicePrincipal creates a new instance of [DataServicePrincipal].
func NewDataServicePrincipal(name string, args DataServicePrincipalArgs) *DataServicePrincipal {
	return &DataServicePrincipal{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicePrincipal)(nil)

// DataServicePrincipal represents the Terraform data resource azuread_service_principal.
type DataServicePrincipal struct {
	Name string
	Args DataServicePrincipalArgs
}

// DataSource returns the Terraform object type for [DataServicePrincipal].
func (sp *DataServicePrincipal) DataSource() string {
	return "azuread_service_principal"
}

// LocalName returns the local name for [DataServicePrincipal].
func (sp *DataServicePrincipal) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataServicePrincipal].
func (sp *DataServicePrincipal) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataServicePrincipal].
func (sp *DataServicePrincipal) Attributes() dataServicePrincipalAttributes {
	return dataServicePrincipalAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataServicePrincipalArgs contains the configurations for azuread_service_principal.
type DataServicePrincipalArgs struct {
	// ApplicationId: string, optional
	ApplicationId terra.StringValue `hcl:"application_id,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// AppRoles: min=0
	AppRoles []dataserviceprincipal.AppRoles `hcl:"app_roles,block" validate:"min=0"`
	// FeatureTags: min=0
	FeatureTags []dataserviceprincipal.FeatureTags `hcl:"feature_tags,block" validate:"min=0"`
	// Features: min=0
	Features []dataserviceprincipal.Features `hcl:"features,block" validate:"min=0"`
	// Oauth2PermissionScopes: min=0
	Oauth2PermissionScopes []dataserviceprincipal.Oauth2PermissionScopes `hcl:"oauth2_permission_scopes,block" validate:"min=0"`
	// SamlSingleSignOn: min=0
	SamlSingleSignOn []dataserviceprincipal.SamlSingleSignOn `hcl:"saml_single_sign_on,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataserviceprincipal.Timeouts `hcl:"timeouts,block"`
}
type dataServicePrincipalAttributes struct {
	ref terra.Reference
}

// AccountEnabled returns a reference to field account_enabled of azuread_service_principal.
func (sp dataServicePrincipalAttributes) AccountEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("account_enabled"))
}

// AlternativeNames returns a reference to field alternative_names of azuread_service_principal.
func (sp dataServicePrincipalAttributes) AlternativeNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("alternative_names"))
}

// AppRoleAssignmentRequired returns a reference to field app_role_assignment_required of azuread_service_principal.
func (sp dataServicePrincipalAttributes) AppRoleAssignmentRequired() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("app_role_assignment_required"))
}

// AppRoleIds returns a reference to field app_role_ids of azuread_service_principal.
func (sp dataServicePrincipalAttributes) AppRoleIds() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("app_role_ids"))
}

// ApplicationId returns a reference to field application_id of azuread_service_principal.
func (sp dataServicePrincipalAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("application_id"))
}

// ApplicationTenantId returns a reference to field application_tenant_id of azuread_service_principal.
func (sp dataServicePrincipalAttributes) ApplicationTenantId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("application_tenant_id"))
}

// Description returns a reference to field description of azuread_service_principal.
func (sp dataServicePrincipalAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_service_principal.
func (sp dataServicePrincipalAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("display_name"))
}

// HomepageUrl returns a reference to field homepage_url of azuread_service_principal.
func (sp dataServicePrincipalAttributes) HomepageUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("homepage_url"))
}

// Id returns a reference to field id of azuread_service_principal.
func (sp dataServicePrincipalAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

// LoginUrl returns a reference to field login_url of azuread_service_principal.
func (sp dataServicePrincipalAttributes) LoginUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("login_url"))
}

// LogoutUrl returns a reference to field logout_url of azuread_service_principal.
func (sp dataServicePrincipalAttributes) LogoutUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("logout_url"))
}

// Notes returns a reference to field notes of azuread_service_principal.
func (sp dataServicePrincipalAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("notes"))
}

// NotificationEmailAddresses returns a reference to field notification_email_addresses of azuread_service_principal.
func (sp dataServicePrincipalAttributes) NotificationEmailAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("notification_email_addresses"))
}

// Oauth2PermissionScopeIds returns a reference to field oauth2_permission_scope_ids of azuread_service_principal.
func (sp dataServicePrincipalAttributes) Oauth2PermissionScopeIds() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sp.ref.Append("oauth2_permission_scope_ids"))
}

// ObjectId returns a reference to field object_id of azuread_service_principal.
func (sp dataServicePrincipalAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("object_id"))
}

// PreferredSingleSignOnMode returns a reference to field preferred_single_sign_on_mode of azuread_service_principal.
func (sp dataServicePrincipalAttributes) PreferredSingleSignOnMode() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("preferred_single_sign_on_mode"))
}

// RedirectUris returns a reference to field redirect_uris of azuread_service_principal.
func (sp dataServicePrincipalAttributes) RedirectUris() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("redirect_uris"))
}

// SamlMetadataUrl returns a reference to field saml_metadata_url of azuread_service_principal.
func (sp dataServicePrincipalAttributes) SamlMetadataUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("saml_metadata_url"))
}

// ServicePrincipalNames returns a reference to field service_principal_names of azuread_service_principal.
func (sp dataServicePrincipalAttributes) ServicePrincipalNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("service_principal_names"))
}

// SignInAudience returns a reference to field sign_in_audience of azuread_service_principal.
func (sp dataServicePrincipalAttributes) SignInAudience() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("sign_in_audience"))
}

// Tags returns a reference to field tags of azuread_service_principal.
func (sp dataServicePrincipalAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("tags"))
}

// Type returns a reference to field type of azuread_service_principal.
func (sp dataServicePrincipalAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("type"))
}

func (sp dataServicePrincipalAttributes) AppRoles() terra.ListValue[dataserviceprincipal.AppRolesAttributes] {
	return terra.ReferenceAsList[dataserviceprincipal.AppRolesAttributes](sp.ref.Append("app_roles"))
}

func (sp dataServicePrincipalAttributes) FeatureTags() terra.ListValue[dataserviceprincipal.FeatureTagsAttributes] {
	return terra.ReferenceAsList[dataserviceprincipal.FeatureTagsAttributes](sp.ref.Append("feature_tags"))
}

func (sp dataServicePrincipalAttributes) Features() terra.ListValue[dataserviceprincipal.FeaturesAttributes] {
	return terra.ReferenceAsList[dataserviceprincipal.FeaturesAttributes](sp.ref.Append("features"))
}

func (sp dataServicePrincipalAttributes) Oauth2PermissionScopes() terra.ListValue[dataserviceprincipal.Oauth2PermissionScopesAttributes] {
	return terra.ReferenceAsList[dataserviceprincipal.Oauth2PermissionScopesAttributes](sp.ref.Append("oauth2_permission_scopes"))
}

func (sp dataServicePrincipalAttributes) SamlSingleSignOn() terra.ListValue[dataserviceprincipal.SamlSingleSignOnAttributes] {
	return terra.ReferenceAsList[dataserviceprincipal.SamlSingleSignOnAttributes](sp.ref.Append("saml_single_sign_on"))
}

func (sp dataServicePrincipalAttributes) Timeouts() dataserviceprincipal.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataserviceprincipal.TimeoutsAttributes](sp.ref.Append("timeouts"))
}
