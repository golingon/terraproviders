// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package iottopicruledestination

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VpcConfiguration struct {
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
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

type VpcConfigurationAttributes struct {
	ref terra.Reference
}

func (vc VpcConfigurationAttributes) InternalRef() terra.Reference {
	return vc.ref
}

func (vc VpcConfigurationAttributes) InternalWithRef(ref terra.Reference) VpcConfigurationAttributes {
	return VpcConfigurationAttributes{ref: ref}
}

func (vc VpcConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return vc.ref.InternalTokens()
}

func (vc VpcConfigurationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("role_arn"))
}

func (vc VpcConfigurationAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("security_groups"))
}

func (vc VpcConfigurationAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](vc.ref.Append("subnet_ids"))
}

func (vc VpcConfigurationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vpc_id"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VpcConfigurationState struct {
	RoleArn        string   `json:"role_arn"`
	SecurityGroups []string `json:"security_groups"`
	SubnetIds      []string `json:"subnet_ids"`
	VpcId          string   `json:"vpc_id"`
}
