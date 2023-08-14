// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	batchschedulingpolicy "github.com/golingon/terraproviders/aws/5.12.0/batchschedulingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBatchSchedulingPolicy creates a new instance of [BatchSchedulingPolicy].
func NewBatchSchedulingPolicy(name string, args BatchSchedulingPolicyArgs) *BatchSchedulingPolicy {
	return &BatchSchedulingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BatchSchedulingPolicy)(nil)

// BatchSchedulingPolicy represents the Terraform resource aws_batch_scheduling_policy.
type BatchSchedulingPolicy struct {
	Name      string
	Args      BatchSchedulingPolicyArgs
	state     *batchSchedulingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BatchSchedulingPolicy].
func (bsp *BatchSchedulingPolicy) Type() string {
	return "aws_batch_scheduling_policy"
}

// LocalName returns the local name for [BatchSchedulingPolicy].
func (bsp *BatchSchedulingPolicy) LocalName() string {
	return bsp.Name
}

// Configuration returns the configuration (args) for [BatchSchedulingPolicy].
func (bsp *BatchSchedulingPolicy) Configuration() interface{} {
	return bsp.Args
}

// DependOn is used for other resources to depend on [BatchSchedulingPolicy].
func (bsp *BatchSchedulingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(bsp)
}

// Dependencies returns the list of resources [BatchSchedulingPolicy] depends_on.
func (bsp *BatchSchedulingPolicy) Dependencies() terra.Dependencies {
	return bsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BatchSchedulingPolicy].
func (bsp *BatchSchedulingPolicy) LifecycleManagement() *terra.Lifecycle {
	return bsp.Lifecycle
}

// Attributes returns the attributes for [BatchSchedulingPolicy].
func (bsp *BatchSchedulingPolicy) Attributes() batchSchedulingPolicyAttributes {
	return batchSchedulingPolicyAttributes{ref: terra.ReferenceResource(bsp)}
}

// ImportState imports the given attribute values into [BatchSchedulingPolicy]'s state.
func (bsp *BatchSchedulingPolicy) ImportState(av io.Reader) error {
	bsp.state = &batchSchedulingPolicyState{}
	if err := json.NewDecoder(av).Decode(bsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bsp.Type(), bsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BatchSchedulingPolicy] has state.
func (bsp *BatchSchedulingPolicy) State() (*batchSchedulingPolicyState, bool) {
	return bsp.state, bsp.state != nil
}

// StateMust returns the state for [BatchSchedulingPolicy]. Panics if the state is nil.
func (bsp *BatchSchedulingPolicy) StateMust() *batchSchedulingPolicyState {
	if bsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bsp.Type(), bsp.LocalName()))
	}
	return bsp.state
}

// BatchSchedulingPolicyArgs contains the configurations for aws_batch_scheduling_policy.
type BatchSchedulingPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// FairSharePolicy: optional
	FairSharePolicy *batchschedulingpolicy.FairSharePolicy `hcl:"fair_share_policy,block"`
}
type batchSchedulingPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_batch_scheduling_policy.
func (bsp batchSchedulingPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bsp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_batch_scheduling_policy.
func (bsp batchSchedulingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bsp.ref.Append("id"))
}

// Name returns a reference to field name of aws_batch_scheduling_policy.
func (bsp batchSchedulingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bsp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_batch_scheduling_policy.
func (bsp batchSchedulingPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bsp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_batch_scheduling_policy.
func (bsp batchSchedulingPolicyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bsp.ref.Append("tags_all"))
}

func (bsp batchSchedulingPolicyAttributes) FairSharePolicy() terra.ListValue[batchschedulingpolicy.FairSharePolicyAttributes] {
	return terra.ReferenceAsList[batchschedulingpolicy.FairSharePolicyAttributes](bsp.ref.Append("fair_share_policy"))
}

type batchSchedulingPolicyState struct {
	Arn             string                                       `json:"arn"`
	Id              string                                       `json:"id"`
	Name            string                                       `json:"name"`
	Tags            map[string]string                            `json:"tags"`
	TagsAll         map[string]string                            `json:"tags_all"`
	FairSharePolicy []batchschedulingpolicy.FairSharePolicyState `json:"fair_share_policy"`
}
