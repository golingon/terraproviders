// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_storage_share

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Acl struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// AclAccessPolicy: min=0
	AccessPolicy []AclAccessPolicy `hcl:"access_policy,block" validate:"min=0"`
}

type AclAccessPolicy struct {
	// Expiry: string, optional
	Expiry terra.StringValue `hcl:"expiry,attr"`
	// Permissions: string, required
	Permissions terra.StringValue `hcl:"permissions,attr" validate:"required"`
	// Start: string, optional
	Start terra.StringValue `hcl:"start,attr"`
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

type AclAttributes struct {
	ref terra.Reference
}

func (a AclAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AclAttributes) InternalWithRef(ref terra.Reference) AclAttributes {
	return AclAttributes{ref: ref}
}

func (a AclAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

func (a AclAttributes) AccessPolicy() terra.ListValue[AclAccessPolicyAttributes] {
	return terra.ReferenceAsList[AclAccessPolicyAttributes](a.ref.Append("access_policy"))
}

type AclAccessPolicyAttributes struct {
	ref terra.Reference
}

func (ap AclAccessPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AclAccessPolicyAttributes) InternalWithRef(ref terra.Reference) AclAccessPolicyAttributes {
	return AclAccessPolicyAttributes{ref: ref}
}

func (ap AclAccessPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AclAccessPolicyAttributes) Expiry() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("expiry"))
}

func (ap AclAccessPolicyAttributes) Permissions() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("permissions"))
}

func (ap AclAccessPolicyAttributes) Start() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("start"))
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

type AclState struct {
	Id           string                 `json:"id"`
	AccessPolicy []AclAccessPolicyState `json:"access_policy"`
}

type AclAccessPolicyState struct {
	Expiry      string `json:"expiry"`
	Permissions string `json:"permissions"`
	Start       string `json:"start"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
