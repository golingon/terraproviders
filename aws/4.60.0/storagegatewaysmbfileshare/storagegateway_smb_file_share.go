// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storagegatewaysmbfileshare

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CacheAttributes struct {
	// CacheStaleTimeoutInSeconds: number, optional
	CacheStaleTimeoutInSeconds terra.NumberValue `hcl:"cache_stale_timeout_in_seconds,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CacheAttributesAttributes struct {
	ref terra.Reference
}

func (ca CacheAttributesAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca CacheAttributesAttributes) InternalWithRef(ref terra.Reference) CacheAttributesAttributes {
	return CacheAttributesAttributes{ref: ref}
}

func (ca CacheAttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca CacheAttributesAttributes) CacheStaleTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("cache_stale_timeout_in_seconds"))
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

type CacheAttributesState struct {
	CacheStaleTimeoutInSeconds float64 `json:"cache_stale_timeout_in_seconds"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
