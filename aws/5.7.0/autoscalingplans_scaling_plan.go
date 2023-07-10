// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	autoscalingplansscalingplan "github.com/golingon/terraproviders/aws/5.7.0/autoscalingplansscalingplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAutoscalingplansScalingPlan creates a new instance of [AutoscalingplansScalingPlan].
func NewAutoscalingplansScalingPlan(name string, args AutoscalingplansScalingPlanArgs) *AutoscalingplansScalingPlan {
	return &AutoscalingplansScalingPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingplansScalingPlan)(nil)

// AutoscalingplansScalingPlan represents the Terraform resource aws_autoscalingplans_scaling_plan.
type AutoscalingplansScalingPlan struct {
	Name      string
	Args      AutoscalingplansScalingPlanArgs
	state     *autoscalingplansScalingPlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingplansScalingPlan].
func (asp *AutoscalingplansScalingPlan) Type() string {
	return "aws_autoscalingplans_scaling_plan"
}

// LocalName returns the local name for [AutoscalingplansScalingPlan].
func (asp *AutoscalingplansScalingPlan) LocalName() string {
	return asp.Name
}

// Configuration returns the configuration (args) for [AutoscalingplansScalingPlan].
func (asp *AutoscalingplansScalingPlan) Configuration() interface{} {
	return asp.Args
}

// DependOn is used for other resources to depend on [AutoscalingplansScalingPlan].
func (asp *AutoscalingplansScalingPlan) DependOn() terra.Reference {
	return terra.ReferenceResource(asp)
}

// Dependencies returns the list of resources [AutoscalingplansScalingPlan] depends_on.
func (asp *AutoscalingplansScalingPlan) Dependencies() terra.Dependencies {
	return asp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingplansScalingPlan].
func (asp *AutoscalingplansScalingPlan) LifecycleManagement() *terra.Lifecycle {
	return asp.Lifecycle
}

// Attributes returns the attributes for [AutoscalingplansScalingPlan].
func (asp *AutoscalingplansScalingPlan) Attributes() autoscalingplansScalingPlanAttributes {
	return autoscalingplansScalingPlanAttributes{ref: terra.ReferenceResource(asp)}
}

// ImportState imports the given attribute values into [AutoscalingplansScalingPlan]'s state.
func (asp *AutoscalingplansScalingPlan) ImportState(av io.Reader) error {
	asp.state = &autoscalingplansScalingPlanState{}
	if err := json.NewDecoder(av).Decode(asp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asp.Type(), asp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingplansScalingPlan] has state.
func (asp *AutoscalingplansScalingPlan) State() (*autoscalingplansScalingPlanState, bool) {
	return asp.state, asp.state != nil
}

// StateMust returns the state for [AutoscalingplansScalingPlan]. Panics if the state is nil.
func (asp *AutoscalingplansScalingPlan) StateMust() *autoscalingplansScalingPlanState {
	if asp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asp.Type(), asp.LocalName()))
	}
	return asp.state
}

// AutoscalingplansScalingPlanArgs contains the configurations for aws_autoscalingplans_scaling_plan.
type AutoscalingplansScalingPlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ApplicationSource: required
	ApplicationSource *autoscalingplansscalingplan.ApplicationSource `hcl:"application_source,block" validate:"required"`
	// ScalingInstruction: min=1
	ScalingInstruction []autoscalingplansscalingplan.ScalingInstruction `hcl:"scaling_instruction,block" validate:"min=1"`
}
type autoscalingplansScalingPlanAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_autoscalingplans_scaling_plan.
func (asp autoscalingplansScalingPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("id"))
}

// Name returns a reference to field name of aws_autoscalingplans_scaling_plan.
func (asp autoscalingplansScalingPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("name"))
}

// ScalingPlanVersion returns a reference to field scaling_plan_version of aws_autoscalingplans_scaling_plan.
func (asp autoscalingplansScalingPlanAttributes) ScalingPlanVersion() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("scaling_plan_version"))
}

func (asp autoscalingplansScalingPlanAttributes) ApplicationSource() terra.ListValue[autoscalingplansscalingplan.ApplicationSourceAttributes] {
	return terra.ReferenceAsList[autoscalingplansscalingplan.ApplicationSourceAttributes](asp.ref.Append("application_source"))
}

func (asp autoscalingplansScalingPlanAttributes) ScalingInstruction() terra.SetValue[autoscalingplansscalingplan.ScalingInstructionAttributes] {
	return terra.ReferenceAsSet[autoscalingplansscalingplan.ScalingInstructionAttributes](asp.ref.Append("scaling_instruction"))
}

type autoscalingplansScalingPlanState struct {
	Id                 string                                                `json:"id"`
	Name               string                                                `json:"name"`
	ScalingPlanVersion float64                                               `json:"scaling_plan_version"`
	ApplicationSource  []autoscalingplansscalingplan.ApplicationSourceState  `json:"application_source"`
	ScalingInstruction []autoscalingplansscalingplan.ScalingInstructionState `json:"scaling_instruction"`
}
