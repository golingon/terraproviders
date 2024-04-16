// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mysql_server

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataIdentityAttributes struct {
	ref terra.Reference
}

func (i DataIdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataIdentityAttributes) InternalWithRef(ref terra.Reference) DataIdentityAttributes {
	return DataIdentityAttributes{ref: ref}
}

func (i DataIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i DataIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i DataIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type DataThreatDetectionPolicyAttributes struct {
	ref terra.Reference
}

func (tdp DataThreatDetectionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tdp.ref, nil
}

func (tdp DataThreatDetectionPolicyAttributes) InternalWithRef(ref terra.Reference) DataThreatDetectionPolicyAttributes {
	return DataThreatDetectionPolicyAttributes{ref: ref}
}

func (tdp DataThreatDetectionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tdp.ref.InternalTokens()
}

func (tdp DataThreatDetectionPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tdp.ref.Append("disabled_alerts"))
}

func (tdp DataThreatDetectionPolicyAttributes) EmailAccountAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(tdp.ref.Append("email_account_admins"))
}

func (tdp DataThreatDetectionPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](tdp.ref.Append("email_addresses"))
}

func (tdp DataThreatDetectionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(tdp.ref.Append("enabled"))
}

func (tdp DataThreatDetectionPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(tdp.ref.Append("retention_days"))
}

func (tdp DataThreatDetectionPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("storage_account_access_key"))
}

func (tdp DataThreatDetectionPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(tdp.ref.Append("storage_endpoint"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataIdentityState struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
}

type DataThreatDetectionPolicyState struct {
	DisabledAlerts          []string `json:"disabled_alerts"`
	EmailAccountAdmins      bool     `json:"email_account_admins"`
	EmailAddresses          []string `json:"email_addresses"`
	Enabled                 bool     `json:"enabled"`
	RetentionDays           float64  `json:"retention_days"`
	StorageAccountAccessKey string   `json:"storage_account_access_key"`
	StorageEndpoint         string   `json:"storage_endpoint"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
