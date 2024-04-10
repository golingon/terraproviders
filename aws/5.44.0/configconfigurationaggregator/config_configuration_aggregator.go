// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package configconfigurationaggregator

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AccountAggregationSource struct {
	// AccountIds: list of string, required
	AccountIds terra.ListValue[terra.StringValue] `hcl:"account_ids,attr" validate:"required"`
	// AllRegions: bool, optional
	AllRegions terra.BoolValue `hcl:"all_regions,attr"`
	// Regions: list of string, optional
	Regions terra.ListValue[terra.StringValue] `hcl:"regions,attr"`
}

type OrganizationAggregationSource struct {
	// AllRegions: bool, optional
	AllRegions terra.BoolValue `hcl:"all_regions,attr"`
	// Regions: list of string, optional
	Regions terra.ListValue[terra.StringValue] `hcl:"regions,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type AccountAggregationSourceAttributes struct {
	ref terra.Reference
}

func (aas AccountAggregationSourceAttributes) InternalRef() (terra.Reference, error) {
	return aas.ref, nil
}

func (aas AccountAggregationSourceAttributes) InternalWithRef(ref terra.Reference) AccountAggregationSourceAttributes {
	return AccountAggregationSourceAttributes{ref: ref}
}

func (aas AccountAggregationSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aas.ref.InternalTokens()
}

func (aas AccountAggregationSourceAttributes) AccountIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aas.ref.Append("account_ids"))
}

func (aas AccountAggregationSourceAttributes) AllRegions() terra.BoolValue {
	return terra.ReferenceAsBool(aas.ref.Append("all_regions"))
}

func (aas AccountAggregationSourceAttributes) Regions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aas.ref.Append("regions"))
}

type OrganizationAggregationSourceAttributes struct {
	ref terra.Reference
}

func (oas OrganizationAggregationSourceAttributes) InternalRef() (terra.Reference, error) {
	return oas.ref, nil
}

func (oas OrganizationAggregationSourceAttributes) InternalWithRef(ref terra.Reference) OrganizationAggregationSourceAttributes {
	return OrganizationAggregationSourceAttributes{ref: ref}
}

func (oas OrganizationAggregationSourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oas.ref.InternalTokens()
}

func (oas OrganizationAggregationSourceAttributes) AllRegions() terra.BoolValue {
	return terra.ReferenceAsBool(oas.ref.Append("all_regions"))
}

func (oas OrganizationAggregationSourceAttributes) Regions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](oas.ref.Append("regions"))
}

func (oas OrganizationAggregationSourceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(oas.ref.Append("role_arn"))
}

type AccountAggregationSourceState struct {
	AccountIds []string `json:"account_ids"`
	AllRegions bool     `json:"all_regions"`
	Regions    []string `json:"regions"`
}

type OrganizationAggregationSourceState struct {
	AllRegions bool     `json:"all_regions"`
	Regions    []string `json:"regions"`
	RoleArn    string   `json:"role_arn"`
}
