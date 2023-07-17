// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computereservation "github.com/golingon/terraproviders/google/4.73.1/computereservation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeReservation creates a new instance of [ComputeReservation].
func NewComputeReservation(name string, args ComputeReservationArgs) *ComputeReservation {
	return &ComputeReservation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeReservation)(nil)

// ComputeReservation represents the Terraform resource google_compute_reservation.
type ComputeReservation struct {
	Name      string
	Args      ComputeReservationArgs
	state     *computeReservationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeReservation].
func (cr *ComputeReservation) Type() string {
	return "google_compute_reservation"
}

// LocalName returns the local name for [ComputeReservation].
func (cr *ComputeReservation) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [ComputeReservation].
func (cr *ComputeReservation) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [ComputeReservation].
func (cr *ComputeReservation) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [ComputeReservation] depends_on.
func (cr *ComputeReservation) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeReservation].
func (cr *ComputeReservation) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [ComputeReservation].
func (cr *ComputeReservation) Attributes() computeReservationAttributes {
	return computeReservationAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [ComputeReservation]'s state.
func (cr *ComputeReservation) ImportState(av io.Reader) error {
	cr.state = &computeReservationState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeReservation] has state.
func (cr *ComputeReservation) State() (*computeReservationState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [ComputeReservation]. Panics if the state is nil.
func (cr *ComputeReservation) StateMust() *computeReservationState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// ComputeReservationArgs contains the configurations for google_compute_reservation.
type ComputeReservationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SpecificReservationRequired: bool, optional
	SpecificReservationRequired terra.BoolValue `hcl:"specific_reservation_required,attr"`
	// Zone: string, required
	Zone terra.StringValue `hcl:"zone,attr" validate:"required"`
	// ShareSettings: optional
	ShareSettings *computereservation.ShareSettings `hcl:"share_settings,block"`
	// SpecificReservation: required
	SpecificReservation *computereservation.SpecificReservation `hcl:"specific_reservation,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computereservation.Timeouts `hcl:"timeouts,block"`
}
type computeReservationAttributes struct {
	ref terra.Reference
}

// Commitment returns a reference to field commitment of google_compute_reservation.
func (cr computeReservationAttributes) Commitment() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("commitment"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_reservation.
func (cr computeReservationAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_reservation.
func (cr computeReservationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_reservation.
func (cr computeReservationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_reservation.
func (cr computeReservationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_reservation.
func (cr computeReservationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_reservation.
func (cr computeReservationAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("self_link"))
}

// SpecificReservationRequired returns a reference to field specific_reservation_required of google_compute_reservation.
func (cr computeReservationAttributes) SpecificReservationRequired() terra.BoolValue {
	return terra.ReferenceAsBool(cr.ref.Append("specific_reservation_required"))
}

// Status returns a reference to field status of google_compute_reservation.
func (cr computeReservationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("status"))
}

// Zone returns a reference to field zone of google_compute_reservation.
func (cr computeReservationAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("zone"))
}

func (cr computeReservationAttributes) ShareSettings() terra.ListValue[computereservation.ShareSettingsAttributes] {
	return terra.ReferenceAsList[computereservation.ShareSettingsAttributes](cr.ref.Append("share_settings"))
}

func (cr computeReservationAttributes) SpecificReservation() terra.ListValue[computereservation.SpecificReservationAttributes] {
	return terra.ReferenceAsList[computereservation.SpecificReservationAttributes](cr.ref.Append("specific_reservation"))
}

func (cr computeReservationAttributes) Timeouts() computereservation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computereservation.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type computeReservationState struct {
	Commitment                  string                                        `json:"commitment"`
	CreationTimestamp           string                                        `json:"creation_timestamp"`
	Description                 string                                        `json:"description"`
	Id                          string                                        `json:"id"`
	Name                        string                                        `json:"name"`
	Project                     string                                        `json:"project"`
	SelfLink                    string                                        `json:"self_link"`
	SpecificReservationRequired bool                                          `json:"specific_reservation_required"`
	Status                      string                                        `json:"status"`
	Zone                        string                                        `json:"zone"`
	ShareSettings               []computereservation.ShareSettingsState       `json:"share_settings"`
	SpecificReservation         []computereservation.SpecificReservationState `json:"specific_reservation"`
	Timeouts                    *computereservation.TimeoutsState             `json:"timeouts"`
}
