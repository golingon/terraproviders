// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package storagebucketobject

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomerEncryption struct {
	// EncryptionAlgorithm: string, optional
	EncryptionAlgorithm terra.StringValue `hcl:"encryption_algorithm,attr"`
	// EncryptionKey: string, required
	EncryptionKey terra.StringValue `hcl:"encryption_key,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type CustomerEncryptionAttributes struct {
	ref terra.Reference
}

func (ce CustomerEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return ce.ref, nil
}

func (ce CustomerEncryptionAttributes) InternalWithRef(ref terra.Reference) CustomerEncryptionAttributes {
	return CustomerEncryptionAttributes{ref: ref}
}

func (ce CustomerEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ce.ref.InternalTokens()
}

func (ce CustomerEncryptionAttributes) EncryptionAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("encryption_algorithm"))
}

func (ce CustomerEncryptionAttributes) EncryptionKey() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("encryption_key"))
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

type CustomerEncryptionState struct {
	EncryptionAlgorithm string `json:"encryption_algorithm"`
	EncryptionKey       string `json:"encryption_key"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}