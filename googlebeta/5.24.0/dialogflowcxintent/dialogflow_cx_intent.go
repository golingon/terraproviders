// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dialogflowcxintent

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Parameters struct {
	// EntityType: string, required
	EntityType terra.StringValue `hcl:"entity_type,attr" validate:"required"`
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// IsList: bool, optional
	IsList terra.BoolValue `hcl:"is_list,attr"`
	// Redact: bool, optional
	Redact terra.BoolValue `hcl:"redact,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TrainingPhrases struct {
	// RepeatCount: number, optional
	RepeatCount terra.NumberValue `hcl:"repeat_count,attr"`
	// Parts: min=1
	Parts []Parts `hcl:"parts,block" validate:"min=1"`
}

type Parts struct {
	// ParameterId: string, optional
	ParameterId terra.StringValue `hcl:"parameter_id,attr"`
	// Text: string, required
	Text terra.StringValue `hcl:"text,attr" validate:"required"`
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) EntityType() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("entity_type"))
}

func (p ParametersAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("id"))
}

func (p ParametersAttributes) IsList() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("is_list"))
}

func (p ParametersAttributes) Redact() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("redact"))
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

type TrainingPhrasesAttributes struct {
	ref terra.Reference
}

func (tp TrainingPhrasesAttributes) InternalRef() (terra.Reference, error) {
	return tp.ref, nil
}

func (tp TrainingPhrasesAttributes) InternalWithRef(ref terra.Reference) TrainingPhrasesAttributes {
	return TrainingPhrasesAttributes{ref: ref}
}

func (tp TrainingPhrasesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tp.ref.InternalTokens()
}

func (tp TrainingPhrasesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tp.ref.Append("id"))
}

func (tp TrainingPhrasesAttributes) RepeatCount() terra.NumberValue {
	return terra.ReferenceAsNumber(tp.ref.Append("repeat_count"))
}

func (tp TrainingPhrasesAttributes) Parts() terra.ListValue[PartsAttributes] {
	return terra.ReferenceAsList[PartsAttributes](tp.ref.Append("parts"))
}

type PartsAttributes struct {
	ref terra.Reference
}

func (p PartsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PartsAttributes) InternalWithRef(ref terra.Reference) PartsAttributes {
	return PartsAttributes{ref: ref}
}

func (p PartsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PartsAttributes) ParameterId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("parameter_id"))
}

func (p PartsAttributes) Text() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("text"))
}

type ParametersState struct {
	EntityType string `json:"entity_type"`
	Id         string `json:"id"`
	IsList     bool   `json:"is_list"`
	Redact     bool   `json:"redact"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type TrainingPhrasesState struct {
	Id          string       `json:"id"`
	RepeatCount float64      `json:"repeat_count"`
	Parts       []PartsState `json:"parts"`
}

type PartsState struct {
	ParameterId string `json:"parameter_id"`
	Text        string `json:"text"`
}
