// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package voiceservicescommunicationsgateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ServiceLocation struct {
	// AllowedMediaSourceAddressPrefixes: set of string, optional
	AllowedMediaSourceAddressPrefixes terra.SetValue[terra.StringValue] `hcl:"allowed_media_source_address_prefixes,attr"`
	// AllowedSignalingSourceAddressPrefixes: set of string, optional
	AllowedSignalingSourceAddressPrefixes terra.SetValue[terra.StringValue] `hcl:"allowed_signaling_source_address_prefixes,attr"`
	// EsrpAddresses: set of string, optional
	EsrpAddresses terra.SetValue[terra.StringValue] `hcl:"esrp_addresses,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// OperatorAddresses: set of string, required
	OperatorAddresses terra.SetValue[terra.StringValue] `hcl:"operator_addresses,attr" validate:"required"`
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

type ServiceLocationAttributes struct {
	ref terra.Reference
}

func (sl ServiceLocationAttributes) InternalRef() terra.Reference {
	return sl.ref
}

func (sl ServiceLocationAttributes) InternalWithRef(ref terra.Reference) ServiceLocationAttributes {
	return ServiceLocationAttributes{ref: ref}
}

func (sl ServiceLocationAttributes) InternalTokens() hclwrite.Tokens {
	return sl.ref.InternalTokens()
}

func (sl ServiceLocationAttributes) AllowedMediaSourceAddressPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sl.ref.Append("allowed_media_source_address_prefixes"))
}

func (sl ServiceLocationAttributes) AllowedSignalingSourceAddressPrefixes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sl.ref.Append("allowed_signaling_source_address_prefixes"))
}

func (sl ServiceLocationAttributes) EsrpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sl.ref.Append("esrp_addresses"))
}

func (sl ServiceLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceString(sl.ref.Append("location"))
}

func (sl ServiceLocationAttributes) OperatorAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sl.ref.Append("operator_addresses"))
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

type ServiceLocationState struct {
	AllowedMediaSourceAddressPrefixes     []string `json:"allowed_media_source_address_prefixes"`
	AllowedSignalingSourceAddressPrefixes []string `json:"allowed_signaling_source_address_prefixes"`
	EsrpAddresses                         []string `json:"esrp_addresses"`
	Location                              string   `json:"location"`
	OperatorAddresses                     []string `json:"operator_addresses"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
