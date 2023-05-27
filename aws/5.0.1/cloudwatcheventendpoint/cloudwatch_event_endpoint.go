// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudwatcheventendpoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EventBus struct {
	// EventBusArn: string, required
	EventBusArn terra.StringValue `hcl:"event_bus_arn,attr" validate:"required"`
}

type ReplicationConfig struct {
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
}

type RoutingConfig struct {
	// FailoverConfig: required
	FailoverConfig *FailoverConfig `hcl:"failover_config,block" validate:"required"`
}

type FailoverConfig struct {
	// Primary: required
	Primary *Primary `hcl:"primary,block" validate:"required"`
	// Secondary: required
	Secondary *Secondary `hcl:"secondary,block" validate:"required"`
}

type Primary struct {
	// HealthCheck: string, optional
	HealthCheck terra.StringValue `hcl:"health_check,attr"`
}

type Secondary struct {
	// Route: string, optional
	Route terra.StringValue `hcl:"route,attr"`
}

type EventBusAttributes struct {
	ref terra.Reference
}

func (eb EventBusAttributes) InternalRef() (terra.Reference, error) {
	return eb.ref, nil
}

func (eb EventBusAttributes) InternalWithRef(ref terra.Reference) EventBusAttributes {
	return EventBusAttributes{ref: ref}
}

func (eb EventBusAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eb.ref.InternalTokens()
}

func (eb EventBusAttributes) EventBusArn() terra.StringValue {
	return terra.ReferenceAsString(eb.ref.Append("event_bus_arn"))
}

type ReplicationConfigAttributes struct {
	ref terra.Reference
}

func (rc ReplicationConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc ReplicationConfigAttributes) InternalWithRef(ref terra.Reference) ReplicationConfigAttributes {
	return ReplicationConfigAttributes{ref: ref}
}

func (rc ReplicationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc ReplicationConfigAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("state"))
}

type RoutingConfigAttributes struct {
	ref terra.Reference
}

func (rc RoutingConfigAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RoutingConfigAttributes) InternalWithRef(ref terra.Reference) RoutingConfigAttributes {
	return RoutingConfigAttributes{ref: ref}
}

func (rc RoutingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RoutingConfigAttributes) FailoverConfig() terra.ListValue[FailoverConfigAttributes] {
	return terra.ReferenceAsList[FailoverConfigAttributes](rc.ref.Append("failover_config"))
}

type FailoverConfigAttributes struct {
	ref terra.Reference
}

func (fc FailoverConfigAttributes) InternalRef() (terra.Reference, error) {
	return fc.ref, nil
}

func (fc FailoverConfigAttributes) InternalWithRef(ref terra.Reference) FailoverConfigAttributes {
	return FailoverConfigAttributes{ref: ref}
}

func (fc FailoverConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fc.ref.InternalTokens()
}

func (fc FailoverConfigAttributes) Primary() terra.ListValue[PrimaryAttributes] {
	return terra.ReferenceAsList[PrimaryAttributes](fc.ref.Append("primary"))
}

func (fc FailoverConfigAttributes) Secondary() terra.ListValue[SecondaryAttributes] {
	return terra.ReferenceAsList[SecondaryAttributes](fc.ref.Append("secondary"))
}

type PrimaryAttributes struct {
	ref terra.Reference
}

func (p PrimaryAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PrimaryAttributes) InternalWithRef(ref terra.Reference) PrimaryAttributes {
	return PrimaryAttributes{ref: ref}
}

func (p PrimaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PrimaryAttributes) HealthCheck() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("health_check"))
}

type SecondaryAttributes struct {
	ref terra.Reference
}

func (s SecondaryAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SecondaryAttributes) InternalWithRef(ref terra.Reference) SecondaryAttributes {
	return SecondaryAttributes{ref: ref}
}

func (s SecondaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SecondaryAttributes) Route() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("route"))
}

type EventBusState struct {
	EventBusArn string `json:"event_bus_arn"`
}

type ReplicationConfigState struct {
	State string `json:"state"`
}

type RoutingConfigState struct {
	FailoverConfig []FailoverConfigState `json:"failover_config"`
}

type FailoverConfigState struct {
	Primary   []PrimaryState   `json:"primary"`
	Secondary []SecondaryState `json:"secondary"`
}

type PrimaryState struct {
	HealthCheck string `json:"health_check"`
}

type SecondaryState struct {
	Route string `json:"route"`
}
