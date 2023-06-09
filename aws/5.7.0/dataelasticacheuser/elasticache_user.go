// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataelasticacheuser

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AuthenticationMode struct {
	// PasswordCount: number, optional
	PasswordCount terra.NumberValue `hcl:"password_count,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type AuthenticationModeAttributes struct {
	ref terra.Reference
}

func (am AuthenticationModeAttributes) InternalRef() (terra.Reference, error) {
	return am.ref, nil
}

func (am AuthenticationModeAttributes) InternalWithRef(ref terra.Reference) AuthenticationModeAttributes {
	return AuthenticationModeAttributes{ref: ref}
}

func (am AuthenticationModeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return am.ref.InternalTokens()
}

func (am AuthenticationModeAttributes) PasswordCount() terra.NumberValue {
	return terra.ReferenceAsNumber(am.ref.Append("password_count"))
}

func (am AuthenticationModeAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(am.ref.Append("type"))
}

type AuthenticationModeState struct {
	PasswordCount float64 `json:"password_count"`
	Type          string  `json:"type"`
}
