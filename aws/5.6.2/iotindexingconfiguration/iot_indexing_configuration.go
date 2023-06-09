// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iotindexingconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ThingGroupIndexingConfiguration struct {
	// ThingGroupIndexingMode: string, required
	ThingGroupIndexingMode terra.StringValue `hcl:"thing_group_indexing_mode,attr" validate:"required"`
	// ThingGroupIndexingConfigurationCustomField: min=0
	CustomField []ThingGroupIndexingConfigurationCustomField `hcl:"custom_field,block" validate:"min=0"`
	// ThingGroupIndexingConfigurationManagedField: min=0
	ManagedField []ThingGroupIndexingConfigurationManagedField `hcl:"managed_field,block" validate:"min=0"`
}

type ThingGroupIndexingConfigurationCustomField struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type ThingGroupIndexingConfigurationManagedField struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type ThingIndexingConfiguration struct {
	// DeviceDefenderIndexingMode: string, optional
	DeviceDefenderIndexingMode terra.StringValue `hcl:"device_defender_indexing_mode,attr"`
	// NamedShadowIndexingMode: string, optional
	NamedShadowIndexingMode terra.StringValue `hcl:"named_shadow_indexing_mode,attr"`
	// ThingConnectivityIndexingMode: string, optional
	ThingConnectivityIndexingMode terra.StringValue `hcl:"thing_connectivity_indexing_mode,attr"`
	// ThingIndexingMode: string, required
	ThingIndexingMode terra.StringValue `hcl:"thing_indexing_mode,attr" validate:"required"`
	// ThingIndexingConfigurationCustomField: min=0
	CustomField []ThingIndexingConfigurationCustomField `hcl:"custom_field,block" validate:"min=0"`
	// ThingIndexingConfigurationManagedField: min=0
	ManagedField []ThingIndexingConfigurationManagedField `hcl:"managed_field,block" validate:"min=0"`
}

type ThingIndexingConfigurationCustomField struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type ThingIndexingConfigurationManagedField struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type ThingGroupIndexingConfigurationAttributes struct {
	ref terra.Reference
}

func (tgic ThingGroupIndexingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return tgic.ref, nil
}

func (tgic ThingGroupIndexingConfigurationAttributes) InternalWithRef(ref terra.Reference) ThingGroupIndexingConfigurationAttributes {
	return ThingGroupIndexingConfigurationAttributes{ref: ref}
}

func (tgic ThingGroupIndexingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tgic.ref.InternalTokens()
}

func (tgic ThingGroupIndexingConfigurationAttributes) ThingGroupIndexingMode() terra.StringValue {
	return terra.ReferenceAsString(tgic.ref.Append("thing_group_indexing_mode"))
}

func (tgic ThingGroupIndexingConfigurationAttributes) CustomField() terra.SetValue[ThingGroupIndexingConfigurationCustomFieldAttributes] {
	return terra.ReferenceAsSet[ThingGroupIndexingConfigurationCustomFieldAttributes](tgic.ref.Append("custom_field"))
}

func (tgic ThingGroupIndexingConfigurationAttributes) ManagedField() terra.SetValue[ThingGroupIndexingConfigurationManagedFieldAttributes] {
	return terra.ReferenceAsSet[ThingGroupIndexingConfigurationManagedFieldAttributes](tgic.ref.Append("managed_field"))
}

type ThingGroupIndexingConfigurationCustomFieldAttributes struct {
	ref terra.Reference
}

func (cf ThingGroupIndexingConfigurationCustomFieldAttributes) InternalRef() (terra.Reference, error) {
	return cf.ref, nil
}

func (cf ThingGroupIndexingConfigurationCustomFieldAttributes) InternalWithRef(ref terra.Reference) ThingGroupIndexingConfigurationCustomFieldAttributes {
	return ThingGroupIndexingConfigurationCustomFieldAttributes{ref: ref}
}

func (cf ThingGroupIndexingConfigurationCustomFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cf.ref.InternalTokens()
}

func (cf ThingGroupIndexingConfigurationCustomFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

func (cf ThingGroupIndexingConfigurationCustomFieldAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("type"))
}

type ThingGroupIndexingConfigurationManagedFieldAttributes struct {
	ref terra.Reference
}

func (mf ThingGroupIndexingConfigurationManagedFieldAttributes) InternalRef() (terra.Reference, error) {
	return mf.ref, nil
}

func (mf ThingGroupIndexingConfigurationManagedFieldAttributes) InternalWithRef(ref terra.Reference) ThingGroupIndexingConfigurationManagedFieldAttributes {
	return ThingGroupIndexingConfigurationManagedFieldAttributes{ref: ref}
}

func (mf ThingGroupIndexingConfigurationManagedFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mf.ref.InternalTokens()
}

func (mf ThingGroupIndexingConfigurationManagedFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mf.ref.Append("name"))
}

func (mf ThingGroupIndexingConfigurationManagedFieldAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mf.ref.Append("type"))
}

type ThingIndexingConfigurationAttributes struct {
	ref terra.Reference
}

func (tic ThingIndexingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return tic.ref, nil
}

func (tic ThingIndexingConfigurationAttributes) InternalWithRef(ref terra.Reference) ThingIndexingConfigurationAttributes {
	return ThingIndexingConfigurationAttributes{ref: ref}
}

func (tic ThingIndexingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tic.ref.InternalTokens()
}

func (tic ThingIndexingConfigurationAttributes) DeviceDefenderIndexingMode() terra.StringValue {
	return terra.ReferenceAsString(tic.ref.Append("device_defender_indexing_mode"))
}

func (tic ThingIndexingConfigurationAttributes) NamedShadowIndexingMode() terra.StringValue {
	return terra.ReferenceAsString(tic.ref.Append("named_shadow_indexing_mode"))
}

func (tic ThingIndexingConfigurationAttributes) ThingConnectivityIndexingMode() terra.StringValue {
	return terra.ReferenceAsString(tic.ref.Append("thing_connectivity_indexing_mode"))
}

func (tic ThingIndexingConfigurationAttributes) ThingIndexingMode() terra.StringValue {
	return terra.ReferenceAsString(tic.ref.Append("thing_indexing_mode"))
}

func (tic ThingIndexingConfigurationAttributes) CustomField() terra.SetValue[ThingIndexingConfigurationCustomFieldAttributes] {
	return terra.ReferenceAsSet[ThingIndexingConfigurationCustomFieldAttributes](tic.ref.Append("custom_field"))
}

func (tic ThingIndexingConfigurationAttributes) ManagedField() terra.SetValue[ThingIndexingConfigurationManagedFieldAttributes] {
	return terra.ReferenceAsSet[ThingIndexingConfigurationManagedFieldAttributes](tic.ref.Append("managed_field"))
}

type ThingIndexingConfigurationCustomFieldAttributes struct {
	ref terra.Reference
}

func (cf ThingIndexingConfigurationCustomFieldAttributes) InternalRef() (terra.Reference, error) {
	return cf.ref, nil
}

func (cf ThingIndexingConfigurationCustomFieldAttributes) InternalWithRef(ref terra.Reference) ThingIndexingConfigurationCustomFieldAttributes {
	return ThingIndexingConfigurationCustomFieldAttributes{ref: ref}
}

func (cf ThingIndexingConfigurationCustomFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cf.ref.InternalTokens()
}

func (cf ThingIndexingConfigurationCustomFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

func (cf ThingIndexingConfigurationCustomFieldAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("type"))
}

type ThingIndexingConfigurationManagedFieldAttributes struct {
	ref terra.Reference
}

func (mf ThingIndexingConfigurationManagedFieldAttributes) InternalRef() (terra.Reference, error) {
	return mf.ref, nil
}

func (mf ThingIndexingConfigurationManagedFieldAttributes) InternalWithRef(ref terra.Reference) ThingIndexingConfigurationManagedFieldAttributes {
	return ThingIndexingConfigurationManagedFieldAttributes{ref: ref}
}

func (mf ThingIndexingConfigurationManagedFieldAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mf.ref.InternalTokens()
}

func (mf ThingIndexingConfigurationManagedFieldAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mf.ref.Append("name"))
}

func (mf ThingIndexingConfigurationManagedFieldAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(mf.ref.Append("type"))
}

type ThingGroupIndexingConfigurationState struct {
	ThingGroupIndexingMode string                                             `json:"thing_group_indexing_mode"`
	CustomField            []ThingGroupIndexingConfigurationCustomFieldState  `json:"custom_field"`
	ManagedField           []ThingGroupIndexingConfigurationManagedFieldState `json:"managed_field"`
}

type ThingGroupIndexingConfigurationCustomFieldState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ThingGroupIndexingConfigurationManagedFieldState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ThingIndexingConfigurationState struct {
	DeviceDefenderIndexingMode    string                                        `json:"device_defender_indexing_mode"`
	NamedShadowIndexingMode       string                                        `json:"named_shadow_indexing_mode"`
	ThingConnectivityIndexingMode string                                        `json:"thing_connectivity_indexing_mode"`
	ThingIndexingMode             string                                        `json:"thing_indexing_mode"`
	CustomField                   []ThingIndexingConfigurationCustomFieldState  `json:"custom_field"`
	ManagedField                  []ThingIndexingConfigurationManagedFieldState `json:"managed_field"`
}

type ThingIndexingConfigurationCustomFieldState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ThingIndexingConfigurationManagedFieldState struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
