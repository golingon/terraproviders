// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkwatcherflowlog

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
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

func (rp RetentionPolicyAttributes) InternalRef() terra.Reference {
	return rp.ref
}

func (rp RetentionPolicyAttributes) InternalWithRef(ref terra.Reference) RetentionPolicyAttributes {
	return RetentionPolicyAttributes{ref: ref}
}

func (rp RetentionPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return rp.ref.InternalTokens()
}

func (rp RetentionPolicyAttributes) Days() terra.NumberValue {
	return terra.ReferenceNumber(rp.ref.Append("days"))
}

func (rp RetentionPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(rp.ref.Append("enabled"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type TrafficAnalyticsAttributes struct {
	ref terra.Reference
}

func (ta TrafficAnalyticsAttributes) InternalRef() terra.Reference {
	return ta.ref
}

func (ta TrafficAnalyticsAttributes) InternalWithRef(ref terra.Reference) TrafficAnalyticsAttributes {
	return TrafficAnalyticsAttributes{ref: ref}
}

func (ta TrafficAnalyticsAttributes) InternalTokens() hclwrite.Tokens {
	return ta.ref.InternalTokens()
}

func (ta TrafficAnalyticsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ta.ref.Append("enabled"))
}

func (ta TrafficAnalyticsAttributes) IntervalInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(ta.ref.Append("interval_in_minutes"))
}

func (ta TrafficAnalyticsAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceString(ta.ref.Append("workspace_id"))
}

func (ta TrafficAnalyticsAttributes) WorkspaceRegion() terra.StringValue {
	return terra.ReferenceString(ta.ref.Append("workspace_region"))
}

func (ta TrafficAnalyticsAttributes) WorkspaceResourceId() terra.StringValue {
	return terra.ReferenceString(ta.ref.Append("workspace_resource_id"))
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
