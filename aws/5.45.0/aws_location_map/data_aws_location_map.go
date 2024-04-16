// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_location_map

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_location_map.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (alm *DataSource) DataSource() string {
	return "aws_location_map"
}

// LocalName returns the local name for [DataSource].
func (alm *DataSource) LocalName() string {
	return alm.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (alm *DataSource) Configuration() interface{} {
	return alm.Args
}

// Attributes returns the attributes for [DataSource].
func (alm *DataSource) Attributes() dataAwsLocationMapAttributes {
	return dataAwsLocationMapAttributes{ref: terra.ReferenceDataSource(alm)}
}

// DataArgs contains the configurations for aws_location_map.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MapName: string, required
	MapName terra.StringValue `hcl:"map_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsLocationMapAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of aws_location_map.
func (alm dataAwsLocationMapAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(alm.ref.Append("create_time"))
}

// Description returns a reference to field description of aws_location_map.
func (alm dataAwsLocationMapAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(alm.ref.Append("description"))
}

// Id returns a reference to field id of aws_location_map.
func (alm dataAwsLocationMapAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alm.ref.Append("id"))
}

// MapArn returns a reference to field map_arn of aws_location_map.
func (alm dataAwsLocationMapAttributes) MapArn() terra.StringValue {
	return terra.ReferenceAsString(alm.ref.Append("map_arn"))
}

// MapName returns a reference to field map_name of aws_location_map.
func (alm dataAwsLocationMapAttributes) MapName() terra.StringValue {
	return terra.ReferenceAsString(alm.ref.Append("map_name"))
}

// Tags returns a reference to field tags of aws_location_map.
func (alm dataAwsLocationMapAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alm.ref.Append("tags"))
}

// UpdateTime returns a reference to field update_time of aws_location_map.
func (alm dataAwsLocationMapAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(alm.ref.Append("update_time"))
}

func (alm dataAwsLocationMapAttributes) Configuration() terra.ListValue[DataConfigurationAttributes] {
	return terra.ReferenceAsList[DataConfigurationAttributes](alm.ref.Append("configuration"))
}
