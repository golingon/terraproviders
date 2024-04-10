// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package devicefarmtestgridproject

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type VpcConfig struct {
	// SecurityGroupIds: set of string, required
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr" validate:"required"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
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

func (vc VpcConfigAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("security_group_ids"))
}

func (vc VpcConfigAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("subnet_ids"))
}

func (vc VpcConfigAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vpc_id"))
}

type VpcConfigState struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
}
