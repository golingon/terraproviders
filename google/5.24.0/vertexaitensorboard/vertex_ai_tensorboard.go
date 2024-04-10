// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vertexaitensorboard

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type EncryptionSpec struct {
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

type EncryptionSpecAttributes struct {
	ref terra.Reference
}

func (es EncryptionSpecAttributes) InternalRef() (terra.Reference, error) {
	return es.ref, nil
}

func (es EncryptionSpecAttributes) InternalWithRef(ref terra.Reference) EncryptionSpecAttributes {
	return EncryptionSpecAttributes{ref: ref}
}

func (es EncryptionSpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return es.ref.InternalTokens()
}

func (es EncryptionSpecAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("kms_key_name"))
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

type EncryptionSpecState struct {
	KmsKeyName string `json:"kms_key_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
