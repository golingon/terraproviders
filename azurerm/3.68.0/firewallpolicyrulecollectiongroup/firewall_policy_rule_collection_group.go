// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package firewallpolicyrulecollectiongroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ApplicationRuleCollection struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// ApplicationRuleCollectionRule: min=1
	Rule []ApplicationRuleCollectionRule `hcl:"rule,block" validate:"min=1"`
}

type ApplicationRuleCollectionRule struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationAddresses: list of string, optional
	DestinationAddresses terra.ListValue[terra.StringValue] `hcl:"destination_addresses,attr"`
	// DestinationFqdnTags: list of string, optional
	DestinationFqdnTags terra.ListValue[terra.StringValue] `hcl:"destination_fqdn_tags,attr"`
	// DestinationFqdns: list of string, optional
	DestinationFqdns terra.ListValue[terra.StringValue] `hcl:"destination_fqdns,attr"`
	// DestinationUrls: list of string, optional
	DestinationUrls terra.ListValue[terra.StringValue] `hcl:"destination_urls,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SourceAddresses: list of string, optional
	SourceAddresses terra.ListValue[terra.StringValue] `hcl:"source_addresses,attr"`
	// SourceIpGroups: list of string, optional
	SourceIpGroups terra.ListValue[terra.StringValue] `hcl:"source_ip_groups,attr"`
	// TerminateTls: bool, optional
	TerminateTls terra.BoolValue `hcl:"terminate_tls,attr"`
	// WebCategories: list of string, optional
	WebCategories terra.ListValue[terra.StringValue] `hcl:"web_categories,attr"`
	// Protocols: min=0
	Protocols []Protocols `hcl:"protocols,block" validate:"min=0"`
}

type Protocols struct {
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type NatRuleCollection struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// NatRuleCollectionRule: min=1
	Rule []NatRuleCollectionRule `hcl:"rule,block" validate:"min=1"`
}

type NatRuleCollectionRule struct {
	// DestinationAddress: string, optional
	DestinationAddress terra.StringValue `hcl:"destination_address,attr"`
	// DestinationPorts: list of string, optional
	DestinationPorts terra.ListValue[terra.StringValue] `hcl:"destination_ports,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Protocols: list of string, required
	Protocols terra.ListValue[terra.StringValue] `hcl:"protocols,attr" validate:"required"`
	// SourceAddresses: list of string, optional
	SourceAddresses terra.ListValue[terra.StringValue] `hcl:"source_addresses,attr"`
	// SourceIpGroups: list of string, optional
	SourceIpGroups terra.ListValue[terra.StringValue] `hcl:"source_ip_groups,attr"`
	// TranslatedAddress: string, optional
	TranslatedAddress terra.StringValue `hcl:"translated_address,attr"`
	// TranslatedFqdn: string, optional
	TranslatedFqdn terra.StringValue `hcl:"translated_fqdn,attr"`
	// TranslatedPort: number, required
	TranslatedPort terra.NumberValue `hcl:"translated_port,attr" validate:"required"`
}

type NetworkRuleCollection struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Priority: number, required
	Priority terra.NumberValue `hcl:"priority,attr" validate:"required"`
	// NetworkRuleCollectionRule: min=1
	Rule []NetworkRuleCollectionRule `hcl:"rule,block" validate:"min=1"`
}

type NetworkRuleCollectionRule struct {
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

type ApplicationRuleCollectionAttributes struct {
	ref terra.Reference
}

func (arc ApplicationRuleCollectionAttributes) InternalRef() (terra.Reference, error) {
	return arc.ref, nil
}

func (arc ApplicationRuleCollectionAttributes) InternalWithRef(ref terra.Reference) ApplicationRuleCollectionAttributes {
	return ApplicationRuleCollectionAttributes{ref: ref}
}

func (arc ApplicationRuleCollectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return arc.ref.InternalTokens()
}

func (arc ApplicationRuleCollectionAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(arc.ref.Append("action"))
}

func (arc ApplicationRuleCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arc.ref.Append("name"))
}

func (arc ApplicationRuleCollectionAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(arc.ref.Append("priority"))
}

func (arc ApplicationRuleCollectionAttributes) Rule() terra.ListValue[ApplicationRuleCollectionRuleAttributes] {
	return terra.ReferenceAsList[ApplicationRuleCollectionRuleAttributes](arc.ref.Append("rule"))
}

type ApplicationRuleCollectionRuleAttributes struct {
	ref terra.Reference
}

func (r ApplicationRuleCollectionRuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ApplicationRuleCollectionRuleAttributes) InternalWithRef(ref terra.Reference) ApplicationRuleCollectionRuleAttributes {
	return ApplicationRuleCollectionRuleAttributes{ref: ref}
}

func (r ApplicationRuleCollectionRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ApplicationRuleCollectionRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r ApplicationRuleCollectionRuleAttributes) DestinationAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_addresses"))
}

func (r ApplicationRuleCollectionRuleAttributes) DestinationFqdnTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_fqdn_tags"))
}

func (r ApplicationRuleCollectionRuleAttributes) DestinationFqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_fqdns"))
}

func (r ApplicationRuleCollectionRuleAttributes) DestinationUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_urls"))
}

func (r ApplicationRuleCollectionRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r ApplicationRuleCollectionRuleAttributes) SourceAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_addresses"))
}

func (r ApplicationRuleCollectionRuleAttributes) SourceIpGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_ip_groups"))
}

func (r ApplicationRuleCollectionRuleAttributes) TerminateTls() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("terminate_tls"))
}

func (r ApplicationRuleCollectionRuleAttributes) WebCategories() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("web_categories"))
}

func (r ApplicationRuleCollectionRuleAttributes) Protocols() terra.ListValue[ProtocolsAttributes] {
	return terra.ReferenceAsList[ProtocolsAttributes](r.ref.Append("protocols"))
}

type ProtocolsAttributes struct {
	ref terra.Reference
}

func (p ProtocolsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ProtocolsAttributes) InternalWithRef(ref terra.Reference) ProtocolsAttributes {
	return ProtocolsAttributes{ref: ref}
}

func (p ProtocolsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ProtocolsAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("port"))
}

func (p ProtocolsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("type"))
}

type NatRuleCollectionAttributes struct {
	ref terra.Reference
}

func (nrc NatRuleCollectionAttributes) InternalRef() (terra.Reference, error) {
	return nrc.ref, nil
}

func (nrc NatRuleCollectionAttributes) InternalWithRef(ref terra.Reference) NatRuleCollectionAttributes {
	return NatRuleCollectionAttributes{ref: ref}
}

func (nrc NatRuleCollectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nrc.ref.InternalTokens()
}

func (nrc NatRuleCollectionAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(nrc.ref.Append("action"))
}

func (nrc NatRuleCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nrc.ref.Append("name"))
}

func (nrc NatRuleCollectionAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(nrc.ref.Append("priority"))
}

func (nrc NatRuleCollectionAttributes) Rule() terra.ListValue[NatRuleCollectionRuleAttributes] {
	return terra.ReferenceAsList[NatRuleCollectionRuleAttributes](nrc.ref.Append("rule"))
}

type NatRuleCollectionRuleAttributes struct {
	ref terra.Reference
}

func (r NatRuleCollectionRuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r NatRuleCollectionRuleAttributes) InternalWithRef(ref terra.Reference) NatRuleCollectionRuleAttributes {
	return NatRuleCollectionRuleAttributes{ref: ref}
}

func (r NatRuleCollectionRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r NatRuleCollectionRuleAttributes) DestinationAddress() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("destination_address"))
}

func (r NatRuleCollectionRuleAttributes) DestinationPorts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_ports"))
}

func (r NatRuleCollectionRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r NatRuleCollectionRuleAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("protocols"))
}

func (r NatRuleCollectionRuleAttributes) SourceAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_addresses"))
}

func (r NatRuleCollectionRuleAttributes) SourceIpGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_ip_groups"))
}

func (r NatRuleCollectionRuleAttributes) TranslatedAddress() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("translated_address"))
}

func (r NatRuleCollectionRuleAttributes) TranslatedFqdn() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("translated_fqdn"))
}

func (r NatRuleCollectionRuleAttributes) TranslatedPort() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("translated_port"))
}

type NetworkRuleCollectionAttributes struct {
	ref terra.Reference
}

func (nrc NetworkRuleCollectionAttributes) InternalRef() (terra.Reference, error) {
	return nrc.ref, nil
}

func (nrc NetworkRuleCollectionAttributes) InternalWithRef(ref terra.Reference) NetworkRuleCollectionAttributes {
	return NetworkRuleCollectionAttributes{ref: ref}
}

func (nrc NetworkRuleCollectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nrc.ref.InternalTokens()
}

func (nrc NetworkRuleCollectionAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(nrc.ref.Append("action"))
}

func (nrc NetworkRuleCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nrc.ref.Append("name"))
}

func (nrc NetworkRuleCollectionAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(nrc.ref.Append("priority"))
}

func (nrc NetworkRuleCollectionAttributes) Rule() terra.ListValue[NetworkRuleCollectionRuleAttributes] {
	return terra.ReferenceAsList[NetworkRuleCollectionRuleAttributes](nrc.ref.Append("rule"))
}

type NetworkRuleCollectionRuleAttributes struct {
	ref terra.Reference
}

func (r NetworkRuleCollectionRuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r NetworkRuleCollectionRuleAttributes) InternalWithRef(ref terra.Reference) NetworkRuleCollectionRuleAttributes {
	return NetworkRuleCollectionRuleAttributes{ref: ref}
}

func (r NetworkRuleCollectionRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r NetworkRuleCollectionRuleAttributes) DestinationAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_addresses"))
}

func (r NetworkRuleCollectionRuleAttributes) DestinationFqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_fqdns"))
}

func (r NetworkRuleCollectionRuleAttributes) DestinationIpGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_ip_groups"))
}

func (r NetworkRuleCollectionRuleAttributes) DestinationPorts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("destination_ports"))
}

func (r NetworkRuleCollectionRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r NetworkRuleCollectionRuleAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("protocols"))
}

func (r NetworkRuleCollectionRuleAttributes) SourceAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("source_addresses"))
}

func (r NetworkRuleCollectionRuleAttributes) SourceIpGroups() terra.ListValue[terra.StringValue] {
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

type ApplicationRuleCollectionState struct {
	Action   string                               `json:"action"`
	Name     string                               `json:"name"`
	Priority float64                              `json:"priority"`
	Rule     []ApplicationRuleCollectionRuleState `json:"rule"`
}

type ApplicationRuleCollectionRuleState struct {
	Description          string           `json:"description"`
	DestinationAddresses []string         `json:"destination_addresses"`
	DestinationFqdnTags  []string         `json:"destination_fqdn_tags"`
	DestinationFqdns     []string         `json:"destination_fqdns"`
	DestinationUrls      []string         `json:"destination_urls"`
	Name                 string           `json:"name"`
	SourceAddresses      []string         `json:"source_addresses"`
	SourceIpGroups       []string         `json:"source_ip_groups"`
	TerminateTls         bool             `json:"terminate_tls"`
	WebCategories        []string         `json:"web_categories"`
	Protocols            []ProtocolsState `json:"protocols"`
}

type ProtocolsState struct {
	Port float64 `json:"port"`
	Type string  `json:"type"`
}

type NatRuleCollectionState struct {
	Action   string                       `json:"action"`
	Name     string                       `json:"name"`
	Priority float64                      `json:"priority"`
	Rule     []NatRuleCollectionRuleState `json:"rule"`
}

type NatRuleCollectionRuleState struct {
	DestinationAddress string   `json:"destination_address"`
	DestinationPorts   []string `json:"destination_ports"`
	Name               string   `json:"name"`
	Protocols          []string `json:"protocols"`
	SourceAddresses    []string `json:"source_addresses"`
	SourceIpGroups     []string `json:"source_ip_groups"`
	TranslatedAddress  string   `json:"translated_address"`
	TranslatedFqdn     string   `json:"translated_fqdn"`
	TranslatedPort     float64  `json:"translated_port"`
}

type NetworkRuleCollectionState struct {
	Action   string                           `json:"action"`
	Name     string                           `json:"name"`
	Priority float64                          `json:"priority"`
	Rule     []NetworkRuleCollectionRuleState `json:"rule"`
}

type NetworkRuleCollectionRuleState struct {
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