// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package siterecoveryreplicationrecoveryplan

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AzureToAzureSettings struct {
	// PrimaryEdgeZone: string, optional
	PrimaryEdgeZone terra.StringValue `hcl:"primary_edge_zone,attr"`
	// PrimaryZone: string, optional
	PrimaryZone terra.StringValue `hcl:"primary_zone,attr"`
	// RecoveryEdgeZone: string, optional
	RecoveryEdgeZone terra.StringValue `hcl:"recovery_edge_zone,attr"`
	// RecoveryZone: string, optional
	RecoveryZone terra.StringValue `hcl:"recovery_zone,attr"`
}

type RecoveryGroup struct {
	// ReplicatedProtectedItems: list of string, optional
	ReplicatedProtectedItems terra.ListValue[terra.StringValue] `hcl:"replicated_protected_items,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// PostAction: min=0
	PostAction []PostAction `hcl:"post_action,block" validate:"min=0"`
	// PreAction: min=0
	PreAction []PreAction `hcl:"pre_action,block" validate:"min=0"`
}

type PostAction struct {
	// FabricLocation: string, optional
	FabricLocation terra.StringValue `hcl:"fabric_location,attr"`
	// FailOverDirections: set of string, required
	FailOverDirections terra.SetValue[terra.StringValue] `hcl:"fail_over_directions,attr" validate:"required"`
	// FailOverTypes: set of string, required
	FailOverTypes terra.SetValue[terra.StringValue] `hcl:"fail_over_types,attr" validate:"required"`
	// ManualActionInstruction: string, optional
	ManualActionInstruction terra.StringValue `hcl:"manual_action_instruction,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RunbookId: string, optional
	RunbookId terra.StringValue `hcl:"runbook_id,attr"`
	// ScriptPath: string, optional
	ScriptPath terra.StringValue `hcl:"script_path,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type PreAction struct {
	// FabricLocation: string, optional
	FabricLocation terra.StringValue `hcl:"fabric_location,attr"`
	// FailOverDirections: set of string, required
	FailOverDirections terra.SetValue[terra.StringValue] `hcl:"fail_over_directions,attr" validate:"required"`
	// FailOverTypes: set of string, required
	FailOverTypes terra.SetValue[terra.StringValue] `hcl:"fail_over_types,attr" validate:"required"`
	// ManualActionInstruction: string, optional
	ManualActionInstruction terra.StringValue `hcl:"manual_action_instruction,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RunbookId: string, optional
	RunbookId terra.StringValue `hcl:"runbook_id,attr"`
	// ScriptPath: string, optional
	ScriptPath terra.StringValue `hcl:"script_path,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
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

type AzureToAzureSettingsAttributes struct {
	ref terra.Reference
}

func (atas AzureToAzureSettingsAttributes) InternalRef() (terra.Reference, error) {
	return atas.ref, nil
}

func (atas AzureToAzureSettingsAttributes) InternalWithRef(ref terra.Reference) AzureToAzureSettingsAttributes {
	return AzureToAzureSettingsAttributes{ref: ref}
}

func (atas AzureToAzureSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return atas.ref.InternalTokens()
}

func (atas AzureToAzureSettingsAttributes) PrimaryEdgeZone() terra.StringValue {
	return terra.ReferenceAsString(atas.ref.Append("primary_edge_zone"))
}

func (atas AzureToAzureSettingsAttributes) PrimaryZone() terra.StringValue {
	return terra.ReferenceAsString(atas.ref.Append("primary_zone"))
}

func (atas AzureToAzureSettingsAttributes) RecoveryEdgeZone() terra.StringValue {
	return terra.ReferenceAsString(atas.ref.Append("recovery_edge_zone"))
}

func (atas AzureToAzureSettingsAttributes) RecoveryZone() terra.StringValue {
	return terra.ReferenceAsString(atas.ref.Append("recovery_zone"))
}

type RecoveryGroupAttributes struct {
	ref terra.Reference
}

func (rg RecoveryGroupAttributes) InternalRef() (terra.Reference, error) {
	return rg.ref, nil
}

func (rg RecoveryGroupAttributes) InternalWithRef(ref terra.Reference) RecoveryGroupAttributes {
	return RecoveryGroupAttributes{ref: ref}
}

func (rg RecoveryGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rg.ref.InternalTokens()
}

func (rg RecoveryGroupAttributes) ReplicatedProtectedItems() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rg.ref.Append("replicated_protected_items"))
}

func (rg RecoveryGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(rg.ref.Append("type"))
}

func (rg RecoveryGroupAttributes) PostAction() terra.SetValue[PostActionAttributes] {
	return terra.ReferenceAsSet[PostActionAttributes](rg.ref.Append("post_action"))
}

func (rg RecoveryGroupAttributes) PreAction() terra.SetValue[PreActionAttributes] {
	return terra.ReferenceAsSet[PreActionAttributes](rg.ref.Append("pre_action"))
}

type PostActionAttributes struct {
	ref terra.Reference
}

func (pa PostActionAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PostActionAttributes) InternalWithRef(ref terra.Reference) PostActionAttributes {
	return PostActionAttributes{ref: ref}
}

func (pa PostActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pa.ref.InternalTokens()
}

func (pa PostActionAttributes) FabricLocation() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("fabric_location"))
}

func (pa PostActionAttributes) FailOverDirections() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pa.ref.Append("fail_over_directions"))
}

func (pa PostActionAttributes) FailOverTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pa.ref.Append("fail_over_types"))
}

func (pa PostActionAttributes) ManualActionInstruction() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("manual_action_instruction"))
}

func (pa PostActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("name"))
}

func (pa PostActionAttributes) RunbookId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("runbook_id"))
}

func (pa PostActionAttributes) ScriptPath() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("script_path"))
}

func (pa PostActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("type"))
}

type PreActionAttributes struct {
	ref terra.Reference
}

func (pa PreActionAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PreActionAttributes) InternalWithRef(ref terra.Reference) PreActionAttributes {
	return PreActionAttributes{ref: ref}
}

func (pa PreActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pa.ref.InternalTokens()
}

func (pa PreActionAttributes) FabricLocation() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("fabric_location"))
}

func (pa PreActionAttributes) FailOverDirections() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pa.ref.Append("fail_over_directions"))
}

func (pa PreActionAttributes) FailOverTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pa.ref.Append("fail_over_types"))
}

func (pa PreActionAttributes) ManualActionInstruction() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("manual_action_instruction"))
}

func (pa PreActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("name"))
}

func (pa PreActionAttributes) RunbookId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("runbook_id"))
}

func (pa PreActionAttributes) ScriptPath() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("script_path"))
}

func (pa PreActionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("type"))
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

type AzureToAzureSettingsState struct {
	PrimaryEdgeZone  string `json:"primary_edge_zone"`
	PrimaryZone      string `json:"primary_zone"`
	RecoveryEdgeZone string `json:"recovery_edge_zone"`
	RecoveryZone     string `json:"recovery_zone"`
}

type RecoveryGroupState struct {
	ReplicatedProtectedItems []string          `json:"replicated_protected_items"`
	Type                     string            `json:"type"`
	PostAction               []PostActionState `json:"post_action"`
	PreAction                []PreActionState  `json:"pre_action"`
}

type PostActionState struct {
	FabricLocation          string   `json:"fabric_location"`
	FailOverDirections      []string `json:"fail_over_directions"`
	FailOverTypes           []string `json:"fail_over_types"`
	ManualActionInstruction string   `json:"manual_action_instruction"`
	Name                    string   `json:"name"`
	RunbookId               string   `json:"runbook_id"`
	ScriptPath              string   `json:"script_path"`
	Type                    string   `json:"type"`
}

type PreActionState struct {
	FabricLocation          string   `json:"fabric_location"`
	FailOverDirections      []string `json:"fail_over_directions"`
	FailOverTypes           []string `json:"fail_over_types"`
	ManualActionInstruction string   `json:"manual_action_instruction"`
	Name                    string   `json:"name"`
	RunbookId               string   `json:"runbook_id"`
	ScriptPath              string   `json:"script_path"`
	Type                    string   `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
