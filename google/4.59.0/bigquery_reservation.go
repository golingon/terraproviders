// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryreservation "github.com/golingon/terraproviders/google/4.59.0/bigqueryreservation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryReservation creates a new instance of [BigqueryReservation].
func NewBigqueryReservation(name string, args BigqueryReservationArgs) *BigqueryReservation {
	return &BigqueryReservation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryReservation)(nil)

// BigqueryReservation represents the Terraform resource google_bigquery_reservation.
type BigqueryReservation struct {
	Name      string
	Args      BigqueryReservationArgs
	state     *bigqueryReservationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryReservation].
func (br *BigqueryReservation) Type() string {
	return "google_bigquery_reservation"
}

// LocalName returns the local name for [BigqueryReservation].
func (br *BigqueryReservation) LocalName() string {
	return br.Name
}

// Configuration returns the configuration (args) for [BigqueryReservation].
func (br *BigqueryReservation) Configuration() interface{} {
	return br.Args
}

// DependOn is used for other resources to depend on [BigqueryReservation].
func (br *BigqueryReservation) DependOn() terra.Reference {
	return terra.ReferenceResource(br)
}

// Dependencies returns the list of resources [BigqueryReservation] depends_on.
func (br *BigqueryReservation) Dependencies() terra.Dependencies {
	return br.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryReservation].
func (br *BigqueryReservation) LifecycleManagement() *terra.Lifecycle {
	return br.Lifecycle
}

// Attributes returns the attributes for [BigqueryReservation].
func (br *BigqueryReservation) Attributes() bigqueryReservationAttributes {
	return bigqueryReservationAttributes{ref: terra.ReferenceResource(br)}
}

// ImportState imports the given attribute values into [BigqueryReservation]'s state.
func (br *BigqueryReservation) ImportState(av io.Reader) error {
	br.state = &bigqueryReservationState{}
	if err := json.NewDecoder(av).Decode(br.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", br.Type(), br.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryReservation] has state.
func (br *BigqueryReservation) State() (*bigqueryReservationState, bool) {
	return br.state, br.state != nil
}

// StateMust returns the state for [BigqueryReservation]. Panics if the state is nil.
func (br *BigqueryReservation) StateMust() *bigqueryReservationState {
	if br.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", br.Type(), br.LocalName()))
	}
	return br.state
}

// BigqueryReservationArgs contains the configurations for google_bigquery_reservation.
type BigqueryReservationArgs struct {
	// Concurrency: number, optional
	Concurrency terra.NumberValue `hcl:"concurrency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreIdleSlots: bool, optional
	IgnoreIdleSlots terra.BoolValue `hcl:"ignore_idle_slots,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// MultiRegionAuxiliary: bool, optional
	MultiRegionAuxiliary terra.BoolValue `hcl:"multi_region_auxiliary,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SlotCapacity: number, required
	SlotCapacity terra.NumberValue `hcl:"slot_capacity,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *bigqueryreservation.Timeouts `hcl:"timeouts,block"`
}
type bigqueryReservationAttributes struct {
	ref terra.Reference
}

// Concurrency returns a reference to field concurrency of google_bigquery_reservation.
func (br bigqueryReservationAttributes) Concurrency() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("concurrency"))
}

// Id returns a reference to field id of google_bigquery_reservation.
func (br bigqueryReservationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("id"))
}

// IgnoreIdleSlots returns a reference to field ignore_idle_slots of google_bigquery_reservation.
func (br bigqueryReservationAttributes) IgnoreIdleSlots() terra.BoolValue {
	return terra.ReferenceAsBool(br.ref.Append("ignore_idle_slots"))
}

// Location returns a reference to field location of google_bigquery_reservation.
func (br bigqueryReservationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("location"))
}

// MultiRegionAuxiliary returns a reference to field multi_region_auxiliary of google_bigquery_reservation.
func (br bigqueryReservationAttributes) MultiRegionAuxiliary() terra.BoolValue {
	return terra.ReferenceAsBool(br.ref.Append("multi_region_auxiliary"))
}

// Name returns a reference to field name of google_bigquery_reservation.
func (br bigqueryReservationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("name"))
}

// Project returns a reference to field project of google_bigquery_reservation.
func (br bigqueryReservationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("project"))
}

// SlotCapacity returns a reference to field slot_capacity of google_bigquery_reservation.
func (br bigqueryReservationAttributes) SlotCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("slot_capacity"))
}

func (br bigqueryReservationAttributes) Timeouts() bigqueryreservation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryreservation.TimeoutsAttributes](br.ref.Append("timeouts"))
}

type bigqueryReservationState struct {
	Concurrency          float64                            `json:"concurrency"`
	Id                   string                             `json:"id"`
	IgnoreIdleSlots      bool                               `json:"ignore_idle_slots"`
	Location             string                             `json:"location"`
	MultiRegionAuxiliary bool                               `json:"multi_region_auxiliary"`
	Name                 string                             `json:"name"`
	Project              string                             `json:"project"`
	SlotCapacity         float64                            `json:"slot_capacity"`
	Timeouts             *bigqueryreservation.TimeoutsState `json:"timeouts"`
}
