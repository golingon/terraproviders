// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package automationsourcecontrol

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Security struct {
	// RefreshToken: string, optional
	RefreshToken terra.StringValue `hcl:"refresh_token,attr"`
	// Token: string, required
	Token terra.StringValue `hcl:"token,attr" validate:"required"`
	// TokenType: string, required
	TokenType terra.StringValue `hcl:"token_type,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SecurityAttributes struct {
	ref terra.Reference
}

func (s SecurityAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecurityAttributes) InternalWithRef(ref terra.Reference) SecurityAttributes {
	return SecurityAttributes{ref: ref}
}

func (s SecurityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecurityAttributes) RefreshToken() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("refresh_token"))
}

func (s SecurityAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("token"))
}

func (s SecurityAttributes) TokenType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("token_type"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type SecurityState struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"token"`
	TokenType    string `json:"token_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
