// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_redshiftserverless_endpoint_access

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type VpcEndpointAttributes struct {
	ref terra.Reference
}

func (ve VpcEndpointAttributes) InternalRef() (terra.Reference, error) {
	return ve.ref, nil
}

func (ve VpcEndpointAttributes) InternalWithRef(ref terra.Reference) VpcEndpointAttributes {
	return VpcEndpointAttributes{ref: ref}
}

func (ve VpcEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ve.ref.InternalTokens()
}

func (ve VpcEndpointAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_endpoint_id"))
}

func (ve VpcEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_id"))
}

func (ve VpcEndpointAttributes) NetworkInterface() terra.ListValue[VpcEndpointNetworkInterfaceAttributes] {
	return terra.ReferenceAsList[VpcEndpointNetworkInterfaceAttributes](ve.ref.Append("network_interface"))
}

type VpcEndpointNetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni VpcEndpointNetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni VpcEndpointNetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) VpcEndpointNetworkInterfaceAttributes {
	return VpcEndpointNetworkInterfaceAttributes{ref: ref}
}

func (ni VpcEndpointNetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni VpcEndpointNetworkInterfaceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("availability_zone"))
}

func (ni VpcEndpointNetworkInterfaceAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_interface_id"))
}

func (ni VpcEndpointNetworkInterfaceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_ip_address"))
}

func (ni VpcEndpointNetworkInterfaceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet_id"))
}

type VpcEndpointState struct {
	VpcEndpointId    string                             `json:"vpc_endpoint_id"`
	VpcId            string                             `json:"vpc_id"`
	NetworkInterface []VpcEndpointNetworkInterfaceState `json:"network_interface"`
}

type VpcEndpointNetworkInterfaceState struct {
	AvailabilityZone   string `json:"availability_zone"`
	NetworkInterfaceId string `json:"network_interface_id"`
	PrivateIpAddress   string `json:"private_ip_address"`
	SubnetId           string `json:"subnet_id"`
}
