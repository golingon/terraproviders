// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigquerycapacitycommitment "github.com/golingon/terraproviders/google/4.72.1/bigquerycapacitycommitment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryCapacityCommitment creates a new instance of [BigqueryCapacityCommitment].
func NewBigqueryCapacityCommitment(name string, args BigqueryCapacityCommitmentArgs) *BigqueryCapacityCommitment {
	return &BigqueryCapacityCommitment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryCapacityCommitment)(nil)

// BigqueryCapacityCommitment represents the Terraform resource google_bigquery_capacity_commitment.
type BigqueryCapacityCommitment struct {
	Name      string
	Args      BigqueryCapacityCommitmentArgs
	state     *bigqueryCapacityCommitmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryCapacityCommitment].
func (bcc *BigqueryCapacityCommitment) Type() string {
	return "google_bigquery_capacity_commitment"
}

// LocalName returns the local name for [BigqueryCapacityCommitment].
func (bcc *BigqueryCapacityCommitment) LocalName() string {
	return bcc.Name
}

// Configuration returns the configuration (args) for [BigqueryCapacityCommitment].
func (bcc *BigqueryCapacityCommitment) Configuration() interface{} {
	return bcc.Args
}

// DependOn is used for other resources to depend on [BigqueryCapacityCommitment].
func (bcc *BigqueryCapacityCommitment) DependOn() terra.Reference {
	return terra.ReferenceResource(bcc)
}

// Dependencies returns the list of resources [BigqueryCapacityCommitment] depends_on.
func (bcc *BigqueryCapacityCommitment) Dependencies() terra.Dependencies {
	return bcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryCapacityCommitment].
func (bcc *BigqueryCapacityCommitment) LifecycleManagement() *terra.Lifecycle {
	return bcc.Lifecycle
}

// Attributes returns the attributes for [BigqueryCapacityCommitment].
func (bcc *BigqueryCapacityCommitment) Attributes() bigqueryCapacityCommitmentAttributes {
	return bigqueryCapacityCommitmentAttributes{ref: terra.ReferenceResource(bcc)}
}

// ImportState imports the given attribute values into [BigqueryCapacityCommitment]'s state.
func (bcc *BigqueryCapacityCommitment) ImportState(av io.Reader) error {
	bcc.state = &bigqueryCapacityCommitmentState{}
	if err := json.NewDecoder(av).Decode(bcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcc.Type(), bcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryCapacityCommitment] has state.
func (bcc *BigqueryCapacityCommitment) State() (*bigqueryCapacityCommitmentState, bool) {
	return bcc.state, bcc.state != nil
}

// StateMust returns the state for [BigqueryCapacityCommitment]. Panics if the state is nil.
func (bcc *BigqueryCapacityCommitment) StateMust() *bigqueryCapacityCommitmentState {
	if bcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcc.Type(), bcc.LocalName()))
	}
	return bcc.state
}

// BigqueryCapacityCommitmentArgs contains the configurations for google_bigquery_capacity_commitment.
type BigqueryCapacityCommitmentArgs struct {
	// CapacityCommitmentId: string, optional
	CapacityCommitmentId terra.StringValue `hcl:"capacity_commitment_id,attr"`
	// Edition: string, optional
	Edition terra.StringValue `hcl:"edition,attr"`
	// EnforceSingleAdminProjectPerOrg: string, optional
	EnforceSingleAdminProjectPerOrg terra.StringValue `hcl:"enforce_single_admin_project_per_org,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Plan: string, required
	Plan terra.StringValue `hcl:"plan,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RenewalPlan: string, optional
	RenewalPlan terra.StringValue `hcl:"renewal_plan,attr"`
	// SlotCount: number, required
	SlotCount terra.NumberValue `hcl:"slot_count,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *bigquerycapacitycommitment.Timeouts `hcl:"timeouts,block"`
}
type bigqueryCapacityCommitmentAttributes struct {
	ref terra.Reference
}

// CapacityCommitmentId returns a reference to field capacity_commitment_id of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) CapacityCommitmentId() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("capacity_commitment_id"))
}

// CommitmentEndTime returns a reference to field commitment_end_time of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) CommitmentEndTime() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("commitment_end_time"))
}

// CommitmentStartTime returns a reference to field commitment_start_time of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) CommitmentStartTime() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("commitment_start_time"))
}

// Edition returns a reference to field edition of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("edition"))
}

// EnforceSingleAdminProjectPerOrg returns a reference to field enforce_single_admin_project_per_org of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) EnforceSingleAdminProjectPerOrg() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("enforce_single_admin_project_per_org"))
}

// Id returns a reference to field id of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("name"))
}

// Plan returns a reference to field plan of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) Plan() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("plan"))
}

// Project returns a reference to field project of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("project"))
}

// RenewalPlan returns a reference to field renewal_plan of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) RenewalPlan() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("renewal_plan"))
}

// SlotCount returns a reference to field slot_count of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) SlotCount() terra.NumberValue {
	return terra.ReferenceAsNumber(bcc.ref.Append("slot_count"))
}

// State returns a reference to field state of google_bigquery_capacity_commitment.
func (bcc bigqueryCapacityCommitmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bcc.ref.Append("state"))
}

func (bcc bigqueryCapacityCommitmentAttributes) Timeouts() bigquerycapacitycommitment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigquerycapacitycommitment.TimeoutsAttributes](bcc.ref.Append("timeouts"))
}

type bigqueryCapacityCommitmentState struct {
	CapacityCommitmentId            string                                    `json:"capacity_commitment_id"`
	CommitmentEndTime               string                                    `json:"commitment_end_time"`
	CommitmentStartTime             string                                    `json:"commitment_start_time"`
	Edition                         string                                    `json:"edition"`
	EnforceSingleAdminProjectPerOrg string                                    `json:"enforce_single_admin_project_per_org"`
	Id                              string                                    `json:"id"`
	Location                        string                                    `json:"location"`
	Name                            string                                    `json:"name"`
	Plan                            string                                    `json:"plan"`
	Project                         string                                    `json:"project"`
	RenewalPlan                     string                                    `json:"renewal_plan"`
	SlotCount                       float64                                   `json:"slot_count"`
	State                           string                                    `json:"state"`
	Timeouts                        *bigquerycapacitycommitment.TimeoutsState `json:"timeouts"`
}
