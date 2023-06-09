// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package bigtableappprofile

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SingleClusterRouting struct {
	// AllowTransactionalWrites: bool, optional
	AllowTransactionalWrites terra.BoolValue `hcl:"allow_transactional_writes,attr"`
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SingleClusterRoutingAttributes struct {
	ref terra.Reference
}

func (scr SingleClusterRoutingAttributes) InternalRef() (terra.Reference, error) {
	return scr.ref, nil
}

func (scr SingleClusterRoutingAttributes) InternalWithRef(ref terra.Reference) SingleClusterRoutingAttributes {
	return SingleClusterRoutingAttributes{ref: ref}
}

func (scr SingleClusterRoutingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scr.ref.InternalTokens()
}

func (scr SingleClusterRoutingAttributes) AllowTransactionalWrites() terra.BoolValue {
	return terra.ReferenceAsBool(scr.ref.Append("allow_transactional_writes"))
}

func (scr SingleClusterRoutingAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(scr.ref.Append("cluster_id"))
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

type SingleClusterRoutingState struct {
	AllowTransactionalWrites bool   `json:"allow_transactional_writes"`
	ClusterId                string `json:"cluster_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
