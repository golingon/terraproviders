// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamemorydbcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClusterEndpoint struct{}

type Shards struct {
	// Nodes: min=0
	Nodes []Nodes `hcl:"nodes,block" validate:"min=0"`
}

type Nodes struct {
	// Endpoint: min=0
	Endpoint []Endpoint `hcl:"endpoint,block" validate:"min=0"`
}

type Endpoint struct{}

type ClusterEndpointAttributes struct {
	ref terra.Reference
}

func (ce ClusterEndpointAttributes) InternalRef() terra.Reference {
	return ce.ref
}

func (ce ClusterEndpointAttributes) InternalWithRef(ref terra.Reference) ClusterEndpointAttributes {
	return ClusterEndpointAttributes{ref: ref}
}

func (ce ClusterEndpointAttributes) InternalTokens() hclwrite.Tokens {
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

func (s ShardsAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s ShardsAttributes) InternalWithRef(ref terra.Reference) ShardsAttributes {
	return ShardsAttributes{ref: ref}
}

func (s ShardsAttributes) InternalTokens() hclwrite.Tokens {
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

func (s ShardsAttributes) Nodes() terra.SetValue[NodesAttributes] {
	return terra.ReferenceAsSet[NodesAttributes](s.ref.Append("nodes"))
}

type NodesAttributes struct {
	ref terra.Reference
}

func (n NodesAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NodesAttributes) InternalWithRef(ref terra.Reference) NodesAttributes {
	return NodesAttributes{ref: ref}
}

func (n NodesAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NodesAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("availability_zone"))
}

func (n NodesAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("create_time"))
}

func (n NodesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("name"))
}

func (n NodesAttributes) Endpoint() terra.ListValue[EndpointAttributes] {
	return terra.ReferenceAsList[EndpointAttributes](n.ref.Append("endpoint"))
}

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

func (e EndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("port"))
}

type ClusterEndpointState struct {
	Address string  `json:"address"`
	Port    float64 `json:"port"`
}

type ShardsState struct {
	Name     string       `json:"name"`
	NumNodes float64      `json:"num_nodes"`
	Slots    string       `json:"slots"`
	Nodes    []NodesState `json:"nodes"`
}

type NodesState struct {
	AvailabilityZone string          `json:"availability_zone"`
	CreateTime       string          `json:"create_time"`
	Name             string          `json:"name"`
	Endpoint         []EndpointState `json:"endpoint"`
}

type EndpointState struct {
	Address string  `json:"address"`
	Port    float64 `json:"port"`
}
