// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDxLocation creates a new instance of [DataDxLocation].
func NewDataDxLocation(name string, args DataDxLocationArgs) *DataDxLocation {
	return &DataDxLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDxLocation)(nil)

// DataDxLocation represents the Terraform data resource aws_dx_location.
type DataDxLocation struct {
	Name string
	Args DataDxLocationArgs
}

// DataSource returns the Terraform object type for [DataDxLocation].
func (dl *DataDxLocation) DataSource() string {
	return "aws_dx_location"
}

// LocalName returns the local name for [DataDxLocation].
func (dl *DataDxLocation) LocalName() string {
	return dl.Name
}

// Configuration returns the configuration (args) for [DataDxLocation].
func (dl *DataDxLocation) Configuration() interface{} {
	return dl.Args
}

// Attributes returns the attributes for [DataDxLocation].
func (dl *DataDxLocation) Attributes() dataDxLocationAttributes {
	return dataDxLocationAttributes{ref: terra.ReferenceDataResource(dl)}
}

// DataDxLocationArgs contains the configurations for aws_dx_location.
type DataDxLocationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationCode: string, required
	LocationCode terra.StringValue `hcl:"location_code,attr" validate:"required"`
}
type dataDxLocationAttributes struct {
	ref terra.Reference
}

// AvailableMacsecPortSpeeds returns a reference to field available_macsec_port_speeds of aws_dx_location.
func (dl dataDxLocationAttributes) AvailableMacsecPortSpeeds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dl.ref.Append("available_macsec_port_speeds"))
}

// AvailablePortSpeeds returns a reference to field available_port_speeds of aws_dx_location.
func (dl dataDxLocationAttributes) AvailablePortSpeeds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dl.ref.Append("available_port_speeds"))
}

// AvailableProviders returns a reference to field available_providers of aws_dx_location.
func (dl dataDxLocationAttributes) AvailableProviders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dl.ref.Append("available_providers"))
}

// Id returns a reference to field id of aws_dx_location.
func (dl dataDxLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("id"))
}

// LocationCode returns a reference to field location_code of aws_dx_location.
func (dl dataDxLocationAttributes) LocationCode() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("location_code"))
}

// LocationName returns a reference to field location_name of aws_dx_location.
func (dl dataDxLocationAttributes) LocationName() terra.StringValue {
	return terra.ReferenceAsString(dl.ref.Append("location_name"))
}