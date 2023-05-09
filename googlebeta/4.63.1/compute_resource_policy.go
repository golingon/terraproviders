// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeresourcepolicy "github.com/golingon/terraproviders/googlebeta/4.63.1/computeresourcepolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeResourcePolicy creates a new instance of [ComputeResourcePolicy].
func NewComputeResourcePolicy(name string, args ComputeResourcePolicyArgs) *ComputeResourcePolicy {
	return &ComputeResourcePolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeResourcePolicy)(nil)

// ComputeResourcePolicy represents the Terraform resource google_compute_resource_policy.
type ComputeResourcePolicy struct {
	Name      string
	Args      ComputeResourcePolicyArgs
	state     *computeResourcePolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeResourcePolicy].
func (crp *ComputeResourcePolicy) Type() string {
	return "google_compute_resource_policy"
}

// LocalName returns the local name for [ComputeResourcePolicy].
func (crp *ComputeResourcePolicy) LocalName() string {
	return crp.Name
}

// Configuration returns the configuration (args) for [ComputeResourcePolicy].
func (crp *ComputeResourcePolicy) Configuration() interface{} {
	return crp.Args
}

// DependOn is used for other resources to depend on [ComputeResourcePolicy].
func (crp *ComputeResourcePolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crp)
}

// Dependencies returns the list of resources [ComputeResourcePolicy] depends_on.
func (crp *ComputeResourcePolicy) Dependencies() terra.Dependencies {
	return crp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeResourcePolicy].
func (crp *ComputeResourcePolicy) LifecycleManagement() *terra.Lifecycle {
	return crp.Lifecycle
}

// Attributes returns the attributes for [ComputeResourcePolicy].
func (crp *ComputeResourcePolicy) Attributes() computeResourcePolicyAttributes {
	return computeResourcePolicyAttributes{ref: terra.ReferenceResource(crp)}
}

// ImportState imports the given attribute values into [ComputeResourcePolicy]'s state.
func (crp *ComputeResourcePolicy) ImportState(av io.Reader) error {
	crp.state = &computeResourcePolicyState{}
	if err := json.NewDecoder(av).Decode(crp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crp.Type(), crp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeResourcePolicy] has state.
func (crp *ComputeResourcePolicy) State() (*computeResourcePolicyState, bool) {
	return crp.state, crp.state != nil
}

// StateMust returns the state for [ComputeResourcePolicy]. Panics if the state is nil.
func (crp *ComputeResourcePolicy) StateMust() *computeResourcePolicyState {
	if crp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crp.Type(), crp.LocalName()))
	}
	return crp.state
}

// ComputeResourcePolicyArgs contains the configurations for google_compute_resource_policy.
type ComputeResourcePolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// GroupPlacementPolicy: optional
	GroupPlacementPolicy *computeresourcepolicy.GroupPlacementPolicy `hcl:"group_placement_policy,block"`
	// InstanceSchedulePolicy: optional
	InstanceSchedulePolicy *computeresourcepolicy.InstanceSchedulePolicy `hcl:"instance_schedule_policy,block"`
	// SnapshotSchedulePolicy: optional
	SnapshotSchedulePolicy *computeresourcepolicy.SnapshotSchedulePolicy `hcl:"snapshot_schedule_policy,block"`
	// Timeouts: optional
	Timeouts *computeresourcepolicy.Timeouts `hcl:"timeouts,block"`
}
type computeResourcePolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_resource_policy.
func (crp computeResourcePolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_resource_policy.
func (crp computeResourcePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_resource_policy.
func (crp computeResourcePolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_resource_policy.
func (crp computeResourcePolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_resource_policy.
func (crp computeResourcePolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_resource_policy.
func (crp computeResourcePolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("self_link"))
}

func (crp computeResourcePolicyAttributes) GroupPlacementPolicy() terra.ListValue[computeresourcepolicy.GroupPlacementPolicyAttributes] {
	return terra.ReferenceAsList[computeresourcepolicy.GroupPlacementPolicyAttributes](crp.ref.Append("group_placement_policy"))
}

func (crp computeResourcePolicyAttributes) InstanceSchedulePolicy() terra.ListValue[computeresourcepolicy.InstanceSchedulePolicyAttributes] {
	return terra.ReferenceAsList[computeresourcepolicy.InstanceSchedulePolicyAttributes](crp.ref.Append("instance_schedule_policy"))
}

func (crp computeResourcePolicyAttributes) SnapshotSchedulePolicy() terra.ListValue[computeresourcepolicy.SnapshotSchedulePolicyAttributes] {
	return terra.ReferenceAsList[computeresourcepolicy.SnapshotSchedulePolicyAttributes](crp.ref.Append("snapshot_schedule_policy"))
}

func (crp computeResourcePolicyAttributes) Timeouts() computeresourcepolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeresourcepolicy.TimeoutsAttributes](crp.ref.Append("timeouts"))
}

type computeResourcePolicyState struct {
	Description            string                                              `json:"description"`
	Id                     string                                              `json:"id"`
	Name                   string                                              `json:"name"`
	Project                string                                              `json:"project"`
	Region                 string                                              `json:"region"`
	SelfLink               string                                              `json:"self_link"`
	GroupPlacementPolicy   []computeresourcepolicy.GroupPlacementPolicyState   `json:"group_placement_policy"`
	InstanceSchedulePolicy []computeresourcepolicy.InstanceSchedulePolicyState `json:"instance_schedule_policy"`
	SnapshotSchedulePolicy []computeresourcepolicy.SnapshotSchedulePolicyState `json:"snapshot_schedule_policy"`
	Timeouts               *computeresourcepolicy.TimeoutsState                `json:"timeouts"`
}
