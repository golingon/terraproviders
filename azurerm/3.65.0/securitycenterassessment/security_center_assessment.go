// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package securitycenterassessment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Status struct {
	// Cause: string, optional
	Cause terra.StringValue `hcl:"cause,attr"`
	// Code: string, required
	Code terra.StringValue `hcl:"code,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
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

type StatusAttributes struct {
	ref terra.Reference
}

func (s StatusAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatusAttributes) InternalWithRef(ref terra.Reference) StatusAttributes {
	return StatusAttributes{ref: ref}
}

func (s StatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatusAttributes) Cause() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("cause"))
}

func (s StatusAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("code"))
}

func (s StatusAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
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

type StatusState struct {
	Cause       string `json:"cause"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
