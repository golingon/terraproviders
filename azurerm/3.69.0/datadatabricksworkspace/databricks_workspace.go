// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadatabricksworkspace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ManagedDiskIdentity struct{}

type StorageAccountIdentity struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ManagedDiskIdentityAttributes struct {
	ref terra.Reference
}

func (mdi ManagedDiskIdentityAttributes) InternalRef() (terra.Reference, error) {
	return mdi.ref, nil
}

func (mdi ManagedDiskIdentityAttributes) InternalWithRef(ref terra.Reference) ManagedDiskIdentityAttributes {
	return ManagedDiskIdentityAttributes{ref: ref}
}

func (mdi ManagedDiskIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mdi.ref.InternalTokens()
}

func (mdi ManagedDiskIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(mdi.ref.Append("principal_id"))
}

func (mdi ManagedDiskIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(mdi.ref.Append("tenant_id"))
}

func (mdi ManagedDiskIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mdi.ref.Append("type"))
}

type StorageAccountIdentityAttributes struct {
	ref terra.Reference
}

func (sai StorageAccountIdentityAttributes) InternalRef() (terra.Reference, error) {
	return sai.ref, nil
}

func (sai StorageAccountIdentityAttributes) InternalWithRef(ref terra.Reference) StorageAccountIdentityAttributes {
	return StorageAccountIdentityAttributes{ref: ref}
}

func (sai StorageAccountIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sai.ref.InternalTokens()
}

func (sai StorageAccountIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(sai.ref.Append("principal_id"))
}

func (sai StorageAccountIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(sai.ref.Append("tenant_id"))
}

func (sai StorageAccountIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sai.ref.Append("type"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ManagedDiskIdentityState struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
}

type StorageAccountIdentityState struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}