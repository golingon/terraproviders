// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storagegatewaynfsfileshare

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CacheAttributes struct {
	// CacheStaleTimeoutInSeconds: number, optional
	CacheStaleTimeoutInSeconds terra.NumberValue `hcl:"cache_stale_timeout_in_seconds,attr"`
}

type NfsFileShareDefaults struct {
	// DirectoryMode: string, optional
	DirectoryMode terra.StringValue `hcl:"directory_mode,attr"`
	// FileMode: string, optional
	FileMode terra.StringValue `hcl:"file_mode,attr"`
	// GroupId: string, optional
	GroupId terra.StringValue `hcl:"group_id,attr"`
	// OwnerId: string, optional
	OwnerId terra.StringValue `hcl:"owner_id,attr"`
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

func (ca CacheAttributesAttributes) InternalTokens() hclwrite.Tokens {
	return ca.ref.InternalTokens()
}

func (ca CacheAttributesAttributes) CacheStaleTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("cache_stale_timeout_in_seconds"))
}

type NfsFileShareDefaultsAttributes struct {
	ref terra.Reference
}

func (nfsd NfsFileShareDefaultsAttributes) InternalRef() (terra.Reference, error) {
	return nfsd.ref, nil
}

func (nfsd NfsFileShareDefaultsAttributes) InternalWithRef(ref terra.Reference) NfsFileShareDefaultsAttributes {
	return NfsFileShareDefaultsAttributes{ref: ref}
}

func (nfsd NfsFileShareDefaultsAttributes) InternalTokens() hclwrite.Tokens {
	return nfsd.ref.InternalTokens()
}

func (nfsd NfsFileShareDefaultsAttributes) DirectoryMode() terra.StringValue {
	return terra.ReferenceAsString(nfsd.ref.Append("directory_mode"))
}

func (nfsd NfsFileShareDefaultsAttributes) FileMode() terra.StringValue {
	return terra.ReferenceAsString(nfsd.ref.Append("file_mode"))
}

func (nfsd NfsFileShareDefaultsAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(nfsd.ref.Append("group_id"))
}

func (nfsd NfsFileShareDefaultsAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(nfsd.ref.Append("owner_id"))
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

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
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

type NfsFileShareDefaultsState struct {
	DirectoryMode string `json:"directory_mode"`
	FileMode      string `json:"file_mode"`
	GroupId       string `json:"group_id"`
	OwnerId       string `json:"owner_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
