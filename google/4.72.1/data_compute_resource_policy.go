// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeresourcepolicy "github.com/golingon/terraproviders/google/4.72.1/datacomputeresourcepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeResourcePolicy creates a new instance of [DataComputeResourcePolicy].
func NewDataComputeResourcePolicy(name string, args DataComputeResourcePolicyArgs) *DataComputeResourcePolicy {
	return &DataComputeResourcePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeResourcePolicy)(nil)

// DataComputeResourcePolicy represents the Terraform data resource google_compute_resource_policy.
type DataComputeResourcePolicy struct {
	Name string
	Args DataComputeResourcePolicyArgs
}

// DataSource returns the Terraform object type for [DataComputeResourcePolicy].
func (crp *DataComputeResourcePolicy) DataSource() string {
	return "google_compute_resource_policy"
}

// LocalName returns the local name for [DataComputeResourcePolicy].
func (crp *DataComputeResourcePolicy) LocalName() string {
	return crp.Name
}

// Configuration returns the configuration (args) for [DataComputeResourcePolicy].
func (crp *DataComputeResourcePolicy) Configuration() interface{} {
	return crp.Args
}

// Attributes returns the attributes for [DataComputeResourcePolicy].
func (crp *DataComputeResourcePolicy) Attributes() dataComputeResourcePolicyAttributes {
	return dataComputeResourcePolicyAttributes{ref: terra.ReferenceDataResource(crp)}
}

// DataComputeResourcePolicyArgs contains the configurations for google_compute_resource_policy.
type DataComputeResourcePolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// GroupPlacementPolicy: min=0
	GroupPlacementPolicy []datacomputeresourcepolicy.GroupPlacementPolicy `hcl:"group_placement_policy,block" validate:"min=0"`
	// InstanceSchedulePolicy: min=0
	InstanceSchedulePolicy []datacomputeresourcepolicy.InstanceSchedulePolicy `hcl:"instance_schedule_policy,block" validate:"min=0"`
	// SnapshotSchedulePolicy: min=0
	SnapshotSchedulePolicy []datacomputeresourcepolicy.SnapshotSchedulePolicy `hcl:"snapshot_schedule_policy,block" validate:"min=0"`
}
type dataComputeResourcePolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_resource_policy.
func (crp dataComputeResourcePolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_resource_policy.
func (crp dataComputeResourcePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_resource_policy.
func (crp dataComputeResourcePolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_resource_policy.
func (crp dataComputeResourcePolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_resource_policy.
func (crp dataComputeResourcePolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_resource_policy.
func (crp dataComputeResourcePolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("self_link"))
}

func (crp dataComputeResourcePolicyAttributes) GroupPlacementPolicy() terra.ListValue[datacomputeresourcepolicy.GroupPlacementPolicyAttributes] {
	return terra.ReferenceAsList[datacomputeresourcepolicy.GroupPlacementPolicyAttributes](crp.ref.Append("group_placement_policy"))
}

func (crp dataComputeResourcePolicyAttributes) InstanceSchedulePolicy() terra.ListValue[datacomputeresourcepolicy.InstanceSchedulePolicyAttributes] {
	return terra.ReferenceAsList[datacomputeresourcepolicy.InstanceSchedulePolicyAttributes](crp.ref.Append("instance_schedule_policy"))
}

func (crp dataComputeResourcePolicyAttributes) SnapshotSchedulePolicy() terra.ListValue[datacomputeresourcepolicy.SnapshotSchedulePolicyAttributes] {
	return terra.ReferenceAsList[datacomputeresourcepolicy.SnapshotSchedulePolicyAttributes](crp.ref.Append("snapshot_schedule_policy"))
}
