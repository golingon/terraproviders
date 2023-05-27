// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSnsTopic creates a new instance of [DataSnsTopic].
func NewDataSnsTopic(name string, args DataSnsTopicArgs) *DataSnsTopic {
	return &DataSnsTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSnsTopic)(nil)

// DataSnsTopic represents the Terraform data resource aws_sns_topic.
type DataSnsTopic struct {
	Name string
	Args DataSnsTopicArgs
}

// DataSource returns the Terraform object type for [DataSnsTopic].
func (st *DataSnsTopic) DataSource() string {
	return "aws_sns_topic"
}

// LocalName returns the local name for [DataSnsTopic].
func (st *DataSnsTopic) LocalName() string {
	return st.Name
}

// Configuration returns the configuration (args) for [DataSnsTopic].
func (st *DataSnsTopic) Configuration() interface{} {
	return st.Args
}

// Attributes returns the attributes for [DataSnsTopic].
func (st *DataSnsTopic) Attributes() dataSnsTopicAttributes {
	return dataSnsTopicAttributes{ref: terra.ReferenceDataResource(st)}
}

// DataSnsTopicArgs contains the configurations for aws_sns_topic.
type DataSnsTopicArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataSnsTopicAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sns_topic.
func (st dataSnsTopicAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("arn"))
}

// Id returns a reference to field id of aws_sns_topic.
func (st dataSnsTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("id"))
}

// Name returns a reference to field name of aws_sns_topic.
func (st dataSnsTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("name"))
}
