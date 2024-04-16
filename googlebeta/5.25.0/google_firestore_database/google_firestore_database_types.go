// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_firestore_database

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CmekConfig struct {
	// KmsKeyName: string, required
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CmekConfigAttributes struct {
	ref terra.Reference
}

func (cc CmekConfigAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CmekConfigAttributes) InternalWithRef(ref terra.Reference) CmekConfigAttributes {
	return CmekConfigAttributes{ref: ref}
}

func (cc CmekConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CmekConfigAttributes) ActiveKeyVersion() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("active_key_version"))
}

func (cc CmekConfigAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("kms_key_name"))
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

type CmekConfigState struct {
	ActiveKeyVersion []string `json:"active_key_version"`
	KmsKeyName       string   `json:"kms_key_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
