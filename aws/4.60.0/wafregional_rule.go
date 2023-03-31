// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalrule "github.com/golingon/terraproviders/aws/4.60.0/wafregionalrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalRule creates a new instance of [WafregionalRule].
func NewWafregionalRule(name string, args WafregionalRuleArgs) *WafregionalRule {
	return &WafregionalRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalRule)(nil)

// WafregionalRule represents the Terraform resource aws_wafregional_rule.
type WafregionalRule struct {
	Name      string
	Args      WafregionalRuleArgs
	state     *wafregionalRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalRule].
func (wr *WafregionalRule) Type() string {
	return "aws_wafregional_rule"
}

// LocalName returns the local name for [WafregionalRule].
func (wr *WafregionalRule) LocalName() string {
	return wr.Name
}

// Configuration returns the configuration (args) for [WafregionalRule].
func (wr *WafregionalRule) Configuration() interface{} {
	return wr.Args
}

// DependOn is used for other resources to depend on [WafregionalRule].
func (wr *WafregionalRule) DependOn() terra.Reference {
	return terra.ReferenceResource(wr)
}

// Dependencies returns the list of resources [WafregionalRule] depends_on.
func (wr *WafregionalRule) Dependencies() terra.Dependencies {
	return wr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalRule].
func (wr *WafregionalRule) LifecycleManagement() *terra.Lifecycle {
	return wr.Lifecycle
}

// Attributes returns the attributes for [WafregionalRule].
func (wr *WafregionalRule) Attributes() wafregionalRuleAttributes {
	return wafregionalRuleAttributes{ref: terra.ReferenceResource(wr)}
}

// ImportState imports the given attribute values into [WafregionalRule]'s state.
func (wr *WafregionalRule) ImportState(av io.Reader) error {
	wr.state = &wafregionalRuleState{}
	if err := json.NewDecoder(av).Decode(wr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wr.Type(), wr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalRule] has state.
func (wr *WafregionalRule) State() (*wafregionalRuleState, bool) {
	return wr.state, wr.state != nil
}

// StateMust returns the state for [WafregionalRule]. Panics if the state is nil.
func (wr *WafregionalRule) StateMust() *wafregionalRuleState {
	if wr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wr.Type(), wr.LocalName()))
	}
	return wr.state
}

// WafregionalRuleArgs contains the configurations for aws_wafregional_rule.
type WafregionalRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Predicate: min=0
	Predicate []wafregionalrule.Predicate `hcl:"predicate,block" validate:"min=0"`
}
type wafregionalRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafregional_rule.
func (wr wafregionalRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("arn"))
}

// Id returns a reference to field id of aws_wafregional_rule.
func (wr wafregionalRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("id"))
}

// MetricName returns a reference to field metric_name of aws_wafregional_rule.
func (wr wafregionalRuleAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("metric_name"))
}

// Name returns a reference to field name of aws_wafregional_rule.
func (wr wafregionalRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wr.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_wafregional_rule.
func (wr wafregionalRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafregional_rule.
func (wr wafregionalRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wr.ref.Append("tags_all"))
}

func (wr wafregionalRuleAttributes) Predicate() terra.SetValue[wafregionalrule.PredicateAttributes] {
	return terra.ReferenceAsSet[wafregionalrule.PredicateAttributes](wr.ref.Append("predicate"))
}

type wafregionalRuleState struct {
	Arn        string                           `json:"arn"`
	Id         string                           `json:"id"`
	MetricName string                           `json:"metric_name"`
	Name       string                           `json:"name"`
	Tags       map[string]string                `json:"tags"`
	TagsAll    map[string]string                `json:"tags_all"`
	Predicate  []wafregionalrule.PredicateState `json:"predicate"`
}
