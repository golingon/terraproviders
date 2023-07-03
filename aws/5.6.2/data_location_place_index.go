// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalocationplaceindex "github.com/golingon/terraproviders/aws/5.6.2/datalocationplaceindex"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLocationPlaceIndex creates a new instance of [DataLocationPlaceIndex].
func NewDataLocationPlaceIndex(name string, args DataLocationPlaceIndexArgs) *DataLocationPlaceIndex {
	return &DataLocationPlaceIndex{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLocationPlaceIndex)(nil)

// DataLocationPlaceIndex represents the Terraform data resource aws_location_place_index.
type DataLocationPlaceIndex struct {
	Name string
	Args DataLocationPlaceIndexArgs
}

// DataSource returns the Terraform object type for [DataLocationPlaceIndex].
func (lpi *DataLocationPlaceIndex) DataSource() string {
	return "aws_location_place_index"
}

// LocalName returns the local name for [DataLocationPlaceIndex].
func (lpi *DataLocationPlaceIndex) LocalName() string {
	return lpi.Name
}

// Configuration returns the configuration (args) for [DataLocationPlaceIndex].
func (lpi *DataLocationPlaceIndex) Configuration() interface{} {
	return lpi.Args
}

// Attributes returns the attributes for [DataLocationPlaceIndex].
func (lpi *DataLocationPlaceIndex) Attributes() dataLocationPlaceIndexAttributes {
	return dataLocationPlaceIndexAttributes{ref: terra.ReferenceDataResource(lpi)}
}

// DataLocationPlaceIndexArgs contains the configurations for aws_location_place_index.
type DataLocationPlaceIndexArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexName: string, required
	IndexName terra.StringValue `hcl:"index_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DataSourceConfiguration: min=0
	DataSourceConfiguration []datalocationplaceindex.DataSourceConfiguration `hcl:"data_source_configuration,block" validate:"min=0"`
}
type dataLocationPlaceIndexAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("create_time"))
}

// DataSource returns a reference to field data_source of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) DataSource() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("data_source"))
}

// Description returns a reference to field description of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("id"))
}

// IndexArn returns a reference to field index_arn of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) IndexArn() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("index_arn"))
}

// IndexName returns a reference to field index_name of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) IndexName() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("index_name"))
}

// Tags returns a reference to field tags of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lpi.ref.Append("tags"))
}

// UpdateTime returns a reference to field update_time of aws_location_place_index.
func (lpi dataLocationPlaceIndexAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lpi.ref.Append("update_time"))
}

func (lpi dataLocationPlaceIndexAttributes) DataSourceConfiguration() terra.ListValue[datalocationplaceindex.DataSourceConfigurationAttributes] {
	return terra.ReferenceAsList[datalocationplaceindex.DataSourceConfigurationAttributes](lpi.ref.Append("data_source_configuration"))
}
