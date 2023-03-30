// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacecostcategory "github.com/golingon/terraproviders/aws/4.60.0/datacecostcategory"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataCeCostCategory(name string, args DataCeCostCategoryArgs) *DataCeCostCategory {
	return &DataCeCostCategory{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCeCostCategory)(nil)

type DataCeCostCategory struct {
	Name string
	Args DataCeCostCategoryArgs
}

func (ccc *DataCeCostCategory) DataSource() string {
	return "aws_ce_cost_category"
}

func (ccc *DataCeCostCategory) LocalName() string {
	return ccc.Name
}

func (ccc *DataCeCostCategory) Configuration() interface{} {
	return ccc.Args
}

func (ccc *DataCeCostCategory) Attributes() dataCeCostCategoryAttributes {
	return dataCeCostCategoryAttributes{ref: terra.ReferenceDataResource(ccc)}
}

type DataCeCostCategoryArgs struct {
	// CostCategoryArn: string, required
	CostCategoryArn terra.StringValue `hcl:"cost_category_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Rule: min=0
	Rule []datacecostcategory.Rule `hcl:"rule,block" validate:"min=0"`
	// SplitChargeRule: min=0
	SplitChargeRule []datacecostcategory.SplitChargeRule `hcl:"split_charge_rule,block" validate:"min=0"`
}
type dataCeCostCategoryAttributes struct {
	ref terra.Reference
}

func (ccc dataCeCostCategoryAttributes) CostCategoryArn() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("cost_category_arn"))
}

func (ccc dataCeCostCategoryAttributes) DefaultValue() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("default_value"))
}

func (ccc dataCeCostCategoryAttributes) EffectiveEnd() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("effective_end"))
}

func (ccc dataCeCostCategoryAttributes) EffectiveStart() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("effective_start"))
}

func (ccc dataCeCostCategoryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("id"))
}

func (ccc dataCeCostCategoryAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("name"))
}

func (ccc dataCeCostCategoryAttributes) RuleVersion() terra.StringValue {
	return terra.ReferenceString(ccc.ref.Append("rule_version"))
}

func (ccc dataCeCostCategoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ccc.ref.Append("tags"))
}

func (ccc dataCeCostCategoryAttributes) Rule() terra.SetValue[datacecostcategory.RuleAttributes] {
	return terra.ReferenceSet[datacecostcategory.RuleAttributes](ccc.ref.Append("rule"))
}

func (ccc dataCeCostCategoryAttributes) SplitChargeRule() terra.SetValue[datacecostcategory.SplitChargeRuleAttributes] {
	return terra.ReferenceSet[datacecostcategory.SplitChargeRuleAttributes](ccc.ref.Append("split_charge_rule"))
}
