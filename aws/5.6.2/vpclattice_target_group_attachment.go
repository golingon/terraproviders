// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpclatticetargetgroupattachment "github.com/golingon/terraproviders/aws/5.6.2/vpclatticetargetgroupattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeTargetGroupAttachment creates a new instance of [VpclatticeTargetGroupAttachment].
func NewVpclatticeTargetGroupAttachment(name string, args VpclatticeTargetGroupAttachmentArgs) *VpclatticeTargetGroupAttachment {
	return &VpclatticeTargetGroupAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeTargetGroupAttachment)(nil)

// VpclatticeTargetGroupAttachment represents the Terraform resource aws_vpclattice_target_group_attachment.
type VpclatticeTargetGroupAttachment struct {
	Name      string
	Args      VpclatticeTargetGroupAttachmentArgs
	state     *vpclatticeTargetGroupAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeTargetGroupAttachment].
func (vtga *VpclatticeTargetGroupAttachment) Type() string {
	return "aws_vpclattice_target_group_attachment"
}

// LocalName returns the local name for [VpclatticeTargetGroupAttachment].
func (vtga *VpclatticeTargetGroupAttachment) LocalName() string {
	return vtga.Name
}

// Configuration returns the configuration (args) for [VpclatticeTargetGroupAttachment].
func (vtga *VpclatticeTargetGroupAttachment) Configuration() interface{} {
	return vtga.Args
}

// DependOn is used for other resources to depend on [VpclatticeTargetGroupAttachment].
func (vtga *VpclatticeTargetGroupAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(vtga)
}

// Dependencies returns the list of resources [VpclatticeTargetGroupAttachment] depends_on.
func (vtga *VpclatticeTargetGroupAttachment) Dependencies() terra.Dependencies {
	return vtga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeTargetGroupAttachment].
func (vtga *VpclatticeTargetGroupAttachment) LifecycleManagement() *terra.Lifecycle {
	return vtga.Lifecycle
}

// Attributes returns the attributes for [VpclatticeTargetGroupAttachment].
func (vtga *VpclatticeTargetGroupAttachment) Attributes() vpclatticeTargetGroupAttachmentAttributes {
	return vpclatticeTargetGroupAttachmentAttributes{ref: terra.ReferenceResource(vtga)}
}

// ImportState imports the given attribute values into [VpclatticeTargetGroupAttachment]'s state.
func (vtga *VpclatticeTargetGroupAttachment) ImportState(av io.Reader) error {
	vtga.state = &vpclatticeTargetGroupAttachmentState{}
	if err := json.NewDecoder(av).Decode(vtga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vtga.Type(), vtga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeTargetGroupAttachment] has state.
func (vtga *VpclatticeTargetGroupAttachment) State() (*vpclatticeTargetGroupAttachmentState, bool) {
	return vtga.state, vtga.state != nil
}

// StateMust returns the state for [VpclatticeTargetGroupAttachment]. Panics if the state is nil.
func (vtga *VpclatticeTargetGroupAttachment) StateMust() *vpclatticeTargetGroupAttachmentState {
	if vtga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vtga.Type(), vtga.LocalName()))
	}
	return vtga.state
}

// VpclatticeTargetGroupAttachmentArgs contains the configurations for aws_vpclattice_target_group_attachment.
type VpclatticeTargetGroupAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TargetGroupIdentifier: string, required
	TargetGroupIdentifier terra.StringValue `hcl:"target_group_identifier,attr" validate:"required"`
	// Target: required
	Target *vpclatticetargetgroupattachment.Target `hcl:"target,block" validate:"required"`
	// Timeouts: optional
	Timeouts *vpclatticetargetgroupattachment.Timeouts `hcl:"timeouts,block"`
}
type vpclatticeTargetGroupAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpclattice_target_group_attachment.
func (vtga vpclatticeTargetGroupAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vtga.ref.Append("id"))
}

// TargetGroupIdentifier returns a reference to field target_group_identifier of aws_vpclattice_target_group_attachment.
func (vtga vpclatticeTargetGroupAttachmentAttributes) TargetGroupIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vtga.ref.Append("target_group_identifier"))
}

func (vtga vpclatticeTargetGroupAttachmentAttributes) Target() terra.ListValue[vpclatticetargetgroupattachment.TargetAttributes] {
	return terra.ReferenceAsList[vpclatticetargetgroupattachment.TargetAttributes](vtga.ref.Append("target"))
}

func (vtga vpclatticeTargetGroupAttachmentAttributes) Timeouts() vpclatticetargetgroupattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpclatticetargetgroupattachment.TimeoutsAttributes](vtga.ref.Append("timeouts"))
}

type vpclatticeTargetGroupAttachmentState struct {
	Id                    string                                         `json:"id"`
	TargetGroupIdentifier string                                         `json:"target_group_identifier"`
	Target                []vpclatticetargetgroupattachment.TargetState  `json:"target"`
	Timeouts              *vpclatticetargetgroupattachment.TimeoutsState `json:"timeouts"`
}
