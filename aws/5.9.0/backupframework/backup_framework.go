// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package backupframework

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Control struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// InputParameter: min=0
	InputParameter []InputParameter `hcl:"input_parameter,block" validate:"min=0"`
	// Scope: optional
	Scope *Scope `hcl:"scope,block"`
}

type InputParameter struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Scope struct {
	// ComplianceResourceIds: set of string, optional
	ComplianceResourceIds terra.SetValue[terra.StringValue] `hcl:"compliance_resource_ids,attr"`
	// ComplianceResourceTypes: set of string, optional
	ComplianceResourceTypes terra.SetValue[terra.StringValue] `hcl:"compliance_resource_types,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ControlAttributes struct {
	ref terra.Reference
}

func (c ControlAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ControlAttributes) InternalWithRef(ref terra.Reference) ControlAttributes {
	return ControlAttributes{ref: ref}
}

func (c ControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ControlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ControlAttributes) InputParameter() terra.SetValue[InputParameterAttributes] {
	return terra.ReferenceAsSet[InputParameterAttributes](c.ref.Append("input_parameter"))
}

func (c ControlAttributes) Scope() terra.ListValue[ScopeAttributes] {
	return terra.ReferenceAsList[ScopeAttributes](c.ref.Append("scope"))
}

type InputParameterAttributes struct {
	ref terra.Reference
}

func (ip InputParameterAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip InputParameterAttributes) InternalWithRef(ref terra.Reference) InputParameterAttributes {
	return InputParameterAttributes{ref: ref}
}

func (ip InputParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip InputParameterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name"))
}

func (ip InputParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("value"))
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

func (s ScopeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScopeAttributes) ComplianceResourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("compliance_resource_ids"))
}

func (s ScopeAttributes) ComplianceResourceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("compliance_resource_types"))
}

func (s ScopeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ControlState struct {
	Name           string                `json:"name"`
	InputParameter []InputParameterState `json:"input_parameter"`
	Scope          []ScopeState          `json:"scope"`
}

type InputParameterState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ScopeState struct {
	ComplianceResourceIds   []string          `json:"compliance_resource_ids"`
	ComplianceResourceTypes []string          `json:"compliance_resource_types"`
	Tags                    map[string]string `json:"tags"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
