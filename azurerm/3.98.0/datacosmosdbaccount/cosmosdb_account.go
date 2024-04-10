// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacosmosdbaccount

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Capabilities struct{}

type ConsistencyPolicy struct{}

type GeoLocation struct{}

type VirtualNetworkRule struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type CapabilitiesAttributes struct {
	ref terra.Reference
}

func (c CapabilitiesAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CapabilitiesAttributes) InternalWithRef(ref terra.Reference) CapabilitiesAttributes {
	return CapabilitiesAttributes{ref: ref}
}

func (c CapabilitiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CapabilitiesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

type ConsistencyPolicyAttributes struct {
	ref terra.Reference
}

func (cp ConsistencyPolicyAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp ConsistencyPolicyAttributes) InternalWithRef(ref terra.Reference) ConsistencyPolicyAttributes {
	return ConsistencyPolicyAttributes{ref: ref}
}

func (cp ConsistencyPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp ConsistencyPolicyAttributes) ConsistencyLevel() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("consistency_level"))
}

func (cp ConsistencyPolicyAttributes) MaxIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_interval_in_seconds"))
}

func (cp ConsistencyPolicyAttributes) MaxStalenessPrefix() terra.NumberValue {
	return terra.ReferenceAsNumber(cp.ref.Append("max_staleness_prefix"))
}

type GeoLocationAttributes struct {
	ref terra.Reference
}

func (gl GeoLocationAttributes) InternalRef() (terra.Reference, error) {
	return gl.ref, nil
}

func (gl GeoLocationAttributes) InternalWithRef(ref terra.Reference) GeoLocationAttributes {
	return GeoLocationAttributes{ref: ref}
}

func (gl GeoLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gl.ref.InternalTokens()
}

func (gl GeoLocationAttributes) FailoverPriority() terra.NumberValue {
	return terra.ReferenceAsNumber(gl.ref.Append("failover_priority"))
}

func (gl GeoLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gl.ref.Append("id"))
}

func (gl GeoLocationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gl.ref.Append("location"))
}

type VirtualNetworkRuleAttributes struct {
	ref terra.Reference
}

func (vnr VirtualNetworkRuleAttributes) InternalRef() (terra.Reference, error) {
	return vnr.ref, nil
}

func (vnr VirtualNetworkRuleAttributes) InternalWithRef(ref terra.Reference) VirtualNetworkRuleAttributes {
	return VirtualNetworkRuleAttributes{ref: ref}
}

func (vnr VirtualNetworkRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vnr.ref.InternalTokens()
}

func (vnr VirtualNetworkRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnr.ref.Append("id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type CapabilitiesState struct {
	Name string `json:"name"`
}

type ConsistencyPolicyState struct {
	ConsistencyLevel     string  `json:"consistency_level"`
	MaxIntervalInSeconds float64 `json:"max_interval_in_seconds"`
	MaxStalenessPrefix   float64 `json:"max_staleness_prefix"`
}

type GeoLocationState struct {
	FailoverPriority float64 `json:"failover_priority"`
	Id               string  `json:"id"`
	Location         string  `json:"location"`
}

type VirtualNetworkRuleState struct {
	Id string `json:"id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
