// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalratebasedrule "github.com/golingon/terraproviders/aws/5.8.0/wafregionalratebasedrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalRateBasedRule creates a new instance of [WafregionalRateBasedRule].
func NewWafregionalRateBasedRule(name string, args WafregionalRateBasedRuleArgs) *WafregionalRateBasedRule {
	return &WafregionalRateBasedRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalRateBasedRule)(nil)

// WafregionalRateBasedRule represents the Terraform resource aws_wafregional_rate_based_rule.
type WafregionalRateBasedRule struct {
	Name      string
	Args      WafregionalRateBasedRuleArgs
	state     *wafregionalRateBasedRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalRateBasedRule].
func (wrbr *WafregionalRateBasedRule) Type() string {
	return "aws_wafregional_rate_based_rule"
}

// LocalName returns the local name for [WafregionalRateBasedRule].
func (wrbr *WafregionalRateBasedRule) LocalName() string {
	return wrbr.Name
}

// Configuration returns the configuration (args) for [WafregionalRateBasedRule].
func (wrbr *WafregionalRateBasedRule) Configuration() interface{} {
	return wrbr.Args
}

// DependOn is used for other resources to depend on [WafregionalRateBasedRule].
func (wrbr *WafregionalRateBasedRule) DependOn() terra.Reference {
	return terra.ReferenceResource(wrbr)
}

// Dependencies returns the list of resources [WafregionalRateBasedRule] depends_on.
func (wrbr *WafregionalRateBasedRule) Dependencies() terra.Dependencies {
	return wrbr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalRateBasedRule].
func (wrbr *WafregionalRateBasedRule) LifecycleManagement() *terra.Lifecycle {
	return wrbr.Lifecycle
}

// Attributes returns the attributes for [WafregionalRateBasedRule].
func (wrbr *WafregionalRateBasedRule) Attributes() wafregionalRateBasedRuleAttributes {
	return wafregionalRateBasedRuleAttributes{ref: terra.ReferenceResource(wrbr)}
}

// ImportState imports the given attribute values into [WafregionalRateBasedRule]'s state.
func (wrbr *WafregionalRateBasedRule) ImportState(av io.Reader) error {
	wrbr.state = &wafregionalRateBasedRuleState{}
	if err := json.NewDecoder(av).Decode(wrbr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrbr.Type(), wrbr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalRateBasedRule] has state.
func (wrbr *WafregionalRateBasedRule) State() (*wafregionalRateBasedRuleState, bool) {
	return wrbr.state, wrbr.state != nil
}

// StateMust returns the state for [WafregionalRateBasedRule]. Panics if the state is nil.
func (wrbr *WafregionalRateBasedRule) StateMust() *wafregionalRateBasedRuleState {
	if wrbr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrbr.Type(), wrbr.LocalName()))
	}
	return wrbr.state
}

// WafregionalRateBasedRuleArgs contains the configurations for aws_wafregional_rate_based_rule.
type WafregionalRateBasedRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MetricName: string, required
	MetricName terra.StringValue `hcl:"metric_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RateKey: string, required
	RateKey terra.StringValue `hcl:"rate_key,attr" validate:"required"`
	// RateLimit: number, required
	RateLimit terra.NumberValue `hcl:"rate_limit,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Predicate: min=0
	Predicate []wafregionalratebasedrule.Predicate `hcl:"predicate,block" validate:"min=0"`
}
type wafregionalRateBasedRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrbr.ref.Append("arn"))
}

// Id returns a reference to field id of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrbr.ref.Append("id"))
}

// MetricName returns a reference to field metric_name of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(wrbr.ref.Append("metric_name"))
}

// Name returns a reference to field name of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrbr.ref.Append("name"))
}

// RateKey returns a reference to field rate_key of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) RateKey() terra.StringValue {
	return terra.ReferenceAsString(wrbr.ref.Append("rate_key"))
}

// RateLimit returns a reference to field rate_limit of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) RateLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(wrbr.ref.Append("rate_limit"))
}

// Tags returns a reference to field tags of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrbr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafregional_rate_based_rule.
func (wrbr wafregionalRateBasedRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrbr.ref.Append("tags_all"))
}

func (wrbr wafregionalRateBasedRuleAttributes) Predicate() terra.SetValue[wafregionalratebasedrule.PredicateAttributes] {
	return terra.ReferenceAsSet[wafregionalratebasedrule.PredicateAttributes](wrbr.ref.Append("predicate"))
}

type wafregionalRateBasedRuleState struct {
	Arn        string                                    `json:"arn"`
	Id         string                                    `json:"id"`
	MetricName string                                    `json:"metric_name"`
	Name       string                                    `json:"name"`
	RateKey    string                                    `json:"rate_key"`
	RateLimit  float64                                   `json:"rate_limit"`
	Tags       map[string]string                         `json:"tags"`
	TagsAll    map[string]string                         `json:"tags_all"`
	Predicate  []wafregionalratebasedrule.PredicateState `json:"predicate"`
}
