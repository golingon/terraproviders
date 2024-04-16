// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_data_protection_backup_vault

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataIdentityAttributes struct {
	ref terra.Reference
}

func (i DataIdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataIdentityAttributes) InternalWithRef(ref terra.Reference) DataIdentityAttributes {
	return DataIdentityAttributes{ref: ref}
}

func (i DataIdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataIdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i DataIdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i DataIdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataIdentityState struct {
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
	Type        string `json:"type"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
