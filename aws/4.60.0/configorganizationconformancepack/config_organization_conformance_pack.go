// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package configorganizationconformancepack

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InputParameter struct {
	// ParameterName: string, required
	ParameterName terra.StringValue `hcl:"parameter_name,attr" validate:"required"`
	// ParameterValue: string, required
	ParameterValue terra.StringValue `hcl:"parameter_value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type InputParameterAttributes struct {
	ref terra.Reference
}

func (ip InputParameterAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip InputParameterAttributes) InternalWithRef(ref terra.Reference) InputParameterAttributes {
	return InputParameterAttributes{ref: ref}
}

func (ip InputParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip InputParameterAttributes) ParameterName() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("parameter_name"))
}

func (ip InputParameterAttributes) ParameterValue() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("parameter_value"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type InputParameterState struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
