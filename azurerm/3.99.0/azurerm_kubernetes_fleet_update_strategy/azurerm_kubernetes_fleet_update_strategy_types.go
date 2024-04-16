// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_kubernetes_fleet_update_strategy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Stage struct {
	// AfterStageWaitInSeconds: number, optional
	AfterStageWaitInSeconds terra.NumberValue `hcl:"after_stage_wait_in_seconds,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StageGroup: min=1
	Group []StageGroup `hcl:"group,block" validate:"min=1"`
}

type StageGroup struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type StageAttributes struct {
	ref terra.Reference
}

func (s StageAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StageAttributes) InternalWithRef(ref terra.Reference) StageAttributes {
	return StageAttributes{ref: ref}
}

func (s StageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StageAttributes) AfterStageWaitInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("after_stage_wait_in_seconds"))
}

func (s StageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s StageAttributes) Group() terra.ListValue[StageGroupAttributes] {
	return terra.ReferenceAsList[StageGroupAttributes](s.ref.Append("group"))
}

type StageGroupAttributes struct {
	ref terra.Reference
}

func (g StageGroupAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g StageGroupAttributes) InternalWithRef(ref terra.Reference) StageGroupAttributes {
	return StageGroupAttributes{ref: ref}
}

func (g StageGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g StageGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("name"))
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

type StageState struct {
	AfterStageWaitInSeconds float64           `json:"after_stage_wait_in_seconds"`
	Name                    string            `json:"name"`
	Group                   []StageGroupState `json:"group"`
}

type StageGroupState struct {
	Name string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
