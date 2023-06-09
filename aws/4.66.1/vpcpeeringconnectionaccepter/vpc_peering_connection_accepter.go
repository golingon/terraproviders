// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpcpeeringconnectionaccepter

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Accepter struct {
	// AllowClassicLinkToRemoteVpc: bool, optional
	AllowClassicLinkToRemoteVpc terra.BoolValue `hcl:"allow_classic_link_to_remote_vpc,attr"`
	// AllowRemoteVpcDnsResolution: bool, optional
	AllowRemoteVpcDnsResolution terra.BoolValue `hcl:"allow_remote_vpc_dns_resolution,attr"`
	// AllowVpcToRemoteClassicLink: bool, optional
	AllowVpcToRemoteClassicLink terra.BoolValue `hcl:"allow_vpc_to_remote_classic_link,attr"`
}

type Requester struct {
	// AllowClassicLinkToRemoteVpc: bool, optional
	AllowClassicLinkToRemoteVpc terra.BoolValue `hcl:"allow_classic_link_to_remote_vpc,attr"`
	// AllowRemoteVpcDnsResolution: bool, optional
	AllowRemoteVpcDnsResolution terra.BoolValue `hcl:"allow_remote_vpc_dns_resolution,attr"`
	// AllowVpcToRemoteClassicLink: bool, optional
	AllowVpcToRemoteClassicLink terra.BoolValue `hcl:"allow_vpc_to_remote_classic_link,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
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

func (a AccepterAttributes) AllowClassicLinkToRemoteVpc() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("allow_classic_link_to_remote_vpc"))
}

func (a AccepterAttributes) AllowRemoteVpcDnsResolution() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("allow_remote_vpc_dns_resolution"))
}

func (a AccepterAttributes) AllowVpcToRemoteClassicLink() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("allow_vpc_to_remote_classic_link"))
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

func (r RequesterAttributes) AllowClassicLinkToRemoteVpc() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("allow_classic_link_to_remote_vpc"))
}

func (r RequesterAttributes) AllowRemoteVpcDnsResolution() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("allow_remote_vpc_dns_resolution"))
}

func (r RequesterAttributes) AllowVpcToRemoteClassicLink() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("allow_vpc_to_remote_classic_link"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AccepterState struct {
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type RequesterState struct {
	AllowClassicLinkToRemoteVpc bool `json:"allow_classic_link_to_remote_vpc"`
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
	AllowVpcToRemoteClassicLink bool `json:"allow_vpc_to_remote_classic_link"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
