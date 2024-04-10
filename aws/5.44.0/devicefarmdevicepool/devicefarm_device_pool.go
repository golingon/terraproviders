// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package devicefarmdevicepool

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Rule struct {
	// Attribute: string, optional
	Attribute terra.StringValue `hcl:"attribute,attr"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Attribute() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("attribute"))
}

func (r RuleAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("operator"))
}

func (r RuleAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("value"))
}

type RuleState struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}
