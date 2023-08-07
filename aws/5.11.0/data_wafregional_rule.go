// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataWafregionalRule creates a new instance of [DataWafregionalRule].
func NewDataWafregionalRule(name string, args DataWafregionalRuleArgs) *DataWafregionalRule {
	return &DataWafregionalRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWafregionalRule)(nil)

// DataWafregionalRule represents the Terraform data resource aws_wafregional_rule.
type DataWafregionalRule struct {
	Name string
	Args DataWafregionalRuleArgs
}

// DataSource returns the Terraform object type for [DataWafregionalRule].
func (wr *DataWafregionalRule) DataSource() string {
	return "aws_wafregional_rule"
}

// LocalName returns the local name for [DataWafregionalRule].
func (wr *DataWafregionalRule) LocalName() string {
	return wr.Name
}

// Configuration returns the configuration (args) for [DataWafregionalRule].
func (wr *DataWafregionalRule) Configuration() interface{} {
	return wr.Args
}

// Attributes returns the attributes for [DataWafregionalRule].
func (wr *DataWafregionalRule) Attributes() dataWafregionalRuleAttributes {
	return dataWafregionalRuleAttributes{ref: terra.ReferenceDataResource(wr)}
}

// DataWafregionalRuleArgs contains the configurations for aws_wafregional_rule.
type DataWafregionalRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataWafregionalRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_wafregional_rule.
func (wr dataWafregionalRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("id"))
}

// Name returns a reference to field name of aws_wafregional_rule.
func (wr dataWafregionalRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("name"))
}
