// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package discoveryenginesearchengine

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CommonConfig struct {
	// CompanyName: string, optional
	CompanyName terra.StringValue `hcl:"company_name,attr"`
}

type SearchEngineConfig struct {
	// SearchAddOns: list of string, optional
	SearchAddOns terra.ListValue[terra.StringValue] `hcl:"search_add_ons,attr"`
	// SearchTier: string, optional
	SearchTier terra.StringValue `hcl:"search_tier,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CommonConfigAttributes struct {
	ref terra.Reference
}

func (cc CommonConfigAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CommonConfigAttributes) InternalWithRef(ref terra.Reference) CommonConfigAttributes {
	return CommonConfigAttributes{ref: ref}
}

func (cc CommonConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CommonConfigAttributes) CompanyName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("company_name"))
}

type SearchEngineConfigAttributes struct {
	ref terra.Reference
}

func (sec SearchEngineConfigAttributes) InternalRef() (terra.Reference, error) {
	return sec.ref, nil
}

func (sec SearchEngineConfigAttributes) InternalWithRef(ref terra.Reference) SearchEngineConfigAttributes {
	return SearchEngineConfigAttributes{ref: ref}
}

func (sec SearchEngineConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sec.ref.InternalTokens()
}

func (sec SearchEngineConfigAttributes) SearchAddOns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sec.ref.Append("search_add_ons"))
}

func (sec SearchEngineConfigAttributes) SearchTier() terra.StringValue {
	return terra.ReferenceAsString(sec.ref.Append("search_tier"))
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

type CommonConfigState struct {
	CompanyName string `json:"company_name"`
}

type SearchEngineConfigState struct {
	SearchAddOns []string `json:"search_add_ons"`
	SearchTier   string   `json:"search_tier"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
