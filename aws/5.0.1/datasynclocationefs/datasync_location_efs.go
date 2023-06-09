// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datasynclocationefs

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Ec2Config struct {
	// SecurityGroupArns: set of string, required
	SecurityGroupArns terra.SetValue[terra.StringValue] `hcl:"security_group_arns,attr" validate:"required"`
	// SubnetArn: string, required
	SubnetArn terra.StringValue `hcl:"subnet_arn,attr" validate:"required"`
}

type Ec2ConfigAttributes struct {
	ref terra.Reference
}

func (ec Ec2ConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec Ec2ConfigAttributes) InternalWithRef(ref terra.Reference) Ec2ConfigAttributes {
	return Ec2ConfigAttributes{ref: ref}
}

func (ec Ec2ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec Ec2ConfigAttributes) SecurityGroupArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ec.ref.Append("security_group_arns"))
}

func (ec Ec2ConfigAttributes) SubnetArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("subnet_arn"))
}

type Ec2ConfigState struct {
	SecurityGroupArns []string `json:"security_group_arns"`
	SubnetArn         string   `json:"subnet_arn"`
}
