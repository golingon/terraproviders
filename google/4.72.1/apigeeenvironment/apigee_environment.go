// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apigeeenvironment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NodeConfig struct {
	// MaxNodeCount: string, optional
	MaxNodeCount terra.StringValue `hcl:"max_node_count,attr"`
	// MinNodeCount: string, optional
	MinNodeCount terra.StringValue `hcl:"min_node_count,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type NodeConfigAttributes struct {
	ref terra.Reference
}

func (nc NodeConfigAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NodeConfigAttributes) InternalWithRef(ref terra.Reference) NodeConfigAttributes {
	return NodeConfigAttributes{ref: ref}
}

func (nc NodeConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NodeConfigAttributes) CurrentAggregateNodeCount() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("current_aggregate_node_count"))
}

func (nc NodeConfigAttributes) MaxNodeCount() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("max_node_count"))
}

func (nc NodeConfigAttributes) MinNodeCount() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("min_node_count"))
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

type NodeConfigState struct {
	CurrentAggregateNodeCount string `json:"current_aggregate_node_count"`
	MaxNodeCount              string `json:"max_node_count"`
	MinNodeCount              string `json:"min_node_count"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
