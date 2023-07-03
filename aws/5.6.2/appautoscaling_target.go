// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppautoscalingTarget creates a new instance of [AppautoscalingTarget].
func NewAppautoscalingTarget(name string, args AppautoscalingTargetArgs) *AppautoscalingTarget {
	return &AppautoscalingTarget{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppautoscalingTarget)(nil)

// AppautoscalingTarget represents the Terraform resource aws_appautoscaling_target.
type AppautoscalingTarget struct {
	Name      string
	Args      AppautoscalingTargetArgs
	state     *appautoscalingTargetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppautoscalingTarget].
func (at *AppautoscalingTarget) Type() string {
	return "aws_appautoscaling_target"
}

// LocalName returns the local name for [AppautoscalingTarget].
func (at *AppautoscalingTarget) LocalName() string {
	return at.Name
}

// Configuration returns the configuration (args) for [AppautoscalingTarget].
func (at *AppautoscalingTarget) Configuration() interface{} {
	return at.Args
}

// DependOn is used for other resources to depend on [AppautoscalingTarget].
func (at *AppautoscalingTarget) DependOn() terra.Reference {
	return terra.ReferenceResource(at)
}

// Dependencies returns the list of resources [AppautoscalingTarget] depends_on.
func (at *AppautoscalingTarget) Dependencies() terra.Dependencies {
	return at.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppautoscalingTarget].
func (at *AppautoscalingTarget) LifecycleManagement() *terra.Lifecycle {
	return at.Lifecycle
}

// Attributes returns the attributes for [AppautoscalingTarget].
func (at *AppautoscalingTarget) Attributes() appautoscalingTargetAttributes {
	return appautoscalingTargetAttributes{ref: terra.ReferenceResource(at)}
}

// ImportState imports the given attribute values into [AppautoscalingTarget]'s state.
func (at *AppautoscalingTarget) ImportState(av io.Reader) error {
	at.state = &appautoscalingTargetState{}
	if err := json.NewDecoder(av).Decode(at.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", at.Type(), at.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppautoscalingTarget] has state.
func (at *AppautoscalingTarget) State() (*appautoscalingTargetState, bool) {
	return at.state, at.state != nil
}

// StateMust returns the state for [AppautoscalingTarget]. Panics if the state is nil.
func (at *AppautoscalingTarget) StateMust() *appautoscalingTargetState {
	if at.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", at.Type(), at.LocalName()))
	}
	return at.state
}

// AppautoscalingTargetArgs contains the configurations for aws_appautoscaling_target.
type AppautoscalingTargetArgs struct {
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
type appautoscalingTargetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("arn"))
}

// Id returns a reference to field id of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("id"))
}

// MaxCapacity returns a reference to field max_capacity of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(at.ref.Append("max_capacity"))
}

// MinCapacity returns a reference to field min_capacity of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(at.ref.Append("min_capacity"))
}

// ResourceId returns a reference to field resource_id of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("resource_id"))
}

// RoleArn returns a reference to field role_arn of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("role_arn"))
}

// ScalableDimension returns a reference to field scalable_dimension of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) ScalableDimension() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("scalable_dimension"))
}

// ServiceNamespace returns a reference to field service_namespace of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) ServiceNamespace() terra.StringValue {
	return terra.ReferenceAsString(at.ref.Append("service_namespace"))
}

// Tags returns a reference to field tags of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](at.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appautoscaling_target.
func (at appautoscalingTargetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](at.ref.Append("tags_all"))
}

type appautoscalingTargetState struct {
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