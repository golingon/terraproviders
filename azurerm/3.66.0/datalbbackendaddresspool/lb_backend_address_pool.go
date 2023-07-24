// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalbbackendaddresspool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BackendAddress struct {
	// InboundNatRulePortMapping: min=0
	InboundNatRulePortMapping []InboundNatRulePortMapping `hcl:"inbound_nat_rule_port_mapping,block" validate:"min=0"`
}

type InboundNatRulePortMapping struct{}

type BackendIpConfigurations struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type BackendAddressAttributes struct {
	ref terra.Reference
}

func (ba BackendAddressAttributes) InternalRef() (terra.Reference, error) {
	return ba.ref, nil
}

func (ba BackendAddressAttributes) InternalWithRef(ref terra.Reference) BackendAddressAttributes {
	return BackendAddressAttributes{ref: ref}
}

func (ba BackendAddressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ba.ref.InternalTokens()
}

func (ba BackendAddressAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("ip_address"))
}

func (ba BackendAddressAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("name"))
}

func (ba BackendAddressAttributes) VirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ba.ref.Append("virtual_network_id"))
}

func (ba BackendAddressAttributes) InboundNatRulePortMapping() terra.ListValue[InboundNatRulePortMappingAttributes] {
	return terra.ReferenceAsList[InboundNatRulePortMappingAttributes](ba.ref.Append("inbound_nat_rule_port_mapping"))
}

type InboundNatRulePortMappingAttributes struct {
	ref terra.Reference
}

func (inrpm InboundNatRulePortMappingAttributes) InternalRef() (terra.Reference, error) {
	return inrpm.ref, nil
}

func (inrpm InboundNatRulePortMappingAttributes) InternalWithRef(ref terra.Reference) InboundNatRulePortMappingAttributes {
	return InboundNatRulePortMappingAttributes{ref: ref}
}

func (inrpm InboundNatRulePortMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return inrpm.ref.InternalTokens()
}

func (inrpm InboundNatRulePortMappingAttributes) BackendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(inrpm.ref.Append("backend_port"))
}

func (inrpm InboundNatRulePortMappingAttributes) FrontendPort() terra.NumberValue {
	return terra.ReferenceAsNumber(inrpm.ref.Append("frontend_port"))
}

func (inrpm InboundNatRulePortMappingAttributes) InboundNatRuleName() terra.StringValue {
	return terra.ReferenceAsString(inrpm.ref.Append("inbound_nat_rule_name"))
}

type BackendIpConfigurationsAttributes struct {
	ref terra.Reference
}

func (bic BackendIpConfigurationsAttributes) InternalRef() (terra.Reference, error) {
	return bic.ref, nil
}

func (bic BackendIpConfigurationsAttributes) InternalWithRef(ref terra.Reference) BackendIpConfigurationsAttributes {
	return BackendIpConfigurationsAttributes{ref: ref}
}

func (bic BackendIpConfigurationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bic.ref.InternalTokens()
}

func (bic BackendIpConfigurationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bic.ref.Append("id"))
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

type BackendAddressState struct {
	IpAddress                 string                           `json:"ip_address"`
	Name                      string                           `json:"name"`
	VirtualNetworkId          string                           `json:"virtual_network_id"`
	InboundNatRulePortMapping []InboundNatRulePortMappingState `json:"inbound_nat_rule_port_mapping"`
}

type InboundNatRulePortMappingState struct {
	BackendPort        float64 `json:"backend_port"`
	FrontendPort       float64 `json:"frontend_port"`
	InboundNatRuleName string  `json:"inbound_nat_rule_name"`
}

type BackendIpConfigurationsState struct {
	Id string `json:"id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
