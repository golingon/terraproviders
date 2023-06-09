// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpcpeeringconnectionoptions

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Accepter struct {
	// AllowRemoteVpcDnsResolution: bool, optional
	AllowRemoteVpcDnsResolution terra.BoolValue `hcl:"allow_remote_vpc_dns_resolution,attr"`
}

type Requester struct {
	// AllowRemoteVpcDnsResolution: bool, optional
	AllowRemoteVpcDnsResolution terra.BoolValue `hcl:"allow_remote_vpc_dns_resolution,attr"`
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

type AccepterState struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
}

type RequesterState struct {
	AllowRemoteVpcDnsResolution bool `json:"allow_remote_vpc_dns_resolution"`
}
