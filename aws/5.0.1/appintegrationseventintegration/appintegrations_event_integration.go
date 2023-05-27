// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appintegrationseventintegration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EventFilter struct {
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
}

type EventFilterAttributes struct {
	ref terra.Reference
}

func (ef EventFilterAttributes) InternalRef() (terra.Reference, error) {
	return ef.ref, nil
}

func (ef EventFilterAttributes) InternalWithRef(ref terra.Reference) EventFilterAttributes {
	return EventFilterAttributes{ref: ref}
}

func (ef EventFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ef.ref.InternalTokens()
}

func (ef EventFilterAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("source"))
}

type EventFilterState struct {
	Source string `json:"source"`
}
