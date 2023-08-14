// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kmscryptokey

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VersionTemplate struct {
	// Algorithm: string, required
	Algorithm terra.StringValue `hcl:"algorithm,attr" validate:"required"`
	// ProtectionLevel: string, optional
	ProtectionLevel terra.StringValue `hcl:"protection_level,attr"`
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

type VersionTemplateAttributes struct {
	ref terra.Reference
}

func (vt VersionTemplateAttributes) InternalRef() (terra.Reference, error) {
	return vt.ref, nil
}

func (vt VersionTemplateAttributes) InternalWithRef(ref terra.Reference) VersionTemplateAttributes {
	return VersionTemplateAttributes{ref: ref}
}

func (vt VersionTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vt.ref.InternalTokens()
}

func (vt VersionTemplateAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(vt.ref.Append("algorithm"))
}

func (vt VersionTemplateAttributes) ProtectionLevel() terra.StringValue {
	return terra.ReferenceAsString(vt.ref.Append("protection_level"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VersionTemplateState struct {
	Algorithm       string `json:"algorithm"`
	ProtectionLevel string `json:"protection_level"`
}
