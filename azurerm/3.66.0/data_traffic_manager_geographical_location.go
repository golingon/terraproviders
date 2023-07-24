// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datatrafficmanagergeographicallocation "github.com/golingon/terraproviders/azurerm/3.66.0/datatrafficmanagergeographicallocation"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataTrafficManagerGeographicalLocation creates a new instance of [DataTrafficManagerGeographicalLocation].
func NewDataTrafficManagerGeographicalLocation(name string, args DataTrafficManagerGeographicalLocationArgs) *DataTrafficManagerGeographicalLocation {
	return &DataTrafficManagerGeographicalLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTrafficManagerGeographicalLocation)(nil)

// DataTrafficManagerGeographicalLocation represents the Terraform data resource azurerm_traffic_manager_geographical_location.
type DataTrafficManagerGeographicalLocation struct {
	Name string
	Args DataTrafficManagerGeographicalLocationArgs
}

// DataSource returns the Terraform object type for [DataTrafficManagerGeographicalLocation].
func (tmgl *DataTrafficManagerGeographicalLocation) DataSource() string {
	return "azurerm_traffic_manager_geographical_location"
}

// LocalName returns the local name for [DataTrafficManagerGeographicalLocation].
func (tmgl *DataTrafficManagerGeographicalLocation) LocalName() string {
	return tmgl.Name
}

// Configuration returns the configuration (args) for [DataTrafficManagerGeographicalLocation].
func (tmgl *DataTrafficManagerGeographicalLocation) Configuration() interface{} {
	return tmgl.Args
}

// Attributes returns the attributes for [DataTrafficManagerGeographicalLocation].
func (tmgl *DataTrafficManagerGeographicalLocation) Attributes() dataTrafficManagerGeographicalLocationAttributes {
	return dataTrafficManagerGeographicalLocationAttributes{ref: terra.ReferenceDataResource(tmgl)}
}

// DataTrafficManagerGeographicalLocationArgs contains the configurations for azurerm_traffic_manager_geographical_location.
type DataTrafficManagerGeographicalLocationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datatrafficmanagergeographicallocation.Timeouts `hcl:"timeouts,block"`
}
type dataTrafficManagerGeographicalLocationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_traffic_manager_geographical_location.
func (tmgl dataTrafficManagerGeographicalLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tmgl.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_traffic_manager_geographical_location.
func (tmgl dataTrafficManagerGeographicalLocationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tmgl.ref.Append("name"))
}

func (tmgl dataTrafficManagerGeographicalLocationAttributes) Timeouts() datatrafficmanagergeographicallocation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datatrafficmanagergeographicallocation.TimeoutsAttributes](tmgl.ref.Append("timeouts"))
}
