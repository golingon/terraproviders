// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataWafv2RuleGroup creates a new instance of [DataWafv2RuleGroup].
func NewDataWafv2RuleGroup(name string, args DataWafv2RuleGroupArgs) *DataWafv2RuleGroup {
	return &DataWafv2RuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWafv2RuleGroup)(nil)

// DataWafv2RuleGroup represents the Terraform data resource aws_wafv2_rule_group.
type DataWafv2RuleGroup struct {
	Name string
	Args DataWafv2RuleGroupArgs
}

// DataSource returns the Terraform object type for [DataWafv2RuleGroup].
func (wrg *DataWafv2RuleGroup) DataSource() string {
	return "aws_wafv2_rule_group"
}

// LocalName returns the local name for [DataWafv2RuleGroup].
func (wrg *DataWafv2RuleGroup) LocalName() string {
	return wrg.Name
}

// Configuration returns the configuration (args) for [DataWafv2RuleGroup].
func (wrg *DataWafv2RuleGroup) Configuration() interface{} {
	return wrg.Args
}

// Attributes returns the attributes for [DataWafv2RuleGroup].
func (wrg *DataWafv2RuleGroup) Attributes() dataWafv2RuleGroupAttributes {
	return dataWafv2RuleGroupAttributes{ref: terra.ReferenceDataResource(wrg)}
}

// DataWafv2RuleGroupArgs contains the configurations for aws_wafv2_rule_group.
type DataWafv2RuleGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Scope: string, required
	Scope terra.StringValue `hcl:"scope,attr" validate:"required"`
}
type dataWafv2RuleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafv2_rule_group.
func (wrg dataWafv2RuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_wafv2_rule_group.
func (wrg dataWafv2RuleGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("description"))
}

// Id returns a reference to field id of aws_wafv2_rule_group.
func (wrg dataWafv2RuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafv2_rule_group.
func (wrg dataWafv2RuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("name"))
}

// Scope returns a reference to field scope of aws_wafv2_rule_group.
func (wrg dataWafv2RuleGroupAttributes) Scope() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("scope"))
}
