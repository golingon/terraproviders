// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_chime_voice_connector_termination_credentials

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Credentials struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type CredentialsAttributes struct {
	ref terra.Reference
}

func (c CredentialsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CredentialsAttributes) InternalWithRef(ref terra.Reference) CredentialsAttributes {
	return CredentialsAttributes{ref: ref}
}

func (c CredentialsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CredentialsAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("password"))
}

func (c CredentialsAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("username"))
}

type CredentialsState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
