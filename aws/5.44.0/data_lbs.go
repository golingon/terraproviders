// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataLbs creates a new instance of [DataLbs].
func NewDataLbs(name string, args DataLbsArgs) *DataLbs {
	return &DataLbs{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbs)(nil)

// DataLbs represents the Terraform data resource aws_lbs.
type DataLbs struct {
	Name string
	Args DataLbsArgs
}

// DataSource returns the Terraform object type for [DataLbs].
func (l *DataLbs) DataSource() string {
	return "aws_lbs"
}

// LocalName returns the local name for [DataLbs].
func (l *DataLbs) LocalName() string {
	return l.Name
}

// Configuration returns the configuration (args) for [DataLbs].
func (l *DataLbs) Configuration() interface{} {
	return l.Args
}

// Attributes returns the attributes for [DataLbs].
func (l *DataLbs) Attributes() dataLbsAttributes {
	return dataLbsAttributes{ref: terra.ReferenceDataResource(l)}
}

// DataLbsArgs contains the configurations for aws_lbs.
type DataLbsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataLbsAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_lbs.
func (l dataLbsAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](l.ref.Append("arns"))
}

// Id returns a reference to field id of aws_lbs.
func (l dataLbsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_lbs.
func (l dataLbsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](l.ref.Append("tags"))
}
