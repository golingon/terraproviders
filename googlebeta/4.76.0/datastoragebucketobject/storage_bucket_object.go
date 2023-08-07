// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datastoragebucketobject

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomerEncryption struct{}

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

type CustomerEncryptionState struct {
	EncryptionAlgorithm string `json:"encryption_algorithm"`
	EncryptionKey       string `json:"encryption_key"`
}
