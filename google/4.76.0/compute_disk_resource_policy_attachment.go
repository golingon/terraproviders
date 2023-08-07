// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computediskresourcepolicyattachment "github.com/golingon/terraproviders/google/4.76.0/computediskresourcepolicyattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeDiskResourcePolicyAttachment creates a new instance of [ComputeDiskResourcePolicyAttachment].
func NewComputeDiskResourcePolicyAttachment(name string, args ComputeDiskResourcePolicyAttachmentArgs) *ComputeDiskResourcePolicyAttachment {
	return &ComputeDiskResourcePolicyAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeDiskResourcePolicyAttachment)(nil)

// ComputeDiskResourcePolicyAttachment represents the Terraform resource google_compute_disk_resource_policy_attachment.
type ComputeDiskResourcePolicyAttachment struct {
	Name      string
	Args      ComputeDiskResourcePolicyAttachmentArgs
	state     *computeDiskResourcePolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeDiskResourcePolicyAttachment].
func (cdrpa *ComputeDiskResourcePolicyAttachment) Type() string {
	return "google_compute_disk_resource_policy_attachment"
}

// LocalName returns the local name for [ComputeDiskResourcePolicyAttachment].
func (cdrpa *ComputeDiskResourcePolicyAttachment) LocalName() string {
	return cdrpa.Name
}

// Configuration returns the configuration (args) for [ComputeDiskResourcePolicyAttachment].
func (cdrpa *ComputeDiskResourcePolicyAttachment) Configuration() interface{} {
	return cdrpa.Args
}

// DependOn is used for other resources to depend on [ComputeDiskResourcePolicyAttachment].
func (cdrpa *ComputeDiskResourcePolicyAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(cdrpa)
}

// Dependencies returns the list of resources [ComputeDiskResourcePolicyAttachment] depends_on.
func (cdrpa *ComputeDiskResourcePolicyAttachment) Dependencies() terra.Dependencies {
	return cdrpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeDiskResourcePolicyAttachment].
func (cdrpa *ComputeDiskResourcePolicyAttachment) LifecycleManagement() *terra.Lifecycle {
	return cdrpa.Lifecycle
}

// Attributes returns the attributes for [ComputeDiskResourcePolicyAttachment].
func (cdrpa *ComputeDiskResourcePolicyAttachment) Attributes() computeDiskResourcePolicyAttachmentAttributes {
	return computeDiskResourcePolicyAttachmentAttributes{ref: terra.ReferenceResource(cdrpa)}
}

// ImportState imports the given attribute values into [ComputeDiskResourcePolicyAttachment]'s state.
func (cdrpa *ComputeDiskResourcePolicyAttachment) ImportState(av io.Reader) error {
	cdrpa.state = &computeDiskResourcePolicyAttachmentState{}
	if err := json.NewDecoder(av).Decode(cdrpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdrpa.Type(), cdrpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeDiskResourcePolicyAttachment] has state.
func (cdrpa *ComputeDiskResourcePolicyAttachment) State() (*computeDiskResourcePolicyAttachmentState, bool) {
	return cdrpa.state, cdrpa.state != nil
}

// StateMust returns the state for [ComputeDiskResourcePolicyAttachment]. Panics if the state is nil.
func (cdrpa *ComputeDiskResourcePolicyAttachment) StateMust() *computeDiskResourcePolicyAttachmentState {
	if cdrpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdrpa.Type(), cdrpa.LocalName()))
	}
	return cdrpa.state
}

// ComputeDiskResourcePolicyAttachmentArgs contains the configurations for google_compute_disk_resource_policy_attachment.
type ComputeDiskResourcePolicyAttachmentArgs struct {
	// Disk: string, required
	Disk terra.StringValue `hcl:"disk,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *computediskresourcepolicyattachment.Timeouts `hcl:"timeouts,block"`
}
type computeDiskResourcePolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Disk returns a reference to field disk of google_compute_disk_resource_policy_attachment.
func (cdrpa computeDiskResourcePolicyAttachmentAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(cdrpa.ref.Append("disk"))
}

// Id returns a reference to field id of google_compute_disk_resource_policy_attachment.
func (cdrpa computeDiskResourcePolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdrpa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_disk_resource_policy_attachment.
func (cdrpa computeDiskResourcePolicyAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdrpa.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_disk_resource_policy_attachment.
func (cdrpa computeDiskResourcePolicyAttachmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cdrpa.ref.Append("project"))
}

// Zone returns a reference to field zone of google_compute_disk_resource_policy_attachment.
func (cdrpa computeDiskResourcePolicyAttachmentAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cdrpa.ref.Append("zone"))
}

func (cdrpa computeDiskResourcePolicyAttachmentAttributes) Timeouts() computediskresourcepolicyattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computediskresourcepolicyattachment.TimeoutsAttributes](cdrpa.ref.Append("timeouts"))
}

type computeDiskResourcePolicyAttachmentState struct {
	Disk     string                                             `json:"disk"`
	Id       string                                             `json:"id"`
	Name     string                                             `json:"name"`
	Project  string                                             `json:"project"`
	Zone     string                                             `json:"zone"`
	Timeouts *computediskresourcepolicyattachment.TimeoutsState `json:"timeouts"`
}
