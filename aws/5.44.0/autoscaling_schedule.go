// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewAutoscalingSchedule creates a new instance of [AutoscalingSchedule].
func NewAutoscalingSchedule(name string, args AutoscalingScheduleArgs) *AutoscalingSchedule {
	return &AutoscalingSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingSchedule)(nil)

// AutoscalingSchedule represents the Terraform resource aws_autoscaling_schedule.
type AutoscalingSchedule struct {
	Name      string
	Args      AutoscalingScheduleArgs
	state     *autoscalingScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingSchedule].
func (as *AutoscalingSchedule) Type() string {
	return "aws_autoscaling_schedule"
}

// LocalName returns the local name for [AutoscalingSchedule].
func (as *AutoscalingSchedule) LocalName() string {
	return as.Name
}

// Configuration returns the configuration (args) for [AutoscalingSchedule].
func (as *AutoscalingSchedule) Configuration() interface{} {
	return as.Args
}

// DependOn is used for other resources to depend on [AutoscalingSchedule].
func (as *AutoscalingSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(as)
}

// Dependencies returns the list of resources [AutoscalingSchedule] depends_on.
func (as *AutoscalingSchedule) Dependencies() terra.Dependencies {
	return as.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingSchedule].
func (as *AutoscalingSchedule) LifecycleManagement() *terra.Lifecycle {
	return as.Lifecycle
}

// Attributes returns the attributes for [AutoscalingSchedule].
func (as *AutoscalingSchedule) Attributes() autoscalingScheduleAttributes {
	return autoscalingScheduleAttributes{ref: terra.ReferenceResource(as)}
}

// ImportState imports the given attribute values into [AutoscalingSchedule]'s state.
func (as *AutoscalingSchedule) ImportState(av io.Reader) error {
	as.state = &autoscalingScheduleState{}
	if err := json.NewDecoder(av).Decode(as.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", as.Type(), as.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingSchedule] has state.
func (as *AutoscalingSchedule) State() (*autoscalingScheduleState, bool) {
	return as.state, as.state != nil
}

// StateMust returns the state for [AutoscalingSchedule]. Panics if the state is nil.
func (as *AutoscalingSchedule) StateMust() *autoscalingScheduleState {
	if as.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", as.Type(), as.LocalName()))
	}
	return as.state
}

// AutoscalingScheduleArgs contains the configurations for aws_autoscaling_schedule.
type AutoscalingScheduleArgs struct {
	// AutoscalingGroupName: string, required
	AutoscalingGroupName terra.StringValue `hcl:"autoscaling_group_name,attr" validate:"required"`
	// DesiredCapacity: number, optional
	DesiredCapacity terra.NumberValue `hcl:"desired_capacity,attr"`
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxSize: number, optional
	MaxSize terra.NumberValue `hcl:"max_size,attr"`
	// MinSize: number, optional
	MinSize terra.NumberValue `hcl:"min_size,attr"`
	// Recurrence: string, optional
	Recurrence terra.StringValue `hcl:"recurrence,attr"`
	// ScheduledActionName: string, required
	ScheduledActionName terra.StringValue `hcl:"scheduled_action_name,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
}
type autoscalingScheduleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("arn"))
}

// AutoscalingGroupName returns a reference to field autoscaling_group_name of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) AutoscalingGroupName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("autoscaling_group_name"))
}

// DesiredCapacity returns a reference to field desired_capacity of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) DesiredCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("desired_capacity"))
}

// EndTime returns a reference to field end_time of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("end_time"))
}

// Id returns a reference to field id of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("id"))
}

// MaxSize returns a reference to field max_size of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) MaxSize() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("max_size"))
}

// MinSize returns a reference to field min_size of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) MinSize() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("min_size"))
}

// Recurrence returns a reference to field recurrence of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) Recurrence() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("recurrence"))
}

// ScheduledActionName returns a reference to field scheduled_action_name of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) ScheduledActionName() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("scheduled_action_name"))
}

// StartTime returns a reference to field start_time of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("start_time"))
}

// TimeZone returns a reference to field time_zone of aws_autoscaling_schedule.
func (as autoscalingScheduleAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(as.ref.Append("time_zone"))
}

type autoscalingScheduleState struct {
	Arn                  string  `json:"arn"`
	AutoscalingGroupName string  `json:"autoscaling_group_name"`
	DesiredCapacity      float64 `json:"desired_capacity"`
	EndTime              string  `json:"end_time"`
	Id                   string  `json:"id"`
	MaxSize              float64 `json:"max_size"`
	MinSize              float64 `json:"min_size"`
	Recurrence           string  `json:"recurrence"`
	ScheduledActionName  string  `json:"scheduled_action_name"`
	StartTime            string  `json:"start_time"`
	TimeZone             string  `json:"time_zone"`
}
