// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_memorydb_cluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ClusterEndpointAttributes struct {
	ref terra.Reference
}

func (ce ClusterEndpointAttributes) InternalRef() (terra.Reference, error) {
	return ce.ref, nil
}

func (ce ClusterEndpointAttributes) InternalWithRef(ref terra.Reference) ClusterEndpointAttributes {
	return ClusterEndpointAttributes{ref: ref}
}

func (ce ClusterEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ce.ref.InternalTokens()
}

func (ce ClusterEndpointAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("address"))
}

func (ce ClusterEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ce.ref.Append("port"))
}

type ShardsAttributes struct {
	ref terra.Reference
}

func (s ShardsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ShardsAttributes) InternalWithRef(ref terra.Reference) ShardsAttributes {
	return ShardsAttributes{ref: ref}
}

func (s ShardsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ShardsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s ShardsAttributes) NumNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("num_nodes"))
}

func (s ShardsAttributes) Slots() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("slots"))
}

func (s ShardsAttributes) Nodes() terra.SetValue[ShardsNodesAttributes] {
	return terra.ReferenceAsSet[ShardsNodesAttributes](s.ref.Append("nodes"))
}

type ShardsNodesAttributes struct {
	ref terra.Reference
}

func (n ShardsNodesAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n ShardsNodesAttributes) InternalWithRef(ref terra.Reference) ShardsNodesAttributes {
	return ShardsNodesAttributes{ref: ref}
}

func (n ShardsNodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n ShardsNodesAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("availability_zone"))
}

func (n ShardsNodesAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("create_time"))
}

func (n ShardsNodesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("name"))
}

func (n ShardsNodesAttributes) Endpoint() terra.ListValue[ShardsNodesEndpointAttributes] {
	return terra.ReferenceAsList[ShardsNodesEndpointAttributes](n.ref.Append("endpoint"))
}

type ShardsNodesEndpointAttributes struct {
	ref terra.Reference
}

func (e ShardsNodesEndpointAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ShardsNodesEndpointAttributes) InternalWithRef(ref terra.Reference) ShardsNodesEndpointAttributes {
	return ShardsNodesEndpointAttributes{ref: ref}
}

func (e ShardsNodesEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ShardsNodesEndpointAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

func (e ShardsNodesEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("port"))
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

type ClusterEndpointState struct {
	Address string  `json:"address"`
	Port    float64 `json:"port"`
}

type ShardsState struct {
	Name     string             `json:"name"`
	NumNodes float64            `json:"num_nodes"`
	Slots    string             `json:"slots"`
	Nodes    []ShardsNodesState `json:"nodes"`
}

type ShardsNodesState struct {
	AvailabilityZone string                     `json:"availability_zone"`
	CreateTime       string                     `json:"create_time"`
	Name             string                     `json:"name"`
	Endpoint         []ShardsNodesEndpointState `json:"endpoint"`
}

type ShardsNodesEndpointState struct {
	Address string  `json:"address"`
	Port    float64 `json:"port"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
