// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lb

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type FrontendIpConfiguration struct {
	// GatewayLoadBalancerFrontendIpConfigurationId: string, optional
	GatewayLoadBalancerFrontendIpConfigurationId terra.StringValue `hcl:"gateway_load_balancer_frontend_ip_configuration_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateIpAddress: string, optional
	PrivateIpAddress terra.StringValue `hcl:"private_ip_address,attr"`
	// PrivateIpAddressAllocation: string, optional
	PrivateIpAddressAllocation terra.StringValue `hcl:"private_ip_address_allocation,attr"`
	// PrivateIpAddressVersion: string, optional
	PrivateIpAddressVersion terra.StringValue `hcl:"private_ip_address_version,attr"`
	// PublicIpAddressId: string, optional
	PublicIpAddressId terra.StringValue `hcl:"public_ip_address_id,attr"`
	// PublicIpPrefixId: string, optional
	PublicIpPrefixId terra.StringValue `hcl:"public_ip_prefix_id,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
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

type FrontendIpConfigurationAttributes struct {
	ref terra.Reference
}

func (fic FrontendIpConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return fic.ref, nil
}

func (fic FrontendIpConfigurationAttributes) InternalWithRef(ref terra.Reference) FrontendIpConfigurationAttributes {
	return FrontendIpConfigurationAttributes{ref: ref}
}

func (fic FrontendIpConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fic.ref.InternalTokens()
}

func (fic FrontendIpConfigurationAttributes) GatewayLoadBalancerFrontendIpConfigurationId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("gateway_load_balancer_frontend_ip_configuration_id"))
}

func (fic FrontendIpConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("id"))
}

func (fic FrontendIpConfigurationAttributes) InboundNatRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fic.ref.Append("inbound_nat_rules"))
}

func (fic FrontendIpConfigurationAttributes) LoadBalancerRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fic.ref.Append("load_balancer_rules"))
}

func (fic FrontendIpConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("name"))
}

func (fic FrontendIpConfigurationAttributes) OutboundRules() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fic.ref.Append("outbound_rules"))
}

func (fic FrontendIpConfigurationAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("private_ip_address"))
}

func (fic FrontendIpConfigurationAttributes) PrivateIpAddressAllocation() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("private_ip_address_allocation"))
}

func (fic FrontendIpConfigurationAttributes) PrivateIpAddressVersion() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("private_ip_address_version"))
}

func (fic FrontendIpConfigurationAttributes) PublicIpAddressId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("public_ip_address_id"))
}

func (fic FrontendIpConfigurationAttributes) PublicIpPrefixId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("public_ip_prefix_id"))
}

func (fic FrontendIpConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(fic.ref.Append("subnet_id"))
}

func (fic FrontendIpConfigurationAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fic.ref.Append("zones"))
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

type FrontendIpConfigurationState struct {
	GatewayLoadBalancerFrontendIpConfigurationId string   `json:"gateway_load_balancer_frontend_ip_configuration_id"`
	Id                                           string   `json:"id"`
	InboundNatRules                              []string `json:"inbound_nat_rules"`
	LoadBalancerRules                            []string `json:"load_balancer_rules"`
	Name                                         string   `json:"name"`
	OutboundRules                                []string `json:"outbound_rules"`
	PrivateIpAddress                             string   `json:"private_ip_address"`
	PrivateIpAddressAllocation                   string   `json:"private_ip_address_allocation"`
	PrivateIpAddressVersion                      string   `json:"private_ip_address_version"`
	PublicIpAddressId                            string   `json:"public_ip_address_id"`
	PublicIpPrefixId                             string   `json:"public_ip_prefix_id"`
	SubnetId                                     string   `json:"subnet_id"`
	Zones                                        []string `json:"zones"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
