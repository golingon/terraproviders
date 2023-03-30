// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appconfigconfigurationprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Validator struct {
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ValidatorAttributes struct {
	ref terra.Reference
}

func (v ValidatorAttributes) InternalRef() terra.Reference {
	return v.ref
}

func (v ValidatorAttributes) InternalWithRef(ref terra.Reference) ValidatorAttributes {
	return ValidatorAttributes{ref: ref}
}

func (v ValidatorAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v ValidatorAttributes) Content() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("content"))
}

func (v ValidatorAttributes) Type() terra.StringValue {
	return terra.ReferenceString(v.ref.Append("type"))
}

type ValidatorState struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}
