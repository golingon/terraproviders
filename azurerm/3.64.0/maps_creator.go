// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mapscreator "github.com/golingon/terraproviders/azurerm/3.64.0/mapscreator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMapsCreator creates a new instance of [MapsCreator].
func NewMapsCreator(name string, args MapsCreatorArgs) *MapsCreator {
	return &MapsCreator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MapsCreator)(nil)

// MapsCreator represents the Terraform resource azurerm_maps_creator.
type MapsCreator struct {
	Name      string
	Args      MapsCreatorArgs
	state     *mapsCreatorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MapsCreator].
func (mc *MapsCreator) Type() string {
	return "azurerm_maps_creator"
}

// LocalName returns the local name for [MapsCreator].
func (mc *MapsCreator) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MapsCreator].
func (mc *MapsCreator) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MapsCreator].
func (mc *MapsCreator) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MapsCreator] depends_on.
func (mc *MapsCreator) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MapsCreator].
func (mc *MapsCreator) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MapsCreator].
func (mc *MapsCreator) Attributes() mapsCreatorAttributes {
	return mapsCreatorAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MapsCreator]'s state.
func (mc *MapsCreator) ImportState(av io.Reader) error {
	mc.state = &mapsCreatorState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MapsCreator] has state.
func (mc *MapsCreator) State() (*mapsCreatorState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MapsCreator]. Panics if the state is nil.
func (mc *MapsCreator) StateMust() *mapsCreatorState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MapsCreatorArgs contains the configurations for azurerm_maps_creator.
type MapsCreatorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MapsAccountId: string, required
	MapsAccountId terra.StringValue `hcl:"maps_account_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageUnits: number, required
	StorageUnits terra.NumberValue `hcl:"storage_units,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *mapscreator.Timeouts `hcl:"timeouts,block"`
}
type mapsCreatorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_maps_creator.
func (mc mapsCreatorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_maps_creator.
func (mc mapsCreatorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("location"))
}

// MapsAccountId returns a reference to field maps_account_id of azurerm_maps_creator.
func (mc mapsCreatorAttributes) MapsAccountId() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("maps_account_id"))
}

// Name returns a reference to field name of azurerm_maps_creator.
func (mc mapsCreatorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// StorageUnits returns a reference to field storage_units of azurerm_maps_creator.
func (mc mapsCreatorAttributes) StorageUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(mc.ref.Append("storage_units"))
}

// Tags returns a reference to field tags of azurerm_maps_creator.
func (mc mapsCreatorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mc.ref.Append("tags"))
}

func (mc mapsCreatorAttributes) Timeouts() mapscreator.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mapscreator.TimeoutsAttributes](mc.ref.Append("timeouts"))
}

type mapsCreatorState struct {
	Id            string                     `json:"id"`
	Location      string                     `json:"location"`
	MapsAccountId string                     `json:"maps_account_id"`
	Name          string                     `json:"name"`
	StorageUnits  float64                    `json:"storage_units"`
	Tags          map[string]string          `json:"tags"`
	Timeouts      *mapscreator.TimeoutsState `json:"timeouts"`
}
