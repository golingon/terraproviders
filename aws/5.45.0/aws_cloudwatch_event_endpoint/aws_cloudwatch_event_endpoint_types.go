// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudwatch_event_endpoint

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
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
	// RoutingConfigFailoverConfig: required
	FailoverConfig *RoutingConfigFailoverConfig `hcl:"failover_config,block" validate:"required"`
}

type RoutingConfigFailoverConfig struct {
	// RoutingConfigFailoverConfigPrimary: required
	Primary *RoutingConfigFailoverConfigPrimary `hcl:"primary,block" validate:"required"`
	// RoutingConfigFailoverConfigSecondary: required
	Secondary *RoutingConfigFailoverConfigSecondary `hcl:"secondary,block" validate:"required"`
}

type RoutingConfigFailoverConfigPrimary struct {
	// HealthCheck: string, optional
	HealthCheck terra.StringValue `hcl:"health_check,attr"`
}

type RoutingConfigFailoverConfigSecondary struct {
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

func (rc RoutingConfigAttributes) FailoverConfig() terra.ListValue[RoutingConfigFailoverConfigAttributes] {
	return terra.ReferenceAsList[RoutingConfigFailoverConfigAttributes](rc.ref.Append("failover_config"))
}

type RoutingConfigFailoverConfigAttributes struct {
	ref terra.Reference
}

func (fc RoutingConfigFailoverConfigAttributes) InternalRef() (terra.Reference, error) {
	return fc.ref, nil
}

func (fc RoutingConfigFailoverConfigAttributes) InternalWithRef(ref terra.Reference) RoutingConfigFailoverConfigAttributes {
	return RoutingConfigFailoverConfigAttributes{ref: ref}
}

func (fc RoutingConfigFailoverConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fc.ref.InternalTokens()
}

func (fc RoutingConfigFailoverConfigAttributes) Primary() terra.ListValue[RoutingConfigFailoverConfigPrimaryAttributes] {
	return terra.ReferenceAsList[RoutingConfigFailoverConfigPrimaryAttributes](fc.ref.Append("primary"))
}

func (fc RoutingConfigFailoverConfigAttributes) Secondary() terra.ListValue[RoutingConfigFailoverConfigSecondaryAttributes] {
	return terra.ReferenceAsList[RoutingConfigFailoverConfigSecondaryAttributes](fc.ref.Append("secondary"))
}

type RoutingConfigFailoverConfigPrimaryAttributes struct {
	ref terra.Reference
}

func (p RoutingConfigFailoverConfigPrimaryAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p RoutingConfigFailoverConfigPrimaryAttributes) InternalWithRef(ref terra.Reference) RoutingConfigFailoverConfigPrimaryAttributes {
	return RoutingConfigFailoverConfigPrimaryAttributes{ref: ref}
}

func (p RoutingConfigFailoverConfigPrimaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p RoutingConfigFailoverConfigPrimaryAttributes) HealthCheck() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("health_check"))
}

type RoutingConfigFailoverConfigSecondaryAttributes struct {
	ref terra.Reference
}

func (s RoutingConfigFailoverConfigSecondaryAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s RoutingConfigFailoverConfigSecondaryAttributes) InternalWithRef(ref terra.Reference) RoutingConfigFailoverConfigSecondaryAttributes {
	return RoutingConfigFailoverConfigSecondaryAttributes{ref: ref}
}

func (s RoutingConfigFailoverConfigSecondaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s RoutingConfigFailoverConfigSecondaryAttributes) Route() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("route"))
}

type EventBusState struct {
	EventBusArn string `json:"event_bus_arn"`
}

type ReplicationConfigState struct {
	State string `json:"state"`
}

type RoutingConfigState struct {
	FailoverConfig []RoutingConfigFailoverConfigState `json:"failover_config"`
}

type RoutingConfigFailoverConfigState struct {
	Primary   []RoutingConfigFailoverConfigPrimaryState   `json:"primary"`
	Secondary []RoutingConfigFailoverConfigSecondaryState `json:"secondary"`
}

type RoutingConfigFailoverConfigPrimaryState struct {
	HealthCheck string `json:"health_check"`
}

type RoutingConfigFailoverConfigSecondaryState struct {
	Route string `json:"route"`
}
