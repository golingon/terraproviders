// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appmeshmesh

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Spec struct {
	// EgressFilter: optional
	EgressFilter *EgressFilter `hcl:"egress_filter,block"`
}

type EgressFilter struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type SpecAttributes struct {
	ref terra.Reference
}

func (s SpecAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SpecAttributes) InternalWithRef(ref terra.Reference) SpecAttributes {
	return SpecAttributes{ref: ref}
}

func (s SpecAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SpecAttributes) EgressFilter() terra.ListValue[EgressFilterAttributes] {
	return terra.ReferenceList[EgressFilterAttributes](s.ref.Append("egress_filter"))
}

type EgressFilterAttributes struct {
	ref terra.Reference
}

func (ef EgressFilterAttributes) InternalRef() terra.Reference {
	return ef.ref
}

func (ef EgressFilterAttributes) InternalWithRef(ref terra.Reference) EgressFilterAttributes {
	return EgressFilterAttributes{ref: ref}
}

func (ef EgressFilterAttributes) InternalTokens() hclwrite.Tokens {
	return ef.ref.InternalTokens()
}

func (ef EgressFilterAttributes) Type() terra.StringValue {
	return terra.ReferenceString(ef.ref.Append("type"))
}

type SpecState struct {
	EgressFilter []EgressFilterState `json:"egress_filter"`
}

type EgressFilterState struct {
	Type string `json:"type"`
}
