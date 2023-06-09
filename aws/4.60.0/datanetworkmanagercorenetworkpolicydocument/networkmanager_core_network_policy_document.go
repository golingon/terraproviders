// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datanetworkmanagercorenetworkpolicydocument

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AttachmentPolicies struct {
	// ConditionLogic: string, optional
	ConditionLogic terra.StringValue `hcl:"condition_logic,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// RuleNumber: number, required
	RuleNumber terra.NumberValue `hcl:"rule_number,attr" validate:"required"`
	// Action: required
	Action *Action `hcl:"action,block" validate:"required"`
	// Conditions: min=1
	Conditions []Conditions `hcl:"conditions,block" validate:"min=1"`
}

type Action struct {
	// AssociationMethod: string, required
	AssociationMethod terra.StringValue `hcl:"association_method,attr" validate:"required"`
	// RequireAcceptance: bool, optional
	RequireAcceptance terra.BoolValue `hcl:"require_acceptance,attr"`
	// Segment: string, optional
	Segment terra.StringValue `hcl:"segment,attr"`
	// TagValueOfKey: string, optional
	TagValueOfKey terra.StringValue `hcl:"tag_value_of_key,attr"`
}

type Conditions struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Operator: string, optional
	Operator terra.StringValue `hcl:"operator,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type CoreNetworkConfiguration struct {
	// AsnRanges: set of string, required
	AsnRanges terra.SetValue[terra.StringValue] `hcl:"asn_ranges,attr" validate:"required"`
	// InsideCidrBlocks: set of string, optional
	InsideCidrBlocks terra.SetValue[terra.StringValue] `hcl:"inside_cidr_blocks,attr"`
	// VpnEcmpSupport: bool, optional
	VpnEcmpSupport terra.BoolValue `hcl:"vpn_ecmp_support,attr"`
	// EdgeLocations: min=1,max=17
	EdgeLocations []EdgeLocations `hcl:"edge_locations,block" validate:"min=1,max=17"`
}

type EdgeLocations struct {
	// Asn: string, optional
	Asn terra.StringValue `hcl:"asn,attr"`
	// InsideCidrBlocks: list of string, optional
	InsideCidrBlocks terra.ListValue[terra.StringValue] `hcl:"inside_cidr_blocks,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
}

type SegmentActions struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestinationCidrBlocks: set of string, optional
	DestinationCidrBlocks terra.SetValue[terra.StringValue] `hcl:"destination_cidr_blocks,attr"`
	// Destinations: set of string, optional
	Destinations terra.SetValue[terra.StringValue] `hcl:"destinations,attr"`
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
	// Segment: string, required
	Segment terra.StringValue `hcl:"segment,attr" validate:"required"`
	// ShareWith: set of string, optional
	ShareWith terra.SetValue[terra.StringValue] `hcl:"share_with,attr"`
	// ShareWithExcept: set of string, optional
	ShareWithExcept terra.SetValue[terra.StringValue] `hcl:"share_with_except,attr"`
}

type Segments struct {
	// AllowFilter: set of string, optional
	AllowFilter terra.SetValue[terra.StringValue] `hcl:"allow_filter,attr"`
	// DenyFilter: set of string, optional
	DenyFilter terra.SetValue[terra.StringValue] `hcl:"deny_filter,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EdgeLocations: set of string, optional
	EdgeLocations terra.SetValue[terra.StringValue] `hcl:"edge_locations,attr"`
	// IsolateAttachments: bool, optional
	IsolateAttachments terra.BoolValue `hcl:"isolate_attachments,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequireAttachmentAcceptance: bool, optional
	RequireAttachmentAcceptance terra.BoolValue `hcl:"require_attachment_acceptance,attr"`
}

type AttachmentPoliciesAttributes struct {
	ref terra.Reference
}

func (ap AttachmentPoliciesAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AttachmentPoliciesAttributes) InternalWithRef(ref terra.Reference) AttachmentPoliciesAttributes {
	return AttachmentPoliciesAttributes{ref: ref}
}

func (ap AttachmentPoliciesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AttachmentPoliciesAttributes) ConditionLogic() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("condition_logic"))
}

func (ap AttachmentPoliciesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("description"))
}

func (ap AttachmentPoliciesAttributes) RuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("rule_number"))
}

func (ap AttachmentPoliciesAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](ap.ref.Append("action"))
}

func (ap AttachmentPoliciesAttributes) Conditions() terra.ListValue[ConditionsAttributes] {
	return terra.ReferenceAsList[ConditionsAttributes](ap.ref.Append("conditions"))
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) AssociationMethod() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("association_method"))
}

func (a ActionAttributes) RequireAcceptance() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("require_acceptance"))
}

func (a ActionAttributes) Segment() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("segment"))
}

func (a ActionAttributes) TagValueOfKey() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("tag_value_of_key"))
}

type ConditionsAttributes struct {
	ref terra.Reference
}

func (c ConditionsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionsAttributes) InternalWithRef(ref terra.Reference) ConditionsAttributes {
	return ConditionsAttributes{ref: ref}
}

func (c ConditionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("key"))
}

func (c ConditionsAttributes) Operator() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("operator"))
}

func (c ConditionsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

func (c ConditionsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("value"))
}

type CoreNetworkConfigurationAttributes struct {
	ref terra.Reference
}

func (cnc CoreNetworkConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cnc.ref, nil
}

func (cnc CoreNetworkConfigurationAttributes) InternalWithRef(ref terra.Reference) CoreNetworkConfigurationAttributes {
	return CoreNetworkConfigurationAttributes{ref: ref}
}

func (cnc CoreNetworkConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cnc.ref.InternalTokens()
}

func (cnc CoreNetworkConfigurationAttributes) AsnRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cnc.ref.Append("asn_ranges"))
}

func (cnc CoreNetworkConfigurationAttributes) InsideCidrBlocks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cnc.ref.Append("inside_cidr_blocks"))
}

func (cnc CoreNetworkConfigurationAttributes) VpnEcmpSupport() terra.BoolValue {
	return terra.ReferenceAsBool(cnc.ref.Append("vpn_ecmp_support"))
}

func (cnc CoreNetworkConfigurationAttributes) EdgeLocations() terra.ListValue[EdgeLocationsAttributes] {
	return terra.ReferenceAsList[EdgeLocationsAttributes](cnc.ref.Append("edge_locations"))
}

type EdgeLocationsAttributes struct {
	ref terra.Reference
}

func (el EdgeLocationsAttributes) InternalRef() (terra.Reference, error) {
	return el.ref, nil
}

func (el EdgeLocationsAttributes) InternalWithRef(ref terra.Reference) EdgeLocationsAttributes {
	return EdgeLocationsAttributes{ref: ref}
}

func (el EdgeLocationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return el.ref.InternalTokens()
}

func (el EdgeLocationsAttributes) Asn() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("asn"))
}

func (el EdgeLocationsAttributes) InsideCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](el.ref.Append("inside_cidr_blocks"))
}

func (el EdgeLocationsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("location"))
}

type SegmentActionsAttributes struct {
	ref terra.Reference
}

func (sa SegmentActionsAttributes) InternalRef() (terra.Reference, error) {
	return sa.ref, nil
}

func (sa SegmentActionsAttributes) InternalWithRef(ref terra.Reference) SegmentActionsAttributes {
	return SegmentActionsAttributes{ref: ref}
}

func (sa SegmentActionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sa.ref.InternalTokens()
}

func (sa SegmentActionsAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("action"))
}

func (sa SegmentActionsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("description"))
}

func (sa SegmentActionsAttributes) DestinationCidrBlocks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sa.ref.Append("destination_cidr_blocks"))
}

func (sa SegmentActionsAttributes) Destinations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sa.ref.Append("destinations"))
}

func (sa SegmentActionsAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("mode"))
}

func (sa SegmentActionsAttributes) Segment() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("segment"))
}

func (sa SegmentActionsAttributes) ShareWith() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sa.ref.Append("share_with"))
}

func (sa SegmentActionsAttributes) ShareWithExcept() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sa.ref.Append("share_with_except"))
}

type SegmentsAttributes struct {
	ref terra.Reference
}

func (s SegmentsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SegmentsAttributes) InternalWithRef(ref terra.Reference) SegmentsAttributes {
	return SegmentsAttributes{ref: ref}
}

func (s SegmentsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SegmentsAttributes) AllowFilter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("allow_filter"))
}

func (s SegmentsAttributes) DenyFilter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("deny_filter"))
}

func (s SegmentsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
}

func (s SegmentsAttributes) EdgeLocations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("edge_locations"))
}

func (s SegmentsAttributes) IsolateAttachments() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("isolate_attachments"))
}

func (s SegmentsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SegmentsAttributes) RequireAttachmentAcceptance() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("require_attachment_acceptance"))
}

type AttachmentPoliciesState struct {
	ConditionLogic string            `json:"condition_logic"`
	Description    string            `json:"description"`
	RuleNumber     float64           `json:"rule_number"`
	Action         []ActionState     `json:"action"`
	Conditions     []ConditionsState `json:"conditions"`
}

type ActionState struct {
	AssociationMethod string `json:"association_method"`
	RequireAcceptance bool   `json:"require_acceptance"`
	Segment           string `json:"segment"`
	TagValueOfKey     string `json:"tag_value_of_key"`
}

type ConditionsState struct {
	Key      string `json:"key"`
	Operator string `json:"operator"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

type CoreNetworkConfigurationState struct {
	AsnRanges        []string             `json:"asn_ranges"`
	InsideCidrBlocks []string             `json:"inside_cidr_blocks"`
	VpnEcmpSupport   bool                 `json:"vpn_ecmp_support"`
	EdgeLocations    []EdgeLocationsState `json:"edge_locations"`
}

type EdgeLocationsState struct {
	Asn              string   `json:"asn"`
	InsideCidrBlocks []string `json:"inside_cidr_blocks"`
	Location         string   `json:"location"`
}

type SegmentActionsState struct {
	Action                string   `json:"action"`
	Description           string   `json:"description"`
	DestinationCidrBlocks []string `json:"destination_cidr_blocks"`
	Destinations          []string `json:"destinations"`
	Mode                  string   `json:"mode"`
	Segment               string   `json:"segment"`
	ShareWith             []string `json:"share_with"`
	ShareWithExcept       []string `json:"share_with_except"`
}

type SegmentsState struct {
	AllowFilter                 []string `json:"allow_filter"`
	DenyFilter                  []string `json:"deny_filter"`
	Description                 string   `json:"description"`
	EdgeLocations               []string `json:"edge_locations"`
	IsolateAttachments          bool     `json:"isolate_attachments"`
	Name                        string   `json:"name"`
	RequireAttachmentAcceptance bool     `json:"require_attachment_acceptance"`
}
