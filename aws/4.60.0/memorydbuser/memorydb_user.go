// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package memorydbuser

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AuthenticationMode struct {
	// Passwords: set of string, required
	Passwords terra.SetValue[terra.StringValue] `hcl:"passwords,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type AuthenticationModeAttributes struct {
	ref terra.Reference
}

func (am AuthenticationModeAttributes) InternalRef() terra.Reference {
	return am.ref
}

func (am AuthenticationModeAttributes) InternalWithRef(ref terra.Reference) AuthenticationModeAttributes {
	return AuthenticationModeAttributes{ref: ref}
}

func (am AuthenticationModeAttributes) InternalTokens() hclwrite.Tokens {
	return am.ref.InternalTokens()
}

func (am AuthenticationModeAttributes) PasswordCount() terra.NumberValue {
	return terra.ReferenceNumber(am.ref.Append("password_count"))
}

func (am AuthenticationModeAttributes) Passwords() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](am.ref.Append("passwords"))
}

func (am AuthenticationModeAttributes) Type() terra.StringValue {
	return terra.ReferenceString(am.ref.Append("type"))
}

type AuthenticationModeState struct {
	PasswordCount float64  `json:"password_count"`
	Passwords     []string `json:"passwords"`
	Type          string   `json:"type"`
}
