// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package firewallnatrulecollection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationAddresses: list of string, required
	DestinationAddresses terra.ListValue[terra.StringValue] `hcl:"destination_addresses,attr" validate:"required"`
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
	// TranslatedAddress: string, required
	TranslatedAddress terra.StringValue `hcl:"translated_address,attr" validate:"required"`
	// TranslatedPort: string, required
	TranslatedPort terra.StringValue `hcl:"translated_port,attr" validate:"required"`
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

func (r RuleAttributes) TranslatedAddress() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("translated_address"))
}

func (r RuleAttributes) TranslatedPort() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("translated_port"))
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
	DestinationPorts     []string `json:"destination_ports"`
	Name                 string   `json:"name"`
	Protocols            []string `json:"protocols"`
	SourceAddresses      []string `json:"source_addresses"`
	SourceIpGroups       []string `json:"source_ip_groups"`
	TranslatedAddress    string   `json:"translated_address"`
	TranslatedPort       string   `json:"translated_port"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
