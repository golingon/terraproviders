// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package wafregionalipset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IpSetDescriptor struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type IpSetDescriptorAttributes struct {
	ref terra.Reference
}

func (isd IpSetDescriptorAttributes) InternalRef() (terra.Reference, error) {
	return isd.ref, nil
}

func (isd IpSetDescriptorAttributes) InternalWithRef(ref terra.Reference) IpSetDescriptorAttributes {
	return IpSetDescriptorAttributes{ref: ref}
}

func (isd IpSetDescriptorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return isd.ref.InternalTokens()
}

func (isd IpSetDescriptorAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(isd.ref.Append("type"))
}

func (isd IpSetDescriptorAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(isd.ref.Append("value"))
}

type IpSetDescriptorState struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}