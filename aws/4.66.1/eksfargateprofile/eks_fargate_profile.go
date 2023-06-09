// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package eksfargateprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Selector struct {
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type SelectorAttributes struct {
	ref terra.Reference
}

func (s SelectorAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SelectorAttributes) InternalWithRef(ref terra.Reference) SelectorAttributes {
	return SelectorAttributes{ref: ref}
}

func (s SelectorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SelectorAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("labels"))
}

func (s SelectorAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("namespace"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type SelectorState struct {
	Labels    map[string]string `json:"labels"`
	Namespace string            `json:"namespace"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
