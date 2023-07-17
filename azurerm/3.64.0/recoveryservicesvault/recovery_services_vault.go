// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package recoveryservicesvault

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Encryption struct {
	// InfrastructureEncryptionEnabled: bool, required
	InfrastructureEncryptionEnabled terra.BoolValue `hcl:"infrastructure_encryption_enabled,attr" validate:"required"`
	// KeyId: string, required
	KeyId terra.StringValue `hcl:"key_id,attr" validate:"required"`
	// UseSystemAssignedIdentity: bool, optional
	UseSystemAssignedIdentity terra.BoolValue `hcl:"use_system_assigned_identity,attr"`
	// UserAssignedIdentityId: string, optional
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Monitoring struct {
	// AlertsForAllJobFailuresEnabled: bool, optional
	AlertsForAllJobFailuresEnabled terra.BoolValue `hcl:"alerts_for_all_job_failures_enabled,attr"`
	// AlertsForCriticalOperationFailuresEnabled: bool, optional
	AlertsForCriticalOperationFailuresEnabled terra.BoolValue `hcl:"alerts_for_critical_operation_failures_enabled,attr"`
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

type EncryptionAttributes struct {
	ref terra.Reference
}

func (e EncryptionAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EncryptionAttributes) InternalWithRef(ref terra.Reference) EncryptionAttributes {
	return EncryptionAttributes{ref: ref}
}

func (e EncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EncryptionAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("infrastructure_encryption_enabled"))
}

func (e EncryptionAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("key_id"))
}

func (e EncryptionAttributes) UseSystemAssignedIdentity() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("use_system_assigned_identity"))
}

func (e EncryptionAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("user_assigned_identity_id"))
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

type MonitoringAttributes struct {
	ref terra.Reference
}

func (m MonitoringAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MonitoringAttributes) InternalWithRef(ref terra.Reference) MonitoringAttributes {
	return MonitoringAttributes{ref: ref}
}

func (m MonitoringAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MonitoringAttributes) AlertsForAllJobFailuresEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(m.ref.Append("alerts_for_all_job_failures_enabled"))
}

func (m MonitoringAttributes) AlertsForCriticalOperationFailuresEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(m.ref.Append("alerts_for_critical_operation_failures_enabled"))
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

type EncryptionState struct {
	InfrastructureEncryptionEnabled bool   `json:"infrastructure_encryption_enabled"`
	KeyId                           string `json:"key_id"`
	UseSystemAssignedIdentity       bool   `json:"use_system_assigned_identity"`
	UserAssignedIdentityId          string `json:"user_assigned_identity_id"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type MonitoringState struct {
	AlertsForAllJobFailuresEnabled            bool `json:"alerts_for_all_job_failures_enabled"`
	AlertsForCriticalOperationFailuresEnabled bool `json:"alerts_for_critical_operation_failures_enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}