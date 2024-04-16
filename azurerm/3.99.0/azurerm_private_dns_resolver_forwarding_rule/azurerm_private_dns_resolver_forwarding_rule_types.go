// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_dns_resolver_forwarding_rule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type TargetDnsServers struct {
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
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

type TargetDnsServersAttributes struct {
	ref terra.Reference
}

func (tds TargetDnsServersAttributes) InternalRef() (terra.Reference, error) {
	return tds.ref, nil
}

func (tds TargetDnsServersAttributes) InternalWithRef(ref terra.Reference) TargetDnsServersAttributes {
	return TargetDnsServersAttributes{ref: ref}
}

func (tds TargetDnsServersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tds.ref.InternalTokens()
}

func (tds TargetDnsServersAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(tds.ref.Append("ip_address"))
}

func (tds TargetDnsServersAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(tds.ref.Append("port"))
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

type TargetDnsServersState struct {
	IpAddress string  `json:"ip_address"`
	Port      float64 `json:"port"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
