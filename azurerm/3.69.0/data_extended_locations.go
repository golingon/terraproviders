// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataextendedlocations "github.com/golingon/terraproviders/azurerm/3.69.0/dataextendedlocations"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataExtendedLocations creates a new instance of [DataExtendedLocations].
func NewDataExtendedLocations(name string, args DataExtendedLocationsArgs) *DataExtendedLocations {
	return &DataExtendedLocations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataExtendedLocations)(nil)

// DataExtendedLocations represents the Terraform data resource azurerm_extended_locations.
type DataExtendedLocations struct {
	Name string
	Args DataExtendedLocationsArgs
}

// DataSource returns the Terraform object type for [DataExtendedLocations].
func (el *DataExtendedLocations) DataSource() string {
	return "azurerm_extended_locations"
}

// LocalName returns the local name for [DataExtendedLocations].
func (el *DataExtendedLocations) LocalName() string {
	return el.Name
}

// Configuration returns the configuration (args) for [DataExtendedLocations].
func (el *DataExtendedLocations) Configuration() interface{} {
	return el.Args
}

// Attributes returns the attributes for [DataExtendedLocations].
func (el *DataExtendedLocations) Attributes() dataExtendedLocationsAttributes {
	return dataExtendedLocationsAttributes{ref: terra.ReferenceDataResource(el)}
}

// DataExtendedLocationsArgs contains the configurations for azurerm_extended_locations.
type DataExtendedLocationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataextendedlocations.Timeouts `hcl:"timeouts,block"`
}
type dataExtendedLocationsAttributes struct {
	ref terra.Reference
}

// ExtendedLocations returns a reference to field extended_locations of azurerm_extended_locations.
func (el dataExtendedLocationsAttributes) ExtendedLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](el.ref.Append("extended_locations"))
}

// Id returns a reference to field id of azurerm_extended_locations.
func (el dataExtendedLocationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_extended_locations.
func (el dataExtendedLocationsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("location"))
}

func (el dataExtendedLocationsAttributes) Timeouts() dataextendedlocations.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataextendedlocations.TimeoutsAttributes](el.ref.Append("timeouts"))
}
