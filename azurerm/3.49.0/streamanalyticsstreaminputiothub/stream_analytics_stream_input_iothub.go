// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package streamanalyticsstreaminputiothub

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Serialization struct {
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// FieldDelimiter: string, optional
	FieldDelimiter terra.StringValue `hcl:"field_delimiter,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
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

type SerializationAttributes struct {
	ref terra.Reference
}

func (s SerializationAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SerializationAttributes) InternalWithRef(ref terra.Reference) SerializationAttributes {
	return SerializationAttributes{ref: ref}
}

func (s SerializationAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SerializationAttributes) Encoding() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("encoding"))
}

func (s SerializationAttributes) FieldDelimiter() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("field_delimiter"))
}

func (s SerializationAttributes) Type() terra.StringValue {
	return terra.ReferenceString(s.ref.Append("type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type SerializationState struct {
	Encoding       string `json:"encoding"`
	FieldDelimiter string `json:"field_delimiter"`
	Type           string `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
