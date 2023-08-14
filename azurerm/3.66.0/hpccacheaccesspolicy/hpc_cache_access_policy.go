// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package hpccacheaccesspolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccessRule struct {
	// Access: string, required
	Access terra.StringValue `hcl:"access,attr" validate:"required"`
	// AnonymousGid: number, optional
	AnonymousGid terra.NumberValue `hcl:"anonymous_gid,attr"`
	// AnonymousUid: number, optional
	AnonymousUid terra.NumberValue `hcl:"anonymous_uid,attr"`
	// Filter: string, optional
	Filter terra.StringValue `hcl:"filter,attr"`
	// RootSquashEnabled: bool, optional
	RootSquashEnabled terra.BoolValue `hcl:"root_squash_enabled,attr"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
	// SubmountAccessEnabled: bool, optional
	SubmountAccessEnabled terra.BoolValue `hcl:"submount_access_enabled,attr"`
	// SuidEnabled: bool, optional
	SuidEnabled terra.BoolValue `hcl:"suid_enabled,attr"`
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

type AccessRuleAttributes struct {
	ref terra.Reference
}

func (ar AccessRuleAttributes) InternalRef() (terra.Reference, error) {
	return ar.ref, nil
}

func (ar AccessRuleAttributes) InternalWithRef(ref terra.Reference) AccessRuleAttributes {
	return AccessRuleAttributes{ref: ref}
}

func (ar AccessRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ar.ref.InternalTokens()
}

func (ar AccessRuleAttributes) Access() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("access"))
}

func (ar AccessRuleAttributes) AnonymousGid() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("anonymous_gid"))
}

func (ar AccessRuleAttributes) AnonymousUid() terra.NumberValue {
	return terra.ReferenceAsNumber(ar.ref.Append("anonymous_uid"))
}

func (ar AccessRuleAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("filter"))
}

func (ar AccessRuleAttributes) RootSquashEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("root_squash_enabled"))
}

func (ar AccessRuleAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("scope"))
}

func (ar AccessRuleAttributes) SubmountAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("submount_access_enabled"))
}

func (ar AccessRuleAttributes) SuidEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ar.ref.Append("suid_enabled"))
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

type AccessRuleState struct {
	Access                string  `json:"access"`
	AnonymousGid          float64 `json:"anonymous_gid"`
	AnonymousUid          float64 `json:"anonymous_uid"`
	Filter                string  `json:"filter"`
	RootSquashEnabled     bool    `json:"root_squash_enabled"`
	Scope                 string  `json:"scope"`
	SubmountAccessEnabled bool    `json:"submount_access_enabled"`
	SuidEnabled           bool    `json:"suid_enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}