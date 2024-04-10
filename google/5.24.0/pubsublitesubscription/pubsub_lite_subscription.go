// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package pubsublitesubscription

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DeliveryConfig struct {
	// DeliveryRequirement: string, required
	DeliveryRequirement terra.StringValue `hcl:"delivery_requirement,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DeliveryConfigAttributes struct {
	ref terra.Reference
}

func (dc DeliveryConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DeliveryConfigAttributes) InternalWithRef(ref terra.Reference) DeliveryConfigAttributes {
	return DeliveryConfigAttributes{ref: ref}
}

func (dc DeliveryConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DeliveryConfigAttributes) DeliveryRequirement() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("delivery_requirement"))
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

type DeliveryConfigState struct {
	DeliveryRequirement string `json:"delivery_requirement"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
