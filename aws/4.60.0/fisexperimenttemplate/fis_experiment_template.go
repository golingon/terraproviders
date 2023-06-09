// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fisexperimenttemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Action struct {
	// ActionId: string, required
	ActionId terra.StringValue `hcl:"action_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StartAfter: set of string, optional
	StartAfter terra.SetValue[terra.StringValue] `hcl:"start_after,attr"`
	// Parameter: min=0
	Parameter []Parameter `hcl:"parameter,block" validate:"min=0"`
	// ActionTarget: optional
	Target *ActionTarget `hcl:"target,block"`
}

type Parameter struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type ActionTarget struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type StopCondition struct {
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type Target struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceArns: set of string, optional
	ResourceArns terra.SetValue[terra.StringValue] `hcl:"resource_arns,attr"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// SelectionMode: string, required
	SelectionMode terra.StringValue `hcl:"selection_mode,attr" validate:"required"`
	// Filter: min=0
	Filter []Filter `hcl:"filter,block" validate:"min=0"`
	// ResourceTag: min=0,max=50
	ResourceTag []ResourceTag `hcl:"resource_tag,block" validate:"min=0,max=50"`
}

type Filter struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type ResourceTag struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

func (a ActionAttributes) ActionId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("action_id"))
}

func (a ActionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("description"))
}

func (a ActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a ActionAttributes) StartAfter() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("start_after"))
}

func (a ActionAttributes) Parameter() terra.SetValue[ParameterAttributes] {
	return terra.ReferenceAsSet[ParameterAttributes](a.ref.Append("parameter"))
}

func (a ActionAttributes) Target() terra.ListValue[ActionTargetAttributes] {
	return terra.ReferenceAsList[ActionTargetAttributes](a.ref.Append("target"))
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("key"))
}

func (p ParameterAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type ActionTargetAttributes struct {
	ref terra.Reference
}

func (t ActionTargetAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t ActionTargetAttributes) InternalWithRef(ref terra.Reference) ActionTargetAttributes {
	return ActionTargetAttributes{ref: ref}
}

func (t ActionTargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t ActionTargetAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t ActionTargetAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("value"))
}

type StopConditionAttributes struct {
	ref terra.Reference
}

func (sc StopConditionAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc StopConditionAttributes) InternalWithRef(ref terra.Reference) StopConditionAttributes {
	return StopConditionAttributes{ref: ref}
}

func (sc StopConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc StopConditionAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("source"))
}

func (sc StopConditionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("value"))
}

type TargetAttributes struct {
	ref terra.Reference
}

func (t TargetAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TargetAttributes) InternalWithRef(ref terra.Reference) TargetAttributes {
	return TargetAttributes{ref: ref}
}

func (t TargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TargetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("name"))
}

func (t TargetAttributes) ResourceArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](t.ref.Append("resource_arns"))
}

func (t TargetAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("resource_type"))
}

func (t TargetAttributes) SelectionMode() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("selection_mode"))
}

func (t TargetAttributes) Filter() terra.ListValue[FilterAttributes] {
	return terra.ReferenceAsList[FilterAttributes](t.ref.Append("filter"))
}

func (t TargetAttributes) ResourceTag() terra.SetValue[ResourceTagAttributes] {
	return terra.ReferenceAsSet[ResourceTagAttributes](t.ref.Append("resource_tag"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("path"))
}

func (f FilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("values"))
}

type ResourceTagAttributes struct {
	ref terra.Reference
}

func (rt ResourceTagAttributes) InternalRef() (terra.Reference, error) {
	return rt.ref, nil
}

func (rt ResourceTagAttributes) InternalWithRef(ref terra.Reference) ResourceTagAttributes {
	return ResourceTagAttributes{ref: ref}
}

func (rt ResourceTagAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rt.ref.InternalTokens()
}

func (rt ResourceTagAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("key"))
}

func (rt ResourceTagAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("value"))
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

type ActionState struct {
	ActionId    string              `json:"action_id"`
	Description string              `json:"description"`
	Name        string              `json:"name"`
	StartAfter  []string            `json:"start_after"`
	Parameter   []ParameterState    `json:"parameter"`
	Target      []ActionTargetState `json:"target"`
}

type ParameterState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ActionTargetState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type StopConditionState struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

type TargetState struct {
	Name          string             `json:"name"`
	ResourceArns  []string           `json:"resource_arns"`
	ResourceType  string             `json:"resource_type"`
	SelectionMode string             `json:"selection_mode"`
	Filter        []FilterState      `json:"filter"`
	ResourceTag   []ResourceTagState `json:"resource_tag"`
}

type FilterState struct {
	Path   string   `json:"path"`
	Values []string `json:"values"`
}

type ResourceTagState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
