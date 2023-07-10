// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediaassetfilter

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PresentationTimeRange struct {
	// EndInUnits: number, optional
	EndInUnits terra.NumberValue `hcl:"end_in_units,attr"`
	// ForceEnd: bool, optional
	ForceEnd terra.BoolValue `hcl:"force_end,attr"`
	// LiveBackoffInUnits: number, optional
	LiveBackoffInUnits terra.NumberValue `hcl:"live_backoff_in_units,attr"`
	// PresentationWindowInUnits: number, optional
	PresentationWindowInUnits terra.NumberValue `hcl:"presentation_window_in_units,attr"`
	// StartInUnits: number, optional
	StartInUnits terra.NumberValue `hcl:"start_in_units,attr"`
	// UnitTimescaleInMiliseconds: number, optional
	UnitTimescaleInMiliseconds terra.NumberValue `hcl:"unit_timescale_in_miliseconds,attr"`
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

type TrackSelection struct {
	// Condition: min=1
	Condition []Condition `hcl:"condition,block" validate:"min=1"`
}

type Condition struct {
	// Operation: string, optional
	Operation terra.StringValue `hcl:"operation,attr"`
	// Property: string, optional
	Property terra.StringValue `hcl:"property,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type PresentationTimeRangeAttributes struct {
	ref terra.Reference
}

func (ptr PresentationTimeRangeAttributes) InternalRef() (terra.Reference, error) {
	return ptr.ref, nil
}

func (ptr PresentationTimeRangeAttributes) InternalWithRef(ref terra.Reference) PresentationTimeRangeAttributes {
	return PresentationTimeRangeAttributes{ref: ref}
}

func (ptr PresentationTimeRangeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ptr.ref.InternalTokens()
}

func (ptr PresentationTimeRangeAttributes) EndInUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(ptr.ref.Append("end_in_units"))
}

func (ptr PresentationTimeRangeAttributes) ForceEnd() terra.BoolValue {
	return terra.ReferenceAsBool(ptr.ref.Append("force_end"))
}

func (ptr PresentationTimeRangeAttributes) LiveBackoffInUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(ptr.ref.Append("live_backoff_in_units"))
}

func (ptr PresentationTimeRangeAttributes) PresentationWindowInUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(ptr.ref.Append("presentation_window_in_units"))
}

func (ptr PresentationTimeRangeAttributes) StartInUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(ptr.ref.Append("start_in_units"))
}

func (ptr PresentationTimeRangeAttributes) UnitTimescaleInMiliseconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ptr.ref.Append("unit_timescale_in_miliseconds"))
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

type TrackSelectionAttributes struct {
	ref terra.Reference
}

func (ts TrackSelectionAttributes) InternalRef() (terra.Reference, error) {
	return ts.ref, nil
}

func (ts TrackSelectionAttributes) InternalWithRef(ref terra.Reference) TrackSelectionAttributes {
	return TrackSelectionAttributes{ref: ref}
}

func (ts TrackSelectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ts.ref.InternalTokens()
}

func (ts TrackSelectionAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](ts.ref.Append("condition"))
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

func (c ConditionAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("operation"))
}

func (c ConditionAttributes) Property() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("property"))
}

func (c ConditionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("value"))
}

type PresentationTimeRangeState struct {
	EndInUnits                 float64 `json:"end_in_units"`
	ForceEnd                   bool    `json:"force_end"`
	LiveBackoffInUnits         float64 `json:"live_backoff_in_units"`
	PresentationWindowInUnits  float64 `json:"presentation_window_in_units"`
	StartInUnits               float64 `json:"start_in_units"`
	UnitTimescaleInMiliseconds float64 `json:"unit_timescale_in_miliseconds"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TrackSelectionState struct {
	Condition []ConditionState `json:"condition"`
}

type ConditionState struct {
	Operation string `json:"operation"`
	Property  string `json:"property"`
	Value     string `json:"value"`
}
