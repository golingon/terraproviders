// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dashboardgrafana

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AzureMonitorWorkspaceIntegrations struct {
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Smtp struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FromAddress: string, required
	FromAddress terra.StringValue `hcl:"from_address,attr" validate:"required"`
	// FromName: string, optional
	FromName terra.StringValue `hcl:"from_name,attr"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// StartTlsPolicy: string, required
	StartTlsPolicy terra.StringValue `hcl:"start_tls_policy,attr" validate:"required"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
	// VerificationSkipEnabled: bool, optional
	VerificationSkipEnabled terra.BoolValue `hcl:"verification_skip_enabled,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AzureMonitorWorkspaceIntegrationsAttributes struct {
	ref terra.Reference
}

func (amwi AzureMonitorWorkspaceIntegrationsAttributes) InternalRef() (terra.Reference, error) {
	return amwi.ref, nil
}

func (amwi AzureMonitorWorkspaceIntegrationsAttributes) InternalWithRef(ref terra.Reference) AzureMonitorWorkspaceIntegrationsAttributes {
	return AzureMonitorWorkspaceIntegrationsAttributes{ref: ref}
}

func (amwi AzureMonitorWorkspaceIntegrationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return amwi.ref.InternalTokens()
}

func (amwi AzureMonitorWorkspaceIntegrationsAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(amwi.ref.Append("resource_id"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type SmtpAttributes struct {
	ref terra.Reference
}

func (s SmtpAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SmtpAttributes) InternalWithRef(ref terra.Reference) SmtpAttributes {
	return SmtpAttributes{ref: ref}
}

func (s SmtpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SmtpAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("enabled"))
}

func (s SmtpAttributes) FromAddress() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("from_address"))
}

func (s SmtpAttributes) FromName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("from_name"))
}

func (s SmtpAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("host"))
}

func (s SmtpAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("password"))
}

func (s SmtpAttributes) StartTlsPolicy() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("start_tls_policy"))
}

func (s SmtpAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("user"))
}

func (s SmtpAttributes) VerificationSkipEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("verification_skip_enabled"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AzureMonitorWorkspaceIntegrationsState struct {
	ResourceId string `json:"resource_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type SmtpState struct {
	Enabled                 bool   `json:"enabled"`
	FromAddress             string `json:"from_address"`
	FromName                string `json:"from_name"`
	Host                    string `json:"host"`
	Password                string `json:"password"`
	StartTlsPolicy          string `json:"start_tls_policy"`
	User                    string `json:"user"`
	VerificationSkipEnabled bool   `json:"verification_skip_enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
