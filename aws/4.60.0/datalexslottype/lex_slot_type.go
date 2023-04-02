// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalexslottype

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EnumerationValue struct{}

type EnumerationValueAttributes struct {
	ref terra.Reference
}

func (ev EnumerationValueAttributes) InternalRef() (terra.Reference, error) {
	return ev.ref, nil
}

func (ev EnumerationValueAttributes) InternalWithRef(ref terra.Reference) EnumerationValueAttributes {
	return EnumerationValueAttributes{ref: ref}
}

func (ev EnumerationValueAttributes) InternalTokens() hclwrite.Tokens {
	return ev.ref.InternalTokens()
}

func (ev EnumerationValueAttributes) Synonyms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ev.ref.Append("synonyms"))
}

func (ev EnumerationValueAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("value"))
}

type EnumerationValueState struct {
	Synonyms []string `json:"synonyms"`
	Value    string   `json:"value"`
}
