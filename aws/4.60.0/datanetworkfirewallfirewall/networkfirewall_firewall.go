// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datanetworkfirewallfirewall

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfiguration struct{}

type FirewallStatus struct {
	// CapacityUsageSummary: min=0
	CapacityUsageSummary []CapacityUsageSummary `hcl:"capacity_usage_summary,block" validate:"min=0"`
	// SyncStates: min=0
	SyncStates []SyncStates `hcl:"sync_states,block" validate:"min=0"`
}

type CapacityUsageSummary struct {
	// Cidrs: min=0
	Cidrs []Cidrs `hcl:"cidrs,block" validate:"min=0"`
}

type Cidrs struct {
	// IpSetReferences: min=0
	IpSetReferences []IpSetReferences `hcl:"ip_set_references,block" validate:"min=0"`
}

type IpSetReferences struct{}

type SyncStates struct {
	// Attachment: min=0
	Attachment []Attachment `hcl:"attachment,block" validate:"min=0"`
}

type Attachment struct{}

type SubnetMapping struct{}

type EncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigurationAttributes) InternalRef() terra.Reference {
	return ec.ref
}

func (ec EncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigurationAttributes {
	return EncryptionConfigurationAttributes{ref: ref}
}

func (ec EncryptionConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigurationAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("key_id"))
}

func (ec EncryptionConfigurationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("type"))
}

type FirewallStatusAttributes struct {
	ref terra.Reference
}

func (fs FirewallStatusAttributes) InternalRef() terra.Reference {
	return fs.ref
}

func (fs FirewallStatusAttributes) InternalWithRef(ref terra.Reference) FirewallStatusAttributes {
	return FirewallStatusAttributes{ref: ref}
}

func (fs FirewallStatusAttributes) InternalTokens() hclwrite.Tokens {
	return fs.ref.InternalTokens()
}

func (fs FirewallStatusAttributes) ConfigurationSyncStateSummary() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("configuration_sync_state_summary"))
}

func (fs FirewallStatusAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("status"))
}

func (fs FirewallStatusAttributes) CapacityUsageSummary() terra.SetValue[CapacityUsageSummaryAttributes] {
	return terra.ReferenceAsSet[CapacityUsageSummaryAttributes](fs.ref.Append("capacity_usage_summary"))
}

func (fs FirewallStatusAttributes) SyncStates() terra.SetValue[SyncStatesAttributes] {
	return terra.ReferenceAsSet[SyncStatesAttributes](fs.ref.Append("sync_states"))
}

type CapacityUsageSummaryAttributes struct {
	ref terra.Reference
}

func (cus CapacityUsageSummaryAttributes) InternalRef() terra.Reference {
	return cus.ref
}

func (cus CapacityUsageSummaryAttributes) InternalWithRef(ref terra.Reference) CapacityUsageSummaryAttributes {
	return CapacityUsageSummaryAttributes{ref: ref}
}

func (cus CapacityUsageSummaryAttributes) InternalTokens() hclwrite.Tokens {
	return cus.ref.InternalTokens()
}

func (cus CapacityUsageSummaryAttributes) Cidrs() terra.SetValue[CidrsAttributes] {
	return terra.ReferenceAsSet[CidrsAttributes](cus.ref.Append("cidrs"))
}

type CidrsAttributes struct {
	ref terra.Reference
}

func (c CidrsAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c CidrsAttributes) InternalWithRef(ref terra.Reference) CidrsAttributes {
	return CidrsAttributes{ref: ref}
}

func (c CidrsAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c CidrsAttributes) AvailableCidrCount() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("available_cidr_count"))
}

func (c CidrsAttributes) UtilizedCidrCount() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("utilized_cidr_count"))
}

func (c CidrsAttributes) IpSetReferences() terra.SetValue[IpSetReferencesAttributes] {
	return terra.ReferenceAsSet[IpSetReferencesAttributes](c.ref.Append("ip_set_references"))
}

type IpSetReferencesAttributes struct {
	ref terra.Reference
}

func (isr IpSetReferencesAttributes) InternalRef() terra.Reference {
	return isr.ref
}

func (isr IpSetReferencesAttributes) InternalWithRef(ref terra.Reference) IpSetReferencesAttributes {
	return IpSetReferencesAttributes{ref: ref}
}

func (isr IpSetReferencesAttributes) InternalTokens() hclwrite.Tokens {
	return isr.ref.InternalTokens()
}

func (isr IpSetReferencesAttributes) ResolvedCidrCount() terra.NumberValue {
	return terra.ReferenceAsNumber(isr.ref.Append("resolved_cidr_count"))
}

type SyncStatesAttributes struct {
	ref terra.Reference
}

func (ss SyncStatesAttributes) InternalRef() terra.Reference {
	return ss.ref
}

func (ss SyncStatesAttributes) InternalWithRef(ref terra.Reference) SyncStatesAttributes {
	return SyncStatesAttributes{ref: ref}
}

func (ss SyncStatesAttributes) InternalTokens() hclwrite.Tokens {
	return ss.ref.InternalTokens()
}

func (ss SyncStatesAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("availability_zone"))
}

func (ss SyncStatesAttributes) Attachment() terra.ListValue[AttachmentAttributes] {
	return terra.ReferenceAsList[AttachmentAttributes](ss.ref.Append("attachment"))
}

type AttachmentAttributes struct {
	ref terra.Reference
}

func (a AttachmentAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AttachmentAttributes) InternalWithRef(ref terra.Reference) AttachmentAttributes {
	return AttachmentAttributes{ref: ref}
}

func (a AttachmentAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AttachmentAttributes) EndpointId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("endpoint_id"))
}

func (a AttachmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("status"))
}

func (a AttachmentAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("subnet_id"))
}

type SubnetMappingAttributes struct {
	ref terra.Reference
}

func (sm SubnetMappingAttributes) InternalRef() terra.Reference {
	return sm.ref
}

func (sm SubnetMappingAttributes) InternalWithRef(ref terra.Reference) SubnetMappingAttributes {
	return SubnetMappingAttributes{ref: ref}
}

func (sm SubnetMappingAttributes) InternalTokens() hclwrite.Tokens {
	return sm.ref.InternalTokens()
}

func (sm SubnetMappingAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(sm.ref.Append("subnet_id"))
}

type EncryptionConfigurationState struct {
	KeyId string `json:"key_id"`
	Type  string `json:"type"`
}

type FirewallStatusState struct {
	ConfigurationSyncStateSummary string                      `json:"configuration_sync_state_summary"`
	Status                        string                      `json:"status"`
	CapacityUsageSummary          []CapacityUsageSummaryState `json:"capacity_usage_summary"`
	SyncStates                    []SyncStatesState           `json:"sync_states"`
}

type CapacityUsageSummaryState struct {
	Cidrs []CidrsState `json:"cidrs"`
}

type CidrsState struct {
	AvailableCidrCount float64                `json:"available_cidr_count"`
	UtilizedCidrCount  float64                `json:"utilized_cidr_count"`
	IpSetReferences    []IpSetReferencesState `json:"ip_set_references"`
}

type IpSetReferencesState struct {
	ResolvedCidrCount float64 `json:"resolved_cidr_count"`
}

type SyncStatesState struct {
	AvailabilityZone string            `json:"availability_zone"`
	Attachment       []AttachmentState `json:"attachment"`
}

type AttachmentState struct {
	EndpointId string `json:"endpoint_id"`
	Status     string `json:"status"`
	SubnetId   string `json:"subnet_id"`
}

type SubnetMappingState struct {
	SubnetId string `json:"subnet_id"`
}
