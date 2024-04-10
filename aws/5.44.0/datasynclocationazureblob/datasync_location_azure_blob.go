// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datasynclocationazureblob

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SasConfiguration struct {
	// Token: string, required
	Token terra.StringValue `hcl:"token,attr" validate:"required"`
}

type SasConfigurationAttributes struct {
	ref terra.Reference
}

func (sc SasConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SasConfigurationAttributes) InternalWithRef(ref terra.Reference) SasConfigurationAttributes {
	return SasConfigurationAttributes{ref: ref}
}

func (sc SasConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SasConfigurationAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("token"))
}

type SasConfigurationState struct {
	Token string `json:"token"`
}
