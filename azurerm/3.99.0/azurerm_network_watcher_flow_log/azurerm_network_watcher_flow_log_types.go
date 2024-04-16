// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_network_watcher_flow_log

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type RetentionPolicy struct {
	// Days: number, required
	Days terra.NumberValue `hcl:"days,attr" validate:"required"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
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

type TrafficAnalytics struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// IntervalInMinutes: number, optional
	IntervalInMinutes terra.NumberValue `hcl:"interval_in_minutes,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// WorkspaceRegion: string, required
	WorkspaceRegion terra.StringValue `hcl:"workspace_region,attr" validate:"required"`
	// WorkspaceResourceId: string, required
	WorkspaceResourceId terra.StringValue `hcl:"workspace_resource_id,attr" validate:"required"`
}

type RetentionPolicyAttributes struct {
	ref terra.Reference
}

func (rp RetentionPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RetentionPolicyAttributes) InternalWithRef(ref terra.Reference) RetentionPolicyAttributes {
	return RetentionPolicyAttributes{ref: ref}
}

func (rp RetentionPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("days"))
}

func (rp RetentionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enabled"))
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

type TrafficAnalyticsAttributes struct {
	ref terra.Reference
}

func (ta TrafficAnalyticsAttributes) InternalRef() (terra.Reference, error) {
	return ta.ref, nil
}

func (ta TrafficAnalyticsAttributes) InternalWithRef(ref terra.Reference) TrafficAnalyticsAttributes {
	return TrafficAnalyticsAttributes{ref: ref}
}

func (ta TrafficAnalyticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ta.ref.InternalTokens()
}

func (ta TrafficAnalyticsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ta.ref.Append("enabled"))
}

func (ta TrafficAnalyticsAttributes) IntervalInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(ta.ref.Append("interval_in_minutes"))
}

func (ta TrafficAnalyticsAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("workspace_id"))
}

func (ta TrafficAnalyticsAttributes) WorkspaceRegion() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("workspace_region"))
}

func (ta TrafficAnalyticsAttributes) WorkspaceResourceId() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("workspace_resource_id"))
}

type RetentionPolicyState struct {
	Days    float64 `json:"days"`
	Enabled bool    `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TrafficAnalyticsState struct {
	Enabled             bool    `json:"enabled"`
	IntervalInMinutes   float64 `json:"interval_in_minutes"`
	WorkspaceId         string  `json:"workspace_id"`
	WorkspaceRegion     string  `json:"workspace_region"`
	WorkspaceResourceId string  `json:"workspace_resource_id"`
}
