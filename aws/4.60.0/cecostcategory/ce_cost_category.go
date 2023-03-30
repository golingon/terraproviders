// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cecostcategory

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Rule struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
	// InheritedValue: optional
	InheritedValue *InheritedValue `hcl:"inherited_value,block"`
	// RuleRule: optional
	Rule *RuleRule `hcl:"rule,block"`
}

type InheritedValue struct {
	// DimensionKey: string, optional
	DimensionKey terra.StringValue `hcl:"dimension_key,attr"`
	// DimensionName: string, optional
	DimensionName terra.StringValue `hcl:"dimension_name,attr"`
}

type RuleRule struct {
	// And: min=0
	And []And `hcl:"and,block" validate:"min=0"`
	// RuleCostCategory: optional
	CostCategory *RuleCostCategory `hcl:"cost_category,block"`
	// RuleDimension: optional
	Dimension *RuleDimension `hcl:"dimension,block"`
	// Not: optional
	Not *Not `hcl:"not,block"`
	// Or: min=0
	Or []Or `hcl:"or,block" validate:"min=0"`
	// RuleTags: optional
	Tags *RuleTags `hcl:"tags,block"`
}

type And struct {
	// AndCostCategory: optional
	CostCategory *AndCostCategory `hcl:"cost_category,block"`
	// AndDimension: optional
	Dimension *AndDimension `hcl:"dimension,block"`
	// AndTags: optional
	Tags *AndTags `hcl:"tags,block"`
}

type AndCostCategory struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type AndDimension struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type AndTags struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type RuleCostCategory struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type RuleDimension struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type Not struct {
	// NotCostCategory: optional
	CostCategory *NotCostCategory `hcl:"cost_category,block"`
	// NotDimension: optional
	Dimension *NotDimension `hcl:"dimension,block"`
	// NotTags: optional
	Tags *NotTags `hcl:"tags,block"`
}

type NotCostCategory struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type NotDimension struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type NotTags struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type Or struct {
	// OrCostCategory: optional
	CostCategory *OrCostCategory `hcl:"cost_category,block"`
	// OrDimension: optional
	Dimension *OrDimension `hcl:"dimension,block"`
	// OrTags: optional
	Tags *OrTags `hcl:"tags,block"`
}

type OrCostCategory struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type OrDimension struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type OrTags struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type RuleTags struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// MatchOptions: set of string, optional
	MatchOptions terra.SetValue[terra.StringValue] `hcl:"match_options,attr"`
	// Values: set of string, optional
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr"`
}

type SplitChargeRule struct {
	// Method: string, required
	Method terra.StringValue `hcl:"method,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Targets: set of string, required
	Targets terra.SetValue[terra.StringValue] `hcl:"targets,attr" validate:"required"`
	// Parameter: min=0
	Parameter []Parameter `hcl:"parameter,block" validate:"min=0"`
}

type Parameter struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Values: list of string, optional
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr"`
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) Type() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("type"))
}

func (r RuleAttributes) Value() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("value"))
}

func (r RuleAttributes) InheritedValue() terra.ListValue[InheritedValueAttributes] {
	return terra.ReferenceList[InheritedValueAttributes](r.ref.Append("inherited_value"))
}

func (r RuleAttributes) Rule() terra.ListValue[RuleRuleAttributes] {
	return terra.ReferenceList[RuleRuleAttributes](r.ref.Append("rule"))
}

type InheritedValueAttributes struct {
	ref terra.Reference
}

func (iv InheritedValueAttributes) InternalRef() terra.Reference {
	return iv.ref
}

func (iv InheritedValueAttributes) InternalWithRef(ref terra.Reference) InheritedValueAttributes {
	return InheritedValueAttributes{ref: ref}
}

func (iv InheritedValueAttributes) InternalTokens() hclwrite.Tokens {
	return iv.ref.InternalTokens()
}

func (iv InheritedValueAttributes) DimensionKey() terra.StringValue {
	return terra.ReferenceString(iv.ref.Append("dimension_key"))
}

func (iv InheritedValueAttributes) DimensionName() terra.StringValue {
	return terra.ReferenceString(iv.ref.Append("dimension_name"))
}

type RuleRuleAttributes struct {
	ref terra.Reference
}

func (r RuleRuleAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RuleRuleAttributes) InternalWithRef(ref terra.Reference) RuleRuleAttributes {
	return RuleRuleAttributes{ref: ref}
}

func (r RuleRuleAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RuleRuleAttributes) And() terra.SetValue[AndAttributes] {
	return terra.ReferenceSet[AndAttributes](r.ref.Append("and"))
}

func (r RuleRuleAttributes) CostCategory() terra.ListValue[RuleCostCategoryAttributes] {
	return terra.ReferenceList[RuleCostCategoryAttributes](r.ref.Append("cost_category"))
}

func (r RuleRuleAttributes) Dimension() terra.ListValue[RuleDimensionAttributes] {
	return terra.ReferenceList[RuleDimensionAttributes](r.ref.Append("dimension"))
}

func (r RuleRuleAttributes) Not() terra.ListValue[NotAttributes] {
	return terra.ReferenceList[NotAttributes](r.ref.Append("not"))
}

func (r RuleRuleAttributes) Or() terra.SetValue[OrAttributes] {
	return terra.ReferenceSet[OrAttributes](r.ref.Append("or"))
}

func (r RuleRuleAttributes) Tags() terra.ListValue[RuleTagsAttributes] {
	return terra.ReferenceList[RuleTagsAttributes](r.ref.Append("tags"))
}

type AndAttributes struct {
	ref terra.Reference
}

func (a AndAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AndAttributes) InternalWithRef(ref terra.Reference) AndAttributes {
	return AndAttributes{ref: ref}
}

func (a AndAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AndAttributes) CostCategory() terra.ListValue[AndCostCategoryAttributes] {
	return terra.ReferenceList[AndCostCategoryAttributes](a.ref.Append("cost_category"))
}

func (a AndAttributes) Dimension() terra.ListValue[AndDimensionAttributes] {
	return terra.ReferenceList[AndDimensionAttributes](a.ref.Append("dimension"))
}

func (a AndAttributes) Tags() terra.ListValue[AndTagsAttributes] {
	return terra.ReferenceList[AndTagsAttributes](a.ref.Append("tags"))
}

type AndCostCategoryAttributes struct {
	ref terra.Reference
}

func (cc AndCostCategoryAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc AndCostCategoryAttributes) InternalWithRef(ref terra.Reference) AndCostCategoryAttributes {
	return AndCostCategoryAttributes{ref: ref}
}

func (cc AndCostCategoryAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc AndCostCategoryAttributes) Key() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("key"))
}

func (cc AndCostCategoryAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("match_options"))
}

func (cc AndCostCategoryAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("values"))
}

type AndDimensionAttributes struct {
	ref terra.Reference
}

func (d AndDimensionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d AndDimensionAttributes) InternalWithRef(ref terra.Reference) AndDimensionAttributes {
	return AndDimensionAttributes{ref: ref}
}

func (d AndDimensionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d AndDimensionAttributes) Key() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("key"))
}

func (d AndDimensionAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("match_options"))
}

func (d AndDimensionAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("values"))
}

type AndTagsAttributes struct {
	ref terra.Reference
}

func (t AndTagsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t AndTagsAttributes) InternalWithRef(ref terra.Reference) AndTagsAttributes {
	return AndTagsAttributes{ref: ref}
}

func (t AndTagsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t AndTagsAttributes) Key() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("key"))
}

func (t AndTagsAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("match_options"))
}

func (t AndTagsAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("values"))
}

type RuleCostCategoryAttributes struct {
	ref terra.Reference
}

func (cc RuleCostCategoryAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc RuleCostCategoryAttributes) InternalWithRef(ref terra.Reference) RuleCostCategoryAttributes {
	return RuleCostCategoryAttributes{ref: ref}
}

func (cc RuleCostCategoryAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc RuleCostCategoryAttributes) Key() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("key"))
}

func (cc RuleCostCategoryAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("match_options"))
}

func (cc RuleCostCategoryAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("values"))
}

type RuleDimensionAttributes struct {
	ref terra.Reference
}

func (d RuleDimensionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d RuleDimensionAttributes) InternalWithRef(ref terra.Reference) RuleDimensionAttributes {
	return RuleDimensionAttributes{ref: ref}
}

func (d RuleDimensionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d RuleDimensionAttributes) Key() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("key"))
}

func (d RuleDimensionAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("match_options"))
}

func (d RuleDimensionAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("values"))
}

type NotAttributes struct {
	ref terra.Reference
}

func (n NotAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NotAttributes) InternalWithRef(ref terra.Reference) NotAttributes {
	return NotAttributes{ref: ref}
}

func (n NotAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NotAttributes) CostCategory() terra.ListValue[NotCostCategoryAttributes] {
	return terra.ReferenceList[NotCostCategoryAttributes](n.ref.Append("cost_category"))
}

func (n NotAttributes) Dimension() terra.ListValue[NotDimensionAttributes] {
	return terra.ReferenceList[NotDimensionAttributes](n.ref.Append("dimension"))
}

func (n NotAttributes) Tags() terra.ListValue[NotTagsAttributes] {
	return terra.ReferenceList[NotTagsAttributes](n.ref.Append("tags"))
}

type NotCostCategoryAttributes struct {
	ref terra.Reference
}

func (cc NotCostCategoryAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc NotCostCategoryAttributes) InternalWithRef(ref terra.Reference) NotCostCategoryAttributes {
	return NotCostCategoryAttributes{ref: ref}
}

func (cc NotCostCategoryAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc NotCostCategoryAttributes) Key() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("key"))
}

func (cc NotCostCategoryAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("match_options"))
}

func (cc NotCostCategoryAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("values"))
}

type NotDimensionAttributes struct {
	ref terra.Reference
}

func (d NotDimensionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d NotDimensionAttributes) InternalWithRef(ref terra.Reference) NotDimensionAttributes {
	return NotDimensionAttributes{ref: ref}
}

func (d NotDimensionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d NotDimensionAttributes) Key() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("key"))
}

func (d NotDimensionAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("match_options"))
}

func (d NotDimensionAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("values"))
}

type NotTagsAttributes struct {
	ref terra.Reference
}

func (t NotTagsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t NotTagsAttributes) InternalWithRef(ref terra.Reference) NotTagsAttributes {
	return NotTagsAttributes{ref: ref}
}

func (t NotTagsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t NotTagsAttributes) Key() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("key"))
}

func (t NotTagsAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("match_options"))
}

func (t NotTagsAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("values"))
}

type OrAttributes struct {
	ref terra.Reference
}

func (o OrAttributes) InternalRef() terra.Reference {
	return o.ref
}

func (o OrAttributes) InternalWithRef(ref terra.Reference) OrAttributes {
	return OrAttributes{ref: ref}
}

func (o OrAttributes) InternalTokens() hclwrite.Tokens {
	return o.ref.InternalTokens()
}

func (o OrAttributes) CostCategory() terra.ListValue[OrCostCategoryAttributes] {
	return terra.ReferenceList[OrCostCategoryAttributes](o.ref.Append("cost_category"))
}

func (o OrAttributes) Dimension() terra.ListValue[OrDimensionAttributes] {
	return terra.ReferenceList[OrDimensionAttributes](o.ref.Append("dimension"))
}

func (o OrAttributes) Tags() terra.ListValue[OrTagsAttributes] {
	return terra.ReferenceList[OrTagsAttributes](o.ref.Append("tags"))
}

type OrCostCategoryAttributes struct {
	ref terra.Reference
}

func (cc OrCostCategoryAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc OrCostCategoryAttributes) InternalWithRef(ref terra.Reference) OrCostCategoryAttributes {
	return OrCostCategoryAttributes{ref: ref}
}

func (cc OrCostCategoryAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc OrCostCategoryAttributes) Key() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("key"))
}

func (cc OrCostCategoryAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("match_options"))
}

func (cc OrCostCategoryAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](cc.ref.Append("values"))
}

type OrDimensionAttributes struct {
	ref terra.Reference
}

func (d OrDimensionAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d OrDimensionAttributes) InternalWithRef(ref terra.Reference) OrDimensionAttributes {
	return OrDimensionAttributes{ref: ref}
}

func (d OrDimensionAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d OrDimensionAttributes) Key() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("key"))
}

func (d OrDimensionAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("match_options"))
}

func (d OrDimensionAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](d.ref.Append("values"))
}

type OrTagsAttributes struct {
	ref terra.Reference
}

func (t OrTagsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t OrTagsAttributes) InternalWithRef(ref terra.Reference) OrTagsAttributes {
	return OrTagsAttributes{ref: ref}
}

func (t OrTagsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t OrTagsAttributes) Key() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("key"))
}

func (t OrTagsAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("match_options"))
}

func (t OrTagsAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("values"))
}

type RuleTagsAttributes struct {
	ref terra.Reference
}

func (t RuleTagsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t RuleTagsAttributes) InternalWithRef(ref terra.Reference) RuleTagsAttributes {
	return RuleTagsAttributes{ref: ref}
}

func (t RuleTagsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t RuleTagsAttributes) Key() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("key"))
}

func (t RuleTagsAttributes) MatchOptions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("match_options"))
}

func (t RuleTagsAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](t.ref.Append("values"))
}

type SplitChargeRuleAttributes struct {
	ref terra.Reference
}

func (scr SplitChargeRuleAttributes) InternalRef() terra.Reference {
	return scr.ref
}

func (scr SplitChargeRuleAttributes) InternalWithRef(ref terra.Reference) SplitChargeRuleAttributes {
	return SplitChargeRuleAttributes{ref: ref}
}

func (scr SplitChargeRuleAttributes) InternalTokens() hclwrite.Tokens {
	return scr.ref.InternalTokens()
}

func (scr SplitChargeRuleAttributes) Method() terra.StringValue {
	return terra.ReferenceString(scr.ref.Append("method"))
}

func (scr SplitChargeRuleAttributes) Source() terra.StringValue {
	return terra.ReferenceString(scr.ref.Append("source"))
}

func (scr SplitChargeRuleAttributes) Targets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](scr.ref.Append("targets"))
}

func (scr SplitChargeRuleAttributes) Parameter() terra.SetValue[ParameterAttributes] {
	return terra.ReferenceSet[ParameterAttributes](scr.ref.Append("parameter"))
}

type ParameterAttributes struct {
	ref terra.Reference
}

func (p ParameterAttributes) InternalRef() terra.Reference {
	return p.ref
}

func (p ParameterAttributes) InternalWithRef(ref terra.Reference) ParameterAttributes {
	return ParameterAttributes{ref: ref}
}

func (p ParameterAttributes) InternalTokens() hclwrite.Tokens {
	return p.ref.InternalTokens()
}

func (p ParameterAttributes) Type() terra.StringValue {
	return terra.ReferenceString(p.ref.Append("type"))
}

func (p ParameterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](p.ref.Append("values"))
}

type RuleState struct {
	Type           string                `json:"type"`
	Value          string                `json:"value"`
	InheritedValue []InheritedValueState `json:"inherited_value"`
	Rule           []RuleRuleState       `json:"rule"`
}

type InheritedValueState struct {
	DimensionKey  string `json:"dimension_key"`
	DimensionName string `json:"dimension_name"`
}

type RuleRuleState struct {
	And          []AndState              `json:"and"`
	CostCategory []RuleCostCategoryState `json:"cost_category"`
	Dimension    []RuleDimensionState    `json:"dimension"`
	Not          []NotState              `json:"not"`
	Or           []OrState               `json:"or"`
	Tags         []RuleTagsState         `json:"tags"`
}

type AndState struct {
	CostCategory []AndCostCategoryState `json:"cost_category"`
	Dimension    []AndDimensionState    `json:"dimension"`
	Tags         []AndTagsState         `json:"tags"`
}

type AndCostCategoryState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type AndDimensionState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type AndTagsState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type RuleCostCategoryState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type RuleDimensionState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type NotState struct {
	CostCategory []NotCostCategoryState `json:"cost_category"`
	Dimension    []NotDimensionState    `json:"dimension"`
	Tags         []NotTagsState         `json:"tags"`
}

type NotCostCategoryState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type NotDimensionState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type NotTagsState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type OrState struct {
	CostCategory []OrCostCategoryState `json:"cost_category"`
	Dimension    []OrDimensionState    `json:"dimension"`
	Tags         []OrTagsState         `json:"tags"`
}

type OrCostCategoryState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type OrDimensionState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type OrTagsState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type RuleTagsState struct {
	Key          string   `json:"key"`
	MatchOptions []string `json:"match_options"`
	Values       []string `json:"values"`
}

type SplitChargeRuleState struct {
	Method    string           `json:"method"`
	Source    string           `json:"source"`
	Targets   []string         `json:"targets"`
	Parameter []ParameterState `json:"parameter"`
}

type ParameterState struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}
