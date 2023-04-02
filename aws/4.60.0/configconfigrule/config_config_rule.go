// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package configconfigrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Scope struct {
	// ComplianceResourceId: string, optional
	ComplianceResourceId terra.StringValue `hcl:"compliance_resource_id,attr"`
	// ComplianceResourceTypes: set of string, optional
	ComplianceResourceTypes terra.SetValue[terra.StringValue] `hcl:"compliance_resource_types,attr"`
	// TagKey: string, optional
	TagKey terra.StringValue `hcl:"tag_key,attr"`
	// TagValue: string, optional
	TagValue terra.StringValue `hcl:"tag_value,attr"`
}

type Source struct {
	// Owner: string, required
	Owner terra.StringValue `hcl:"owner,attr" validate:"required"`
	// SourceIdentifier: string, optional
	SourceIdentifier terra.StringValue `hcl:"source_identifier,attr"`
	// CustomPolicyDetails: optional
	CustomPolicyDetails *CustomPolicyDetails `hcl:"custom_policy_details,block"`
	// SourceDetail: min=0,max=25
	SourceDetail []SourceDetail `hcl:"source_detail,block" validate:"min=0,max=25"`
}

type CustomPolicyDetails struct {
	// EnableDebugLogDelivery: bool, optional
	EnableDebugLogDelivery terra.BoolValue `hcl:"enable_debug_log_delivery,attr"`
	// PolicyRuntime: string, required
	PolicyRuntime terra.StringValue `hcl:"policy_runtime,attr" validate:"required"`
	// PolicyText: string, required
	PolicyText terra.StringValue `hcl:"policy_text,attr" validate:"required"`
}

type SourceDetail struct {
	// EventSource: string, optional
	EventSource terra.StringValue `hcl:"event_source,attr"`
	// MaximumExecutionFrequency: string, optional
	MaximumExecutionFrequency terra.StringValue `hcl:"maximum_execution_frequency,attr"`
	// MessageType: string, optional
	MessageType terra.StringValue `hcl:"message_type,attr"`
}

type ScopeAttributes struct {
	ref terra.Reference
}

func (s ScopeAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScopeAttributes) InternalWithRef(ref terra.Reference) ScopeAttributes {
	return ScopeAttributes{ref: ref}
}

func (s ScopeAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s ScopeAttributes) ComplianceResourceId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("compliance_resource_id"))
}

func (s ScopeAttributes) ComplianceResourceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("compliance_resource_types"))
}

func (s ScopeAttributes) TagKey() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tag_key"))
}

func (s ScopeAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tag_value"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("owner"))
}

func (s SourceAttributes) SourceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("source_identifier"))
}

func (s SourceAttributes) CustomPolicyDetails() terra.ListValue[CustomPolicyDetailsAttributes] {
	return terra.ReferenceAsList[CustomPolicyDetailsAttributes](s.ref.Append("custom_policy_details"))
}

func (s SourceAttributes) SourceDetail() terra.SetValue[SourceDetailAttributes] {
	return terra.ReferenceAsSet[SourceDetailAttributes](s.ref.Append("source_detail"))
}

type CustomPolicyDetailsAttributes struct {
	ref terra.Reference
}

func (cpd CustomPolicyDetailsAttributes) InternalRef() (terra.Reference, error) {
	return cpd.ref, nil
}

func (cpd CustomPolicyDetailsAttributes) InternalWithRef(ref terra.Reference) CustomPolicyDetailsAttributes {
	return CustomPolicyDetailsAttributes{ref: ref}
}

func (cpd CustomPolicyDetailsAttributes) InternalTokens() hclwrite.Tokens {
	return cpd.ref.InternalTokens()
}

func (cpd CustomPolicyDetailsAttributes) EnableDebugLogDelivery() terra.BoolValue {
	return terra.ReferenceAsBool(cpd.ref.Append("enable_debug_log_delivery"))
}

func (cpd CustomPolicyDetailsAttributes) PolicyRuntime() terra.StringValue {
	return terra.ReferenceAsString(cpd.ref.Append("policy_runtime"))
}

func (cpd CustomPolicyDetailsAttributes) PolicyText() terra.StringValue {
	return terra.ReferenceAsString(cpd.ref.Append("policy_text"))
}

type SourceDetailAttributes struct {
	ref terra.Reference
}

func (sd SourceDetailAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd SourceDetailAttributes) InternalWithRef(ref terra.Reference) SourceDetailAttributes {
	return SourceDetailAttributes{ref: ref}
}

func (sd SourceDetailAttributes) InternalTokens() hclwrite.Tokens {
	return sd.ref.InternalTokens()
}

func (sd SourceDetailAttributes) EventSource() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("event_source"))
}

func (sd SourceDetailAttributes) MaximumExecutionFrequency() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("maximum_execution_frequency"))
}

func (sd SourceDetailAttributes) MessageType() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("message_type"))
}

type ScopeState struct {
	ComplianceResourceId    string   `json:"compliance_resource_id"`
	ComplianceResourceTypes []string `json:"compliance_resource_types"`
	TagKey                  string   `json:"tag_key"`
	TagValue                string   `json:"tag_value"`
}

type SourceState struct {
	Owner               string                     `json:"owner"`
	SourceIdentifier    string                     `json:"source_identifier"`
	CustomPolicyDetails []CustomPolicyDetailsState `json:"custom_policy_details"`
	SourceDetail        []SourceDetailState        `json:"source_detail"`
}

type CustomPolicyDetailsState struct {
	EnableDebugLogDelivery bool   `json:"enable_debug_log_delivery"`
	PolicyRuntime          string `json:"policy_runtime"`
	PolicyText             string `json:"policy_text"`
}

type SourceDetailState struct {
	EventSource               string `json:"event_source"`
	MaximumExecutionFrequency string `json:"maximum_execution_frequency"`
	MessageType               string `json:"message_type"`
}
