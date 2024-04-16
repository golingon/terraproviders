// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_reservation

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

// Resource represents the Terraform resource google_compute_reservation.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeReservationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcr *Resource) Type() string {
	return "google_compute_reservation"
}

// LocalName returns the local name for [Resource].
func (gcr *Resource) LocalName() string {
	return gcr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcr *Resource) Configuration() interface{} {
	return gcr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcr *Resource) Dependencies() terra.Dependencies {
	return gcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcr *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcr *Resource) Attributes() googleComputeReservationAttributes {
	return googleComputeReservationAttributes{ref: terra.ReferenceResource(gcr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcr *Resource) ImportState(state io.Reader) error {
	gcr.state = &googleComputeReservationState{}
	if err := json.NewDecoder(state).Decode(gcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcr.Type(), gcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcr *Resource) State() (*googleComputeReservationState, bool) {
	return gcr.state, gcr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcr *Resource) StateMust() *googleComputeReservationState {
	if gcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcr.Type(), gcr.LocalName()))
	}
	return gcr.state
}

// Args contains the configurations for google_compute_reservation.
type Args struct {
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
	ShareSettings *ShareSettings `hcl:"share_settings,block"`
	// SpecificReservation: required
	SpecificReservation *SpecificReservation `hcl:"specific_reservation,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeReservationAttributes struct {
	ref terra.Reference
}

// Commitment returns a reference to field commitment of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Commitment() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("commitment"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_reservation.
func (gcr googleComputeReservationAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_reservation.
func (gcr googleComputeReservationAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("self_link"))
}

// SpecificReservationRequired returns a reference to field specific_reservation_required of google_compute_reservation.
func (gcr googleComputeReservationAttributes) SpecificReservationRequired() terra.BoolValue {
	return terra.ReferenceAsBool(gcr.ref.Append("specific_reservation_required"))
}

// Status returns a reference to field status of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("status"))
}

// Zone returns a reference to field zone of google_compute_reservation.
func (gcr googleComputeReservationAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(gcr.ref.Append("zone"))
}

func (gcr googleComputeReservationAttributes) ShareSettings() terra.ListValue[ShareSettingsAttributes] {
	return terra.ReferenceAsList[ShareSettingsAttributes](gcr.ref.Append("share_settings"))
}

func (gcr googleComputeReservationAttributes) SpecificReservation() terra.ListValue[SpecificReservationAttributes] {
	return terra.ReferenceAsList[SpecificReservationAttributes](gcr.ref.Append("specific_reservation"))
}

func (gcr googleComputeReservationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcr.ref.Append("timeouts"))
}

type googleComputeReservationState struct {
	Commitment                  string                     `json:"commitment"`
	CreationTimestamp           string                     `json:"creation_timestamp"`
	Description                 string                     `json:"description"`
	Id                          string                     `json:"id"`
	Name                        string                     `json:"name"`
	Project                     string                     `json:"project"`
	SelfLink                    string                     `json:"self_link"`
	SpecificReservationRequired bool                       `json:"specific_reservation_required"`
	Status                      string                     `json:"status"`
	Zone                        string                     `json:"zone"`
	ShareSettings               []ShareSettingsState       `json:"share_settings"`
	SpecificReservation         []SpecificReservationState `json:"specific_reservation"`
	Timeouts                    *TimeoutsState             `json:"timeouts"`
}
