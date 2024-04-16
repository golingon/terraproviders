// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_appautoscaling_target

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_appautoscaling_target.
type Resource struct {
	Name      string
	Args      Args
	state     *awsAppautoscalingTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aat *Resource) Type() string {
	return "aws_appautoscaling_target"
}

// LocalName returns the local name for [Resource].
func (aat *Resource) LocalName() string {
	return aat.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aat *Resource) Configuration() interface{} {
	return aat.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aat *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aat)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aat *Resource) Dependencies() terra.Dependencies {
	return aat.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aat *Resource) LifecycleManagement() *terra.Lifecycle {
	return aat.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aat *Resource) Attributes() awsAppautoscalingTargetAttributes {
	return awsAppautoscalingTargetAttributes{ref: terra.ReferenceResource(aat)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aat *Resource) ImportState(state io.Reader) error {
	aat.state = &awsAppautoscalingTargetState{}
	if err := json.NewDecoder(state).Decode(aat.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aat.Type(), aat.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aat *Resource) State() (*awsAppautoscalingTargetState, bool) {
	return aat.state, aat.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aat *Resource) StateMust() *awsAppautoscalingTargetState {
	if aat.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aat.Type(), aat.LocalName()))
	}
	return aat.state
}

// Args contains the configurations for aws_appautoscaling_target.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxCapacity: number, required
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr" validate:"required"`
	// MinCapacity: number, required
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr" validate:"required"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// ScalableDimension: string, required
	ScalableDimension terra.StringValue `hcl:"scalable_dimension,attr" validate:"required"`
	// ServiceNamespace: string, required
	ServiceNamespace terra.StringValue `hcl:"service_namespace,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsAppautoscalingTargetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aat.ref.Append("arn"))
}

// Id returns a reference to field id of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aat.ref.Append("id"))
}

// MaxCapacity returns a reference to field max_capacity of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(aat.ref.Append("max_capacity"))
}

// MinCapacity returns a reference to field min_capacity of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(aat.ref.Append("min_capacity"))
}

// ResourceId returns a reference to field resource_id of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(aat.ref.Append("resource_id"))
}

// RoleArn returns a reference to field role_arn of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(aat.ref.Append("role_arn"))
}

// ScalableDimension returns a reference to field scalable_dimension of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) ScalableDimension() terra.StringValue {
	return terra.ReferenceAsString(aat.ref.Append("scalable_dimension"))
}

// ServiceNamespace returns a reference to field service_namespace of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(aat.ref.Append("service_namespace"))
}

// Tags returns a reference to field tags of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aat.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appautoscaling_target.
func (aat awsAppautoscalingTargetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aat.ref.Append("tags_all"))
}

type awsAppautoscalingTargetState struct {
	Arn               string            `json:"arn"`
	Id                string            `json:"id"`
	MaxCapacity       float64           `json:"max_capacity"`
	MinCapacity       float64           `json:"min_capacity"`
	ResourceId        string            `json:"resource_id"`
	RoleArn           string            `json:"role_arn"`
	ScalableDimension string            `json:"scalable_dimension"`
	ServiceNamespace  string            `json:"service_namespace"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
}
