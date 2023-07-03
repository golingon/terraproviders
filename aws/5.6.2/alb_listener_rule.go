// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	alblistenerrule "github.com/golingon/terraproviders/aws/5.6.2/alblistenerrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlbListenerRule creates a new instance of [AlbListenerRule].
func NewAlbListenerRule(name string, args AlbListenerRuleArgs) *AlbListenerRule {
	return &AlbListenerRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlbListenerRule)(nil)

// AlbListenerRule represents the Terraform resource aws_alb_listener_rule.
type AlbListenerRule struct {
	Name      string
	Args      AlbListenerRuleArgs
	state     *albListenerRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AlbListenerRule].
func (alr *AlbListenerRule) Type() string {
	return "aws_alb_listener_rule"
}

// LocalName returns the local name for [AlbListenerRule].
func (alr *AlbListenerRule) LocalName() string {
	return alr.Name
}

// Configuration returns the configuration (args) for [AlbListenerRule].
func (alr *AlbListenerRule) Configuration() interface{} {
	return alr.Args
}

// DependOn is used for other resources to depend on [AlbListenerRule].
func (alr *AlbListenerRule) DependOn() terra.Reference {
	return terra.ReferenceResource(alr)
}

// Dependencies returns the list of resources [AlbListenerRule] depends_on.
func (alr *AlbListenerRule) Dependencies() terra.Dependencies {
	return alr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AlbListenerRule].
func (alr *AlbListenerRule) LifecycleManagement() *terra.Lifecycle {
	return alr.Lifecycle
}

// Attributes returns the attributes for [AlbListenerRule].
func (alr *AlbListenerRule) Attributes() albListenerRuleAttributes {
	return albListenerRuleAttributes{ref: terra.ReferenceResource(alr)}
}

// ImportState imports the given attribute values into [AlbListenerRule]'s state.
func (alr *AlbListenerRule) ImportState(av io.Reader) error {
	alr.state = &albListenerRuleState{}
	if err := json.NewDecoder(av).Decode(alr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alr.Type(), alr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AlbListenerRule] has state.
func (alr *AlbListenerRule) State() (*albListenerRuleState, bool) {
	return alr.state, alr.state != nil
}

// StateMust returns the state for [AlbListenerRule]. Panics if the state is nil.
func (alr *AlbListenerRule) StateMust() *albListenerRuleState {
	if alr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alr.Type(), alr.LocalName()))
	}
	return alr.state
}

// AlbListenerRuleArgs contains the configurations for aws_alb_listener_rule.
type AlbListenerRuleArgs struct {
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
	Action []alblistenerrule.Action `hcl:"action,block" validate:"min=1"`
	// Condition: min=1
	Condition []alblistenerrule.Condition `hcl:"condition,block" validate:"min=1"`
}
type albListenerRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_alb_listener_rule.
func (alr albListenerRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(alr.ref.Append("arn"))
}

// Id returns a reference to field id of aws_alb_listener_rule.
func (alr albListenerRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alr.ref.Append("id"))
}

// ListenerArn returns a reference to field listener_arn of aws_alb_listener_rule.
func (alr albListenerRuleAttributes) ListenerArn() terra.StringValue {
	return terra.ReferenceAsString(alr.ref.Append("listener_arn"))
}

// Priority returns a reference to field priority of aws_alb_listener_rule.
func (alr albListenerRuleAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(alr.ref.Append("priority"))
}

// Tags returns a reference to field tags of aws_alb_listener_rule.
func (alr albListenerRuleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_alb_listener_rule.
func (alr albListenerRuleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alr.ref.Append("tags_all"))
}

func (alr albListenerRuleAttributes) Action() terra.ListValue[alblistenerrule.ActionAttributes] {
	return terra.ReferenceAsList[alblistenerrule.ActionAttributes](alr.ref.Append("action"))
}

func (alr albListenerRuleAttributes) Condition() terra.SetValue[alblistenerrule.ConditionAttributes] {
	return terra.ReferenceAsSet[alblistenerrule.ConditionAttributes](alr.ref.Append("condition"))
}

type albListenerRuleState struct {
	Arn         string                           `json:"arn"`
	Id          string                           `json:"id"`
	ListenerArn string                           `json:"listener_arn"`
	Priority    float64                          `json:"priority"`
	Tags        map[string]string                `json:"tags"`
	TagsAll     map[string]string                `json:"tags_all"`
	Action      []alblistenerrule.ActionState    `json:"action"`
	Condition   []alblistenerrule.ConditionState `json:"condition"`
}
