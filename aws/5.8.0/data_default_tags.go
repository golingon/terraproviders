// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDefaultTags creates a new instance of [DataDefaultTags].
func NewDataDefaultTags(name string, args DataDefaultTagsArgs) *DataDefaultTags {
	return &DataDefaultTags{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDefaultTags)(nil)

// DataDefaultTags represents the Terraform data resource aws_default_tags.
type DataDefaultTags struct {
	Name string
	Args DataDefaultTagsArgs
}

// DataSource returns the Terraform object type for [DataDefaultTags].
func (dt *DataDefaultTags) DataSource() string {
	return "aws_default_tags"
}

// LocalName returns the local name for [DataDefaultTags].
func (dt *DataDefaultTags) LocalName() string {
	return dt.Name
}

// Configuration returns the configuration (args) for [DataDefaultTags].
func (dt *DataDefaultTags) Configuration() interface{} {
	return dt.Args
}

// Attributes returns the attributes for [DataDefaultTags].
func (dt *DataDefaultTags) Attributes() dataDefaultTagsAttributes {
	return dataDefaultTagsAttributes{ref: terra.ReferenceDataResource(dt)}
}

// DataDefaultTagsArgs contains the configurations for aws_default_tags.
type DataDefaultTagsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataDefaultTagsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_default_tags.
func (dt dataDefaultTagsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dt.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_default_tags.
func (dt dataDefaultTagsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dt.ref.Append("tags"))
}
