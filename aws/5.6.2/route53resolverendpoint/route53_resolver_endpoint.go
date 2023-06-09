// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package route53resolverendpoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IpAddress struct {
	// Ip: string, optional
	Ip terra.StringValue `hcl:"ip,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type IpAddressAttributes struct {
	ref terra.Reference
}

func (ia IpAddressAttributes) InternalRef() (terra.Reference, error) {
	return ia.ref, nil
}

func (ia IpAddressAttributes) InternalWithRef(ref terra.Reference) IpAddressAttributes {
	return IpAddressAttributes{ref: ref}
}

func (ia IpAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ia.ref.InternalTokens()
}

func (ia IpAddressAttributes) Ip() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("ip"))
}

func (ia IpAddressAttributes) IpId() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("ip_id"))
}

func (ia IpAddressAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("subnet_id"))
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

type IpAddressState struct {
	Ip       string `json:"ip"`
	IpId     string `json:"ip_id"`
	SubnetId string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
