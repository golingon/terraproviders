// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssoadminpermissionsboundaryattachment "github.com/golingon/terraproviders/aws/5.11.0/ssoadminpermissionsboundaryattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsoadminPermissionsBoundaryAttachment creates a new instance of [SsoadminPermissionsBoundaryAttachment].
func NewSsoadminPermissionsBoundaryAttachment(name string, args SsoadminPermissionsBoundaryAttachmentArgs) *SsoadminPermissionsBoundaryAttachment {
	return &SsoadminPermissionsBoundaryAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsoadminPermissionsBoundaryAttachment)(nil)

// SsoadminPermissionsBoundaryAttachment represents the Terraform resource aws_ssoadmin_permissions_boundary_attachment.
type SsoadminPermissionsBoundaryAttachment struct {
	Name      string
	Args      SsoadminPermissionsBoundaryAttachmentArgs
	state     *ssoadminPermissionsBoundaryAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsoadminPermissionsBoundaryAttachment].
func (spba *SsoadminPermissionsBoundaryAttachment) Type() string {
	return "aws_ssoadmin_permissions_boundary_attachment"
}

// LocalName returns the local name for [SsoadminPermissionsBoundaryAttachment].
func (spba *SsoadminPermissionsBoundaryAttachment) LocalName() string {
	return spba.Name
}

// Configuration returns the configuration (args) for [SsoadminPermissionsBoundaryAttachment].
func (spba *SsoadminPermissionsBoundaryAttachment) Configuration() interface{} {
	return spba.Args
}

// DependOn is used for other resources to depend on [SsoadminPermissionsBoundaryAttachment].
func (spba *SsoadminPermissionsBoundaryAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(spba)
}

// Dependencies returns the list of resources [SsoadminPermissionsBoundaryAttachment] depends_on.
func (spba *SsoadminPermissionsBoundaryAttachment) Dependencies() terra.Dependencies {
	return spba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsoadminPermissionsBoundaryAttachment].
func (spba *SsoadminPermissionsBoundaryAttachment) LifecycleManagement() *terra.Lifecycle {
	return spba.Lifecycle
}

// Attributes returns the attributes for [SsoadminPermissionsBoundaryAttachment].
func (spba *SsoadminPermissionsBoundaryAttachment) Attributes() ssoadminPermissionsBoundaryAttachmentAttributes {
	return ssoadminPermissionsBoundaryAttachmentAttributes{ref: terra.ReferenceResource(spba)}
}

// ImportState imports the given attribute values into [SsoadminPermissionsBoundaryAttachment]'s state.
func (spba *SsoadminPermissionsBoundaryAttachment) ImportState(av io.Reader) error {
	spba.state = &ssoadminPermissionsBoundaryAttachmentState{}
	if err := json.NewDecoder(av).Decode(spba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spba.Type(), spba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsoadminPermissionsBoundaryAttachment] has state.
func (spba *SsoadminPermissionsBoundaryAttachment) State() (*ssoadminPermissionsBoundaryAttachmentState, bool) {
	return spba.state, spba.state != nil
}

// StateMust returns the state for [SsoadminPermissionsBoundaryAttachment]. Panics if the state is nil.
func (spba *SsoadminPermissionsBoundaryAttachment) StateMust() *ssoadminPermissionsBoundaryAttachmentState {
	if spba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spba.Type(), spba.LocalName()))
	}
	return spba.state
}

// SsoadminPermissionsBoundaryAttachmentArgs contains the configurations for aws_ssoadmin_permissions_boundary_attachment.
type SsoadminPermissionsBoundaryAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceArn: string, required
	InstanceArn terra.StringValue `hcl:"instance_arn,attr" validate:"required"`
	// PermissionSetArn: string, required
	PermissionSetArn terra.StringValue `hcl:"permission_set_arn,attr" validate:"required"`
	// PermissionsBoundary: required
	PermissionsBoundary *ssoadminpermissionsboundaryattachment.PermissionsBoundary `hcl:"permissions_boundary,block" validate:"required"`
}
type ssoadminPermissionsBoundaryAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssoadmin_permissions_boundary_attachment.
func (spba ssoadminPermissionsBoundaryAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(spba.ref.Append("id"))
}

// InstanceArn returns a reference to field instance_arn of aws_ssoadmin_permissions_boundary_attachment.
func (spba ssoadminPermissionsBoundaryAttachmentAttributes) InstanceArn() terra.StringValue {
	return terra.ReferenceAsString(spba.ref.Append("instance_arn"))
}

// PermissionSetArn returns a reference to field permission_set_arn of aws_ssoadmin_permissions_boundary_attachment.
func (spba ssoadminPermissionsBoundaryAttachmentAttributes) PermissionSetArn() terra.StringValue {
	return terra.ReferenceAsString(spba.ref.Append("permission_set_arn"))
}

func (spba ssoadminPermissionsBoundaryAttachmentAttributes) PermissionsBoundary() terra.ListValue[ssoadminpermissionsboundaryattachment.PermissionsBoundaryAttributes] {
	return terra.ReferenceAsList[ssoadminpermissionsboundaryattachment.PermissionsBoundaryAttributes](spba.ref.Append("permissions_boundary"))
}

type ssoadminPermissionsBoundaryAttachmentState struct {
	Id                  string                                                           `json:"id"`
	InstanceArn         string                                                           `json:"instance_arn"`
	PermissionSetArn    string                                                           `json:"permission_set_arn"`
	PermissionsBoundary []ssoadminpermissionsboundaryattachment.PermissionsBoundaryState `json:"permissions_boundary"`
}
