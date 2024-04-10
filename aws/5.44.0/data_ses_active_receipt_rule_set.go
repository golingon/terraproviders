// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataSesActiveReceiptRuleSet creates a new instance of [DataSesActiveReceiptRuleSet].
func NewDataSesActiveReceiptRuleSet(name string, args DataSesActiveReceiptRuleSetArgs) *DataSesActiveReceiptRuleSet {
	return &DataSesActiveReceiptRuleSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSesActiveReceiptRuleSet)(nil)

// DataSesActiveReceiptRuleSet represents the Terraform data resource aws_ses_active_receipt_rule_set.
type DataSesActiveReceiptRuleSet struct {
	Name string
	Args DataSesActiveReceiptRuleSetArgs
}

// DataSource returns the Terraform object type for [DataSesActiveReceiptRuleSet].
func (sarrs *DataSesActiveReceiptRuleSet) DataSource() string {
	return "aws_ses_active_receipt_rule_set"
}

// LocalName returns the local name for [DataSesActiveReceiptRuleSet].
func (sarrs *DataSesActiveReceiptRuleSet) LocalName() string {
	return sarrs.Name
}

// Configuration returns the configuration (args) for [DataSesActiveReceiptRuleSet].
func (sarrs *DataSesActiveReceiptRuleSet) Configuration() interface{} {
	return sarrs.Args
}

// Attributes returns the attributes for [DataSesActiveReceiptRuleSet].
func (sarrs *DataSesActiveReceiptRuleSet) Attributes() dataSesActiveReceiptRuleSetAttributes {
	return dataSesActiveReceiptRuleSetAttributes{ref: terra.ReferenceDataResource(sarrs)}
}

// DataSesActiveReceiptRuleSetArgs contains the configurations for aws_ses_active_receipt_rule_set.
type DataSesActiveReceiptRuleSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataSesActiveReceiptRuleSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ses_active_receipt_rule_set.
func (sarrs dataSesActiveReceiptRuleSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sarrs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ses_active_receipt_rule_set.
func (sarrs dataSesActiveReceiptRuleSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarrs.ref.Append("id"))
}

// RuleSetName returns a reference to field rule_set_name of aws_ses_active_receipt_rule_set.
func (sarrs dataSesActiveReceiptRuleSetAttributes) RuleSetName() terra.StringValue {
	return terra.ReferenceAsString(sarrs.ref.Append("rule_set_name"))
}
