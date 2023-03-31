// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkinterface

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IpConfiguration struct {
	// GatewayLoadBalancerFrontendIpConfigurationId: string, optional
	GatewayLoadBalancerFrontendIpConfigurationId terra.StringValue `hcl:"gateway_load_balancer_frontend_ip_configuration_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Primary: bool, optional
	Primary terra.BoolValue `hcl:"primary,attr"`
	// PrivateIpAddress: string, optional
	PrivateIpAddress terra.StringValue `hcl:"private_ip_address,attr"`
	// PrivateIpAddressAllocation: string, required
	PrivateIpAddressAllocation terra.StringValue `hcl:"private_ip_address_allocation,attr" validate:"required"`
	// PrivateIpAddressVersion: string, optional
	PrivateIpAddressVersion terra.StringValue `hcl:"private_ip_address_version,attr"`
	// PublicIpAddressId: string, optional
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
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

type IpConfigurationAttributes struct {
	ref terra.Reference
}

func (ic IpConfigurationAttributes) InternalRef() terra.Reference {
	return ic.ref
}

func (ic IpConfigurationAttributes) InternalWithRef(ref terra.Reference) IpConfigurationAttributes {
	return IpConfigurationAttributes{ref: ref}
}

func (ic IpConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ic.ref.InternalTokens()
}

func (ic IpConfigurationAttributes) GatewayLoadBalancerFrontendIpConfigurationId() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("gateway_load_balancer_frontend_ip_configuration_id"))
}

func (ic IpConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("name"))
}

func (ic IpConfigurationAttributes) Primary() terra.BoolValue {
	return terra.ReferenceBool(ic.ref.Append("primary"))
}

func (ic IpConfigurationAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("private_ip_address"))
}

func (ic IpConfigurationAttributes) PrivateIpAddressAllocation() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("private_ip_address_allocation"))
}

func (ic IpConfigurationAttributes) PrivateIpAddressVersion() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("private_ip_address_version"))
}

func (ic IpConfigurationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("public_ip_address_id"))
}

func (ic IpConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("subnet_id"))
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

type IpConfigurationState struct {
	GatewayLoadBalancerFrontendIpConfigurationId string `json:"gateway_load_balancer_frontend_ip_configuration_id"`
	Name                                         string `json:"name"`
	Primary                                      bool   `json:"primary"`
	PrivateIpAddress                             string `json:"private_ip_address"`
	PrivateIpAddressAllocation                   string `json:"private_ip_address_allocation"`
	PrivateIpAddressVersion                      string `json:"private_ip_address_version"`
	PublicIpAddressId                            string `json:"public_ip_address_id"`
	SubnetId                                     string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
