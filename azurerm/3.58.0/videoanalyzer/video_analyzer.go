// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package videoanalyzer

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct {
	// IdentityIds: set of string, required
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type StorageAccount struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// UserAssignedIdentityId: string, required
	UserAssignedIdentityId terra.StringValue `hcl:"user_assigned_identity_id,attr" validate:"required"`
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

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type StorageAccountAttributes struct {
	ref terra.Reference
}

func (sa StorageAccountAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa StorageAccountAttributes) InternalWithRef(ref terra.Reference) StorageAccountAttributes {
	return StorageAccountAttributes{ref: ref}
}

func (sa StorageAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa StorageAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

func (sa StorageAccountAttributes) UserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("user_assigned_identity_id"))
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

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	Type        string   `json:"type"`
}

type StorageAccountState struct {
	Id                     string `json:"id"`
	UserAssignedIdentityId string `json:"user_assigned_identity_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
