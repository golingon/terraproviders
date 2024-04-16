// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_region_disk_resource_policy_attachment

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_compute_region_disk_resource_policy_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRegionDiskResourcePolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrdrpa *Resource) Type() string {
	return "google_compute_region_disk_resource_policy_attachment"
}

// LocalName returns the local name for [Resource].
func (gcrdrpa *Resource) LocalName() string {
	return gcrdrpa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrdrpa *Resource) Configuration() interface{} {
	return gcrdrpa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrdrpa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrdrpa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrdrpa *Resource) Dependencies() terra.Dependencies {
	return gcrdrpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrdrpa *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrdrpa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrdrpa *Resource) Attributes() googleComputeRegionDiskResourcePolicyAttachmentAttributes {
	return googleComputeRegionDiskResourcePolicyAttachmentAttributes{ref: terra.ReferenceResource(gcrdrpa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrdrpa *Resource) ImportState(state io.Reader) error {
	gcrdrpa.state = &googleComputeRegionDiskResourcePolicyAttachmentState{}
	if err := json.NewDecoder(state).Decode(gcrdrpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrdrpa.Type(), gcrdrpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrdrpa *Resource) State() (*googleComputeRegionDiskResourcePolicyAttachmentState, bool) {
	return gcrdrpa.state, gcrdrpa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrdrpa *Resource) StateMust() *googleComputeRegionDiskResourcePolicyAttachmentState {
	if gcrdrpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrdrpa.Type(), gcrdrpa.LocalName()))
	}
	return gcrdrpa.state
}

// Args contains the configurations for google_compute_region_disk_resource_policy_attachment.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeRegionDiskResourcePolicyAttachmentAttributes struct {
	ref terra.Reference
}

// Disk returns a reference to field disk of google_compute_region_disk_resource_policy_attachment.
func (gcrdrpa googleComputeRegionDiskResourcePolicyAttachmentAttributes) Disk() terra.StringValue {
	return terra.ReferenceAsString(gcrdrpa.ref.Append("disk"))
}

// Id returns a reference to field id of google_compute_region_disk_resource_policy_attachment.
func (gcrdrpa googleComputeRegionDiskResourcePolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrdrpa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_disk_resource_policy_attachment.
func (gcrdrpa googleComputeRegionDiskResourcePolicyAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrdrpa.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_disk_resource_policy_attachment.
func (gcrdrpa googleComputeRegionDiskResourcePolicyAttachmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrdrpa.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk_resource_policy_attachment.
func (gcrdrpa googleComputeRegionDiskResourcePolicyAttachmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrdrpa.ref.Append("region"))
}

func (gcrdrpa googleComputeRegionDiskResourcePolicyAttachmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrdrpa.ref.Append("timeouts"))
}

type googleComputeRegionDiskResourcePolicyAttachmentState struct {
	Disk     string         `json:"disk"`
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Project  string         `json:"project"`
	Region   string         `json:"region"`
	Timeouts *TimeoutsState `json:"timeouts"`
}
