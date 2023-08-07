// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sentinelthreatintelligenceindicator

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ParsedPattern struct {
	// PatternTypeValues: min=0
	PatternTypeValues []PatternTypeValues `hcl:"pattern_type_values,block" validate:"min=0"`
}

type PatternTypeValues struct{}

type ExternalReference struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Hashes: map of string, optional
	Hashes terra.MapValue[terra.StringValue] `hcl:"hashes,attr"`
	// SourceName: string, optional
	SourceName terra.StringValue `hcl:"source_name,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
}

type GranularMarking struct {
	// Language: string, optional
	Language terra.StringValue `hcl:"language,attr"`
	// MarkingRef: string, optional
	MarkingRef terra.StringValue `hcl:"marking_ref,attr"`
	// Selectors: list of string, optional
	Selectors terra.ListValue[terra.StringValue] `hcl:"selectors,attr"`
}

type KillChainPhase struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
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

type ParsedPatternAttributes struct {
	ref terra.Reference
}

func (pp ParsedPatternAttributes) InternalRef() (terra.Reference, error) {
	return pp.ref, nil
}

func (pp ParsedPatternAttributes) InternalWithRef(ref terra.Reference) ParsedPatternAttributes {
	return ParsedPatternAttributes{ref: ref}
}

func (pp ParsedPatternAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pp.ref.InternalTokens()
}

func (pp ParsedPatternAttributes) PatternTypeKey() terra.StringValue {
	return terra.ReferenceAsString(pp.ref.Append("pattern_type_key"))
}

func (pp ParsedPatternAttributes) PatternTypeValues() terra.ListValue[PatternTypeValuesAttributes] {
	return terra.ReferenceAsList[PatternTypeValuesAttributes](pp.ref.Append("pattern_type_values"))
}

type PatternTypeValuesAttributes struct {
	ref terra.Reference
}

func (ptv PatternTypeValuesAttributes) InternalRef() (terra.Reference, error) {
	return ptv.ref, nil
}

func (ptv PatternTypeValuesAttributes) InternalWithRef(ref terra.Reference) PatternTypeValuesAttributes {
	return PatternTypeValuesAttributes{ref: ref}
}

func (ptv PatternTypeValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ptv.ref.InternalTokens()
}

func (ptv PatternTypeValuesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ptv.ref.Append("value"))
}

func (ptv PatternTypeValuesAttributes) ValueType() terra.StringValue {
	return terra.ReferenceAsString(ptv.ref.Append("value_type"))
}

type ExternalReferenceAttributes struct {
	ref terra.Reference
}

func (er ExternalReferenceAttributes) InternalRef() (terra.Reference, error) {
	return er.ref, nil
}

func (er ExternalReferenceAttributes) InternalWithRef(ref terra.Reference) ExternalReferenceAttributes {
	return ExternalReferenceAttributes{ref: ref}
}

func (er ExternalReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return er.ref.InternalTokens()
}

func (er ExternalReferenceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("description"))
}

func (er ExternalReferenceAttributes) Hashes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](er.ref.Append("hashes"))
}

func (er ExternalReferenceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("id"))
}

func (er ExternalReferenceAttributes) SourceName() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("source_name"))
}

func (er ExternalReferenceAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("url"))
}

type GranularMarkingAttributes struct {
	ref terra.Reference
}

func (gm GranularMarkingAttributes) InternalRef() (terra.Reference, error) {
	return gm.ref, nil
}

func (gm GranularMarkingAttributes) InternalWithRef(ref terra.Reference) GranularMarkingAttributes {
	return GranularMarkingAttributes{ref: ref}
}

func (gm GranularMarkingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gm.ref.InternalTokens()
}

func (gm GranularMarkingAttributes) Language() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("language"))
}

func (gm GranularMarkingAttributes) MarkingRef() terra.StringValue {
	return terra.ReferenceAsString(gm.ref.Append("marking_ref"))
}

func (gm GranularMarkingAttributes) Selectors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gm.ref.Append("selectors"))
}

type KillChainPhaseAttributes struct {
	ref terra.Reference
}

func (kcp KillChainPhaseAttributes) InternalRef() (terra.Reference, error) {
	return kcp.ref, nil
}

func (kcp KillChainPhaseAttributes) InternalWithRef(ref terra.Reference) KillChainPhaseAttributes {
	return KillChainPhaseAttributes{ref: ref}
}

func (kcp KillChainPhaseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kcp.ref.InternalTokens()
}

func (kcp KillChainPhaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kcp.ref.Append("name"))
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

type ParsedPatternState struct {
	PatternTypeKey    string                   `json:"pattern_type_key"`
	PatternTypeValues []PatternTypeValuesState `json:"pattern_type_values"`
}

type PatternTypeValuesState struct {
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
}

type ExternalReferenceState struct {
	Description string            `json:"description"`
	Hashes      map[string]string `json:"hashes"`
	Id          string            `json:"id"`
	SourceName  string            `json:"source_name"`
	Url         string            `json:"url"`
}

type GranularMarkingState struct {
	Language   string   `json:"language"`
	MarkingRef string   `json:"marking_ref"`
	Selectors  []string `json:"selectors"`
}

type KillChainPhaseState struct {
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
