// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dx_locations

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_dx_locations.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adl *DataSource) DataSource() string {
	return "aws_dx_locations"
}

// LocalName returns the local name for [DataSource].
func (adl *DataSource) LocalName() string {
	return adl.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adl *DataSource) Configuration() interface{} {
	return adl.Args
}

// Attributes returns the attributes for [DataSource].
func (adl *DataSource) Attributes() dataAwsDxLocationsAttributes {
	return dataAwsDxLocationsAttributes{ref: terra.ReferenceDataSource(adl)}
}

// DataArgs contains the configurations for aws_dx_locations.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataAwsDxLocationsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_dx_locations.
func (adl dataAwsDxLocationsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adl.ref.Append("id"))
}

// LocationCodes returns a reference to field location_codes of aws_dx_locations.
func (adl dataAwsDxLocationsAttributes) LocationCodes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adl.ref.Append("location_codes"))
}
