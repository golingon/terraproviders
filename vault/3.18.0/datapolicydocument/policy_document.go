// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapolicydocument

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// Capabilities: list of string, required
	Capabilities terra.ListValue[terra.StringValue] `hcl:"capabilities,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// MaxWrappingTtl: string, optional
	MaxWrappingTtl terra.StringValue `hcl:"max_wrapping_ttl,attr"`
	// MinWrappingTtl: string, optional
	MinWrappingTtl terra.StringValue `hcl:"min_wrapping_ttl,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// RequiredParameters: list of string, optional
	RequiredParameters terra.ListValue[terra.StringValue] `hcl:"required_parameters,attr"`
	// AllowedParameter: min=0
	AllowedParameter []AllowedParameter `hcl:"allowed_parameter,block" validate:"min=0"`
	// DeniedParameter: min=0
	DeniedParameter []DeniedParameter `hcl:"denied_parameter,block" validate:"min=0"`
}

type AllowedParameter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: list of string, required
	Value terra.ListValue[terra.StringValue] `hcl:"value,attr" validate:"required"`
}

type DeniedParameter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: list of string, required
	Value terra.ListValue[terra.StringValue] `hcl:"value,attr" validate:"required"`
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Capabilities() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("capabilities"))
}

func (r RuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r RuleAttributes) MaxWrappingTtl() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("max_wrapping_ttl"))
}

func (r RuleAttributes) MinWrappingTtl() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("min_wrapping_ttl"))
}

func (r RuleAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("path"))
}

func (r RuleAttributes) RequiredParameters() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("required_parameters"))
}

func (r RuleAttributes) AllowedParameter() terra.ListValue[AllowedParameterAttributes] {
	return terra.ReferenceAsList[AllowedParameterAttributes](r.ref.Append("allowed_parameter"))
}

func (r RuleAttributes) DeniedParameter() terra.ListValue[DeniedParameterAttributes] {
	return terra.ReferenceAsList[DeniedParameterAttributes](r.ref.Append("denied_parameter"))
}

type AllowedParameterAttributes struct {
	ref terra.Reference
}

func (ap AllowedParameterAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AllowedParameterAttributes) InternalWithRef(ref terra.Reference) AllowedParameterAttributes {
	return AllowedParameterAttributes{ref: ref}
}

func (ap AllowedParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AllowedParameterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("key"))
}

func (ap AllowedParameterAttributes) Value() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ap.ref.Append("value"))
}

type DeniedParameterAttributes struct {
	ref terra.Reference
}

func (dp DeniedParameterAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DeniedParameterAttributes) InternalWithRef(ref terra.Reference) DeniedParameterAttributes {
	return DeniedParameterAttributes{ref: ref}
}

func (dp DeniedParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DeniedParameterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("key"))
}

func (dp DeniedParameterAttributes) Value() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dp.ref.Append("value"))
}

type RuleState struct {
	Capabilities       []string                `json:"capabilities"`
	Description        string                  `json:"description"`
	MaxWrappingTtl     string                  `json:"max_wrapping_ttl"`
	MinWrappingTtl     string                  `json:"min_wrapping_ttl"`
	Path               string                  `json:"path"`
	RequiredParameters []string                `json:"required_parameters"`
	AllowedParameter   []AllowedParameterState `json:"allowed_parameter"`
	DeniedParameter    []DeniedParameterState  `json:"denied_parameter"`
}

type AllowedParameterState struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

type DeniedParameterState struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}