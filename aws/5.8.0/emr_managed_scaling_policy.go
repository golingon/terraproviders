// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrmanagedscalingpolicy "github.com/golingon/terraproviders/aws/5.8.0/emrmanagedscalingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrManagedScalingPolicy creates a new instance of [EmrManagedScalingPolicy].
func NewEmrManagedScalingPolicy(name string, args EmrManagedScalingPolicyArgs) *EmrManagedScalingPolicy {
	return &EmrManagedScalingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrManagedScalingPolicy)(nil)

// EmrManagedScalingPolicy represents the Terraform resource aws_emr_managed_scaling_policy.
type EmrManagedScalingPolicy struct {
	Name      string
	Args      EmrManagedScalingPolicyArgs
	state     *emrManagedScalingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrManagedScalingPolicy].
func (emsp *EmrManagedScalingPolicy) Type() string {
	return "aws_emr_managed_scaling_policy"
}

// LocalName returns the local name for [EmrManagedScalingPolicy].
func (emsp *EmrManagedScalingPolicy) LocalName() string {
	return emsp.Name
}

// Configuration returns the configuration (args) for [EmrManagedScalingPolicy].
func (emsp *EmrManagedScalingPolicy) Configuration() interface{} {
	return emsp.Args
}

// DependOn is used for other resources to depend on [EmrManagedScalingPolicy].
func (emsp *EmrManagedScalingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(emsp)
}

// Dependencies returns the list of resources [EmrManagedScalingPolicy] depends_on.
func (emsp *EmrManagedScalingPolicy) Dependencies() terra.Dependencies {
	return emsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrManagedScalingPolicy].
func (emsp *EmrManagedScalingPolicy) LifecycleManagement() *terra.Lifecycle {
	return emsp.Lifecycle
}

// Attributes returns the attributes for [EmrManagedScalingPolicy].
func (emsp *EmrManagedScalingPolicy) Attributes() emrManagedScalingPolicyAttributes {
	return emrManagedScalingPolicyAttributes{ref: terra.ReferenceResource(emsp)}
}

// ImportState imports the given attribute values into [EmrManagedScalingPolicy]'s state.
func (emsp *EmrManagedScalingPolicy) ImportState(av io.Reader) error {
	emsp.state = &emrManagedScalingPolicyState{}
	if err := json.NewDecoder(av).Decode(emsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", emsp.Type(), emsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrManagedScalingPolicy] has state.
func (emsp *EmrManagedScalingPolicy) State() (*emrManagedScalingPolicyState, bool) {
	return emsp.state, emsp.state != nil
}

// StateMust returns the state for [EmrManagedScalingPolicy]. Panics if the state is nil.
func (emsp *EmrManagedScalingPolicy) StateMust() *emrManagedScalingPolicyState {
	if emsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", emsp.Type(), emsp.LocalName()))
	}
	return emsp.state
}

// EmrManagedScalingPolicyArgs contains the configurations for aws_emr_managed_scaling_policy.
type EmrManagedScalingPolicyArgs struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ComputeLimits: min=1
	ComputeLimits []emrmanagedscalingpolicy.ComputeLimits `hcl:"compute_limits,block" validate:"min=1"`
}
type emrManagedScalingPolicyAttributes struct {
	ref terra.Reference
}

// ClusterId returns a reference to field cluster_id of aws_emr_managed_scaling_policy.
func (emsp emrManagedScalingPolicyAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(emsp.ref.Append("cluster_id"))
}

// Id returns a reference to field id of aws_emr_managed_scaling_policy.
func (emsp emrManagedScalingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(emsp.ref.Append("id"))
}

func (emsp emrManagedScalingPolicyAttributes) ComputeLimits() terra.SetValue[emrmanagedscalingpolicy.ComputeLimitsAttributes] {
	return terra.ReferenceAsSet[emrmanagedscalingpolicy.ComputeLimitsAttributes](emsp.ref.Append("compute_limits"))
}

type emrManagedScalingPolicyState struct {
	ClusterId     string                                       `json:"cluster_id"`
	Id            string                                       `json:"id"`
	ComputeLimits []emrmanagedscalingpolicy.ComputeLimitsState `json:"compute_limits"`
}
