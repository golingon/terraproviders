// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kmsgrant

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Constraints struct {
	// EncryptionContextEquals: map of string, optional
	EncryptionContextEquals terra.MapValue[terra.StringValue] `hcl:"encryption_context_equals,attr"`
	// EncryptionContextSubset: map of string, optional
	EncryptionContextSubset terra.MapValue[terra.StringValue] `hcl:"encryption_context_subset,attr"`
}

type ConstraintsAttributes struct {
	ref terra.Reference
}

func (c ConstraintsAttributes) InternalRef() terra.Reference {
	return c.ref
}

func (c ConstraintsAttributes) InternalWithRef(ref terra.Reference) ConstraintsAttributes {
	return ConstraintsAttributes{ref: ref}
}

func (c ConstraintsAttributes) InternalTokens() hclwrite.Tokens {
	return c.ref.InternalTokens()
}

func (c ConstraintsAttributes) EncryptionContextEquals() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("encryption_context_equals"))
}

func (c ConstraintsAttributes) EncryptionContextSubset() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("encryption_context_subset"))
}

type ConstraintsState struct {
	EncryptionContextEquals map[string]string `json:"encryption_context_equals"`
	EncryptionContextSubset map[string]string `json:"encryption_context_subset"`
}
