// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	schedulerschedule "github.com/golingon/terraproviders/aws/4.63.0/schedulerschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSchedulerSchedule creates a new instance of [SchedulerSchedule].
func NewSchedulerSchedule(name string, args SchedulerScheduleArgs) *SchedulerSchedule {
	return &SchedulerSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SchedulerSchedule)(nil)

// SchedulerSchedule represents the Terraform resource aws_scheduler_schedule.
type SchedulerSchedule struct {
	Name      string
	Args      SchedulerScheduleArgs
	state     *schedulerScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SchedulerSchedule].
func (ss *SchedulerSchedule) Type() string {
	return "aws_scheduler_schedule"
}

// LocalName returns the local name for [SchedulerSchedule].
func (ss *SchedulerSchedule) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [SchedulerSchedule].
func (ss *SchedulerSchedule) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [SchedulerSchedule].
func (ss *SchedulerSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [SchedulerSchedule] depends_on.
func (ss *SchedulerSchedule) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SchedulerSchedule].
func (ss *SchedulerSchedule) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [SchedulerSchedule].
func (ss *SchedulerSchedule) Attributes() schedulerScheduleAttributes {
	return schedulerScheduleAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [SchedulerSchedule]'s state.
func (ss *SchedulerSchedule) ImportState(av io.Reader) error {
	ss.state = &schedulerScheduleState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SchedulerSchedule] has state.
func (ss *SchedulerSchedule) State() (*schedulerScheduleState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [SchedulerSchedule]. Panics if the state is nil.
func (ss *SchedulerSchedule) StateMust() *schedulerScheduleState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// SchedulerScheduleArgs contains the configurations for aws_scheduler_schedule.
type SchedulerScheduleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EndDate: string, optional
	EndDate terra.StringValue `hcl:"end_date,attr"`
	// GroupName: string, optional
	GroupName terra.StringValue `hcl:"group_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// ScheduleExpression: string, required
	ScheduleExpression terra.StringValue `hcl:"schedule_expression,attr" validate:"required"`
	// ScheduleExpressionTimezone: string, optional
	ScheduleExpressionTimezone terra.StringValue `hcl:"schedule_expression_timezone,attr"`
	// StartDate: string, optional
	StartDate terra.StringValue `hcl:"start_date,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// FlexibleTimeWindow: required
	FlexibleTimeWindow *schedulerschedule.FlexibleTimeWindow `hcl:"flexible_time_window,block" validate:"required"`
	// Target: required
	Target *schedulerschedule.Target `hcl:"target,block" validate:"required"`
}
type schedulerScheduleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("arn"))
}

// Description returns a reference to field description of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("description"))
}

// EndDate returns a reference to field end_date of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("end_date"))
}

// GroupName returns a reference to field group_name of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("group_name"))
}

// Id returns a reference to field id of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("kms_key_arn"))
}

// Name returns a reference to field name of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name_prefix"))
}

// ScheduleExpression returns a reference to field schedule_expression of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) ScheduleExpression() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("schedule_expression"))
}

// ScheduleExpressionTimezone returns a reference to field schedule_expression_timezone of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) ScheduleExpressionTimezone() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("schedule_expression_timezone"))
}

// StartDate returns a reference to field start_date of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("start_date"))
}

// State returns a reference to field state of aws_scheduler_schedule.
func (ss schedulerScheduleAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("state"))
}

func (ss schedulerScheduleAttributes) FlexibleTimeWindow() terra.ListValue[schedulerschedule.FlexibleTimeWindowAttributes] {
	return terra.ReferenceAsList[schedulerschedule.FlexibleTimeWindowAttributes](ss.ref.Append("flexible_time_window"))
}

func (ss schedulerScheduleAttributes) Target() terra.ListValue[schedulerschedule.TargetAttributes] {
	return terra.ReferenceAsList[schedulerschedule.TargetAttributes](ss.ref.Append("target"))
}

type schedulerScheduleState struct {
	Arn                        string                                      `json:"arn"`
	Description                string                                      `json:"description"`
	EndDate                    string                                      `json:"end_date"`
	GroupName                  string                                      `json:"group_name"`
	Id                         string                                      `json:"id"`
	KmsKeyArn                  string                                      `json:"kms_key_arn"`
	Name                       string                                      `json:"name"`
	NamePrefix                 string                                      `json:"name_prefix"`
	ScheduleExpression         string                                      `json:"schedule_expression"`
	ScheduleExpressionTimezone string                                      `json:"schedule_expression_timezone"`
	StartDate                  string                                      `json:"start_date"`
	State                      string                                      `json:"state"`
	FlexibleTimeWindow         []schedulerschedule.FlexibleTimeWindowState `json:"flexible_time_window"`
	Target                     []schedulerschedule.TargetState             `json:"target"`
}
