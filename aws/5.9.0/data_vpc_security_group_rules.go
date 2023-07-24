// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcsecuritygrouprules "github.com/golingon/terraproviders/aws/5.9.0/datavpcsecuritygrouprules"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcSecurityGroupRules creates a new instance of [DataVpcSecurityGroupRules].
func NewDataVpcSecurityGroupRules(name string, args DataVpcSecurityGroupRulesArgs) *DataVpcSecurityGroupRules {
	return &DataVpcSecurityGroupRules{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcSecurityGroupRules)(nil)

// DataVpcSecurityGroupRules represents the Terraform data resource aws_vpc_security_group_rules.
type DataVpcSecurityGroupRules struct {
	Name string
	Args DataVpcSecurityGroupRulesArgs
}

// DataSource returns the Terraform object type for [DataVpcSecurityGroupRules].
func (vsgr *DataVpcSecurityGroupRules) DataSource() string {
	return "aws_vpc_security_group_rules"
}

// LocalName returns the local name for [DataVpcSecurityGroupRules].
func (vsgr *DataVpcSecurityGroupRules) LocalName() string {
	return vsgr.Name
}

// Configuration returns the configuration (args) for [DataVpcSecurityGroupRules].
func (vsgr *DataVpcSecurityGroupRules) Configuration() interface{} {
	return vsgr.Args
}

// Attributes returns the attributes for [DataVpcSecurityGroupRules].
func (vsgr *DataVpcSecurityGroupRules) Attributes() dataVpcSecurityGroupRulesAttributes {
	return dataVpcSecurityGroupRulesAttributes{ref: terra.ReferenceDataResource(vsgr)}
}

// DataVpcSecurityGroupRulesArgs contains the configurations for aws_vpc_security_group_rules.
type DataVpcSecurityGroupRulesArgs struct {
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpcsecuritygrouprules.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataVpcSecurityGroupRulesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_security_group_rules.
func (vsgr dataVpcSecurityGroupRulesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vsgr.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_vpc_security_group_rules.
func (vsgr dataVpcSecurityGroupRulesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vsgr.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_vpc_security_group_rules.
func (vsgr dataVpcSecurityGroupRulesAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsgr.ref.Append("tags"))
}

func (vsgr dataVpcSecurityGroupRulesAttributes) Filter() terra.SetValue[datavpcsecuritygrouprules.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcsecuritygrouprules.FilterAttributes](vsgr.ref.Append("filter"))
}
