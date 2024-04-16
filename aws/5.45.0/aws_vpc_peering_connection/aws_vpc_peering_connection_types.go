// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpc_peering_connection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Accepter struct {
	// AllowRemoteVpcDnsResolution: bool, optional
	AllowRemoteVpcDnsResolution terra.BoolValue `hcl:"allow_remote_vpc_dns_resolution,attr"`
}

type Requester struct {
	// AllowRemoteVpcDnsResolution: bool, optional
	AllowRemoteVpcDnsResolution terra.BoolValue `hcl:"allow_remote_vpc_dns_resolution,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AccepterAttributes struct {
	ref terra.Reference
}

func (a AccepterAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AccepterAttributes) InternalWithRef(ref terra.Reference) AccepterAttributes {
	return AccepterAttributes{ref: ref}
}

func (a AccepterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AccepterAttributes) AllowRemoteVpcDnsResolution() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("allow_remote_vpc_dns_resolution"))
}

type RequesterAttributes struct {
	ref terra.Reference
}

func (r RequesterAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequesterAttributes) InternalWithRef(ref terra.Reference) RequesterAttributes {
	return RequesterAttributes{ref: ref}
}

func (r RequesterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequesterAttributes) AllowRemoteVpcDnsResolution() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("allow_remote_vpc_dns_resolution"))
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

type AccepterState struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
}

type RequesterState struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
