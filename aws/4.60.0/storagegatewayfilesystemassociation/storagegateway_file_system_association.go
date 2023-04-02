// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storagegatewayfilesystemassociation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CacheAttributes struct {
	// CacheStaleTimeoutInSeconds: number, optional
	CacheStaleTimeoutInSeconds terra.NumberValue `hcl:"cache_stale_timeout_in_seconds,attr"`
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

func (ca CacheAttributesAttributes) InternalTokens() hclwrite.Tokens {
	return ca.ref.InternalTokens()
}

func (ca CacheAttributesAttributes) CacheStaleTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("cache_stale_timeout_in_seconds"))
}

type CacheAttributesState struct {
	CacheStaleTimeoutInSeconds float64 `json:"cache_stale_timeout_in_seconds"`
}
