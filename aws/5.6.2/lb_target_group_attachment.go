// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbTargetGroupAttachment creates a new instance of [LbTargetGroupAttachment].
func NewLbTargetGroupAttachment(name string, args LbTargetGroupAttachmentArgs) *LbTargetGroupAttachment {
	return &LbTargetGroupAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbTargetGroupAttachment)(nil)

// LbTargetGroupAttachment represents the Terraform resource aws_lb_target_group_attachment.
type LbTargetGroupAttachment struct {
	Name      string
	Args      LbTargetGroupAttachmentArgs
	state     *lbTargetGroupAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbTargetGroupAttachment].
func (ltga *LbTargetGroupAttachment) Type() string {
	return "aws_lb_target_group_attachment"
}

// LocalName returns the local name for [LbTargetGroupAttachment].
func (ltga *LbTargetGroupAttachment) LocalName() string {
	return ltga.Name
}

// Configuration returns the configuration (args) for [LbTargetGroupAttachment].
func (ltga *LbTargetGroupAttachment) Configuration() interface{} {
	return ltga.Args
}

// DependOn is used for other resources to depend on [LbTargetGroupAttachment].
func (ltga *LbTargetGroupAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(ltga)
}

// Dependencies returns the list of resources [LbTargetGroupAttachment] depends_on.
func (ltga *LbTargetGroupAttachment) Dependencies() terra.Dependencies {
	return ltga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbTargetGroupAttachment].
func (ltga *LbTargetGroupAttachment) LifecycleManagement() *terra.Lifecycle {
	return ltga.Lifecycle
}

// Attributes returns the attributes for [LbTargetGroupAttachment].
func (ltga *LbTargetGroupAttachment) Attributes() lbTargetGroupAttachmentAttributes {
	return lbTargetGroupAttachmentAttributes{ref: terra.ReferenceResource(ltga)}
}

// ImportState imports the given attribute values into [LbTargetGroupAttachment]'s state.
func (ltga *LbTargetGroupAttachment) ImportState(av io.Reader) error {
	ltga.state = &lbTargetGroupAttachmentState{}
	if err := json.NewDecoder(av).Decode(ltga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ltga.Type(), ltga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbTargetGroupAttachment] has state.
func (ltga *LbTargetGroupAttachment) State() (*lbTargetGroupAttachmentState, bool) {
	return ltga.state, ltga.state != nil
}

// StateMust returns the state for [LbTargetGroupAttachment]. Panics if the state is nil.
func (ltga *LbTargetGroupAttachment) StateMust() *lbTargetGroupAttachmentState {
	if ltga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ltga.Type(), ltga.LocalName()))
	}
	return ltga.state
}

// LbTargetGroupAttachmentArgs contains the configurations for aws_lb_target_group_attachment.
type LbTargetGroupAttachmentArgs struct {
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// TargetGroupArn: string, required
	TargetGroupArn terra.StringValue `hcl:"target_group_arn,attr" validate:"required"`
	// TargetId: string, required
	TargetId terra.StringValue `hcl:"target_id,attr" validate:"required"`
}
type lbTargetGroupAttachmentAttributes struct {
	ref terra.Reference
}

// AvailabilityZone returns a reference to field availability_zone of aws_lb_target_group_attachment.
func (ltga lbTargetGroupAttachmentAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(ltga.ref.Append("availability_zone"))
}

// Id returns a reference to field id of aws_lb_target_group_attachment.
func (ltga lbTargetGroupAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ltga.ref.Append("id"))
}

// Port returns a reference to field port of aws_lb_target_group_attachment.
func (ltga lbTargetGroupAttachmentAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ltga.ref.Append("port"))
}

// TargetGroupArn returns a reference to field target_group_arn of aws_lb_target_group_attachment.
func (ltga lbTargetGroupAttachmentAttributes) TargetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(ltga.ref.Append("target_group_arn"))
}

// TargetId returns a reference to field target_id of aws_lb_target_group_attachment.
func (ltga lbTargetGroupAttachmentAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(ltga.ref.Append("target_id"))
}

type lbTargetGroupAttachmentState struct {
	AvailabilityZone string  `json:"availability_zone"`
	Id               string  `json:"id"`
	Port             float64 `json:"port"`
	TargetGroupArn   string  `json:"target_group_arn"`
	TargetId         string  `json:"target_id"`
}
