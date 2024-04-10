// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package privatednsmxrecord

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Record struct {
	// Exchange: string, required
	Exchange terra.StringValue `hcl:"exchange,attr" validate:"required"`
	// Preference: number, required
	Preference terra.NumberValue `hcl:"preference,attr" validate:"required"`
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

type RecordAttributes struct {
	ref terra.Reference
}

func (r RecordAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RecordAttributes) InternalWithRef(ref terra.Reference) RecordAttributes {
	return RecordAttributes{ref: ref}
}

func (r RecordAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RecordAttributes) Exchange() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("exchange"))
}

func (r RecordAttributes) Preference() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("preference"))
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

type RecordState struct {
	Exchange   string  `json:"exchange"`
	Preference float64 `json:"preference"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
