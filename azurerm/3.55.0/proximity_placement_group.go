// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	proximityplacementgroup "github.com/golingon/terraproviders/azurerm/3.55.0/proximityplacementgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProximityPlacementGroup creates a new instance of [ProximityPlacementGroup].
func NewProximityPlacementGroup(name string, args ProximityPlacementGroupArgs) *ProximityPlacementGroup {
	return &ProximityPlacementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProximityPlacementGroup)(nil)

// ProximityPlacementGroup represents the Terraform resource azurerm_proximity_placement_group.
type ProximityPlacementGroup struct {
	Name      string
	Args      ProximityPlacementGroupArgs
	state     *proximityPlacementGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProximityPlacementGroup].
func (ppg *ProximityPlacementGroup) Type() string {
	return "azurerm_proximity_placement_group"
}

// LocalName returns the local name for [ProximityPlacementGroup].
func (ppg *ProximityPlacementGroup) LocalName() string {
	return ppg.Name
}

// Configuration returns the configuration (args) for [ProximityPlacementGroup].
func (ppg *ProximityPlacementGroup) Configuration() interface{} {
	return ppg.Args
}

// DependOn is used for other resources to depend on [ProximityPlacementGroup].
func (ppg *ProximityPlacementGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ppg)
}

// Dependencies returns the list of resources [ProximityPlacementGroup] depends_on.
func (ppg *ProximityPlacementGroup) Dependencies() terra.Dependencies {
	return ppg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProximityPlacementGroup].
func (ppg *ProximityPlacementGroup) LifecycleManagement() *terra.Lifecycle {
	return ppg.Lifecycle
}

// Attributes returns the attributes for [ProximityPlacementGroup].
func (ppg *ProximityPlacementGroup) Attributes() proximityPlacementGroupAttributes {
	return proximityPlacementGroupAttributes{ref: terra.ReferenceResource(ppg)}
}

// ImportState imports the given attribute values into [ProximityPlacementGroup]'s state.
func (ppg *ProximityPlacementGroup) ImportState(av io.Reader) error {
	ppg.state = &proximityPlacementGroupState{}
	if err := json.NewDecoder(av).Decode(ppg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ppg.Type(), ppg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProximityPlacementGroup] has state.
func (ppg *ProximityPlacementGroup) State() (*proximityPlacementGroupState, bool) {
	return ppg.state, ppg.state != nil
}

// StateMust returns the state for [ProximityPlacementGroup]. Panics if the state is nil.
func (ppg *ProximityPlacementGroup) StateMust() *proximityPlacementGroupState {
	if ppg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ppg.Type(), ppg.LocalName()))
	}
	return ppg.state
}

// ProximityPlacementGroupArgs contains the configurations for azurerm_proximity_placement_group.
type ProximityPlacementGroupArgs struct {
	// AllowedVmSizes: set of string, optional
	AllowedVmSizes terra.SetValue[terra.StringValue] `hcl:"allowed_vm_sizes,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// Timeouts: optional
	Timeouts *proximityplacementgroup.Timeouts `hcl:"timeouts,block"`
}
type proximityPlacementGroupAttributes struct {
	ref terra.Reference
}

// AllowedVmSizes returns a reference to field allowed_vm_sizes of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) AllowedVmSizes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ppg.ref.Append("allowed_vm_sizes"))
}

// Id returns a reference to field id of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ppg.ref.Append("tags"))
}

// Zone returns a reference to field zone of azurerm_proximity_placement_group.
func (ppg proximityPlacementGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("zone"))
}

func (ppg proximityPlacementGroupAttributes) Timeouts() proximityplacementgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[proximityplacementgroup.TimeoutsAttributes](ppg.ref.Append("timeouts"))
}

type proximityPlacementGroupState struct {
	AllowedVmSizes    []string                               `json:"allowed_vm_sizes"`
	Id                string                                 `json:"id"`
	Location          string                                 `json:"location"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Tags              map[string]string                      `json:"tags"`
	Zone              string                                 `json:"zone"`
	Timeouts          *proximityplacementgroup.TimeoutsState `json:"timeouts"`
}
