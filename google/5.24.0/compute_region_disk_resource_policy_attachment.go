// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computeregiondiskresourcepolicyattachment "github.com/golingon/terraproviders/google/5.24.0/computeregiondiskresourcepolicyattachment"
	"io"
)

// NewComputeRegionDiskResourcePolicyAttachment creates a new instance of [ComputeRegionDiskResourcePolicyAttachment].
func NewComputeRegionDiskResourcePolicyAttachment(name string, args ComputeRegionDiskResourcePolicyAttachmentArgs) *ComputeRegionDiskResourcePolicyAttachment {
	return &ComputeRegionDiskResourcePolicyAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionDiskResourcePolicyAttachment)(nil)

// ComputeRegionDiskResourcePolicyAttachment represents the Terraform resource google_compute_region_disk_resource_policy_attachment.
type ComputeRegionDiskResourcePolicyAttachment struct {
	Name      string
	Args      ComputeRegionDiskResourcePolicyAttachmentArgs
	state     *computeRegionDiskResourcePolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionDiskResourcePolicyAttachment].
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Type() string {
	return "google_compute_region_disk_resource_policy_attachment"
}

// LocalName returns the local name for [ComputeRegionDiskResourcePolicyAttachment].
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) LocalName() string {
	return crdrpa.Name
}

// Configuration returns the configuration (args) for [ComputeRegionDiskResourcePolicyAttachment].
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Configuration() interface{} {
	return crdrpa.Args
}

// DependOn is used for other resources to depend on [ComputeRegionDiskResourcePolicyAttachment].
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(crdrpa)
}

// Dependencies returns the list of resources [ComputeRegionDiskResourcePolicyAttachment] depends_on.
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Dependencies() terra.Dependencies {
	return crdrpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionDiskResourcePolicyAttachment].
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) LifecycleManagement() *terra.Lifecycle {
	return crdrpa.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionDiskResourcePolicyAttachment].
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Attributes() computeRegionDiskResourcePolicyAttachmentAttributes {
	return computeRegionDiskResourcePolicyAttachmentAttributes{ref: terra.ReferenceResource(crdrpa)}
}

// ImportState imports the given attribute values into [ComputeRegionDiskResourcePolicyAttachment]'s state.
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) ImportState(av io.Reader) error {
	crdrpa.state = &computeRegionDiskResourcePolicyAttachmentState{}
	if err := json.NewDecoder(av).Decode(crdrpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crdrpa.Type(), crdrpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionDiskResourcePolicyAttachment] has state.
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) State() (*computeRegionDiskResourcePolicyAttachmentState, bool) {
	return crdrpa.state, crdrpa.state != nil
}

// StateMust returns the state for [ComputeRegionDiskResourcePolicyAttachment]. Panics if the state is nil.
func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) StateMust() *computeRegionDiskResourcePolicyAttachmentState {
	if crdrpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crdrpa.Type(), crdrpa.LocalName()))
	}
	return crdrpa.state
}

// ComputeRegionDiskResourcePolicyAttachmentArgs contains the configurations for google_compute_region_disk_resource_policy_attachment.
type ComputeRegionDiskResourcePolicyAttachmentArgs struct {
	// Disk: string, required
	Disk terra.StringValue `hcl:"disk,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Timeouts: optional
	Timeouts *computeregiondiskresourcepolicyattachment.Timeouts `hcl:"timeouts,block"`
}
type computeRegionDiskResourcePolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Disk returns a reference to field disk of google_compute_region_disk_resource_policy_attachment.
func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(crdrpa.ref.Append("disk"))
}

// Id returns a reference to field id of google_compute_region_disk_resource_policy_attachment.
func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crdrpa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_disk_resource_policy_attachment.
func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crdrpa.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_disk_resource_policy_attachment.
func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crdrpa.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk_resource_policy_attachment.
func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crdrpa.ref.Append("region"))
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Timeouts() computeregiondiskresourcepolicyattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregiondiskresourcepolicyattachment.TimeoutsAttributes](crdrpa.ref.Append("timeouts"))
}

type computeRegionDiskResourcePolicyAttachmentState struct {
	Disk     string                                                   `json:"disk"`
	Id       string                                                   `json:"id"`
	Name     string                                                   `json:"name"`
	Project  string                                                   `json:"project"`
	Region   string                                                   `json:"region"`
	Timeouts *computeregiondiskresourcepolicyattachment.TimeoutsState `json:"timeouts"`
}
