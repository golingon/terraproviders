// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalocationmap "github.com/golingon/terraproviders/aws/5.10.0/datalocationmap"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLocationMap creates a new instance of [DataLocationMap].
func NewDataLocationMap(name string, args DataLocationMapArgs) *DataLocationMap {
	return &DataLocationMap{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLocationMap)(nil)

// DataLocationMap represents the Terraform data resource aws_location_map.
type DataLocationMap struct {
	Name string
	Args DataLocationMapArgs
}

// DataSource returns the Terraform object type for [DataLocationMap].
func (lm *DataLocationMap) DataSource() string {
	return "aws_location_map"
}

// LocalName returns the local name for [DataLocationMap].
func (lm *DataLocationMap) LocalName() string {
	return lm.Name
}

// Configuration returns the configuration (args) for [DataLocationMap].
func (lm *DataLocationMap) Configuration() interface{} {
	return lm.Args
}

// Attributes returns the attributes for [DataLocationMap].
func (lm *DataLocationMap) Attributes() dataLocationMapAttributes {
	return dataLocationMapAttributes{ref: terra.ReferenceDataResource(lm)}
}

// DataLocationMapArgs contains the configurations for aws_location_map.
type DataLocationMapArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MapName: string, required
	MapName terra.StringValue `hcl:"map_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Configuration: min=0
	Configuration []datalocationmap.Configuration `hcl:"configuration,block" validate:"min=0"`
}
type dataLocationMapAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of aws_location_map.
func (lm dataLocationMapAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("create_time"))
}

// Description returns a reference to field description of aws_location_map.
func (lm dataLocationMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_map.
func (lm dataLocationMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("id"))
}

// MapArn returns a reference to field map_arn of aws_location_map.
func (lm dataLocationMapAttributes) MapArn() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("map_arn"))
}

// MapName returns a reference to field map_name of aws_location_map.
func (lm dataLocationMapAttributes) MapName() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("map_name"))
}

// Tags returns a reference to field tags of aws_location_map.
func (lm dataLocationMapAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lm.ref.Append("tags"))
}

// UpdateTime returns a reference to field update_time of aws_location_map.
func (lm dataLocationMapAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(lm.ref.Append("update_time"))
}

func (lm dataLocationMapAttributes) Configuration() terra.ListValue[datalocationmap.ConfigurationAttributes] {
	return terra.ReferenceAsList[datalocationmap.ConfigurationAttributes](lm.ref.Append("configuration"))
}
