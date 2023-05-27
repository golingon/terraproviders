// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	dataalloydblocations "github.com/golingon/terraproviders/googlebeta/4.66.0/dataalloydblocations"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAlloydbLocations creates a new instance of [DataAlloydbLocations].
func NewDataAlloydbLocations(name string, args DataAlloydbLocationsArgs) *DataAlloydbLocations {
	return &DataAlloydbLocations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAlloydbLocations)(nil)

// DataAlloydbLocations represents the Terraform data resource google_alloydb_locations.
type DataAlloydbLocations struct {
	Name string
	Args DataAlloydbLocationsArgs
}

// DataSource returns the Terraform object type for [DataAlloydbLocations].
func (al *DataAlloydbLocations) DataSource() string {
	return "google_alloydb_locations"
}

// LocalName returns the local name for [DataAlloydbLocations].
func (al *DataAlloydbLocations) LocalName() string {
	return al.Name
}

// Configuration returns the configuration (args) for [DataAlloydbLocations].
func (al *DataAlloydbLocations) Configuration() interface{} {
	return al.Args
}

// Attributes returns the attributes for [DataAlloydbLocations].
func (al *DataAlloydbLocations) Attributes() dataAlloydbLocationsAttributes {
	return dataAlloydbLocationsAttributes{ref: terra.ReferenceDataResource(al)}
}

// DataAlloydbLocationsArgs contains the configurations for google_alloydb_locations.
type DataAlloydbLocationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Locations: min=0
	Locations []dataalloydblocations.Locations `hcl:"locations,block" validate:"min=0"`
}
type dataAlloydbLocationsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_alloydb_locations.
func (al dataAlloydbLocationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("id"))
}

// Project returns a reference to field project of google_alloydb_locations.
func (al dataAlloydbLocationsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(al.ref.Append("project"))
}

func (al dataAlloydbLocationsAttributes) Locations() terra.ListValue[dataalloydblocations.LocationsAttributes] {
	return terra.ReferenceAsList[dataalloydblocations.LocationsAttributes](al.ref.Append("locations"))
}
