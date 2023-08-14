// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkservicestlsroute

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rules struct {
	// Action: required
	Action *Action `hcl:"action,block" validate:"required"`
	// Matches: min=1
	Matches []Matches `hcl:"matches,block" validate:"min=1"`
}

type Action struct {
	// Destinations: min=0
	Destinations []Destinations `hcl:"destinations,block" validate:"min=0"`
}

type Destinations struct {
	// ServiceName: string, optional
	ServiceName terra.StringValue `hcl:"service_name,attr"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type Matches struct {
	// Alpn: list of string, optional
	Alpn terra.ListValue[terra.StringValue] `hcl:"alpn,attr"`
	// SniHost: list of string, optional
	SniHost terra.ListValue[terra.StringValue] `hcl:"sni_host,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RulesAttributes struct {
	ref terra.Reference
}

func (r RulesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RulesAttributes) InternalWithRef(ref terra.Reference) RulesAttributes {
	return RulesAttributes{ref: ref}
}

func (r RulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RulesAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](r.ref.Append("action"))
}

func (r RulesAttributes) Matches() terra.ListValue[MatchesAttributes] {
	return terra.ReferenceAsList[MatchesAttributes](r.ref.Append("matches"))
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

func (a ActionAttributes) Destinations() terra.ListValue[DestinationsAttributes] {
	return terra.ReferenceAsList[DestinationsAttributes](a.ref.Append("destinations"))
}

type DestinationsAttributes struct {
	ref terra.Reference
}

func (d DestinationsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DestinationsAttributes) InternalWithRef(ref terra.Reference) DestinationsAttributes {
	return DestinationsAttributes{ref: ref}
}

func (d DestinationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DestinationsAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("service_name"))
}

func (d DestinationsAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("weight"))
}

type MatchesAttributes struct {
	ref terra.Reference
}

func (m MatchesAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MatchesAttributes) InternalWithRef(ref terra.Reference) MatchesAttributes {
	return MatchesAttributes{ref: ref}
}

func (m MatchesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MatchesAttributes) Alpn() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](m.ref.Append("alpn"))
}

func (m MatchesAttributes) SniHost() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](m.ref.Append("sni_host"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type RulesState struct {
	Action  []ActionState  `json:"action"`
	Matches []MatchesState `json:"matches"`
}

type ActionState struct {
	Destinations []DestinationsState `json:"destinations"`
}

type DestinationsState struct {
	ServiceName string  `json:"service_name"`
	Weight      float64 `json:"weight"`
}

type MatchesState struct {
	Alpn    []string `json:"alpn"`
	SniHost []string `json:"sni_host"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}