// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigqueryreservationassignment "github.com/golingon/terraproviders/googlebeta/4.71.0/bigqueryreservationassignment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryReservationAssignment creates a new instance of [BigqueryReservationAssignment].
func NewBigqueryReservationAssignment(name string, args BigqueryReservationAssignmentArgs) *BigqueryReservationAssignment {
	return &BigqueryReservationAssignment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryReservationAssignment)(nil)

// BigqueryReservationAssignment represents the Terraform resource google_bigquery_reservation_assignment.
type BigqueryReservationAssignment struct {
	Name      string
	Args      BigqueryReservationAssignmentArgs
	state     *bigqueryReservationAssignmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryReservationAssignment].
func (bra *BigqueryReservationAssignment) Type() string {
	return "google_bigquery_reservation_assignment"
}

// LocalName returns the local name for [BigqueryReservationAssignment].
func (bra *BigqueryReservationAssignment) LocalName() string {
	return bra.Name
}

// Configuration returns the configuration (args) for [BigqueryReservationAssignment].
func (bra *BigqueryReservationAssignment) Configuration() interface{} {
	return bra.Args
}

// DependOn is used for other resources to depend on [BigqueryReservationAssignment].
func (bra *BigqueryReservationAssignment) DependOn() terra.Reference {
	return terra.ReferenceResource(bra)
}

// Dependencies returns the list of resources [BigqueryReservationAssignment] depends_on.
func (bra *BigqueryReservationAssignment) Dependencies() terra.Dependencies {
	return bra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryReservationAssignment].
func (bra *BigqueryReservationAssignment) LifecycleManagement() *terra.Lifecycle {
	return bra.Lifecycle
}

// Attributes returns the attributes for [BigqueryReservationAssignment].
func (bra *BigqueryReservationAssignment) Attributes() bigqueryReservationAssignmentAttributes {
	return bigqueryReservationAssignmentAttributes{ref: terra.ReferenceResource(bra)}
}

// ImportState imports the given attribute values into [BigqueryReservationAssignment]'s state.
func (bra *BigqueryReservationAssignment) ImportState(av io.Reader) error {
	bra.state = &bigqueryReservationAssignmentState{}
	if err := json.NewDecoder(av).Decode(bra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bra.Type(), bra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryReservationAssignment] has state.
func (bra *BigqueryReservationAssignment) State() (*bigqueryReservationAssignmentState, bool) {
	return bra.state, bra.state != nil
}

// StateMust returns the state for [BigqueryReservationAssignment]. Panics if the state is nil.
func (bra *BigqueryReservationAssignment) StateMust() *bigqueryReservationAssignmentState {
	if bra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bra.Type(), bra.LocalName()))
	}
	return bra.state
}

// BigqueryReservationAssignmentArgs contains the configurations for google_bigquery_reservation_assignment.
type BigqueryReservationAssignmentArgs struct {
	// Assignee: string, required
	Assignee terra.StringValue `hcl:"assignee,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobType: string, required
	JobType terra.StringValue `hcl:"job_type,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Reservation: string, required
	Reservation terra.StringValue `hcl:"reservation,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *bigqueryreservationassignment.Timeouts `hcl:"timeouts,block"`
}
type bigqueryReservationAssignmentAttributes struct {
	ref terra.Reference
}

// Assignee returns a reference to field assignee of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) Assignee() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("assignee"))
}

// Id returns a reference to field id of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("id"))
}

// JobType returns a reference to field job_type of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) JobType() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("job_type"))
}

// Location returns a reference to field location of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("location"))
}

// Name returns a reference to field name of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("name"))
}

// Project returns a reference to field project of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("project"))
}

// Reservation returns a reference to field reservation of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) Reservation() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("reservation"))
}

// State returns a reference to field state of google_bigquery_reservation_assignment.
func (bra bigqueryReservationAssignmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bra.ref.Append("state"))
}

func (bra bigqueryReservationAssignmentAttributes) Timeouts() bigqueryreservationassignment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryreservationassignment.TimeoutsAttributes](bra.ref.Append("timeouts"))
}

type bigqueryReservationAssignmentState struct {
	Assignee    string                                       `json:"assignee"`
	Id          string                                       `json:"id"`
	JobType     string                                       `json:"job_type"`
	Location    string                                       `json:"location"`
	Name        string                                       `json:"name"`
	Project     string                                       `json:"project"`
	Reservation string                                       `json:"reservation"`
	State       string                                       `json:"state"`
	Timeouts    *bigqueryreservationassignment.TimeoutsState `json:"timeouts"`
}
