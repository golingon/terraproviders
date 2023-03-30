// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package workspacesworkspace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type WorkspaceProperties struct {
	// ComputeTypeName: string, optional
	ComputeTypeName terra.StringValue `hcl:"compute_type_name,attr"`
	// RootVolumeSizeGib: number, optional
	RootVolumeSizeGib terra.NumberValue `hcl:"root_volume_size_gib,attr"`
	// RunningMode: string, optional
	RunningMode terra.StringValue `hcl:"running_mode,attr"`
	// RunningModeAutoStopTimeoutInMinutes: number, optional
	RunningModeAutoStopTimeoutInMinutes terra.NumberValue `hcl:"running_mode_auto_stop_timeout_in_minutes,attr"`
	// UserVolumeSizeGib: number, optional
	UserVolumeSizeGib terra.NumberValue `hcl:"user_volume_size_gib,attr"`
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type WorkspacePropertiesAttributes struct {
	ref terra.Reference
}

func (wp WorkspacePropertiesAttributes) InternalRef() terra.Reference {
	return wp.ref
}

func (wp WorkspacePropertiesAttributes) InternalWithRef(ref terra.Reference) WorkspacePropertiesAttributes {
	return WorkspacePropertiesAttributes{ref: ref}
}

func (wp WorkspacePropertiesAttributes) InternalTokens() hclwrite.Tokens {
	return wp.ref.InternalTokens()
}

func (wp WorkspacePropertiesAttributes) ComputeTypeName() terra.StringValue {
	return terra.ReferenceString(wp.ref.Append("compute_type_name"))
}

func (wp WorkspacePropertiesAttributes) RootVolumeSizeGib() terra.NumberValue {
	return terra.ReferenceNumber(wp.ref.Append("root_volume_size_gib"))
}

func (wp WorkspacePropertiesAttributes) RunningMode() terra.StringValue {
	return terra.ReferenceString(wp.ref.Append("running_mode"))
}

func (wp WorkspacePropertiesAttributes) RunningModeAutoStopTimeoutInMinutes() terra.NumberValue {
	return terra.ReferenceNumber(wp.ref.Append("running_mode_auto_stop_timeout_in_minutes"))
}

func (wp WorkspacePropertiesAttributes) UserVolumeSizeGib() terra.NumberValue {
	return terra.ReferenceNumber(wp.ref.Append("user_volume_size_gib"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type WorkspacePropertiesState struct {
	ComputeTypeName                     string  `json:"compute_type_name"`
	RootVolumeSizeGib                   float64 `json:"root_volume_size_gib"`
	RunningMode                         string  `json:"running_mode"`
	RunningModeAutoStopTimeoutInMinutes float64 `json:"running_mode_auto_stop_timeout_in_minutes"`
	UserVolumeSizeGib                   float64 `json:"user_volume_size_gib"`
}
