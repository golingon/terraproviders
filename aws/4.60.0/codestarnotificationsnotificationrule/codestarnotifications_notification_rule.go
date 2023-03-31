// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package codestarnotificationsnotificationrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Target struct {
	// Address: string, required
	Address terra.StringValue `hcl:"address,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type TargetAttributes struct {
	ref terra.Reference
}

func (t TargetAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TargetAttributes) InternalWithRef(ref terra.Reference) TargetAttributes {
	return TargetAttributes{ref: ref}
}

func (t TargetAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TargetAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("address"))
}

func (t TargetAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("status"))
}

func (t TargetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("type"))
}

type TargetState struct {
	Address string `json:"address"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}
