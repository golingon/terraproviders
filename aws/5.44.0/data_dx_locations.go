// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataDxLocations creates a new instance of [DataDxLocations].
func NewDataDxLocations(name string, args DataDxLocationsArgs) *DataDxLocations {
	return &DataDxLocations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDxLocations)(nil)

// DataDxLocations represents the Terraform data resource aws_dx_locations.
type DataDxLocations struct {
	Name string
	Args DataDxLocationsArgs
}

// DataSource returns the Terraform object type for [DataDxLocations].
func (dl *DataDxLocations) DataSource() string {
	return "aws_dx_locations"
}

// LocalName returns the local name for [DataDxLocations].
func (dl *DataDxLocations) LocalName() string {
	return dl.Name
}

// Configuration returns the configuration (args) for [DataDxLocations].
func (dl *DataDxLocations) Configuration() interface{} {
	return dl.Args
}

// Attributes returns the attributes for [DataDxLocations].
func (dl *DataDxLocations) Attributes() dataDxLocationsAttributes {
	return dataDxLocationsAttributes{ref: terra.ReferenceDataResource(dl)}
}

// DataDxLocationsArgs contains the configurations for aws_dx_locations.
type DataDxLocationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataDxLocationsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_dx_locations.
func (dl dataDxLocationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("id"))
}

// LocationCodes returns a reference to field location_codes of aws_dx_locations.
func (dl dataDxLocationsAttributes) LocationCodes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dl.ref.Append("location_codes"))
}
