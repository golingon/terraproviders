// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package clouddeploydeliverypipeline

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Condition struct {
	// PipelineReadyCondition: min=0
	PipelineReadyCondition []PipelineReadyCondition `hcl:"pipeline_ready_condition,block" validate:"min=0"`
	// TargetsPresentCondition: min=0
	TargetsPresentCondition []TargetsPresentCondition `hcl:"targets_present_condition,block" validate:"min=0"`
	// TargetsTypeCondition: min=0
	TargetsTypeCondition []TargetsTypeCondition `hcl:"targets_type_condition,block" validate:"min=0"`
}

type PipelineReadyCondition struct{}

type TargetsPresentCondition struct{}

type TargetsTypeCondition struct{}

type SerialPipeline struct {
	// Stages: min=0
	Stages []Stages `hcl:"stages,block" validate:"min=0"`
}

type Stages struct {
	// Profiles: list of string, optional
	Profiles terra.ListValue[terra.StringValue] `hcl:"profiles,attr"`
	// TargetId: string, optional
	TargetId terra.StringValue `hcl:"target_id,attr"`
	// Strategy: optional
	Strategy *Strategy `hcl:"strategy,block"`
}

type Strategy struct {
	// Standard: optional
	Standard *Standard `hcl:"standard,block"`
}

type Standard struct {
	// Verify: bool, optional
	Verify terra.BoolValue `hcl:"verify,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ConditionAttributes struct {
	ref terra.Reference
}

func (c ConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConditionAttributes) InternalWithRef(ref terra.Reference) ConditionAttributes {
	return ConditionAttributes{ref: ref}
}

func (c ConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConditionAttributes) PipelineReadyCondition() terra.ListValue[PipelineReadyConditionAttributes] {
	return terra.ReferenceAsList[PipelineReadyConditionAttributes](c.ref.Append("pipeline_ready_condition"))
}

func (c ConditionAttributes) TargetsPresentCondition() terra.ListValue[TargetsPresentConditionAttributes] {
	return terra.ReferenceAsList[TargetsPresentConditionAttributes](c.ref.Append("targets_present_condition"))
}

func (c ConditionAttributes) TargetsTypeCondition() terra.ListValue[TargetsTypeConditionAttributes] {
	return terra.ReferenceAsList[TargetsTypeConditionAttributes](c.ref.Append("targets_type_condition"))
}

type PipelineReadyConditionAttributes struct {
	ref terra.Reference
}

func (prc PipelineReadyConditionAttributes) InternalRef() (terra.Reference, error) {
	return prc.ref, nil
}

func (prc PipelineReadyConditionAttributes) InternalWithRef(ref terra.Reference) PipelineReadyConditionAttributes {
	return PipelineReadyConditionAttributes{ref: ref}
}

func (prc PipelineReadyConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return prc.ref.InternalTokens()
}

func (prc PipelineReadyConditionAttributes) Status() terra.BoolValue {
	return terra.ReferenceAsBool(prc.ref.Append("status"))
}

func (prc PipelineReadyConditionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(prc.ref.Append("update_time"))
}

type TargetsPresentConditionAttributes struct {
	ref terra.Reference
}

func (tpc TargetsPresentConditionAttributes) InternalRef() (terra.Reference, error) {
	return tpc.ref, nil
}

func (tpc TargetsPresentConditionAttributes) InternalWithRef(ref terra.Reference) TargetsPresentConditionAttributes {
	return TargetsPresentConditionAttributes{ref: ref}
}

func (tpc TargetsPresentConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tpc.ref.InternalTokens()
}

func (tpc TargetsPresentConditionAttributes) MissingTargets() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tpc.ref.Append("missing_targets"))
}

func (tpc TargetsPresentConditionAttributes) Status() terra.BoolValue {
	return terra.ReferenceAsBool(tpc.ref.Append("status"))
}

func (tpc TargetsPresentConditionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(tpc.ref.Append("update_time"))
}

type TargetsTypeConditionAttributes struct {
	ref terra.Reference
}

func (ttc TargetsTypeConditionAttributes) InternalRef() (terra.Reference, error) {
	return ttc.ref, nil
}

func (ttc TargetsTypeConditionAttributes) InternalWithRef(ref terra.Reference) TargetsTypeConditionAttributes {
	return TargetsTypeConditionAttributes{ref: ref}
}

func (ttc TargetsTypeConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ttc.ref.InternalTokens()
}

func (ttc TargetsTypeConditionAttributes) ErrorDetails() terra.StringValue {
	return terra.ReferenceAsString(ttc.ref.Append("error_details"))
}

func (ttc TargetsTypeConditionAttributes) Status() terra.BoolValue {
	return terra.ReferenceAsBool(ttc.ref.Append("status"))
}

type SerialPipelineAttributes struct {
	ref terra.Reference
}

func (sp SerialPipelineAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp SerialPipelineAttributes) InternalWithRef(ref terra.Reference) SerialPipelineAttributes {
	return SerialPipelineAttributes{ref: ref}
}

func (sp SerialPipelineAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp SerialPipelineAttributes) Stages() terra.ListValue[StagesAttributes] {
	return terra.ReferenceAsList[StagesAttributes](sp.ref.Append("stages"))
}

type StagesAttributes struct {
	ref terra.Reference
}

func (s StagesAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StagesAttributes) InternalWithRef(ref terra.Reference) StagesAttributes {
	return StagesAttributes{ref: ref}
}

func (s StagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StagesAttributes) Profiles() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("profiles"))
}

func (s StagesAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("target_id"))
}

func (s StagesAttributes) Strategy() terra.ListValue[StrategyAttributes] {
	return terra.ReferenceAsList[StrategyAttributes](s.ref.Append("strategy"))
}

type StrategyAttributes struct {
	ref terra.Reference
}

func (s StrategyAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StrategyAttributes) InternalWithRef(ref terra.Reference) StrategyAttributes {
	return StrategyAttributes{ref: ref}
}

func (s StrategyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StrategyAttributes) Standard() terra.ListValue[StandardAttributes] {
	return terra.ReferenceAsList[StandardAttributes](s.ref.Append("standard"))
}

type StandardAttributes struct {
	ref terra.Reference
}

func (s StandardAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StandardAttributes) InternalWithRef(ref terra.Reference) StandardAttributes {
	return StandardAttributes{ref: ref}
}

func (s StandardAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StandardAttributes) Verify() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("verify"))
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

type ConditionState struct {
	PipelineReadyCondition  []PipelineReadyConditionState  `json:"pipeline_ready_condition"`
	TargetsPresentCondition []TargetsPresentConditionState `json:"targets_present_condition"`
	TargetsTypeCondition    []TargetsTypeConditionState    `json:"targets_type_condition"`
}

type PipelineReadyConditionState struct {
	Status     bool   `json:"status"`
	UpdateTime string `json:"update_time"`
}

type TargetsPresentConditionState struct {
	MissingTargets []string `json:"missing_targets"`
	Status         bool     `json:"status"`
	UpdateTime     string   `json:"update_time"`
}

type TargetsTypeConditionState struct {
	ErrorDetails string `json:"error_details"`
	Status       bool   `json:"status"`
}

type SerialPipelineState struct {
	Stages []StagesState `json:"stages"`
}

type StagesState struct {
	Profiles []string        `json:"profiles"`
	TargetId string          `json:"target_id"`
	Strategy []StrategyState `json:"strategy"`
}

type StrategyState struct {
	Standard []StandardState `json:"standard"`
}

type StandardState struct {
	Verify bool `json:"verify"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
