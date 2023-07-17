// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputehavpngateway

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type VpnInterfaces struct{}

type VpnInterfacesAttributes struct {
	ref terra.Reference
}

func (vi VpnInterfacesAttributes) InternalRef() (terra.Reference, error) {
	return vi.ref, nil
}

func (vi VpnInterfacesAttributes) InternalWithRef(ref terra.Reference) VpnInterfacesAttributes {
	return VpnInterfacesAttributes{ref: ref}
}

func (vi VpnInterfacesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vi.ref.InternalTokens()
}

func (vi VpnInterfacesAttributes) Id() terra.NumberValue {
	return terra.ReferenceAsNumber(vi.ref.Append("id"))
}

func (vi VpnInterfacesAttributes) InterconnectAttachment() terra.StringValue {
	return terra.ReferenceAsString(vi.ref.Append("interconnect_attachment"))
}

func (vi VpnInterfacesAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(vi.ref.Append("ip_address"))
}

type VpnInterfacesState struct {
	Id                     float64 `json:"id"`
	InterconnectAttachment string  `json:"interconnect_attachment"`
	IpAddress              string  `json:"ip_address"`
}