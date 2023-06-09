// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkmanagerconnectivityconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AppliesToGroup struct {
	// GlobalMeshEnabled: bool, optional
	GlobalMeshEnabled terra.BoolValue `hcl:"global_mesh_enabled,attr"`
	// GroupConnectivity: string, required
	GroupConnectivity terra.StringValue `hcl:"group_connectivity,attr" validate:"required"`
	// NetworkGroupId: string, required
	NetworkGroupId terra.StringValue `hcl:"network_group_id,attr" validate:"required"`
	// UseHubGateway: bool, optional
	UseHubGateway terra.BoolValue `hcl:"use_hub_gateway,attr"`
}

type Hub struct {
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
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

type AppliesToGroupAttributes struct {
	ref terra.Reference
}

func (atg AppliesToGroupAttributes) InternalRef() (terra.Reference, error) {
	return atg.ref, nil
}

func (atg AppliesToGroupAttributes) InternalWithRef(ref terra.Reference) AppliesToGroupAttributes {
	return AppliesToGroupAttributes{ref: ref}
}

func (atg AppliesToGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return atg.ref.InternalTokens()
}

func (atg AppliesToGroupAttributes) GlobalMeshEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(atg.ref.Append("global_mesh_enabled"))
}

func (atg AppliesToGroupAttributes) GroupConnectivity() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("group_connectivity"))
}

func (atg AppliesToGroupAttributes) NetworkGroupId() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("network_group_id"))
}

func (atg AppliesToGroupAttributes) UseHubGateway() terra.BoolValue {
	return terra.ReferenceAsBool(atg.ref.Append("use_hub_gateway"))
}

type HubAttributes struct {
	ref terra.Reference
}

func (h HubAttributes) InternalRef() (terra.Reference, error) {
	return h.ref, nil
}

func (h HubAttributes) InternalWithRef(ref terra.Reference) HubAttributes {
	return HubAttributes{ref: ref}
}

func (h HubAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return h.ref.InternalTokens()
}

func (h HubAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("resource_id"))
}

func (h HubAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("resource_type"))
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

type AppliesToGroupState struct {
	GlobalMeshEnabled bool   `json:"global_mesh_enabled"`
	GroupConnectivity string `json:"group_connectivity"`
	NetworkGroupId    string `json:"network_group_id"`
	UseHubGateway     bool   `json:"use_hub_gateway"`
}

type HubState struct {
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
