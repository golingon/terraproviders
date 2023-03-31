// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dedicatedhardwaresecuritymodule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ManagementNetworkProfile struct {
	// NetworkInterfacePrivateIpAddresses: set of string, required
	NetworkInterfacePrivateIpAddresses terra.SetValue[terra.StringValue] `hcl:"network_interface_private_ip_addresses,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type NetworkProfile struct {
	// NetworkInterfacePrivateIpAddresses: set of string, required
	NetworkInterfacePrivateIpAddresses terra.SetValue[terra.StringValue] `hcl:"network_interface_private_ip_addresses,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
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

type ManagementNetworkProfileAttributes struct {
	ref terra.Reference
}

func (mnp ManagementNetworkProfileAttributes) InternalRef() terra.Reference {
	return mnp.ref
}

func (mnp ManagementNetworkProfileAttributes) InternalWithRef(ref terra.Reference) ManagementNetworkProfileAttributes {
	return ManagementNetworkProfileAttributes{ref: ref}
}

func (mnp ManagementNetworkProfileAttributes) InternalTokens() hclwrite.Tokens {
	return mnp.ref.InternalTokens()
}

func (mnp ManagementNetworkProfileAttributes) NetworkInterfacePrivateIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](mnp.ref.Append("network_interface_private_ip_addresses"))
}

func (mnp ManagementNetworkProfileAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(mnp.ref.Append("subnet_id"))
}

type NetworkProfileAttributes struct {
	ref terra.Reference
}

func (np NetworkProfileAttributes) InternalRef() terra.Reference {
	return np.ref
}

func (np NetworkProfileAttributes) InternalWithRef(ref terra.Reference) NetworkProfileAttributes {
	return NetworkProfileAttributes{ref: ref}
}

func (np NetworkProfileAttributes) InternalTokens() hclwrite.Tokens {
	return np.ref.InternalTokens()
}

func (np NetworkProfileAttributes) NetworkInterfacePrivateIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](np.ref.Append("network_interface_private_ip_addresses"))
}

func (np NetworkProfileAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(np.ref.Append("subnet_id"))
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

type ManagementNetworkProfileState struct {
	NetworkInterfacePrivateIpAddresses []string `json:"network_interface_private_ip_addresses"`
	SubnetId                           string   `json:"subnet_id"`
}

type NetworkProfileState struct {
	NetworkInterfacePrivateIpAddresses []string `json:"network_interface_private_ip_addresses"`
	SubnetId                           string   `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
