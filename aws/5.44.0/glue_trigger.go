// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	gluetrigger "github.com/golingon/terraproviders/aws/5.44.0/gluetrigger"
	"io"
)

// NewGlueTrigger creates a new instance of [GlueTrigger].
func NewGlueTrigger(name string, args GlueTriggerArgs) *GlueTrigger {
	return &GlueTrigger{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueTrigger)(nil)

// GlueTrigger represents the Terraform resource aws_glue_trigger.
type GlueTrigger struct {
	Name      string
	Args      GlueTriggerArgs
	state     *glueTriggerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueTrigger].
func (gt *GlueTrigger) Type() string {
	return "aws_glue_trigger"
}

// LocalName returns the local name for [GlueTrigger].
func (gt *GlueTrigger) LocalName() string {
	return gt.Name
}

// Configuration returns the configuration (args) for [GlueTrigger].
func (gt *GlueTrigger) Configuration() interface{} {
	return gt.Args
}

// DependOn is used for other resources to depend on [GlueTrigger].
func (gt *GlueTrigger) DependOn() terra.Reference {
	return terra.ReferenceResource(gt)
}

// Dependencies returns the list of resources [GlueTrigger] depends_on.
func (gt *GlueTrigger) Dependencies() terra.Dependencies {
	return gt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueTrigger].
func (gt *GlueTrigger) LifecycleManagement() *terra.Lifecycle {
	return gt.Lifecycle
}

// Attributes returns the attributes for [GlueTrigger].
func (gt *GlueTrigger) Attributes() glueTriggerAttributes {
	return glueTriggerAttributes{ref: terra.ReferenceResource(gt)}
}

// ImportState imports the given attribute values into [GlueTrigger]'s state.
func (gt *GlueTrigger) ImportState(av io.Reader) error {
	gt.state = &glueTriggerState{}
	if err := json.NewDecoder(av).Decode(gt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gt.Type(), gt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueTrigger] has state.
func (gt *GlueTrigger) State() (*glueTriggerState, bool) {
	return gt.state, gt.state != nil
}

// StateMust returns the state for [GlueTrigger]. Panics if the state is nil.
func (gt *GlueTrigger) StateMust() *glueTriggerState {
	if gt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gt.Type(), gt.LocalName()))
	}
	return gt.state
}

// GlueTriggerArgs contains the configurations for aws_glue_trigger.
type GlueTriggerArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// StartOnCreation: bool, optional
	StartOnCreation terra.BoolValue `hcl:"start_on_creation,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// WorkflowName: string, optional
	WorkflowName terra.StringValue `hcl:"workflow_name,attr"`
	// Actions: min=1
	Actions []gluetrigger.Actions `hcl:"actions,block" validate:"min=1"`
	// EventBatchingCondition: min=0
	EventBatchingCondition []gluetrigger.EventBatchingCondition `hcl:"event_batching_condition,block" validate:"min=0"`
	// Predicate: optional
	Predicate *gluetrigger.Predicate `hcl:"predicate,block"`
	// Timeouts: optional
	Timeouts *gluetrigger.Timeouts `hcl:"timeouts,block"`
}
type glueTriggerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_glue_trigger.
func (gt glueTriggerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("arn"))
}

// Description returns a reference to field description of aws_glue_trigger.
func (gt glueTriggerAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_glue_trigger.
func (gt glueTriggerAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(gt.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_glue_trigger.
func (gt glueTriggerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("id"))
}

// Name returns a reference to field name of aws_glue_trigger.
func (gt glueTriggerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("name"))
}

// Schedule returns a reference to field schedule of aws_glue_trigger.
func (gt glueTriggerAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("schedule"))
}

// StartOnCreation returns a reference to field start_on_creation of aws_glue_trigger.
func (gt glueTriggerAttributes) StartOnCreation() terra.BoolValue {
	return terra.ReferenceAsBool(gt.ref.Append("start_on_creation"))
}

// State returns a reference to field state of aws_glue_trigger.
func (gt glueTriggerAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_glue_trigger.
func (gt glueTriggerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_glue_trigger.
func (gt glueTriggerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gt.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_glue_trigger.
func (gt glueTriggerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("type"))
}

// WorkflowName returns a reference to field workflow_name of aws_glue_trigger.
func (gt glueTriggerAttributes) WorkflowName() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("workflow_name"))
}

func (gt glueTriggerAttributes) Actions() terra.ListValue[gluetrigger.ActionsAttributes] {
	return terra.ReferenceAsList[gluetrigger.ActionsAttributes](gt.ref.Append("actions"))
}

func (gt glueTriggerAttributes) EventBatchingCondition() terra.ListValue[gluetrigger.EventBatchingConditionAttributes] {
	return terra.ReferenceAsList[gluetrigger.EventBatchingConditionAttributes](gt.ref.Append("event_batching_condition"))
}

func (gt glueTriggerAttributes) Predicate() terra.ListValue[gluetrigger.PredicateAttributes] {
	return terra.ReferenceAsList[gluetrigger.PredicateAttributes](gt.ref.Append("predicate"))
}

func (gt glueTriggerAttributes) Timeouts() gluetrigger.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gluetrigger.TimeoutsAttributes](gt.ref.Append("timeouts"))
}

type glueTriggerState struct {
	Arn                    string                                    `json:"arn"`
	Description            string                                    `json:"description"`
	Enabled                bool                                      `json:"enabled"`
	Id                     string                                    `json:"id"`
	Name                   string                                    `json:"name"`
	Schedule               string                                    `json:"schedule"`
	StartOnCreation        bool                                      `json:"start_on_creation"`
	State                  string                                    `json:"state"`
	Tags                   map[string]string                         `json:"tags"`
	TagsAll                map[string]string                         `json:"tags_all"`
	Type                   string                                    `json:"type"`
	WorkflowName           string                                    `json:"workflow_name"`
	Actions                []gluetrigger.ActionsState                `json:"actions"`
	EventBatchingCondition []gluetrigger.EventBatchingConditionState `json:"event_batching_condition"`
	Predicate              []gluetrigger.PredicateState              `json:"predicate"`
	Timeouts               *gluetrigger.TimeoutsState                `json:"timeouts"`
}
