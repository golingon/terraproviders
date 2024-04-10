// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package firewallnetworkrulecollection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Rule struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationAddresses: list of string, optional
	DestinationAddresses terra.ListValue[terra.StringValue] `hcl:"destination_addresses,attr"`
	// DestinationFqdns: list of string, optional
	DestinationFqdns terra.ListValue[terra.StringValue] `hcl:"destination_fqdns,attr"`
	// DestinationIpGroups: list of string, optional
	DestinationIpGroups terra.ListValue[terra.StringValue] `hcl:"destination_ip_groups,attr"`
	// DestinationPorts: list of string, required
	DestinationPorts terra.ListValue[terra.StringValue] `hcl:"destination_ports,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocols: list of string, required
	Protocols terra.ListValue[terra.StringValue] `hcl:"protocols,attr" validate:"required"`
	// SourceAddresses: list of string, optional
	SourceAddresses terra.ListValue[terra.StringValue] `hcl:"source_addresses,attr"`
	// SourceIpGroups: list of string, optional
	SourceIpGroups terra.ListValue[terra.StringValue] `hcl:"source_ip_groups,attr"`
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

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r RuleAttributes) DestinationAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_addresses"))
}

func (r RuleAttributes) DestinationFqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_fqdns"))
}

func (r RuleAttributes) DestinationIpGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_ip_groups"))
}

func (r RuleAttributes) DestinationPorts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_ports"))
}

func (r RuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r RuleAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("protocols"))
}

func (r RuleAttributes) SourceAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_addresses"))
}

func (r RuleAttributes) SourceIpGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_ip_groups"))
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

type RuleState struct {
	Description          string   `json:"description"`
	DestinationAddresses []string `json:"destination_addresses"`
	DestinationFqdns     []string `json:"destination_fqdns"`
	DestinationIpGroups  []string `json:"destination_ip_groups"`
	DestinationPorts     []string `json:"destination_ports"`
	Name                 string   `json:"name"`
	Protocols            []string `json:"protocols"`
	SourceAddresses      []string `json:"source_addresses"`
	SourceIpGroups       []string `json:"source_ip_groups"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
