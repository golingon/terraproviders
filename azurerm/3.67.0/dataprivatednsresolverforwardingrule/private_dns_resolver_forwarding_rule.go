// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataprivatednsresolverforwardingrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type TargetDnsServers struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type TargetDnsServersState struct {
	IpAddress string  `json:"ip_address"`
	Port      float64 `json:"port"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
