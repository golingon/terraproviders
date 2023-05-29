// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcsecuritygrouprule "github.com/golingon/terraproviders/aws/4.66.1/datavpcsecuritygrouprule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcSecurityGroupRule creates a new instance of [DataVpcSecurityGroupRule].
func NewDataVpcSecurityGroupRule(name string, args DataVpcSecurityGroupRuleArgs) *DataVpcSecurityGroupRule {
	return &DataVpcSecurityGroupRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcSecurityGroupRule)(nil)

// DataVpcSecurityGroupRule represents the Terraform data resource aws_vpc_security_group_rule.
type DataVpcSecurityGroupRule struct {
	Name string
	Args DataVpcSecurityGroupRuleArgs
}

// DataSource returns the Terraform object type for [DataVpcSecurityGroupRule].
func (vsgr *DataVpcSecurityGroupRule) DataSource() string {
	return "aws_vpc_security_group_rule"
}

// LocalName returns the local name for [DataVpcSecurityGroupRule].
func (vsgr *DataVpcSecurityGroupRule) LocalName() string {
	return vsgr.Name
}

// Configuration returns the configuration (args) for [DataVpcSecurityGroupRule].
func (vsgr *DataVpcSecurityGroupRule) Configuration() interface{} {
	return vsgr.Args
}

// Attributes returns the attributes for [DataVpcSecurityGroupRule].
func (vsgr *DataVpcSecurityGroupRule) Attributes() dataVpcSecurityGroupRuleAttributes {
	return dataVpcSecurityGroupRuleAttributes{ref: terra.ReferenceDataResource(vsgr)}
}

// DataVpcSecurityGroupRuleArgs contains the configurations for aws_vpc_security_group_rule.
type DataVpcSecurityGroupRuleArgs struct {
	// SecurityGroupRuleId: string, optional
	SecurityGroupRuleId terra.StringValue `hcl:"security_group_rule_id,attr"`
	// Filter: min=0
	Filter []datavpcsecuritygrouprule.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataVpcSecurityGroupRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("arn"))
}

// CidrIpv4 returns a reference to field cidr_ipv4 of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) CidrIpv4() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("cidr_ipv4"))
}

// CidrIpv6 returns a reference to field cidr_ipv6 of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) CidrIpv6() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("cidr_ipv6"))
}

// Description returns a reference to field description of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("description"))
}

// FromPort returns a reference to field from_port of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(vsgr.ref.Append("from_port"))
}

// Id returns a reference to field id of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("id"))
}

// IpProtocol returns a reference to field ip_protocol of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("ip_protocol"))
}

// IsEgress returns a reference to field is_egress of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) IsEgress() terra.BoolValue {
	return terra.ReferenceAsBool(vsgr.ref.Append("is_egress"))
}

// PrefixListId returns a reference to field prefix_list_id of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) PrefixListId() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("prefix_list_id"))
}

// ReferencedSecurityGroupId returns a reference to field referenced_security_group_id of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) ReferencedSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("referenced_security_group_id"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("security_group_id"))
}

// SecurityGroupRuleId returns a reference to field security_group_rule_id of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) SecurityGroupRuleId() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("security_group_rule_id"))
}

// Tags returns a reference to field tags of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsgr.ref.Append("tags"))
}

// ToPort returns a reference to field to_port of aws_vpc_security_group_rule.
func (vsgr dataVpcSecurityGroupRuleAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(vsgr.ref.Append("to_port"))
}

func (vsgr dataVpcSecurityGroupRuleAttributes) Filter() terra.SetValue[datavpcsecuritygrouprule.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcsecuritygrouprule.FilterAttributes](vsgr.ref.Append("filter"))
}