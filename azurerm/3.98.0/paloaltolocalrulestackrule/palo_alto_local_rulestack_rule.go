// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package paloaltolocalrulestackrule

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Category struct {
	// CustomUrls: list of string, required
	CustomUrls terra.ListValue[terra.StringValue] `hcl:"custom_urls,attr" validate:"required"`
	// Feeds: list of string, optional
	Feeds terra.ListValue[terra.StringValue] `hcl:"feeds,attr"`
}

type Destination struct {
	// Cidrs: list of string, optional
	Cidrs terra.ListValue[terra.StringValue] `hcl:"cidrs,attr"`
	// Countries: list of string, optional
	Countries terra.ListValue[terra.StringValue] `hcl:"countries,attr"`
	// Feeds: list of string, optional
	Feeds terra.ListValue[terra.StringValue] `hcl:"feeds,attr"`
	// LocalRulestackFqdnListIds: list of string, optional
	LocalRulestackFqdnListIds terra.ListValue[terra.StringValue] `hcl:"local_rulestack_fqdn_list_ids,attr"`
	// LocalRulestackPrefixListIds: list of string, optional
	LocalRulestackPrefixListIds terra.ListValue[terra.StringValue] `hcl:"local_rulestack_prefix_list_ids,attr"`
}

type Source struct {
	// Cidrs: list of string, optional
	Cidrs terra.ListValue[terra.StringValue] `hcl:"cidrs,attr"`
	// Countries: list of string, optional
	Countries terra.ListValue[terra.StringValue] `hcl:"countries,attr"`
	// Feeds: list of string, optional
	Feeds terra.ListValue[terra.StringValue] `hcl:"feeds,attr"`
	// LocalRulestackPrefixListIds: list of string, optional
	LocalRulestackPrefixListIds terra.ListValue[terra.StringValue] `hcl:"local_rulestack_prefix_list_ids,attr"`
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

type CategoryAttributes struct {
	ref terra.Reference
}

func (c CategoryAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CategoryAttributes) InternalWithRef(ref terra.Reference) CategoryAttributes {
	return CategoryAttributes{ref: ref}
}

func (c CategoryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CategoryAttributes) CustomUrls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("custom_urls"))
}

func (c CategoryAttributes) Feeds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("feeds"))
}

type DestinationAttributes struct {
	ref terra.Reference
}

func (d DestinationAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationAttributes) InternalWithRef(ref terra.Reference) DestinationAttributes {
	return DestinationAttributes{ref: ref}
}

func (d DestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationAttributes) Cidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("cidrs"))
}

func (d DestinationAttributes) Countries() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("countries"))
}

func (d DestinationAttributes) Feeds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("feeds"))
}

func (d DestinationAttributes) LocalRulestackFqdnListIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("local_rulestack_fqdn_list_ids"))
}

func (d DestinationAttributes) LocalRulestackPrefixListIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("local_rulestack_prefix_list_ids"))
}

type SourceAttributes struct {
	ref terra.Reference
}

func (s SourceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourceAttributes) InternalWithRef(ref terra.Reference) SourceAttributes {
	return SourceAttributes{ref: ref}
}

func (s SourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourceAttributes) Cidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("cidrs"))
}

func (s SourceAttributes) Countries() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("countries"))
}

func (s SourceAttributes) Feeds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("feeds"))
}

func (s SourceAttributes) LocalRulestackPrefixListIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("local_rulestack_prefix_list_ids"))
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

type CategoryState struct {
	CustomUrls []string `json:"custom_urls"`
	Feeds      []string `json:"feeds"`
}

type DestinationState struct {
	Cidrs                       []string `json:"cidrs"`
	Countries                   []string `json:"countries"`
	Feeds                       []string `json:"feeds"`
	LocalRulestackFqdnListIds   []string `json:"local_rulestack_fqdn_list_ids"`
	LocalRulestackPrefixListIds []string `json:"local_rulestack_prefix_list_ids"`
}

type SourceState struct {
	Cidrs                       []string `json:"cidrs"`
	Countries                   []string `json:"countries"`
	Feeds                       []string `json:"feeds"`
	LocalRulestackPrefixListIds []string `json:"local_rulestack_prefix_list_ids"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
