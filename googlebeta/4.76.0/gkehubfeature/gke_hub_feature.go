// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gkehubfeature

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ResourceState struct{}

type State struct {
	// StateState: min=0
	State []StateState `hcl:"state,block" validate:"min=0"`
}

type StateState struct{}

type Spec struct {
	// Fleetobservability: optional
	Fleetobservability *Fleetobservability `hcl:"fleetobservability,block"`
	// Multiclusteringress: optional
	Multiclusteringress *Multiclusteringress `hcl:"multiclusteringress,block"`
}

type Fleetobservability struct {
	// LoggingConfig: optional
	LoggingConfig *LoggingConfig `hcl:"logging_config,block"`
}

type LoggingConfig struct {
	// DefaultConfig: optional
	DefaultConfig *DefaultConfig `hcl:"default_config,block"`
	// FleetScopeLogsConfig: optional
	FleetScopeLogsConfig *FleetScopeLogsConfig `hcl:"fleet_scope_logs_config,block"`
}

type DefaultConfig struct {
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
}

type FleetScopeLogsConfig struct {
	// Mode: string, optional
	Mode terra.StringValue `hcl:"mode,attr"`
}

type Multiclusteringress struct {
	// ConfigMembership: string, required
	ConfigMembership terra.StringValue `hcl:"config_membership,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ResourceStateAttributes struct {
	ref terra.Reference
}

func (rs ResourceStateAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ResourceStateAttributes) InternalWithRef(ref terra.Reference) ResourceStateAttributes {
	return ResourceStateAttributes{ref: ref}
}

func (rs ResourceStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ResourceStateAttributes) HasResources() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("has_resources"))
}

func (rs ResourceStateAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("state"))
}

type StateAttributes struct {
	ref terra.Reference
}

func (s StateAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StateAttributes) InternalWithRef(ref terra.Reference) StateAttributes {
	return StateAttributes{ref: ref}
}

func (s StateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StateAttributes) State() terra.ListValue[StateStateAttributes] {
	return terra.ReferenceAsList[StateStateAttributes](s.ref.Append("state"))
}

type StateStateAttributes struct {
	ref terra.Reference
}

func (s StateStateAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StateStateAttributes) InternalWithRef(ref terra.Reference) StateStateAttributes {
	return StateStateAttributes{ref: ref}
}

func (s StateStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StateStateAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("code"))
}

func (s StateStateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
}

func (s StateStateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("update_time"))
}

type SpecAttributes struct {
	ref terra.Reference
}

func (s SpecAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SpecAttributes) InternalWithRef(ref terra.Reference) SpecAttributes {
	return SpecAttributes{ref: ref}
}

func (s SpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SpecAttributes) Fleetobservability() terra.ListValue[FleetobservabilityAttributes] {
	return terra.ReferenceAsList[FleetobservabilityAttributes](s.ref.Append("fleetobservability"))
}

func (s SpecAttributes) Multiclusteringress() terra.ListValue[MulticlusteringressAttributes] {
	return terra.ReferenceAsList[MulticlusteringressAttributes](s.ref.Append("multiclusteringress"))
}

type FleetobservabilityAttributes struct {
	ref terra.Reference
}

func (f FleetobservabilityAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FleetobservabilityAttributes) InternalWithRef(ref terra.Reference) FleetobservabilityAttributes {
	return FleetobservabilityAttributes{ref: ref}
}

func (f FleetobservabilityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FleetobservabilityAttributes) LoggingConfig() terra.ListValue[LoggingConfigAttributes] {
	return terra.ReferenceAsList[LoggingConfigAttributes](f.ref.Append("logging_config"))
}

type LoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LoggingConfigAttributes) InternalWithRef(ref terra.Reference) LoggingConfigAttributes {
	return LoggingConfigAttributes{ref: ref}
}

func (lc LoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigAttributes) DefaultConfig() terra.ListValue[DefaultConfigAttributes] {
	return terra.ReferenceAsList[DefaultConfigAttributes](lc.ref.Append("default_config"))
}

func (lc LoggingConfigAttributes) FleetScopeLogsConfig() terra.ListValue[FleetScopeLogsConfigAttributes] {
	return terra.ReferenceAsList[FleetScopeLogsConfigAttributes](lc.ref.Append("fleet_scope_logs_config"))
}

type DefaultConfigAttributes struct {
	ref terra.Reference
}

func (dc DefaultConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DefaultConfigAttributes) InternalWithRef(ref terra.Reference) DefaultConfigAttributes {
	return DefaultConfigAttributes{ref: ref}
}

func (dc DefaultConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DefaultConfigAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("mode"))
}

type FleetScopeLogsConfigAttributes struct {
	ref terra.Reference
}

func (fslc FleetScopeLogsConfigAttributes) InternalRef() (terra.Reference, error) {
	return fslc.ref, nil
}

func (fslc FleetScopeLogsConfigAttributes) InternalWithRef(ref terra.Reference) FleetScopeLogsConfigAttributes {
	return FleetScopeLogsConfigAttributes{ref: ref}
}

func (fslc FleetScopeLogsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fslc.ref.InternalTokens()
}

func (fslc FleetScopeLogsConfigAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(fslc.ref.Append("mode"))
}

type MulticlusteringressAttributes struct {
	ref terra.Reference
}

func (m MulticlusteringressAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MulticlusteringressAttributes) InternalWithRef(ref terra.Reference) MulticlusteringressAttributes {
	return MulticlusteringressAttributes{ref: ref}
}

func (m MulticlusteringressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MulticlusteringressAttributes) ConfigMembership() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("config_membership"))
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

type ResourceStateState struct {
	HasResources bool   `json:"has_resources"`
	State        string `json:"state"`
}

type StateState struct {
	State []StateStateState `json:"state"`
}

type StateStateState struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	UpdateTime  string `json:"update_time"`
}

type SpecState struct {
	Fleetobservability  []FleetobservabilityState  `json:"fleetobservability"`
	Multiclusteringress []MulticlusteringressState `json:"multiclusteringress"`
}

type FleetobservabilityState struct {
	LoggingConfig []LoggingConfigState `json:"logging_config"`
}

type LoggingConfigState struct {
	DefaultConfig        []DefaultConfigState        `json:"default_config"`
	FleetScopeLogsConfig []FleetScopeLogsConfigState `json:"fleet_scope_logs_config"`
}

type DefaultConfigState struct {
	Mode string `json:"mode"`
}

type FleetScopeLogsConfigState struct {
	Mode string `json:"mode"`
}

type MulticlusteringressState struct {
	ConfigMembership string `json:"config_membership"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
