// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appstreamfleet

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ComputeCapacity struct {
	// DesiredInstances: number, required
	DesiredInstances terra.NumberValue `hcl:"desired_instances,attr" validate:"required"`
}

type DomainJoinInfo struct {
	// DirectoryName: string, optional
	DirectoryName terra.StringValue `hcl:"directory_name,attr"`
	// OrganizationalUnitDistinguishedName: string, optional
	OrganizationalUnitDistinguishedName terra.StringValue `hcl:"organizational_unit_distinguished_name,attr"`
}

type VpcConfig struct {
	// SecurityGroupIds: list of string, optional
	SecurityGroupIds terra.ListValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetIds: list of string, optional
	SubnetIds terra.ListValue[terra.StringValue] `hcl:"subnet_ids,attr"`
}

type ComputeCapacityAttributes struct {
	ref terra.Reference
}

func (cc ComputeCapacityAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc ComputeCapacityAttributes) InternalWithRef(ref terra.Reference) ComputeCapacityAttributes {
	return ComputeCapacityAttributes{ref: ref}
}

func (cc ComputeCapacityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc ComputeCapacityAttributes) Available() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("available"))
}

func (cc ComputeCapacityAttributes) DesiredInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("desired_instances"))
}

func (cc ComputeCapacityAttributes) InUse() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("in_use"))
}

func (cc ComputeCapacityAttributes) Running() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("running"))
}

type DomainJoinInfoAttributes struct {
	ref terra.Reference
}

func (dji DomainJoinInfoAttributes) InternalRef() (terra.Reference, error) {
	return dji.ref, nil
}

func (dji DomainJoinInfoAttributes) InternalWithRef(ref terra.Reference) DomainJoinInfoAttributes {
	return DomainJoinInfoAttributes{ref: ref}
}

func (dji DomainJoinInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dji.ref.InternalTokens()
}

func (dji DomainJoinInfoAttributes) DirectoryName() terra.StringValue {
	return terra.ReferenceAsString(dji.ref.Append("directory_name"))
}

func (dji DomainJoinInfoAttributes) OrganizationalUnitDistinguishedName() terra.StringValue {
	return terra.ReferenceAsString(dji.ref.Append("organizational_unit_distinguished_name"))
}

type VpcConfigAttributes struct {
	ref terra.Reference
}

func (vc VpcConfigAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc VpcConfigAttributes) InternalWithRef(ref terra.Reference) VpcConfigAttributes {
	return VpcConfigAttributes{ref: ref}
}

func (vc VpcConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc VpcConfigAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vc.ref.Append("security_group_ids"))
}

func (vc VpcConfigAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vc.ref.Append("subnet_ids"))
}

type ComputeCapacityState struct {
	Available        float64 `json:"available"`
	DesiredInstances float64 `json:"desired_instances"`
	InUse            float64 `json:"in_use"`
	Running          float64 `json:"running"`
}

type DomainJoinInfoState struct {
	DirectoryName                       string `json:"directory_name"`
	OrganizationalUnitDistinguishedName string `json:"organizational_unit_distinguished_name"`
}

type VpcConfigState struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
}
