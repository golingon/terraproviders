// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mysql_server

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Identity struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ThreatDetectionPolicy struct {
	// DisabledAlerts: set of string, optional
	DisabledAlerts terra.SetValue[terra.StringValue] `hcl:"disabled_alerts,attr"`
	// EmailAccountAdmins: bool, optional
	EmailAccountAdmins terra.BoolValue `hcl:"email_account_admins,attr"`
	// EmailAddresses: set of string, optional
	EmailAddresses terra.SetValue[terra.StringValue] `hcl:"email_addresses,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
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

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type ThreatDetectionPolicyAttributes struct {
	ref terra.Reference
}

func (tdp ThreatDetectionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tdp.ref, nil
}

func (tdp ThreatDetectionPolicyAttributes) InternalWithRef(ref terra.Reference) ThreatDetectionPolicyAttributes {
	return ThreatDetectionPolicyAttributes{ref: ref}
}

func (tdp ThreatDetectionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tdp.ref.InternalTokens()
}

func (tdp ThreatDetectionPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tdp.ref.Append("disabled_alerts"))
}

func (tdp ThreatDetectionPolicyAttributes) EmailAccountAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(tdp.ref.Append("email_account_admins"))
}

func (tdp ThreatDetectionPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tdp.ref.Append("email_addresses"))
}

func (tdp ThreatDetectionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tdp.ref.Append("enabled"))
}

func (tdp ThreatDetectionPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(tdp.ref.Append("retention_days"))
}

func (tdp ThreatDetectionPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("storage_account_access_key"))
}

func (tdp ThreatDetectionPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("storage_endpoint"))
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

type IdentityState struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
}

type ThreatDetectionPolicyState struct {
	DisabledAlerts          []string `json:"disabled_alerts"`
	EmailAccountAdmins      bool     `json:"email_account_admins"`
	EmailAddresses          []string `json:"email_addresses"`
	Enabled                 bool     `json:"enabled"`
	RetentionDays           float64  `json:"retention_days"`
	StorageAccountAccessKey string   `json:"storage_account_access_key"`
	StorageEndpoint         string   `json:"storage_endpoint"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
