// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package spannerdatabase

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfig struct {
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

type EncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigAttributes {
	return EncryptionConfigAttributes{ref: ref}
}

func (ec EncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key_name"))
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

type EncryptionConfigState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
