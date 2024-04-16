// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53domains_delegation_signer_record

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SigningAttributes struct {
	// Algorithm: number, required
	Algorithm terra.NumberValue `hcl:"algorithm,attr" validate:"required"`
	// Flags: number, required
	Flags terra.NumberValue `hcl:"flags,attr" validate:"required"`
	// PublicKey: string, required
	PublicKey terra.StringValue `hcl:"public_key,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type SigningAttributesAttributes struct {
	ref terra.Reference
}

func (sa SigningAttributesAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa SigningAttributesAttributes) InternalWithRef(ref terra.Reference) SigningAttributesAttributes {
	return SigningAttributesAttributes{ref: ref}
}

func (sa SigningAttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa SigningAttributesAttributes) Algorithm() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("algorithm"))
}

func (sa SigningAttributesAttributes) Flags() terra.NumberValue {
	return terra.ReferenceAsNumber(sa.ref.Append("flags"))
}

func (sa SigningAttributesAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("public_key"))
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

type SigningAttributesState struct {
	Algorithm float64 `json:"algorithm"`
	Flags     float64 `json:"flags"`
	PublicKey string  `json:"public_key"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
