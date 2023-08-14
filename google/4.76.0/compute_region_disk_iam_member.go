// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregiondiskiammember "github.com/golingon/terraproviders/google/4.76.0/computeregiondiskiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionDiskIamMember creates a new instance of [ComputeRegionDiskIamMember].
func NewComputeRegionDiskIamMember(name string, args ComputeRegionDiskIamMemberArgs) *ComputeRegionDiskIamMember {
	return &ComputeRegionDiskIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionDiskIamMember)(nil)

// ComputeRegionDiskIamMember represents the Terraform resource google_compute_region_disk_iam_member.
type ComputeRegionDiskIamMember struct {
	Name      string
	Args      ComputeRegionDiskIamMemberArgs
	state     *computeRegionDiskIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionDiskIamMember].
func (crdim *ComputeRegionDiskIamMember) Type() string {
	return "google_compute_region_disk_iam_member"
}

// LocalName returns the local name for [ComputeRegionDiskIamMember].
func (crdim *ComputeRegionDiskIamMember) LocalName() string {
	return crdim.Name
}

// Configuration returns the configuration (args) for [ComputeRegionDiskIamMember].
func (crdim *ComputeRegionDiskIamMember) Configuration() interface{} {
	return crdim.Args
}

// DependOn is used for other resources to depend on [ComputeRegionDiskIamMember].
func (crdim *ComputeRegionDiskIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(crdim)
}

// Dependencies returns the list of resources [ComputeRegionDiskIamMember] depends_on.
func (crdim *ComputeRegionDiskIamMember) Dependencies() terra.Dependencies {
	return crdim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionDiskIamMember].
func (crdim *ComputeRegionDiskIamMember) LifecycleManagement() *terra.Lifecycle {
	return crdim.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionDiskIamMember].
func (crdim *ComputeRegionDiskIamMember) Attributes() computeRegionDiskIamMemberAttributes {
	return computeRegionDiskIamMemberAttributes{ref: terra.ReferenceResource(crdim)}
}

// ImportState imports the given attribute values into [ComputeRegionDiskIamMember]'s state.
func (crdim *ComputeRegionDiskIamMember) ImportState(av io.Reader) error {
	crdim.state = &computeRegionDiskIamMemberState{}
	if err := json.NewDecoder(av).Decode(crdim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crdim.Type(), crdim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionDiskIamMember] has state.
func (crdim *ComputeRegionDiskIamMember) State() (*computeRegionDiskIamMemberState, bool) {
	return crdim.state, crdim.state != nil
}

// StateMust returns the state for [ComputeRegionDiskIamMember]. Panics if the state is nil.
func (crdim *ComputeRegionDiskIamMember) StateMust() *computeRegionDiskIamMemberState {
	if crdim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crdim.Type(), crdim.LocalName()))
	}
	return crdim.state
}

// ComputeRegionDiskIamMemberArgs contains the configurations for google_compute_region_disk_iam_member.
type ComputeRegionDiskIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *computeregiondiskiammember.Condition `hcl:"condition,block"`
}
type computeRegionDiskIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("id"))
}

// Member returns a reference to field member of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("member"))
}

// Name returns a reference to field name of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_region_disk_iam_member.
func (crdim computeRegionDiskIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(crdim.ref.Append("role"))
}

func (crdim computeRegionDiskIamMemberAttributes) Condition() terra.ListValue[computeregiondiskiammember.ConditionAttributes] {
	return terra.ReferenceAsList[computeregiondiskiammember.ConditionAttributes](crdim.ref.Append("condition"))
}

type computeRegionDiskIamMemberState struct {
	Etag      string                                      `json:"etag"`
	Id        string                                      `json:"id"`
	Member    string                                      `json:"member"`
	Name      string                                      `json:"name"`
	Project   string                                      `json:"project"`
	Region    string                                      `json:"region"`
	Role      string                                      `json:"role"`
	Condition []computeregiondiskiammember.ConditionState `json:"condition"`
}