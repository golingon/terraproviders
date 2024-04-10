// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vpcendpoint

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DnsEntry struct{}

type DnsOptions struct {
	// DnsRecordIpType: string, optional
	DnsRecordIpType terra.StringValue `hcl:"dns_record_ip_type,attr"`
	// PrivateDnsOnlyForInboundResolverEndpoint: bool, optional
	PrivateDnsOnlyForInboundResolverEndpoint terra.BoolValue `hcl:"private_dns_only_for_inbound_resolver_endpoint,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DnsEntryAttributes struct {
	ref terra.Reference
}

func (de DnsEntryAttributes) InternalRef() (terra.Reference, error) {
	return de.ref, nil
}

func (de DnsEntryAttributes) InternalWithRef(ref terra.Reference) DnsEntryAttributes {
	return DnsEntryAttributes{ref: ref}
}

func (de DnsEntryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return de.ref.InternalTokens()
}

func (de DnsEntryAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("dns_name"))
}

func (de DnsEntryAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("hosted_zone_id"))
}

type DnsOptionsAttributes struct {
	ref terra.Reference
}

func (do DnsOptionsAttributes) InternalRef() (terra.Reference, error) {
	return do.ref, nil
}

func (do DnsOptionsAttributes) InternalWithRef(ref terra.Reference) DnsOptionsAttributes {
	return DnsOptionsAttributes{ref: ref}
}

func (do DnsOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return do.ref.InternalTokens()
}

func (do DnsOptionsAttributes) DnsRecordIpType() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("dns_record_ip_type"))
}

func (do DnsOptionsAttributes) PrivateDnsOnlyForInboundResolverEndpoint() terra.BoolValue {
	return terra.ReferenceAsBool(do.ref.Append("private_dns_only_for_inbound_resolver_endpoint"))
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

type DnsEntryState struct {
	DnsName      string `json:"dns_name"`
	HostedZoneId string `json:"hosted_zone_id"`
}

type DnsOptionsState struct {
	DnsRecordIpType                          string `json:"dns_record_ip_type"`
	PrivateDnsOnlyForInboundResolverEndpoint bool   `json:"private_dns_only_for_inbound_resolver_endpoint"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
