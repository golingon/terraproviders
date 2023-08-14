// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSqsQueuePolicy creates a new instance of [SqsQueuePolicy].
func NewSqsQueuePolicy(name string, args SqsQueuePolicyArgs) *SqsQueuePolicy {
	return &SqsQueuePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqsQueuePolicy)(nil)

// SqsQueuePolicy represents the Terraform resource aws_sqs_queue_policy.
type SqsQueuePolicy struct {
	Name      string
	Args      SqsQueuePolicyArgs
	state     *sqsQueuePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqsQueuePolicy].
func (sqp *SqsQueuePolicy) Type() string {
	return "aws_sqs_queue_policy"
}

// LocalName returns the local name for [SqsQueuePolicy].
func (sqp *SqsQueuePolicy) LocalName() string {
	return sqp.Name
}

// Configuration returns the configuration (args) for [SqsQueuePolicy].
func (sqp *SqsQueuePolicy) Configuration() interface{} {
	return sqp.Args
}

// DependOn is used for other resources to depend on [SqsQueuePolicy].
func (sqp *SqsQueuePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sqp)
}

// Dependencies returns the list of resources [SqsQueuePolicy] depends_on.
func (sqp *SqsQueuePolicy) Dependencies() terra.Dependencies {
	return sqp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqsQueuePolicy].
func (sqp *SqsQueuePolicy) LifecycleManagement() *terra.Lifecycle {
	return sqp.Lifecycle
}

// Attributes returns the attributes for [SqsQueuePolicy].
func (sqp *SqsQueuePolicy) Attributes() sqsQueuePolicyAttributes {
	return sqsQueuePolicyAttributes{ref: terra.ReferenceResource(sqp)}
}

// ImportState imports the given attribute values into [SqsQueuePolicy]'s state.
func (sqp *SqsQueuePolicy) ImportState(av io.Reader) error {
	sqp.state = &sqsQueuePolicyState{}
	if err := json.NewDecoder(av).Decode(sqp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sqp.Type(), sqp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqsQueuePolicy] has state.
func (sqp *SqsQueuePolicy) State() (*sqsQueuePolicyState, bool) {
	return sqp.state, sqp.state != nil
}

// StateMust returns the state for [SqsQueuePolicy]. Panics if the state is nil.
func (sqp *SqsQueuePolicy) StateMust() *sqsQueuePolicyState {
	if sqp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sqp.Type(), sqp.LocalName()))
	}
	return sqp.state
}

// SqsQueuePolicyArgs contains the configurations for aws_sqs_queue_policy.
type SqsQueuePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// QueueUrl: string, required
	QueueUrl terra.StringValue `hcl:"queue_url,attr" validate:"required"`
}
type sqsQueuePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_sqs_queue_policy.
func (sqp sqsQueuePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sqp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_sqs_queue_policy.
func (sqp sqsQueuePolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(sqp.ref.Append("policy"))
}

// QueueUrl returns a reference to field queue_url of aws_sqs_queue_policy.
func (sqp sqsQueuePolicyAttributes) QueueUrl() terra.StringValue {
	return terra.ReferenceAsString(sqp.ref.Append("queue_url"))
}

type sqsQueuePolicyState struct {
	Id       string `json:"id"`
	Policy   string `json:"policy"`
	QueueUrl string `json:"queue_url"`
}
