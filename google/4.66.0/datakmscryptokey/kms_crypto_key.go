// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakmscryptokey

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type VersionTemplate struct{}

type VersionTemplateAttributes struct {
	ref terra.Reference
}

func (vt VersionTemplateAttributes) InternalRef() (terra.Reference, error) {
	return vt.ref, nil
}

func (vt VersionTemplateAttributes) InternalWithRef(ref terra.Reference) VersionTemplateAttributes {
	return VersionTemplateAttributes{ref: ref}
}

func (vt VersionTemplateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vt.ref.InternalTokens()
}

func (vt VersionTemplateAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(vt.ref.Append("algorithm"))
}

func (vt VersionTemplateAttributes) ProtectionLevel() terra.StringValue {
	return terra.ReferenceAsString(vt.ref.Append("protection_level"))
}

type VersionTemplateState struct {
	Algorithm       string `json:"algorithm"`
	ProtectionLevel string `json:"protection_level"`
}
