// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	capacityreservation "github.com/golingon/terraproviders/azurerm/3.69.0/capacityreservation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCapacityReservation creates a new instance of [CapacityReservation].
func NewCapacityReservation(name string, args CapacityReservationArgs) *CapacityReservation {
	return &CapacityReservation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CapacityReservation)(nil)

// CapacityReservation represents the Terraform resource azurerm_capacity_reservation.
type CapacityReservation struct {
	Name      string
	Args      CapacityReservationArgs
	state     *capacityReservationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CapacityReservation].
func (cr *CapacityReservation) Type() string {
	return "azurerm_capacity_reservation"
}

// LocalName returns the local name for [CapacityReservation].
func (cr *CapacityReservation) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [CapacityReservation].
func (cr *CapacityReservation) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [CapacityReservation].
func (cr *CapacityReservation) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [CapacityReservation] depends_on.
func (cr *CapacityReservation) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CapacityReservation].
func (cr *CapacityReservation) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [CapacityReservation].
func (cr *CapacityReservation) Attributes() capacityReservationAttributes {
	return capacityReservationAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [CapacityReservation]'s state.
func (cr *CapacityReservation) ImportState(av io.Reader) error {
	cr.state = &capacityReservationState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CapacityReservation] has state.
func (cr *CapacityReservation) State() (*capacityReservationState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [CapacityReservation]. Panics if the state is nil.
func (cr *CapacityReservation) StateMust() *capacityReservationState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// CapacityReservationArgs contains the configurations for azurerm_capacity_reservation.
type CapacityReservationArgs struct {
	// CapacityReservationGroupId: string, required
	CapacityReservationGroupId terra.StringValue `hcl:"capacity_reservation_group_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Sku: required
	Sku *capacityreservation.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *capacityreservation.Timeouts `hcl:"timeouts,block"`
}
type capacityReservationAttributes struct {
	ref terra.Reference
}

// CapacityReservationGroupId returns a reference to field capacity_reservation_group_id of azurerm_capacity_reservation.
func (cr capacityReservationAttributes) CapacityReservationGroupId() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("capacity_reservation_group_id"))
}

// Id returns a reference to field id of azurerm_capacity_reservation.
func (cr capacityReservationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_capacity_reservation.
func (cr capacityReservationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_capacity_reservation.
func (cr capacityReservationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cr.ref.Append("tags"))
}

// Zone returns a reference to field zone of azurerm_capacity_reservation.
func (cr capacityReservationAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("zone"))
}

func (cr capacityReservationAttributes) Sku() terra.ListValue[capacityreservation.SkuAttributes] {
	return terra.ReferenceAsList[capacityreservation.SkuAttributes](cr.ref.Append("sku"))
}

func (cr capacityReservationAttributes) Timeouts() capacityreservation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[capacityreservation.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type capacityReservationState struct {
	CapacityReservationGroupId string                             `json:"capacity_reservation_group_id"`
	Id                         string                             `json:"id"`
	Name                       string                             `json:"name"`
	Tags                       map[string]string                  `json:"tags"`
	Zone                       string                             `json:"zone"`
	Sku                        []capacityreservation.SkuState     `json:"sku"`
	Timeouts                   *capacityreservation.TimeoutsState `json:"timeouts"`
}
