// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appautoscalingscheduledaction "github.com/golingon/terraproviders/aws/4.66.1/appautoscalingscheduledaction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppautoscalingScheduledAction creates a new instance of [AppautoscalingScheduledAction].
func NewAppautoscalingScheduledAction(name string, args AppautoscalingScheduledActionArgs) *AppautoscalingScheduledAction {
	return &AppautoscalingScheduledAction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppautoscalingScheduledAction)(nil)

// AppautoscalingScheduledAction represents the Terraform resource aws_appautoscaling_scheduled_action.
type AppautoscalingScheduledAction struct {
	Name      string
	Args      AppautoscalingScheduledActionArgs
	state     *appautoscalingScheduledActionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppautoscalingScheduledAction].
func (asa *AppautoscalingScheduledAction) Type() string {
	return "aws_appautoscaling_scheduled_action"
}

// LocalName returns the local name for [AppautoscalingScheduledAction].
func (asa *AppautoscalingScheduledAction) LocalName() string {
	return asa.Name
}

// Configuration returns the configuration (args) for [AppautoscalingScheduledAction].
func (asa *AppautoscalingScheduledAction) Configuration() interface{} {
	return asa.Args
}

// DependOn is used for other resources to depend on [AppautoscalingScheduledAction].
func (asa *AppautoscalingScheduledAction) DependOn() terra.Reference {
	return terra.ReferenceResource(asa)
}

// Dependencies returns the list of resources [AppautoscalingScheduledAction] depends_on.
func (asa *AppautoscalingScheduledAction) Dependencies() terra.Dependencies {
	return asa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppautoscalingScheduledAction].
func (asa *AppautoscalingScheduledAction) LifecycleManagement() *terra.Lifecycle {
	return asa.Lifecycle
}

// Attributes returns the attributes for [AppautoscalingScheduledAction].
func (asa *AppautoscalingScheduledAction) Attributes() appautoscalingScheduledActionAttributes {
	return appautoscalingScheduledActionAttributes{ref: terra.ReferenceResource(asa)}
}

// ImportState imports the given attribute values into [AppautoscalingScheduledAction]'s state.
func (asa *AppautoscalingScheduledAction) ImportState(av io.Reader) error {
	asa.state = &appautoscalingScheduledActionState{}
	if err := json.NewDecoder(av).Decode(asa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asa.Type(), asa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppautoscalingScheduledAction] has state.
func (asa *AppautoscalingScheduledAction) State() (*appautoscalingScheduledActionState, bool) {
	return asa.state, asa.state != nil
}

// StateMust returns the state for [AppautoscalingScheduledAction]. Panics if the state is nil.
func (asa *AppautoscalingScheduledAction) StateMust() *appautoscalingScheduledActionState {
	if asa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asa.Type(), asa.LocalName()))
	}
	return asa.state
}

// AppautoscalingScheduledActionArgs contains the configurations for aws_appautoscaling_scheduled_action.
type AppautoscalingScheduledActionArgs struct {
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// ScalableDimension: string, required
	ScalableDimension terra.StringValue `hcl:"scalable_dimension,attr" validate:"required"`
	// Schedule: string, required
	Schedule terra.StringValue `hcl:"schedule,attr" validate:"required"`
	// ServiceNamespace: string, required
	ServiceNamespace terra.StringValue `hcl:"service_namespace,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// Timezone: string, optional
	Timezone terra.StringValue `hcl:"timezone,attr"`
	// ScalableTargetAction: required
	ScalableTargetAction *appautoscalingscheduledaction.ScalableTargetAction `hcl:"scalable_target_action,block" validate:"required"`
}
type appautoscalingScheduledActionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("arn"))
}

// EndTime returns a reference to field end_time of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("end_time"))
}

// Id returns a reference to field id of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("id"))
}

// Name returns a reference to field name of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("name"))
}

// ResourceId returns a reference to field resource_id of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("resource_id"))
}

// ScalableDimension returns a reference to field scalable_dimension of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) ScalableDimension() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("scalable_dimension"))
}

// Schedule returns a reference to field schedule of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("schedule"))
}

// ServiceNamespace returns a reference to field service_namespace of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("service_namespace"))
}

// StartTime returns a reference to field start_time of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("start_time"))
}

// Timezone returns a reference to field timezone of aws_appautoscaling_scheduled_action.
func (asa appautoscalingScheduledActionAttributes) Timezone() terra.StringValue {
	return terra.ReferenceAsString(asa.ref.Append("timezone"))
}

func (asa appautoscalingScheduledActionAttributes) ScalableTargetAction() terra.ListValue[appautoscalingscheduledaction.ScalableTargetActionAttributes] {
	return terra.ReferenceAsList[appautoscalingscheduledaction.ScalableTargetActionAttributes](asa.ref.Append("scalable_target_action"))
}

type appautoscalingScheduledActionState struct {
	Arn                  string                                                    `json:"arn"`
	EndTime              string                                                    `json:"end_time"`
	Id                   string                                                    `json:"id"`
	Name                 string                                                    `json:"name"`
	ResourceId           string                                                    `json:"resource_id"`
	ScalableDimension    string                                                    `json:"scalable_dimension"`
	Schedule             string                                                    `json:"schedule"`
	ServiceNamespace     string                                                    `json:"service_namespace"`
	StartTime            string                                                    `json:"start_time"`
	Timezone             string                                                    `json:"timezone"`
	ScalableTargetAction []appautoscalingscheduledaction.ScalableTargetActionState `json:"scalable_target_action"`
}
