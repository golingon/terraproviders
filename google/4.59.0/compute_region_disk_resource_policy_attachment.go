// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregiondiskresourcepolicyattachment "github.com/golingon/terraproviders/google/4.59.0/computeregiondiskresourcepolicyattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeRegionDiskResourcePolicyAttachment(name string, args ComputeRegionDiskResourcePolicyAttachmentArgs) *ComputeRegionDiskResourcePolicyAttachment {
	return &ComputeRegionDiskResourcePolicyAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionDiskResourcePolicyAttachment)(nil)

type ComputeRegionDiskResourcePolicyAttachment struct {
	Name  string
	Args  ComputeRegionDiskResourcePolicyAttachmentArgs
	state *computeRegionDiskResourcePolicyAttachmentState
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Type() string {
	return "google_compute_region_disk_resource_policy_attachment"
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) LocalName() string {
	return crdrpa.Name
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Configuration() interface{} {
	return crdrpa.Args
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) Attributes() computeRegionDiskResourcePolicyAttachmentAttributes {
	return computeRegionDiskResourcePolicyAttachmentAttributes{ref: terra.ReferenceResource(crdrpa)}
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) ImportState(av io.Reader) error {
	crdrpa.state = &computeRegionDiskResourcePolicyAttachmentState{}
	if err := json.NewDecoder(av).Decode(crdrpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crdrpa.Type(), crdrpa.LocalName(), err)
	}
	return nil
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) State() (*computeRegionDiskResourcePolicyAttachmentState, bool) {
	return crdrpa.state, crdrpa.state != nil
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) StateMust() *computeRegionDiskResourcePolicyAttachmentState {
	if crdrpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crdrpa.Type(), crdrpa.LocalName()))
	}
	return crdrpa.state
}

func (crdrpa *ComputeRegionDiskResourcePolicyAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(crdrpa)
}

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
	// DependsOn contains resources that ComputeRegionDiskResourcePolicyAttachment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeRegionDiskResourcePolicyAttachmentAttributes struct {
	ref terra.Reference
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Disk() terra.StringValue {
	return terra.ReferenceString(crdrpa.ref.Append("disk"))
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crdrpa.ref.Append("id"))
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crdrpa.ref.Append("name"))
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Project() terra.StringValue {
	return terra.ReferenceString(crdrpa.ref.Append("project"))
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Region() terra.StringValue {
	return terra.ReferenceString(crdrpa.ref.Append("region"))
}

func (crdrpa computeRegionDiskResourcePolicyAttachmentAttributes) Timeouts() computeregiondiskresourcepolicyattachment.TimeoutsAttributes {
	return terra.ReferenceSingle[computeregiondiskresourcepolicyattachment.TimeoutsAttributes](crdrpa.ref.Append("timeouts"))
}

type computeRegionDiskResourcePolicyAttachmentState struct {
	Disk     string                                                   `json:"disk"`
	Id       string                                                   `json:"id"`
	Name     string                                                   `json:"name"`
	Project  string                                                   `json:"project"`
	Region   string                                                   `json:"region"`
	Timeouts *computeregiondiskresourcepolicyattachment.TimeoutsState `json:"timeouts"`
}
