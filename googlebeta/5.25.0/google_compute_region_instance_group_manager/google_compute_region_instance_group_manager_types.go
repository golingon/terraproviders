// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_instance_group_manager

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AllInstancesConfig struct {
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
}

type AutoHealingPolicies struct {
	// HealthCheck: string, required
	HealthCheck terra.StringValue `hcl:"health_check,attr" validate:"required"`
	// InitialDelaySec: number, required
	InitialDelaySec terra.NumberValue `hcl:"initial_delay_sec,attr" validate:"required"`
}

type InstanceLifecyclePolicy struct {
	// DefaultActionOnFailure: string, optional
	DefaultActionOnFailure terra.StringValue `hcl:"default_action_on_failure,attr"`
	// ForceUpdateOnRepair: string, optional
	ForceUpdateOnRepair terra.StringValue `hcl:"force_update_on_repair,attr"`
}

type NamedPort struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type StatefulDisk struct {
	// DeleteRule: string, optional
	DeleteRule terra.StringValue `hcl:"delete_rule,attr"`
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
}

type StatefulExternalIp struct {
	// DeleteRule: string, optional
	DeleteRule terra.StringValue `hcl:"delete_rule,attr"`
	// InterfaceName: string, optional
	InterfaceName terra.StringValue `hcl:"interface_name,attr"`
}

type StatefulInternalIp struct {
	// DeleteRule: string, optional
	DeleteRule terra.StringValue `hcl:"delete_rule,attr"`
	// InterfaceName: string, optional
	InterfaceName terra.StringValue `hcl:"interface_name,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UpdatePolicy struct {
	// InstanceRedistributionType: string, optional
	InstanceRedistributionType terra.StringValue `hcl:"instance_redistribution_type,attr"`
	// MaxSurgeFixed: number, optional
	MaxSurgeFixed terra.NumberValue `hcl:"max_surge_fixed,attr"`
	// MaxSurgePercent: number, optional
	MaxSurgePercent terra.NumberValue `hcl:"max_surge_percent,attr"`
	// MaxUnavailableFixed: number, optional
	MaxUnavailableFixed terra.NumberValue `hcl:"max_unavailable_fixed,attr"`
	// MaxUnavailablePercent: number, optional
	MaxUnavailablePercent terra.NumberValue `hcl:"max_unavailable_percent,attr"`
	// MinReadySec: number, optional
	MinReadySec terra.NumberValue `hcl:"min_ready_sec,attr"`
	// MinimalAction: string, required
	MinimalAction terra.StringValue `hcl:"minimal_action,attr" validate:"required"`
	// MostDisruptiveAllowedAction: string, optional
	MostDisruptiveAllowedAction terra.StringValue `hcl:"most_disruptive_allowed_action,attr"`
	// ReplacementMethod: string, optional
	ReplacementMethod terra.StringValue `hcl:"replacement_method,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Version struct {
	// InstanceTemplate: string, required
	InstanceTemplate terra.StringValue `hcl:"instance_template,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// VersionTargetSize: optional
	TargetSize *VersionTargetSize `hcl:"target_size,block"`
}

type VersionTargetSize struct {
	// Fixed: number, optional
	Fixed terra.NumberValue `hcl:"fixed,attr"`
	// Percent: number, optional
	Percent terra.NumberValue `hcl:"percent,attr"`
}

type StatusAttributes struct {
	ref terra.Reference
}

func (s StatusAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatusAttributes) InternalWithRef(ref terra.Reference) StatusAttributes {
	return StatusAttributes{ref: ref}
}

func (s StatusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatusAttributes) IsStable() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("is_stable"))
}

func (s StatusAttributes) AllInstancesConfig() terra.ListValue[StatusAllInstancesConfigAttributes] {
	return terra.ReferenceAsList[StatusAllInstancesConfigAttributes](s.ref.Append("all_instances_config"))
}

func (s StatusAttributes) Stateful() terra.ListValue[StatusStatefulAttributes] {
	return terra.ReferenceAsList[StatusStatefulAttributes](s.ref.Append("stateful"))
}

func (s StatusAttributes) VersionTarget() terra.ListValue[StatusVersionTargetAttributes] {
	return terra.ReferenceAsList[StatusVersionTargetAttributes](s.ref.Append("version_target"))
}

type StatusAllInstancesConfigAttributes struct {
	ref terra.Reference
}

func (aic StatusAllInstancesConfigAttributes) InternalRef() (terra.Reference, error) {
	return aic.ref, nil
}

func (aic StatusAllInstancesConfigAttributes) InternalWithRef(ref terra.Reference) StatusAllInstancesConfigAttributes {
	return StatusAllInstancesConfigAttributes{ref: ref}
}

func (aic StatusAllInstancesConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aic.ref.InternalTokens()
}

func (aic StatusAllInstancesConfigAttributes) CurrentRevision() terra.StringValue {
	return terra.ReferenceAsString(aic.ref.Append("current_revision"))
}

func (aic StatusAllInstancesConfigAttributes) Effective() terra.BoolValue {
	return terra.ReferenceAsBool(aic.ref.Append("effective"))
}

type StatusStatefulAttributes struct {
	ref terra.Reference
}

func (s StatusStatefulAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StatusStatefulAttributes) InternalWithRef(ref terra.Reference) StatusStatefulAttributes {
	return StatusStatefulAttributes{ref: ref}
}

func (s StatusStatefulAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StatusStatefulAttributes) HasStatefulConfig() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("has_stateful_config"))
}

func (s StatusStatefulAttributes) PerInstanceConfigs() terra.ListValue[StatusStatefulPerInstanceConfigsAttributes] {
	return terra.ReferenceAsList[StatusStatefulPerInstanceConfigsAttributes](s.ref.Append("per_instance_configs"))
}

type StatusStatefulPerInstanceConfigsAttributes struct {
	ref terra.Reference
}

func (pic StatusStatefulPerInstanceConfigsAttributes) InternalRef() (terra.Reference, error) {
	return pic.ref, nil
}

func (pic StatusStatefulPerInstanceConfigsAttributes) InternalWithRef(ref terra.Reference) StatusStatefulPerInstanceConfigsAttributes {
	return StatusStatefulPerInstanceConfigsAttributes{ref: ref}
}

func (pic StatusStatefulPerInstanceConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pic.ref.InternalTokens()
}

func (pic StatusStatefulPerInstanceConfigsAttributes) AllEffective() terra.BoolValue {
	return terra.ReferenceAsBool(pic.ref.Append("all_effective"))
}

type StatusVersionTargetAttributes struct {
	ref terra.Reference
}

func (vt StatusVersionTargetAttributes) InternalRef() (terra.Reference, error) {
	return vt.ref, nil
}

func (vt StatusVersionTargetAttributes) InternalWithRef(ref terra.Reference) StatusVersionTargetAttributes {
	return StatusVersionTargetAttributes{ref: ref}
}

func (vt StatusVersionTargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vt.ref.InternalTokens()
}

func (vt StatusVersionTargetAttributes) IsReached() terra.BoolValue {
	return terra.ReferenceAsBool(vt.ref.Append("is_reached"))
}

type AllInstancesConfigAttributes struct {
	ref terra.Reference
}

func (aic AllInstancesConfigAttributes) InternalRef() (terra.Reference, error) {
	return aic.ref, nil
}

func (aic AllInstancesConfigAttributes) InternalWithRef(ref terra.Reference) AllInstancesConfigAttributes {
	return AllInstancesConfigAttributes{ref: ref}
}

func (aic AllInstancesConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aic.ref.InternalTokens()
}

func (aic AllInstancesConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aic.ref.Append("labels"))
}

func (aic AllInstancesConfigAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aic.ref.Append("metadata"))
}

type AutoHealingPoliciesAttributes struct {
	ref terra.Reference
}

func (ahp AutoHealingPoliciesAttributes) InternalRef() (terra.Reference, error) {
	return ahp.ref, nil
}

func (ahp AutoHealingPoliciesAttributes) InternalWithRef(ref terra.Reference) AutoHealingPoliciesAttributes {
	return AutoHealingPoliciesAttributes{ref: ref}
}

func (ahp AutoHealingPoliciesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ahp.ref.InternalTokens()
}

func (ahp AutoHealingPoliciesAttributes) HealthCheck() terra.StringValue {
	return terra.ReferenceAsString(ahp.ref.Append("health_check"))
}

func (ahp AutoHealingPoliciesAttributes) InitialDelaySec() terra.NumberValue {
	return terra.ReferenceAsNumber(ahp.ref.Append("initial_delay_sec"))
}

type InstanceLifecyclePolicyAttributes struct {
	ref terra.Reference
}

func (ilp InstanceLifecyclePolicyAttributes) InternalRef() (terra.Reference, error) {
	return ilp.ref, nil
}

func (ilp InstanceLifecyclePolicyAttributes) InternalWithRef(ref terra.Reference) InstanceLifecyclePolicyAttributes {
	return InstanceLifecyclePolicyAttributes{ref: ref}
}

func (ilp InstanceLifecyclePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ilp.ref.InternalTokens()
}

func (ilp InstanceLifecyclePolicyAttributes) DefaultActionOnFailure() terra.StringValue {
	return terra.ReferenceAsString(ilp.ref.Append("default_action_on_failure"))
}

func (ilp InstanceLifecyclePolicyAttributes) ForceUpdateOnRepair() terra.StringValue {
	return terra.ReferenceAsString(ilp.ref.Append("force_update_on_repair"))
}

type NamedPortAttributes struct {
	ref terra.Reference
}

func (np NamedPortAttributes) InternalRef() (terra.Reference, error) {
	return np.ref, nil
}

func (np NamedPortAttributes) InternalWithRef(ref terra.Reference) NamedPortAttributes {
	return NamedPortAttributes{ref: ref}
}

func (np NamedPortAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return np.ref.InternalTokens()
}

func (np NamedPortAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("name"))
}

func (np NamedPortAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(np.ref.Append("port"))
}

type StatefulDiskAttributes struct {
	ref terra.Reference
}

func (sd StatefulDiskAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd StatefulDiskAttributes) InternalWithRef(ref terra.Reference) StatefulDiskAttributes {
	return StatefulDiskAttributes{ref: ref}
}

func (sd StatefulDiskAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd StatefulDiskAttributes) DeleteRule() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("delete_rule"))
}

func (sd StatefulDiskAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("device_name"))
}

type StatefulExternalIpAttributes struct {
	ref terra.Reference
}

func (sei StatefulExternalIpAttributes) InternalRef() (terra.Reference, error) {
	return sei.ref, nil
}

func (sei StatefulExternalIpAttributes) InternalWithRef(ref terra.Reference) StatefulExternalIpAttributes {
	return StatefulExternalIpAttributes{ref: ref}
}

func (sei StatefulExternalIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sei.ref.InternalTokens()
}

func (sei StatefulExternalIpAttributes) DeleteRule() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("delete_rule"))
}

func (sei StatefulExternalIpAttributes) InterfaceName() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("interface_name"))
}

type StatefulInternalIpAttributes struct {
	ref terra.Reference
}

func (sii StatefulInternalIpAttributes) InternalRef() (terra.Reference, error) {
	return sii.ref, nil
}

func (sii StatefulInternalIpAttributes) InternalWithRef(ref terra.Reference) StatefulInternalIpAttributes {
	return StatefulInternalIpAttributes{ref: ref}
}

func (sii StatefulInternalIpAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sii.ref.InternalTokens()
}

func (sii StatefulInternalIpAttributes) DeleteRule() terra.StringValue {
	return terra.ReferenceAsString(sii.ref.Append("delete_rule"))
}

func (sii StatefulInternalIpAttributes) InterfaceName() terra.StringValue {
	return terra.ReferenceAsString(sii.ref.Append("interface_name"))
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

type UpdatePolicyAttributes struct {
	ref terra.Reference
}

func (up UpdatePolicyAttributes) InternalRef() (terra.Reference, error) {
	return up.ref, nil
}

func (up UpdatePolicyAttributes) InternalWithRef(ref terra.Reference) UpdatePolicyAttributes {
	return UpdatePolicyAttributes{ref: ref}
}

func (up UpdatePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return up.ref.InternalTokens()
}

func (up UpdatePolicyAttributes) InstanceRedistributionType() terra.StringValue {
	return terra.ReferenceAsString(up.ref.Append("instance_redistribution_type"))
}

func (up UpdatePolicyAttributes) MaxSurgeFixed() terra.NumberValue {
	return terra.ReferenceAsNumber(up.ref.Append("max_surge_fixed"))
}

func (up UpdatePolicyAttributes) MaxSurgePercent() terra.NumberValue {
	return terra.ReferenceAsNumber(up.ref.Append("max_surge_percent"))
}

func (up UpdatePolicyAttributes) MaxUnavailableFixed() terra.NumberValue {
	return terra.ReferenceAsNumber(up.ref.Append("max_unavailable_fixed"))
}

func (up UpdatePolicyAttributes) MaxUnavailablePercent() terra.NumberValue {
	return terra.ReferenceAsNumber(up.ref.Append("max_unavailable_percent"))
}

func (up UpdatePolicyAttributes) MinReadySec() terra.NumberValue {
	return terra.ReferenceAsNumber(up.ref.Append("min_ready_sec"))
}

func (up UpdatePolicyAttributes) MinimalAction() terra.StringValue {
	return terra.ReferenceAsString(up.ref.Append("minimal_action"))
}

func (up UpdatePolicyAttributes) MostDisruptiveAllowedAction() terra.StringValue {
	return terra.ReferenceAsString(up.ref.Append("most_disruptive_allowed_action"))
}

func (up UpdatePolicyAttributes) ReplacementMethod() terra.StringValue {
	return terra.ReferenceAsString(up.ref.Append("replacement_method"))
}

func (up UpdatePolicyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(up.ref.Append("type"))
}

type VersionAttributes struct {
	ref terra.Reference
}

func (v VersionAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VersionAttributes) InternalWithRef(ref terra.Reference) VersionAttributes {
	return VersionAttributes{ref: ref}
}

func (v VersionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return v.ref.InternalTokens()
}

func (v VersionAttributes) InstanceTemplate() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("instance_template"))
}

func (v VersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("name"))
}

func (v VersionAttributes) TargetSize() terra.ListValue[VersionTargetSizeAttributes] {
	return terra.ReferenceAsList[VersionTargetSizeAttributes](v.ref.Append("target_size"))
}

type VersionTargetSizeAttributes struct {
	ref terra.Reference
}

func (ts VersionTargetSizeAttributes) InternalRef() (terra.Reference, error) {
	return ts.ref, nil
}

func (ts VersionTargetSizeAttributes) InternalWithRef(ref terra.Reference) VersionTargetSizeAttributes {
	return VersionTargetSizeAttributes{ref: ref}
}

func (ts VersionTargetSizeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ts.ref.InternalTokens()
}

func (ts VersionTargetSizeAttributes) Fixed() terra.NumberValue {
	return terra.ReferenceAsNumber(ts.ref.Append("fixed"))
}

func (ts VersionTargetSizeAttributes) Percent() terra.NumberValue {
	return terra.ReferenceAsNumber(ts.ref.Append("percent"))
}

type StatusState struct {
	IsStable           bool                            `json:"is_stable"`
	AllInstancesConfig []StatusAllInstancesConfigState `json:"all_instances_config"`
	Stateful           []StatusStatefulState           `json:"stateful"`
	VersionTarget      []StatusVersionTargetState      `json:"version_target"`
}

type StatusAllInstancesConfigState struct {
	CurrentRevision string `json:"current_revision"`
	Effective       bool   `json:"effective"`
}

type StatusStatefulState struct {
	HasStatefulConfig  bool                                    `json:"has_stateful_config"`
	PerInstanceConfigs []StatusStatefulPerInstanceConfigsState `json:"per_instance_configs"`
}

type StatusStatefulPerInstanceConfigsState struct {
	AllEffective bool `json:"all_effective"`
}

type StatusVersionTargetState struct {
	IsReached bool `json:"is_reached"`
}

type AllInstancesConfigState struct {
	Labels   map[string]string `json:"labels"`
	Metadata map[string]string `json:"metadata"`
}

type AutoHealingPoliciesState struct {
	HealthCheck     string  `json:"health_check"`
	InitialDelaySec float64 `json:"initial_delay_sec"`
}

type InstanceLifecyclePolicyState struct {
	DefaultActionOnFailure string `json:"default_action_on_failure"`
	ForceUpdateOnRepair    string `json:"force_update_on_repair"`
}

type NamedPortState struct {
	Name string  `json:"name"`
	Port float64 `json:"port"`
}

type StatefulDiskState struct {
	DeleteRule string `json:"delete_rule"`
	DeviceName string `json:"device_name"`
}

type StatefulExternalIpState struct {
	DeleteRule    string `json:"delete_rule"`
	InterfaceName string `json:"interface_name"`
}

type StatefulInternalIpState struct {
	DeleteRule    string `json:"delete_rule"`
	InterfaceName string `json:"interface_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type UpdatePolicyState struct {
	InstanceRedistributionType  string  `json:"instance_redistribution_type"`
	MaxSurgeFixed               float64 `json:"max_surge_fixed"`
	MaxSurgePercent             float64 `json:"max_surge_percent"`
	MaxUnavailableFixed         float64 `json:"max_unavailable_fixed"`
	MaxUnavailablePercent       float64 `json:"max_unavailable_percent"`
	MinReadySec                 float64 `json:"min_ready_sec"`
	MinimalAction               string  `json:"minimal_action"`
	MostDisruptiveAllowedAction string  `json:"most_disruptive_allowed_action"`
	ReplacementMethod           string  `json:"replacement_method"`
	Type                        string  `json:"type"`
}

type VersionState struct {
	InstanceTemplate string                   `json:"instance_template"`
	Name             string                   `json:"name"`
	TargetSize       []VersionTargetSizeState `json:"target_size"`
}

type VersionTargetSizeState struct {
	Fixed   float64 `json:"fixed"`
	Percent float64 `json:"percent"`
}
