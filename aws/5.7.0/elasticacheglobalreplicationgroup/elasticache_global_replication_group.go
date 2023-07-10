// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elasticacheglobalreplicationgroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type GlobalNodeGroups struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type GlobalNodeGroupsAttributes struct {
	ref terra.Reference
}

func (gng GlobalNodeGroupsAttributes) InternalRef() (terra.Reference, error) {
	return gng.ref, nil
}

func (gng GlobalNodeGroupsAttributes) InternalWithRef(ref terra.Reference) GlobalNodeGroupsAttributes {
	return GlobalNodeGroupsAttributes{ref: ref}
}

func (gng GlobalNodeGroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gng.ref.InternalTokens()
}

func (gng GlobalNodeGroupsAttributes) GlobalNodeGroupId() terra.StringValue {
	return terra.ReferenceAsString(gng.ref.Append("global_node_group_id"))
}

func (gng GlobalNodeGroupsAttributes) Slots() terra.StringValue {
	return terra.ReferenceAsString(gng.ref.Append("slots"))
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

type GlobalNodeGroupsState struct {
	GlobalNodeGroupId string `json:"global_node_group_id"`
	Slots             string `json:"slots"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
