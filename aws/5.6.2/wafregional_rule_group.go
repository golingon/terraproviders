// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafregionalrulegroup "github.com/golingon/terraproviders/aws/5.6.2/wafregionalrulegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafregionalRuleGroup creates a new instance of [WafregionalRuleGroup].
func NewWafregionalRuleGroup(name string, args WafregionalRuleGroupArgs) *WafregionalRuleGroup {
	return &WafregionalRuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafregionalRuleGroup)(nil)

// WafregionalRuleGroup represents the Terraform resource aws_wafregional_rule_group.
type WafregionalRuleGroup struct {
	Name      string
	Args      WafregionalRuleGroupArgs
	state     *wafregionalRuleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafregionalRuleGroup].
func (wrg *WafregionalRuleGroup) Type() string {
	return "aws_wafregional_rule_group"
}

// LocalName returns the local name for [WafregionalRuleGroup].
func (wrg *WafregionalRuleGroup) LocalName() string {
	return wrg.Name
}

// Configuration returns the configuration (args) for [WafregionalRuleGroup].
func (wrg *WafregionalRuleGroup) Configuration() interface{} {
	return wrg.Args
}

// DependOn is used for other resources to depend on [WafregionalRuleGroup].
func (wrg *WafregionalRuleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(wrg)
}

// Dependencies returns the list of resources [WafregionalRuleGroup] depends_on.
func (wrg *WafregionalRuleGroup) Dependencies() terra.Dependencies {
	return wrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafregionalRuleGroup].
func (wrg *WafregionalRuleGroup) LifecycleManagement() *terra.Lifecycle {
	return wrg.Lifecycle
}

// Attributes returns the attributes for [WafregionalRuleGroup].
func (wrg *WafregionalRuleGroup) Attributes() wafregionalRuleGroupAttributes {
	return wafregionalRuleGroupAttributes{ref: terra.ReferenceResource(wrg)}
}

// ImportState imports the given attribute values into [WafregionalRuleGroup]'s state.
func (wrg *WafregionalRuleGroup) ImportState(av io.Reader) error {
	wrg.state = &wafregionalRuleGroupState{}
	if err := json.NewDecoder(av).Decode(wrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrg.Type(), wrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafregionalRuleGroup] has state.
func (wrg *WafregionalRuleGroup) State() (*wafregionalRuleGroupState, bool) {
	return wrg.state, wrg.state != nil
}

// StateMust returns the state for [WafregionalRuleGroup]. Panics if the state is nil.
func (wrg *WafregionalRuleGroup) StateMust() *wafregionalRuleGroupState {
	if wrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrg.Type(), wrg.LocalName()))
	}
	return wrg.state
}

// WafregionalRuleGroupArgs contains the configurations for aws_wafregional_rule_group.
type WafregionalRuleGroupArgs struct {
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
	// ActivatedRule: min=0
	ActivatedRule []wafregionalrulegroup.ActivatedRule `hcl:"activated_rule,block" validate:"min=0"`
}
type wafregionalRuleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_wafregional_rule_group.
func (wrg wafregionalRuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("arn"))
}

// Id returns a reference to field id of aws_wafregional_rule_group.
func (wrg wafregionalRuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("id"))
}

// MetricName returns a reference to field metric_name of aws_wafregional_rule_group.
func (wrg wafregionalRuleGroupAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("metric_name"))
}

// Name returns a reference to field name of aws_wafregional_rule_group.
func (wrg wafregionalRuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_wafregional_rule_group.
func (wrg wafregionalRuleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_wafregional_rule_group.
func (wrg wafregionalRuleGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrg.ref.Append("tags_all"))
}

func (wrg wafregionalRuleGroupAttributes) ActivatedRule() terra.SetValue[wafregionalrulegroup.ActivatedRuleAttributes] {
	return terra.ReferenceAsSet[wafregionalrulegroup.ActivatedRuleAttributes](wrg.ref.Append("activated_rule"))
}

type wafregionalRuleGroupState struct {
	Arn           string                                    `json:"arn"`
	Id            string                                    `json:"id"`
	MetricName    string                                    `json:"metric_name"`
	Name          string                                    `json:"name"`
	Tags          map[string]string                         `json:"tags"`
	TagsAll       map[string]string                         `json:"tags_all"`
	ActivatedRule []wafregionalrulegroup.ActivatedRuleState `json:"activated_rule"`
}
