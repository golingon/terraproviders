// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lbbackendaddresspooladdress

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InboundNatRulePortMapping struct{}

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

type InboundNatRulePortMappingState struct {
	BackendPort        float64 `json:"backend_port"`
	FrontendPort       float64 `json:"frontend_port"`
	InboundNatRuleName string  `json:"inbound_nat_rule_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}