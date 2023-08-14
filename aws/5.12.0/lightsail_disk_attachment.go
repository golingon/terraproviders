// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailDiskAttachment creates a new instance of [LightsailDiskAttachment].
func NewLightsailDiskAttachment(name string, args LightsailDiskAttachmentArgs) *LightsailDiskAttachment {
	return &LightsailDiskAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailDiskAttachment)(nil)

// LightsailDiskAttachment represents the Terraform resource aws_lightsail_disk_attachment.
type LightsailDiskAttachment struct {
	Name      string
	Args      LightsailDiskAttachmentArgs
	state     *lightsailDiskAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailDiskAttachment].
func (lda *LightsailDiskAttachment) Type() string {
	return "aws_lightsail_disk_attachment"
}

// LocalName returns the local name for [LightsailDiskAttachment].
func (lda *LightsailDiskAttachment) LocalName() string {
	return lda.Name
}

// Configuration returns the configuration (args) for [LightsailDiskAttachment].
func (lda *LightsailDiskAttachment) Configuration() interface{} {
	return lda.Args
}

// DependOn is used for other resources to depend on [LightsailDiskAttachment].
func (lda *LightsailDiskAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(lda)
}

// Dependencies returns the list of resources [LightsailDiskAttachment] depends_on.
func (lda *LightsailDiskAttachment) Dependencies() terra.Dependencies {
	return lda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailDiskAttachment].
func (lda *LightsailDiskAttachment) LifecycleManagement() *terra.Lifecycle {
	return lda.Lifecycle
}

// Attributes returns the attributes for [LightsailDiskAttachment].
func (lda *LightsailDiskAttachment) Attributes() lightsailDiskAttachmentAttributes {
	return lightsailDiskAttachmentAttributes{ref: terra.ReferenceResource(lda)}
}

// ImportState imports the given attribute values into [LightsailDiskAttachment]'s state.
func (lda *LightsailDiskAttachment) ImportState(av io.Reader) error {
	lda.state = &lightsailDiskAttachmentState{}
	if err := json.NewDecoder(av).Decode(lda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lda.Type(), lda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailDiskAttachment] has state.
func (lda *LightsailDiskAttachment) State() (*lightsailDiskAttachmentState, bool) {
	return lda.state, lda.state != nil
}

// StateMust returns the state for [LightsailDiskAttachment]. Panics if the state is nil.
func (lda *LightsailDiskAttachment) StateMust() *lightsailDiskAttachmentState {
	if lda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lda.Type(), lda.LocalName()))
	}
	return lda.state
}

// LightsailDiskAttachmentArgs contains the configurations for aws_lightsail_disk_attachment.
type LightsailDiskAttachmentArgs struct {
	// DiskName: string, required
	DiskName terra.StringValue `hcl:"disk_name,attr" validate:"required"`
	// DiskPath: string, required
	DiskPath terra.StringValue `hcl:"disk_path,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
}
type lightsailDiskAttachmentAttributes struct {
	ref terra.Reference
}

// DiskName returns a reference to field disk_name of aws_lightsail_disk_attachment.
func (lda lightsailDiskAttachmentAttributes) DiskName() terra.StringValue {
	return terra.ReferenceAsString(lda.ref.Append("disk_name"))
}

// DiskPath returns a reference to field disk_path of aws_lightsail_disk_attachment.
func (lda lightsailDiskAttachmentAttributes) DiskPath() terra.StringValue {
	return terra.ReferenceAsString(lda.ref.Append("disk_path"))
}

// Id returns a reference to field id of aws_lightsail_disk_attachment.
func (lda lightsailDiskAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lda.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of aws_lightsail_disk_attachment.
func (lda lightsailDiskAttachmentAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(lda.ref.Append("instance_name"))
}

type lightsailDiskAttachmentState struct {
	DiskName     string `json:"disk_name"`
	DiskPath     string `json:"disk_path"`
	Id           string `json:"id"`
	InstanceName string `json:"instance_name"`
}