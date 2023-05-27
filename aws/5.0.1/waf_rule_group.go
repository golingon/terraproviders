// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	wafrulegroup "github.com/golingon/terraproviders/aws/5.0.1/wafrulegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWafRuleGroup creates a new instance of [WafRuleGroup].
func NewWafRuleGroup(name string, args WafRuleGroupArgs) *WafRuleGroup {
	return &WafRuleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WafRuleGroup)(nil)

// WafRuleGroup represents the Terraform resource aws_waf_rule_group.
type WafRuleGroup struct {
	Name      string
	Args      WafRuleGroupArgs
	state     *wafRuleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WafRuleGroup].
func (wrg *WafRuleGroup) Type() string {
	return "aws_waf_rule_group"
}

// LocalName returns the local name for [WafRuleGroup].
func (wrg *WafRuleGroup) LocalName() string {
	return wrg.Name
}

// Configuration returns the configuration (args) for [WafRuleGroup].
func (wrg *WafRuleGroup) Configuration() interface{} {
	return wrg.Args
}

// DependOn is used for other resources to depend on [WafRuleGroup].
func (wrg *WafRuleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(wrg)
}

// Dependencies returns the list of resources [WafRuleGroup] depends_on.
func (wrg *WafRuleGroup) Dependencies() terra.Dependencies {
	return wrg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WafRuleGroup].
func (wrg *WafRuleGroup) LifecycleManagement() *terra.Lifecycle {
	return wrg.Lifecycle
}

// Attributes returns the attributes for [WafRuleGroup].
func (wrg *WafRuleGroup) Attributes() wafRuleGroupAttributes {
	return wafRuleGroupAttributes{ref: terra.ReferenceResource(wrg)}
}

// ImportState imports the given attribute values into [WafRuleGroup]'s state.
func (wrg *WafRuleGroup) ImportState(av io.Reader) error {
	wrg.state = &wafRuleGroupState{}
	if err := json.NewDecoder(av).Decode(wrg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wrg.Type(), wrg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WafRuleGroup] has state.
func (wrg *WafRuleGroup) State() (*wafRuleGroupState, bool) {
	return wrg.state, wrg.state != nil
}

// StateMust returns the state for [WafRuleGroup]. Panics if the state is nil.
func (wrg *WafRuleGroup) StateMust() *wafRuleGroupState {
	if wrg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wrg.Type(), wrg.LocalName()))
	}
	return wrg.state
}

// WafRuleGroupArgs contains the configurations for aws_waf_rule_group.
type WafRuleGroupArgs struct {
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
	ActivatedRule []wafrulegroup.ActivatedRule `hcl:"activated_rule,block" validate:"min=0"`
}
type wafRuleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_waf_rule_group.
func (wrg wafRuleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("arn"))
}

// Id returns a reference to field id of aws_waf_rule_group.
func (wrg wafRuleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("id"))
}

// MetricName returns a reference to field metric_name of aws_waf_rule_group.
func (wrg wafRuleGroupAttributes) MetricName() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("metric_name"))
}

// Name returns a reference to field name of aws_waf_rule_group.
func (wrg wafRuleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wrg.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_waf_rule_group.
func (wrg wafRuleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_waf_rule_group.
func (wrg wafRuleGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wrg.ref.Append("tags_all"))
}

func (wrg wafRuleGroupAttributes) ActivatedRule() terra.SetValue[wafrulegroup.ActivatedRuleAttributes] {
	return terra.ReferenceAsSet[wafrulegroup.ActivatedRuleAttributes](wrg.ref.Append("activated_rule"))
}

type wafRuleGroupState struct {
	Arn           string                            `json:"arn"`
	Id            string                            `json:"id"`
	MetricName    string                            `json:"metric_name"`
	Name          string                            `json:"name"`
	Tags          map[string]string                 `json:"tags"`
	TagsAll       map[string]string                 `json:"tags_all"`
	ActivatedRule []wafrulegroup.ActivatedRuleState `json:"activated_rule"`
}
