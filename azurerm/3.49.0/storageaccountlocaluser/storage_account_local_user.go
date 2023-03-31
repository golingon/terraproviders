// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storageaccountlocaluser

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PermissionScope struct {
	// ResourceName: string, required
	ResourceName terra.StringValue `hcl:"resource_name,attr" validate:"required"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
	// Permissions: required
	Permissions *Permissions `hcl:"permissions,block" validate:"required"`
}

type Permissions struct {
	// Create: bool, optional
	Create terra.BoolValue `hcl:"create,attr"`
	// Delete: bool, optional
	Delete terra.BoolValue `hcl:"delete,attr"`
	// List: bool, optional
	List terra.BoolValue `hcl:"list,attr"`
	// Read: bool, optional
	Read terra.BoolValue `hcl:"read,attr"`
	// Write: bool, optional
	Write terra.BoolValue `hcl:"write,attr"`
}

type SshAuthorizedKey struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
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

type PermissionScopeAttributes struct {
	ref terra.Reference
}

func (ps PermissionScopeAttributes) InternalRef() terra.Reference {
	return ps.ref
}

func (ps PermissionScopeAttributes) InternalWithRef(ref terra.Reference) PermissionScopeAttributes {
	return PermissionScopeAttributes{ref: ref}
}

func (ps PermissionScopeAttributes) InternalTokens() hclwrite.Tokens {
	return ps.ref.InternalTokens()
}

func (ps PermissionScopeAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceString(ps.ref.Append("resource_name"))
}

func (ps PermissionScopeAttributes) Service() terra.StringValue {
	return terra.ReferenceString(ps.ref.Append("service"))
}

func (ps PermissionScopeAttributes) Permissions() terra.ListValue[PermissionsAttributes] {
	return terra.ReferenceList[PermissionsAttributes](ps.ref.Append("permissions"))
}

type PermissionsAttributes struct {
	ref terra.Reference
}

func (p PermissionsAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p PermissionsAttributes) InternalWithRef(ref terra.Reference) PermissionsAttributes {
	return PermissionsAttributes{ref: ref}
}

func (p PermissionsAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p PermissionsAttributes) Create() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("create"))
}

func (p PermissionsAttributes) Delete() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("delete"))
}

func (p PermissionsAttributes) List() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("list"))
}

func (p PermissionsAttributes) Read() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("read"))
}

func (p PermissionsAttributes) Write() terra.BoolValue {
	return terra.ReferenceBool(p.ref.Append("write"))
}

type SshAuthorizedKeyAttributes struct {
	ref terra.Reference
}

func (sak SshAuthorizedKeyAttributes) InternalRef() terra.Reference {
	return sak.ref
}

func (sak SshAuthorizedKeyAttributes) InternalWithRef(ref terra.Reference) SshAuthorizedKeyAttributes {
	return SshAuthorizedKeyAttributes{ref: ref}
}

func (sak SshAuthorizedKeyAttributes) InternalTokens() hclwrite.Tokens {
	return sak.ref.InternalTokens()
}

func (sak SshAuthorizedKeyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("description"))
}

func (sak SshAuthorizedKeyAttributes) Key() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("key"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type PermissionScopeState struct {
	ResourceName string             `json:"resource_name"`
	Service      string             `json:"service"`
	Permissions  []PermissionsState `json:"permissions"`
}

type PermissionsState struct {
	Create bool `json:"create"`
	Delete bool `json:"delete"`
	List   bool `json:"list"`
	Read   bool `json:"read"`
	Write  bool `json:"write"`
}

type SshAuthorizedKeyState struct {
	Description string `json:"description"`
	Key         string `json:"key"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
