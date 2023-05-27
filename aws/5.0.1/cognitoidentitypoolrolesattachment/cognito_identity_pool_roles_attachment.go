// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cognitoidentitypoolrolesattachment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RoleMapping struct {
	// AmbiguousRoleResolution: string, optional
	AmbiguousRoleResolution terra.StringValue `hcl:"ambiguous_role_resolution,attr"`
	// IdentityProvider: string, required
	IdentityProvider terra.StringValue `hcl:"identity_provider,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// MappingRule: min=0,max=25
	MappingRule []MappingRule `hcl:"mapping_rule,block" validate:"min=0,max=25"`
}

type MappingRule struct {
	// Claim: string, required
	Claim terra.StringValue `hcl:"claim,attr" validate:"required"`
	// MatchType: string, required
	MatchType terra.StringValue `hcl:"match_type,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type RoleMappingAttributes struct {
	ref terra.Reference
}

func (rm RoleMappingAttributes) InternalRef() (terra.Reference, error) {
	return rm.ref, nil
}

func (rm RoleMappingAttributes) InternalWithRef(ref terra.Reference) RoleMappingAttributes {
	return RoleMappingAttributes{ref: ref}
}

func (rm RoleMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rm.ref.InternalTokens()
}

func (rm RoleMappingAttributes) AmbiguousRoleResolution() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("ambiguous_role_resolution"))
}

func (rm RoleMappingAttributes) IdentityProvider() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("identity_provider"))
}

func (rm RoleMappingAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(rm.ref.Append("type"))
}

func (rm RoleMappingAttributes) MappingRule() terra.ListValue[MappingRuleAttributes] {
	return terra.ReferenceAsList[MappingRuleAttributes](rm.ref.Append("mapping_rule"))
}

type MappingRuleAttributes struct {
	ref terra.Reference
}

func (mr MappingRuleAttributes) InternalRef() (terra.Reference, error) {
	return mr.ref, nil
}

func (mr MappingRuleAttributes) InternalWithRef(ref terra.Reference) MappingRuleAttributes {
	return MappingRuleAttributes{ref: ref}
}

func (mr MappingRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mr.ref.InternalTokens()
}

func (mr MappingRuleAttributes) Claim() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("claim"))
}

func (mr MappingRuleAttributes) MatchType() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("match_type"))
}

func (mr MappingRuleAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("role_arn"))
}

func (mr MappingRuleAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(mr.ref.Append("value"))
}

type RoleMappingState struct {
	AmbiguousRoleResolution string             `json:"ambiguous_role_resolution"`
	IdentityProvider        string             `json:"identity_provider"`
	Type                    string             `json:"type"`
	MappingRule             []MappingRuleState `json:"mapping_rule"`
}

type MappingRuleState struct {
	Claim     string `json:"claim"`
	MatchType string `json:"match_type"`
	RoleArn   string `json:"role_arn"`
	Value     string `json:"value"`
}
