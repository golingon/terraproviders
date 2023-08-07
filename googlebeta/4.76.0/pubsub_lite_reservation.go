// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsublitereservation "github.com/golingon/terraproviders/googlebeta/4.76.0/pubsublitereservation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubLiteReservation creates a new instance of [PubsubLiteReservation].
func NewPubsubLiteReservation(name string, args PubsubLiteReservationArgs) *PubsubLiteReservation {
	return &PubsubLiteReservation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubLiteReservation)(nil)

// PubsubLiteReservation represents the Terraform resource google_pubsub_lite_reservation.
type PubsubLiteReservation struct {
	Name      string
	Args      PubsubLiteReservationArgs
	state     *pubsubLiteReservationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubLiteReservation].
func (plr *PubsubLiteReservation) Type() string {
	return "google_pubsub_lite_reservation"
}

// LocalName returns the local name for [PubsubLiteReservation].
func (plr *PubsubLiteReservation) LocalName() string {
	return plr.Name
}

// Configuration returns the configuration (args) for [PubsubLiteReservation].
func (plr *PubsubLiteReservation) Configuration() interface{} {
	return plr.Args
}

// DependOn is used for other resources to depend on [PubsubLiteReservation].
func (plr *PubsubLiteReservation) DependOn() terra.Reference {
	return terra.ReferenceResource(plr)
}

// Dependencies returns the list of resources [PubsubLiteReservation] depends_on.
func (plr *PubsubLiteReservation) Dependencies() terra.Dependencies {
	return plr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubLiteReservation].
func (plr *PubsubLiteReservation) LifecycleManagement() *terra.Lifecycle {
	return plr.Lifecycle
}

// Attributes returns the attributes for [PubsubLiteReservation].
func (plr *PubsubLiteReservation) Attributes() pubsubLiteReservationAttributes {
	return pubsubLiteReservationAttributes{ref: terra.ReferenceResource(plr)}
}

// ImportState imports the given attribute values into [PubsubLiteReservation]'s state.
func (plr *PubsubLiteReservation) ImportState(av io.Reader) error {
	plr.state = &pubsubLiteReservationState{}
	if err := json.NewDecoder(av).Decode(plr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", plr.Type(), plr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubLiteReservation] has state.
func (plr *PubsubLiteReservation) State() (*pubsubLiteReservationState, bool) {
	return plr.state, plr.state != nil
}

// StateMust returns the state for [PubsubLiteReservation]. Panics if the state is nil.
func (plr *PubsubLiteReservation) StateMust() *pubsubLiteReservationState {
	if plr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", plr.Type(), plr.LocalName()))
	}
	return plr.state
}

// PubsubLiteReservationArgs contains the configurations for google_pubsub_lite_reservation.
type PubsubLiteReservationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ThroughputCapacity: number, required
	ThroughputCapacity terra.NumberValue `hcl:"throughput_capacity,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *pubsublitereservation.Timeouts `hcl:"timeouts,block"`
}
type pubsubLiteReservationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_pubsub_lite_reservation.
func (plr pubsubLiteReservationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(plr.ref.Append("id"))
}

// Name returns a reference to field name of google_pubsub_lite_reservation.
func (plr pubsubLiteReservationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(plr.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_lite_reservation.
func (plr pubsubLiteReservationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(plr.ref.Append("project"))
}

// Region returns a reference to field region of google_pubsub_lite_reservation.
func (plr pubsubLiteReservationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(plr.ref.Append("region"))
}

// ThroughputCapacity returns a reference to field throughput_capacity of google_pubsub_lite_reservation.
func (plr pubsubLiteReservationAttributes) ThroughputCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(plr.ref.Append("throughput_capacity"))
}

func (plr pubsubLiteReservationAttributes) Timeouts() pubsublitereservation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pubsublitereservation.TimeoutsAttributes](plr.ref.Append("timeouts"))
}

type pubsubLiteReservationState struct {
	Id                 string                               `json:"id"`
	Name               string                               `json:"name"`
	Project            string                               `json:"project"`
	Region             string                               `json:"region"`
	ThroughputCapacity float64                              `json:"throughput_capacity"`
	Timeouts           *pubsublitereservation.TimeoutsState `json:"timeouts"`
}
