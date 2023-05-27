// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakermonitoringschedule "github.com/golingon/terraproviders/aws/5.0.1/sagemakermonitoringschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerMonitoringSchedule creates a new instance of [SagemakerMonitoringSchedule].
func NewSagemakerMonitoringSchedule(name string, args SagemakerMonitoringScheduleArgs) *SagemakerMonitoringSchedule {
	return &SagemakerMonitoringSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerMonitoringSchedule)(nil)

// SagemakerMonitoringSchedule represents the Terraform resource aws_sagemaker_monitoring_schedule.
type SagemakerMonitoringSchedule struct {
	Name      string
	Args      SagemakerMonitoringScheduleArgs
	state     *sagemakerMonitoringScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerMonitoringSchedule].
func (sms *SagemakerMonitoringSchedule) Type() string {
	return "aws_sagemaker_monitoring_schedule"
}

// LocalName returns the local name for [SagemakerMonitoringSchedule].
func (sms *SagemakerMonitoringSchedule) LocalName() string {
	return sms.Name
}

// Configuration returns the configuration (args) for [SagemakerMonitoringSchedule].
func (sms *SagemakerMonitoringSchedule) Configuration() interface{} {
	return sms.Args
}

// DependOn is used for other resources to depend on [SagemakerMonitoringSchedule].
func (sms *SagemakerMonitoringSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(sms)
}

// Dependencies returns the list of resources [SagemakerMonitoringSchedule] depends_on.
func (sms *SagemakerMonitoringSchedule) Dependencies() terra.Dependencies {
	return sms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerMonitoringSchedule].
func (sms *SagemakerMonitoringSchedule) LifecycleManagement() *terra.Lifecycle {
	return sms.Lifecycle
}

// Attributes returns the attributes for [SagemakerMonitoringSchedule].
func (sms *SagemakerMonitoringSchedule) Attributes() sagemakerMonitoringScheduleAttributes {
	return sagemakerMonitoringScheduleAttributes{ref: terra.ReferenceResource(sms)}
}

// ImportState imports the given attribute values into [SagemakerMonitoringSchedule]'s state.
func (sms *SagemakerMonitoringSchedule) ImportState(av io.Reader) error {
	sms.state = &sagemakerMonitoringScheduleState{}
	if err := json.NewDecoder(av).Decode(sms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sms.Type(), sms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerMonitoringSchedule] has state.
func (sms *SagemakerMonitoringSchedule) State() (*sagemakerMonitoringScheduleState, bool) {
	return sms.state, sms.state != nil
}

// StateMust returns the state for [SagemakerMonitoringSchedule]. Panics if the state is nil.
func (sms *SagemakerMonitoringSchedule) StateMust() *sagemakerMonitoringScheduleState {
	if sms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sms.Type(), sms.LocalName()))
	}
	return sms.state
}

// SagemakerMonitoringScheduleArgs contains the configurations for aws_sagemaker_monitoring_schedule.
type SagemakerMonitoringScheduleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// MonitoringScheduleConfig: required
	MonitoringScheduleConfig *sagemakermonitoringschedule.MonitoringScheduleConfig `hcl:"monitoring_schedule_config,block" validate:"required"`
}
type sagemakerMonitoringScheduleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_monitoring_schedule.
func (sms sagemakerMonitoringScheduleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sagemaker_monitoring_schedule.
func (sms sagemakerMonitoringScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("id"))
}

// Name returns a reference to field name of aws_sagemaker_monitoring_schedule.
func (sms sagemakerMonitoringScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sms.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_sagemaker_monitoring_schedule.
func (sms sagemakerMonitoringScheduleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sms.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_monitoring_schedule.
func (sms sagemakerMonitoringScheduleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sms.ref.Append("tags_all"))
}

func (sms sagemakerMonitoringScheduleAttributes) MonitoringScheduleConfig() terra.ListValue[sagemakermonitoringschedule.MonitoringScheduleConfigAttributes] {
	return terra.ReferenceAsList[sagemakermonitoringschedule.MonitoringScheduleConfigAttributes](sms.ref.Append("monitoring_schedule_config"))
}

type sagemakerMonitoringScheduleState struct {
	Arn                      string                                                      `json:"arn"`
	Id                       string                                                      `json:"id"`
	Name                     string                                                      `json:"name"`
	Tags                     map[string]string                                           `json:"tags"`
	TagsAll                  map[string]string                                           `json:"tags_all"`
	MonitoringScheduleConfig []sagemakermonitoringschedule.MonitoringScheduleConfigState `json:"monitoring_schedule_config"`
}
