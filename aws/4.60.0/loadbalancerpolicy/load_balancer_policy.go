// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package loadbalancerpolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PolicyAttribute struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type PolicyAttributeAttributes struct {
	ref terra.Reference
}

func (pa PolicyAttributeAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PolicyAttributeAttributes) InternalWithRef(ref terra.Reference) PolicyAttributeAttributes {
	return PolicyAttributeAttributes{ref: ref}
}

func (pa PolicyAttributeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pa.ref.InternalTokens()
}

func (pa PolicyAttributeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("name"))
}

func (pa PolicyAttributeAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("value"))
}

type PolicyAttributeState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
