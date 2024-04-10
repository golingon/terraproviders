// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package daxcluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Nodes struct{}

type ServerSideEncryption struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type NodesAttributes struct {
	ref terra.Reference
}

func (n NodesAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NodesAttributes) InternalWithRef(ref terra.Reference) NodesAttributes {
	return NodesAttributes{ref: ref}
}

func (n NodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NodesAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("address"))
}

func (n NodesAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("availability_zone"))
}

func (n NodesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("id"))
}

func (n NodesAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(n.ref.Append("port"))
}

type ServerSideEncryptionAttributes struct {
	ref terra.Reference
}

func (sse ServerSideEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return sse.ref, nil
}

func (sse ServerSideEncryptionAttributes) InternalWithRef(ref terra.Reference) ServerSideEncryptionAttributes {
	return ServerSideEncryptionAttributes{ref: ref}
}

func (sse ServerSideEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sse.ref.InternalTokens()
}

func (sse ServerSideEncryptionAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(sse.ref.Append("enabled"))
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

type NodesState struct {
	Address          string  `json:"address"`
	AvailabilityZone string  `json:"availability_zone"`
	Id               string  `json:"id"`
	Port             float64 `json:"port"`
}

type ServerSideEncryptionState struct {
	Enabled bool `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
