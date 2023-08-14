// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataservicediscoveryservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DnsConfig struct {
	// DnsRecords: min=0
	DnsRecords []DnsRecords `hcl:"dns_records,block" validate:"min=0"`
}

type DnsRecords struct{}

type HealthCheckConfig struct{}

type HealthCheckCustomConfig struct{}

type DnsConfigAttributes struct {
	ref terra.Reference
}

func (dc DnsConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DnsConfigAttributes) InternalWithRef(ref terra.Reference) DnsConfigAttributes {
	return DnsConfigAttributes{ref: ref}
}

func (dc DnsConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DnsConfigAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("namespace_id"))
}

func (dc DnsConfigAttributes) RoutingPolicy() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("routing_policy"))
}

func (dc DnsConfigAttributes) DnsRecords() terra.ListValue[DnsRecordsAttributes] {
	return terra.ReferenceAsList[DnsRecordsAttributes](dc.ref.Append("dns_records"))
}

type DnsRecordsAttributes struct {
	ref terra.Reference
}

func (dr DnsRecordsAttributes) InternalRef() (terra.Reference, error) {
	return dr.ref, nil
}

func (dr DnsRecordsAttributes) InternalWithRef(ref terra.Reference) DnsRecordsAttributes {
	return DnsRecordsAttributes{ref: ref}
}

func (dr DnsRecordsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dr.ref.InternalTokens()
}

func (dr DnsRecordsAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dr.ref.Append("ttl"))
}

func (dr DnsRecordsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dr.ref.Append("type"))
}

type HealthCheckConfigAttributes struct {
	ref terra.Reference
}

func (hcc HealthCheckConfigAttributes) InternalRef() (terra.Reference, error) {
	return hcc.ref, nil
}

func (hcc HealthCheckConfigAttributes) InternalWithRef(ref terra.Reference) HealthCheckConfigAttributes {
	return HealthCheckConfigAttributes{ref: ref}
}

func (hcc HealthCheckConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hcc.ref.InternalTokens()
}

func (hcc HealthCheckConfigAttributes) FailureThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hcc.ref.Append("failure_threshold"))
}

func (hcc HealthCheckConfigAttributes) ResourcePath() terra.StringValue {
	return terra.ReferenceAsString(hcc.ref.Append("resource_path"))
}

func (hcc HealthCheckConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(hcc.ref.Append("type"))
}

type HealthCheckCustomConfigAttributes struct {
	ref terra.Reference
}

func (hccc HealthCheckCustomConfigAttributes) InternalRef() (terra.Reference, error) {
	return hccc.ref, nil
}

func (hccc HealthCheckCustomConfigAttributes) InternalWithRef(ref terra.Reference) HealthCheckCustomConfigAttributes {
	return HealthCheckCustomConfigAttributes{ref: ref}
}

func (hccc HealthCheckCustomConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hccc.ref.InternalTokens()
}

func (hccc HealthCheckCustomConfigAttributes) FailureThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hccc.ref.Append("failure_threshold"))
}

type DnsConfigState struct {
	NamespaceId   string            `json:"namespace_id"`
	RoutingPolicy string            `json:"routing_policy"`
	DnsRecords    []DnsRecordsState `json:"dns_records"`
}

type DnsRecordsState struct {
	Ttl  float64 `json:"ttl"`
	Type string  `json:"type"`
}

type HealthCheckConfigState struct {
	FailureThreshold float64 `json:"failure_threshold"`
	ResourcePath     string  `json:"resource_path"`
	Type             string  `json:"type"`
}

type HealthCheckCustomConfigState struct {
	FailureThreshold float64 `json:"failure_threshold"`
}
