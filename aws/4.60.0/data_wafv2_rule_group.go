// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataWafv2RuleGroup(name string, args DataWafv2RuleGroupArgs) *DataWafv2RuleGroup {
	return &DataWafv2RuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWafv2RuleGroup)(nil)

type DataWafv2RuleGroup struct {
	Name string
	Args DataWafv2RuleGroupArgs
}

func (wrg *DataWafv2RuleGroup) DataSource() string {
	return "aws_wafv2_rule_group"
}

func (wrg *DataWafv2RuleGroup) LocalName() string {
	return wrg.Name
}

func (wrg *DataWafv2RuleGroup) Configuration() interface{} {
	return wrg.Args
}

func (wrg *DataWafv2RuleGroup) Attributes() dataWafv2RuleGroupAttributes {
	return dataWafv2RuleGroupAttributes{ref: terra.ReferenceDataResource(wrg)}
}

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

func (wrg dataWafv2RuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(wrg.ref.Append("arn"))
}

func (wrg dataWafv2RuleGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(wrg.ref.Append("description"))
}

func (wrg dataWafv2RuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(wrg.ref.Append("id"))
}

func (wrg dataWafv2RuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(wrg.ref.Append("name"))
}

func (wrg dataWafv2RuleGroupAttributes) Scope() terra.StringValue {
	return terra.ReferenceString(wrg.ref.Append("scope"))
}
