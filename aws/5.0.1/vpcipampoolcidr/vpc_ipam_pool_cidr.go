// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpcipampoolcidr

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CidrAuthorizationContext struct {
	// Message: string, optional
	Message terra.StringValue `hcl:"message,attr"`
	// Signature: string, optional
	Signature terra.StringValue `hcl:"signature,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type CidrAuthorizationContextAttributes struct {
	ref terra.Reference
}

func (cac CidrAuthorizationContextAttributes) InternalRef() (terra.Reference, error) {
	return cac.ref, nil
}

func (cac CidrAuthorizationContextAttributes) InternalWithRef(ref terra.Reference) CidrAuthorizationContextAttributes {
	return CidrAuthorizationContextAttributes{ref: ref}
}

func (cac CidrAuthorizationContextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cac.ref.InternalTokens()
}

func (cac CidrAuthorizationContextAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("message"))
}

func (cac CidrAuthorizationContextAttributes) Signature() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("signature"))
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

type CidrAuthorizationContextState struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
