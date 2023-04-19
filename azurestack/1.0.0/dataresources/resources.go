// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataresources

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Resources struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ResourcesAttributes struct {
	ref terra.Reference
}

func (r ResourcesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResourcesAttributes) InternalWithRef(ref terra.Reference) ResourcesAttributes {
	return ResourcesAttributes{ref: ref}
}

func (r ResourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResourcesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("id"))
}

func (r ResourcesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("location"))
}

func (r ResourcesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r ResourcesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("tags"))
}

func (r ResourcesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ResourcesState struct {
	Id       string            `json:"id"`
	Location string            `json:"location"`
	Name     string            `json:"name"`
	Tags     map[string]string `json:"tags"`
	Type     string            `json:"type"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
