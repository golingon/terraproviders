// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataserviceprincipals

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ServicePrincipals struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ServicePrincipalsAttributes struct {
	ref terra.Reference
}

func (sp ServicePrincipalsAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp ServicePrincipalsAttributes) InternalWithRef(ref terra.Reference) ServicePrincipalsAttributes {
	return ServicePrincipalsAttributes{ref: ref}
}

func (sp ServicePrincipalsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp ServicePrincipalsAttributes) AccountEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("account_enabled"))
}

func (sp ServicePrincipalsAttributes) AppRoleAssignmentRequired() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("app_role_assignment_required"))
}

func (sp ServicePrincipalsAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("application_id"))
}

func (sp ServicePrincipalsAttributes) ApplicationTenantId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("application_tenant_id"))
}

func (sp ServicePrincipalsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("display_name"))
}

func (sp ServicePrincipalsAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("object_id"))
}

func (sp ServicePrincipalsAttributes) PreferredSingleSignOnMode() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("preferred_single_sign_on_mode"))
}

func (sp ServicePrincipalsAttributes) SamlMetadataUrl() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("saml_metadata_url"))
}

func (sp ServicePrincipalsAttributes) ServicePrincipalNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("service_principal_names"))
}

func (sp ServicePrincipalsAttributes) SignInAudience() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("sign_in_audience"))
}

func (sp ServicePrincipalsAttributes) Tags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sp.ref.Append("tags"))
}

func (sp ServicePrincipalsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ServicePrincipalsState struct {
	AccountEnabled            bool     `json:"account_enabled"`
	AppRoleAssignmentRequired bool     `json:"app_role_assignment_required"`
	ApplicationId             string   `json:"application_id"`
	ApplicationTenantId       string   `json:"application_tenant_id"`
	DisplayName               string   `json:"display_name"`
	ObjectId                  string   `json:"object_id"`
	PreferredSingleSignOnMode string   `json:"preferred_single_sign_on_mode"`
	SamlMetadataUrl           string   `json:"saml_metadata_url"`
	ServicePrincipalNames     []string `json:"service_principal_names"`
	SignInAudience            string   `json:"sign_in_audience"`
	Tags                      []string `json:"tags"`
	Type                      string   `json:"type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
