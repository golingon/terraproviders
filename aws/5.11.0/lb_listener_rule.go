// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lblistenerrule "github.com/golingon/terraproviders/aws/5.11.0/lblistenerrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbListenerRule creates a new instance of [LbListenerRule].
func NewLbListenerRule(name string, args LbListenerRuleArgs) *LbListenerRule {
	return &LbListenerRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbListenerRule)(nil)

// LbListenerRule represents the Terraform resource aws_lb_listener_rule.
type LbListenerRule struct {
	Name      string
	Args      LbListenerRuleArgs
	state     *lbListenerRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbListenerRule].
func (llr *LbListenerRule) Type() string {
	return "aws_lb_listener_rule"
}

// LocalName returns the local name for [LbListenerRule].
func (llr *LbListenerRule) LocalName() string {
	return llr.Name
}

// Configuration returns the configuration (args) for [LbListenerRule].
func (llr *LbListenerRule) Configuration() interface{} {
	return llr.Args
}

// DependOn is used for other resources to depend on [LbListenerRule].
func (llr *LbListenerRule) DependOn() terra.Reference {
	return terra.ReferenceResource(llr)
}

// Dependencies returns the list of resources [LbListenerRule] depends_on.
func (llr *LbListenerRule) Dependencies() terra.Dependencies {
	return llr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbListenerRule].
func (llr *LbListenerRule) LifecycleManagement() *terra.Lifecycle {
	return llr.Lifecycle
}

// Attributes returns the attributes for [LbListenerRule].
func (llr *LbListenerRule) Attributes() lbListenerRuleAttributes {
	return lbListenerRuleAttributes{ref: terra.ReferenceResource(llr)}
}

// ImportState imports the given attribute values into [LbListenerRule]'s state.
func (llr *LbListenerRule) ImportState(av io.Reader) error {
	llr.state = &lbListenerRuleState{}
	if err := json.NewDecoder(av).Decode(llr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llr.Type(), llr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbListenerRule] has state.
func (llr *LbListenerRule) State() (*lbListenerRuleState, bool) {
	return llr.state, llr.state != nil
}

// StateMust returns the state for [LbListenerRule]. Panics if the state is nil.
func (llr *LbListenerRule) StateMust() *lbListenerRuleState {
	if llr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llr.Type(), llr.LocalName()))
	}
	return llr.state
}

// LbListenerRuleArgs contains the configurations for aws_lb_listener_rule.
type LbListenerRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListenerArn: string, required
	ListenerArn terra.StringValue `hcl:"listener_arn,attr" validate:"required"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Action: min=1
	Action []lblistenerrule.Action `hcl:"action,block" validate:"min=1"`
	// Condition: min=1
	Condition []lblistenerrule.Condition `hcl:"condition,block" validate:"min=1"`
}
type lbListenerRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lb_listener_rule.
func (llr lbListenerRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(llr.ref.Append("arn"))
}

// Id returns a reference to field id of aws_lb_listener_rule.
func (llr lbListenerRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(llr.ref.Append("id"))
}

// ListenerArn returns a reference to field listener_arn of aws_lb_listener_rule.
func (llr lbListenerRuleAttributes) ListenerArn() terra.StringValue {
	return terra.ReferenceAsString(llr.ref.Append("listener_arn"))
}

// Priority returns a reference to field priority of aws_lb_listener_rule.
func (llr lbListenerRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(llr.ref.Append("priority"))
}

// Tags returns a reference to field tags of aws_lb_listener_rule.
func (llr lbListenerRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](llr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lb_listener_rule.
func (llr lbListenerRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](llr.ref.Append("tags_all"))
}

func (llr lbListenerRuleAttributes) Action() terra.ListValue[lblistenerrule.ActionAttributes] {
	return terra.ReferenceAsList[lblistenerrule.ActionAttributes](llr.ref.Append("action"))
}

func (llr lbListenerRuleAttributes) Condition() terra.SetValue[lblistenerrule.ConditionAttributes] {
	return terra.ReferenceAsSet[lblistenerrule.ConditionAttributes](llr.ref.Append("condition"))
}

type lbListenerRuleState struct {
	Arn         string                          `json:"arn"`
	Id          string                          `json:"id"`
	ListenerArn string                          `json:"listener_arn"`
	Priority    float64                         `json:"priority"`
	Tags        map[string]string               `json:"tags"`
	TagsAll     map[string]string               `json:"tags_all"`
	Action      []lblistenerrule.ActionState    `json:"action"`
	Condition   []lblistenerrule.ConditionState `json:"condition"`
}
