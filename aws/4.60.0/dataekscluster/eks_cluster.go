// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataekscluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CertificateAuthority struct{}

type Identity struct {
	// Oidc: min=0
	Oidc []Oidc `hcl:"oidc,block" validate:"min=0"`
}

type Oidc struct{}

type KubernetesNetworkConfig struct{}

type OutpostConfig struct {
	// ControlPlanePlacement: min=0
	ControlPlanePlacement []ControlPlanePlacement `hcl:"control_plane_placement,block" validate:"min=0"`
}

type ControlPlanePlacement struct{}

type VpcConfig struct{}

type CertificateAuthorityAttributes struct {
	ref terra.Reference
}

func (ca CertificateAuthorityAttributes) InternalRef() terra.Reference {
	return ca.ref
}

func (ca CertificateAuthorityAttributes) InternalWithRef(ref terra.Reference) CertificateAuthorityAttributes {
	return CertificateAuthorityAttributes{ref: ref}
}

func (ca CertificateAuthorityAttributes) InternalTokens() hclwrite.Tokens {
	return ca.ref.InternalTokens()
}

func (ca CertificateAuthorityAttributes) Data() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("data"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) Oidc() terra.ListValue[OidcAttributes] {
	return terra.ReferenceAsList[OidcAttributes](i.ref.Append("oidc"))
}

type OidcAttributes struct {
	ref terra.Reference
}

func (o OidcAttributes) InternalRef() terra.Reference {
	return o.ref
}

func (o OidcAttributes) InternalWithRef(ref terra.Reference) OidcAttributes {
	return OidcAttributes{ref: ref}
}

func (o OidcAttributes) InternalTokens() hclwrite.Tokens {
	return o.ref.InternalTokens()
}

func (o OidcAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("issuer"))
}

type KubernetesNetworkConfigAttributes struct {
	ref terra.Reference
}

func (knc KubernetesNetworkConfigAttributes) InternalRef() terra.Reference {
	return knc.ref
}

func (knc KubernetesNetworkConfigAttributes) InternalWithRef(ref terra.Reference) KubernetesNetworkConfigAttributes {
	return KubernetesNetworkConfigAttributes{ref: ref}
}

func (knc KubernetesNetworkConfigAttributes) InternalTokens() hclwrite.Tokens {
	return knc.ref.InternalTokens()
}

func (knc KubernetesNetworkConfigAttributes) IpFamily() terra.StringValue {
	return terra.ReferenceAsString(knc.ref.Append("ip_family"))
}

func (knc KubernetesNetworkConfigAttributes) ServiceIpv4Cidr() terra.StringValue {
	return terra.ReferenceAsString(knc.ref.Append("service_ipv4_cidr"))
}

func (knc KubernetesNetworkConfigAttributes) ServiceIpv6Cidr() terra.StringValue {
	return terra.ReferenceAsString(knc.ref.Append("service_ipv6_cidr"))
}

type OutpostConfigAttributes struct {
	ref terra.Reference
}

func (oc OutpostConfigAttributes) InternalRef() terra.Reference {
	return oc.ref
}

func (oc OutpostConfigAttributes) InternalWithRef(ref terra.Reference) OutpostConfigAttributes {
	return OutpostConfigAttributes{ref: ref}
}

func (oc OutpostConfigAttributes) InternalTokens() hclwrite.Tokens {
	return oc.ref.InternalTokens()
}

func (oc OutpostConfigAttributes) ControlPlaneInstanceType() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("control_plane_instance_type"))
}

func (oc OutpostConfigAttributes) OutpostArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](oc.ref.Append("outpost_arns"))
}

func (oc OutpostConfigAttributes) ControlPlanePlacement() terra.ListValue[ControlPlanePlacementAttributes] {
	return terra.ReferenceAsList[ControlPlanePlacementAttributes](oc.ref.Append("control_plane_placement"))
}

type ControlPlanePlacementAttributes struct {
	ref terra.Reference
}

func (cpp ControlPlanePlacementAttributes) InternalRef() terra.Reference {
	return cpp.ref
}

func (cpp ControlPlanePlacementAttributes) InternalWithRef(ref terra.Reference) ControlPlanePlacementAttributes {
	return ControlPlanePlacementAttributes{ref: ref}
}

func (cpp ControlPlanePlacementAttributes) InternalTokens() hclwrite.Tokens {
	return cpp.ref.InternalTokens()
}

func (cpp ControlPlanePlacementAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(cpp.ref.Append("group_name"))
}

type VpcConfigAttributes struct {
	ref terra.Reference
}

func (vc VpcConfigAttributes) InternalRef() terra.Reference {
	return vc.ref
}

func (vc VpcConfigAttributes) InternalWithRef(ref terra.Reference) VpcConfigAttributes {
	return VpcConfigAttributes{ref: ref}
}

func (vc VpcConfigAttributes) InternalTokens() hclwrite.Tokens {
	return vc.ref.InternalTokens()
}

func (vc VpcConfigAttributes) ClusterSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("cluster_security_group_id"))
}

func (vc VpcConfigAttributes) EndpointPrivateAccess() terra.BoolValue {
	return terra.ReferenceAsBool(vc.ref.Append("endpoint_private_access"))
}

func (vc VpcConfigAttributes) EndpointPublicAccess() terra.BoolValue {
	return terra.ReferenceAsBool(vc.ref.Append("endpoint_public_access"))
}

func (vc VpcConfigAttributes) PublicAccessCidrs() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("public_access_cidrs"))
}

func (vc VpcConfigAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("security_group_ids"))
}

func (vc VpcConfigAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("subnet_ids"))
}

func (vc VpcConfigAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vpc_id"))
}

type CertificateAuthorityState struct {
	Data string `json:"data"`
}

type IdentityState struct {
	Oidc []OidcState `json:"oidc"`
}

type OidcState struct {
	Issuer string `json:"issuer"`
}

type KubernetesNetworkConfigState struct {
	IpFamily        string `json:"ip_family"`
	ServiceIpv4Cidr string `json:"service_ipv4_cidr"`
	ServiceIpv6Cidr string `json:"service_ipv6_cidr"`
}

type OutpostConfigState struct {
	ControlPlaneInstanceType string                       `json:"control_plane_instance_type"`
	OutpostArns              []string                     `json:"outpost_arns"`
	ControlPlanePlacement    []ControlPlanePlacementState `json:"control_plane_placement"`
}

type ControlPlanePlacementState struct {
	GroupName string `json:"group_name"`
}

type VpcConfigState struct {
	ClusterSecurityGroupId string   `json:"cluster_security_group_id"`
	EndpointPrivateAccess  bool     `json:"endpoint_private_access"`
	EndpointPublicAccess   bool     `json:"endpoint_public_access"`
	PublicAccessCidrs      []string `json:"public_access_cidrs"`
	SecurityGroupIds       []string `json:"security_group_ids"`
	SubnetIds              []string `json:"subnet_ids"`
	VpcId                  string   `json:"vpc_id"`
}
