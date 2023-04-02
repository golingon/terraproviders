// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package redshiftserverlessworkgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Endpoint struct {
	// VpcEndpoint: min=0
	VpcEndpoint []VpcEndpoint `hcl:"vpc_endpoint,block" validate:"min=0"`
}

type VpcEndpoint struct {
	// NetworkInterface: min=0
	NetworkInterface []NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
}

type NetworkInterface struct{}

type ConfigParameter struct {
	// ParameterKey: string, required
	ParameterKey terra.StringValue `hcl:"parameter_key,attr" validate:"required"`
	// ParameterValue: string, required
	ParameterValue terra.StringValue `hcl:"parameter_value,attr" validate:"required"`
}

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

func (e EndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("port"))
}

func (e EndpointAttributes) VpcEndpoint() terra.ListValue[VpcEndpointAttributes] {
	return terra.ReferenceAsList[VpcEndpointAttributes](e.ref.Append("vpc_endpoint"))
}

type VpcEndpointAttributes struct {
	ref terra.Reference
}

func (ve VpcEndpointAttributes) InternalRef() (terra.Reference, error) {
	return ve.ref, nil
}

func (ve VpcEndpointAttributes) InternalWithRef(ref terra.Reference) VpcEndpointAttributes {
	return VpcEndpointAttributes{ref: ref}
}

func (ve VpcEndpointAttributes) InternalTokens() hclwrite.Tokens {
	return ve.ref.InternalTokens()
}

func (ve VpcEndpointAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_endpoint_id"))
}

func (ve VpcEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ve.ref.Append("vpc_id"))
}

func (ve VpcEndpointAttributes) NetworkInterface() terra.ListValue[NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[NetworkInterfaceAttributes](ve.ref.Append("network_interface"))
}

type NetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni NetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni NetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAttributes {
	return NetworkInterfaceAttributes{ref: ref}
}

func (ni NetworkInterfaceAttributes) InternalTokens() hclwrite.Tokens {
	return ni.ref.InternalTokens()
}

func (ni NetworkInterfaceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("availability_zone"))
}

func (ni NetworkInterfaceAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("network_interface_id"))
}

func (ni NetworkInterfaceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("private_ip_address"))
}

func (ni NetworkInterfaceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet_id"))
}

type ConfigParameterAttributes struct {
	ref terra.Reference
}

func (cp ConfigParameterAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ConfigParameterAttributes) InternalWithRef(ref terra.Reference) ConfigParameterAttributes {
	return ConfigParameterAttributes{ref: ref}
}

func (cp ConfigParameterAttributes) InternalTokens() hclwrite.Tokens {
	return cp.ref.InternalTokens()
}

func (cp ConfigParameterAttributes) ParameterKey() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("parameter_key"))
}

func (cp ConfigParameterAttributes) ParameterValue() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("parameter_value"))
}

type EndpointState struct {
	Address     string             `json:"address"`
	Port        float64            `json:"port"`
	VpcEndpoint []VpcEndpointState `json:"vpc_endpoint"`
}

type VpcEndpointState struct {
	VpcEndpointId    string                  `json:"vpc_endpoint_id"`
	VpcId            string                  `json:"vpc_id"`
	NetworkInterface []NetworkInterfaceState `json:"network_interface"`
}

type NetworkInterfaceState struct {
	AvailabilityZone   string `json:"availability_zone"`
	NetworkInterfaceId string `json:"network_interface_id"`
	PrivateIpAddress   string `json:"private_ip_address"`
	SubnetId           string `json:"subnet_id"`
}

type ConfigParameterState struct {
	ParameterKey   string `json:"parameter_key"`
	ParameterValue string `json:"parameter_value"`
}
