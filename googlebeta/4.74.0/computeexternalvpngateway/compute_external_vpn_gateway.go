// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computeexternalvpngateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Interface struct {
	// Id: number, optional
	Id terra.NumberValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type InterfaceAttributes struct {
	ref terra.Reference
}

func (i InterfaceAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i InterfaceAttributes) InternalWithRef(ref terra.Reference) InterfaceAttributes {
	return InterfaceAttributes{ref: ref}
}

func (i InterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i InterfaceAttributes) Id() terra.NumberValue {
	return terra.ReferenceAsNumber(i.ref.Append("id"))
}

func (i InterfaceAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("ip_address"))
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

type InterfaceState struct {
	Id        float64 `json:"id"`
	IpAddress string  `json:"ip_address"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
