// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cecostcategory "github.com/golingon/terraproviders/aws/5.0.1/cecostcategory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCeCostCategory creates a new instance of [CeCostCategory].
func NewCeCostCategory(name string, args CeCostCategoryArgs) *CeCostCategory {
	return &CeCostCategory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CeCostCategory)(nil)

// CeCostCategory represents the Terraform resource aws_ce_cost_category.
type CeCostCategory struct {
	Name      string
	Args      CeCostCategoryArgs
	state     *ceCostCategoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CeCostCategory].
func (ccc *CeCostCategory) Type() string {
	return "aws_ce_cost_category"
}

// LocalName returns the local name for [CeCostCategory].
func (ccc *CeCostCategory) LocalName() string {
	return ccc.Name
}

// Configuration returns the configuration (args) for [CeCostCategory].
func (ccc *CeCostCategory) Configuration() interface{} {
	return ccc.Args
}

// DependOn is used for other resources to depend on [CeCostCategory].
func (ccc *CeCostCategory) DependOn() terra.Reference {
	return terra.ReferenceResource(ccc)
}

// Dependencies returns the list of resources [CeCostCategory] depends_on.
func (ccc *CeCostCategory) Dependencies() terra.Dependencies {
	return ccc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CeCostCategory].
func (ccc *CeCostCategory) LifecycleManagement() *terra.Lifecycle {
	return ccc.Lifecycle
}

// Attributes returns the attributes for [CeCostCategory].
func (ccc *CeCostCategory) Attributes() ceCostCategoryAttributes {
	return ceCostCategoryAttributes{ref: terra.ReferenceResource(ccc)}
}

// ImportState imports the given attribute values into [CeCostCategory]'s state.
func (ccc *CeCostCategory) ImportState(av io.Reader) error {
	ccc.state = &ceCostCategoryState{}
	if err := json.NewDecoder(av).Decode(ccc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccc.Type(), ccc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CeCostCategory] has state.
func (ccc *CeCostCategory) State() (*ceCostCategoryState, bool) {
	return ccc.state, ccc.state != nil
}

// StateMust returns the state for [CeCostCategory]. Panics if the state is nil.
func (ccc *CeCostCategory) StateMust() *ceCostCategoryState {
	if ccc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccc.Type(), ccc.LocalName()))
	}
	return ccc.state
}

// CeCostCategoryArgs contains the configurations for aws_ce_cost_category.
type CeCostCategoryArgs struct {
	// DefaultValue: string, optional
	DefaultValue terra.StringValue `hcl:"default_value,attr"`
	// EffectiveStart: string, optional
	EffectiveStart terra.StringValue `hcl:"effective_start,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RuleVersion: string, required
	RuleVersion terra.StringValue `hcl:"rule_version,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Rule: min=1
	Rule []cecostcategory.Rule `hcl:"rule,block" validate:"min=1"`
	// SplitChargeRule: min=0
	SplitChargeRule []cecostcategory.SplitChargeRule `hcl:"split_charge_rule,block" validate:"min=0"`
}
type ceCostCategoryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("arn"))
}

// DefaultValue returns a reference to field default_value of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) DefaultValue() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("default_value"))
}

// EffectiveEnd returns a reference to field effective_end of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) EffectiveEnd() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("effective_end"))
}

// EffectiveStart returns a reference to field effective_start of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) EffectiveStart() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("effective_start"))
}

// Id returns a reference to field id of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("id"))
}

// Name returns a reference to field name of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("name"))
}

// RuleVersion returns a reference to field rule_version of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) RuleVersion() terra.StringValue {
	return terra.ReferenceAsString(ccc.ref.Append("rule_version"))
}

// Tags returns a reference to field tags of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ce_cost_category.
func (ccc ceCostCategoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ccc.ref.Append("tags_all"))
}

func (ccc ceCostCategoryAttributes) Rule() terra.SetValue[cecostcategory.RuleAttributes] {
	return terra.ReferenceAsSet[cecostcategory.RuleAttributes](ccc.ref.Append("rule"))
}

func (ccc ceCostCategoryAttributes) SplitChargeRule() terra.SetValue[cecostcategory.SplitChargeRuleAttributes] {
	return terra.ReferenceAsSet[cecostcategory.SplitChargeRuleAttributes](ccc.ref.Append("split_charge_rule"))
}

type ceCostCategoryState struct {
	Arn             string                                `json:"arn"`
	DefaultValue    string                                `json:"default_value"`
	EffectiveEnd    string                                `json:"effective_end"`
	EffectiveStart  string                                `json:"effective_start"`
	Id              string                                `json:"id"`
	Name            string                                `json:"name"`
	RuleVersion     string                                `json:"rule_version"`
	Tags            map[string]string                     `json:"tags"`
	TagsAll         map[string]string                     `json:"tags_all"`
	Rule            []cecostcategory.RuleState            `json:"rule"`
	SplitChargeRule []cecostcategory.SplitChargeRuleState `json:"split_charge_rule"`
}
