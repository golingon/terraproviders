// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package logzsubaccounttagrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type TagFilter struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TagFilterAttributes struct {
	ref terra.Reference
}

func (tf TagFilterAttributes) InternalRef() (terra.Reference, error) {
	return tf.ref, nil
}

func (tf TagFilterAttributes) InternalWithRef(ref terra.Reference) TagFilterAttributes {
	return TagFilterAttributes{ref: ref}
}

func (tf TagFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tf.ref.InternalTokens()
}

func (tf TagFilterAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("action"))
}

func (tf TagFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("name"))
}

func (tf TagFilterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(tf.ref.Append("value"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type TagFilterState struct {
	Action string `json:"action"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
