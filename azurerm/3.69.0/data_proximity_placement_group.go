// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataproximityplacementgroup "github.com/golingon/terraproviders/azurerm/3.69.0/dataproximityplacementgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataProximityPlacementGroup creates a new instance of [DataProximityPlacementGroup].
func NewDataProximityPlacementGroup(name string, args DataProximityPlacementGroupArgs) *DataProximityPlacementGroup {
	return &DataProximityPlacementGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataProximityPlacementGroup)(nil)

// DataProximityPlacementGroup represents the Terraform data resource azurerm_proximity_placement_group.
type DataProximityPlacementGroup struct {
	Name string
	Args DataProximityPlacementGroupArgs
}

// DataSource returns the Terraform object type for [DataProximityPlacementGroup].
func (ppg *DataProximityPlacementGroup) DataSource() string {
	return "azurerm_proximity_placement_group"
}

// LocalName returns the local name for [DataProximityPlacementGroup].
func (ppg *DataProximityPlacementGroup) LocalName() string {
	return ppg.Name
}

// Configuration returns the configuration (args) for [DataProximityPlacementGroup].
func (ppg *DataProximityPlacementGroup) Configuration() interface{} {
	return ppg.Args
}

// Attributes returns the attributes for [DataProximityPlacementGroup].
func (ppg *DataProximityPlacementGroup) Attributes() dataProximityPlacementGroupAttributes {
	return dataProximityPlacementGroupAttributes{ref: terra.ReferenceDataResource(ppg)}
}

// DataProximityPlacementGroupArgs contains the configurations for azurerm_proximity_placement_group.
type DataProximityPlacementGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataproximityplacementgroup.Timeouts `hcl:"timeouts,block"`
}
type dataProximityPlacementGroupAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_proximity_placement_group.
func (ppg dataProximityPlacementGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_proximity_placement_group.
func (ppg dataProximityPlacementGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_proximity_placement_group.
func (ppg dataProximityPlacementGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_proximity_placement_group.
func (ppg dataProximityPlacementGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ppg.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_proximity_placement_group.
func (ppg dataProximityPlacementGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ppg.ref.Append("tags"))
}

func (ppg dataProximityPlacementGroupAttributes) Timeouts() dataproximityplacementgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataproximityplacementgroup.TimeoutsAttributes](ppg.ref.Append("timeouts"))
}
