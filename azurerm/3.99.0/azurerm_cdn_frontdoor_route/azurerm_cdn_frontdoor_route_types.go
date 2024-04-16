// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_route

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Cache struct {
	// CompressionEnabled: bool, optional
	CompressionEnabled terra.BoolValue `hcl:"compression_enabled,attr"`
	// ContentTypesToCompress: list of string, optional
	ContentTypesToCompress terra.ListValue[terra.StringValue] `hcl:"content_types_to_compress,attr"`
	// QueryStringCachingBehavior: string, optional
	QueryStringCachingBehavior terra.StringValue `hcl:"query_string_caching_behavior,attr"`
	// QueryStrings: list of string, optional
	QueryStrings terra.ListValue[terra.StringValue] `hcl:"query_strings,attr"`
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

type CacheAttributes struct {
	ref terra.Reference
}

func (c CacheAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CacheAttributes) InternalWithRef(ref terra.Reference) CacheAttributes {
	return CacheAttributes{ref: ref}
}

func (c CacheAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CacheAttributes) CompressionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(c.ref.Append("compression_enabled"))
}

func (c CacheAttributes) ContentTypesToCompress() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("content_types_to_compress"))
}

func (c CacheAttributes) QueryStringCachingBehavior() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("query_string_caching_behavior"))
}

func (c CacheAttributes) QueryStrings() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("query_strings"))
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

type CacheState struct {
	CompressionEnabled         bool     `json:"compression_enabled"`
	ContentTypesToCompress     []string `json:"content_types_to_compress"`
	QueryStringCachingBehavior string   `json:"query_string_caching_behavior"`
	QueryStrings               []string `json:"query_strings"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
