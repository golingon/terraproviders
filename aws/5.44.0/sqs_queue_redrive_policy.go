// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSqsQueueRedrivePolicy creates a new instance of [SqsQueueRedrivePolicy].
func NewSqsQueueRedrivePolicy(name string, args SqsQueueRedrivePolicyArgs) *SqsQueueRedrivePolicy {
	return &SqsQueueRedrivePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SqsQueueRedrivePolicy)(nil)

// SqsQueueRedrivePolicy represents the Terraform resource aws_sqs_queue_redrive_policy.
type SqsQueueRedrivePolicy struct {
	Name      string
	Args      SqsQueueRedrivePolicyArgs
	state     *sqsQueueRedrivePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SqsQueueRedrivePolicy].
func (sqrp *SqsQueueRedrivePolicy) Type() string {
	return "aws_sqs_queue_redrive_policy"
}

// LocalName returns the local name for [SqsQueueRedrivePolicy].
func (sqrp *SqsQueueRedrivePolicy) LocalName() string {
	return sqrp.Name
}

// Configuration returns the configuration (args) for [SqsQueueRedrivePolicy].
func (sqrp *SqsQueueRedrivePolicy) Configuration() interface{} {
	return sqrp.Args
}

// DependOn is used for other resources to depend on [SqsQueueRedrivePolicy].
func (sqrp *SqsQueueRedrivePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sqrp)
}

// Dependencies returns the list of resources [SqsQueueRedrivePolicy] depends_on.
func (sqrp *SqsQueueRedrivePolicy) Dependencies() terra.Dependencies {
	return sqrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SqsQueueRedrivePolicy].
func (sqrp *SqsQueueRedrivePolicy) LifecycleManagement() *terra.Lifecycle {
	return sqrp.Lifecycle
}

// Attributes returns the attributes for [SqsQueueRedrivePolicy].
func (sqrp *SqsQueueRedrivePolicy) Attributes() sqsQueueRedrivePolicyAttributes {
	return sqsQueueRedrivePolicyAttributes{ref: terra.ReferenceResource(sqrp)}
}

// ImportState imports the given attribute values into [SqsQueueRedrivePolicy]'s state.
func (sqrp *SqsQueueRedrivePolicy) ImportState(av io.Reader) error {
	sqrp.state = &sqsQueueRedrivePolicyState{}
	if err := json.NewDecoder(av).Decode(sqrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sqrp.Type(), sqrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SqsQueueRedrivePolicy] has state.
func (sqrp *SqsQueueRedrivePolicy) State() (*sqsQueueRedrivePolicyState, bool) {
	return sqrp.state, sqrp.state != nil
}

// StateMust returns the state for [SqsQueueRedrivePolicy]. Panics if the state is nil.
func (sqrp *SqsQueueRedrivePolicy) StateMust() *sqsQueueRedrivePolicyState {
	if sqrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sqrp.Type(), sqrp.LocalName()))
	}
	return sqrp.state
}

// SqsQueueRedrivePolicyArgs contains the configurations for aws_sqs_queue_redrive_policy.
type SqsQueueRedrivePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// QueueUrl: string, required
	QueueUrl terra.StringValue `hcl:"queue_url,attr" validate:"required"`
	// RedrivePolicy: string, required
	RedrivePolicy terra.StringValue `hcl:"redrive_policy,attr" validate:"required"`
}
type sqsQueueRedrivePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_sqs_queue_redrive_policy.
func (sqrp sqsQueueRedrivePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sqrp.ref.Append("id"))
}

// QueueUrl returns a reference to field queue_url of aws_sqs_queue_redrive_policy.
func (sqrp sqsQueueRedrivePolicyAttributes) QueueUrl() terra.StringValue {
	return terra.ReferenceAsString(sqrp.ref.Append("queue_url"))
}

// RedrivePolicy returns a reference to field redrive_policy of aws_sqs_queue_redrive_policy.
func (sqrp sqsQueueRedrivePolicyAttributes) RedrivePolicy() terra.StringValue {
	return terra.ReferenceAsString(sqrp.ref.Append("redrive_policy"))
}

type sqsQueueRedrivePolicyState struct {
	Id            string `json:"id"`
	QueueUrl      string `json:"queue_url"`
	RedrivePolicy string `json:"redrive_policy"`
}
