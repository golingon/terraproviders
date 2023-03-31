// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudwatchEventConnection creates a new instance of [DataCloudwatchEventConnection].
func NewDataCloudwatchEventConnection(name string, args DataCloudwatchEventConnectionArgs) *DataCloudwatchEventConnection {
	return &DataCloudwatchEventConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudwatchEventConnection)(nil)

// DataCloudwatchEventConnection represents the Terraform data resource aws_cloudwatch_event_connection.
type DataCloudwatchEventConnection struct {
	Name string
	Args DataCloudwatchEventConnectionArgs
}

// DataSource returns the Terraform object type for [DataCloudwatchEventConnection].
func (cec *DataCloudwatchEventConnection) DataSource() string {
	return "aws_cloudwatch_event_connection"
}

// LocalName returns the local name for [DataCloudwatchEventConnection].
func (cec *DataCloudwatchEventConnection) LocalName() string {
	return cec.Name
}

// Configuration returns the configuration (args) for [DataCloudwatchEventConnection].
func (cec *DataCloudwatchEventConnection) Configuration() interface{} {
	return cec.Args
}

// Attributes returns the attributes for [DataCloudwatchEventConnection].
func (cec *DataCloudwatchEventConnection) Attributes() dataCloudwatchEventConnectionAttributes {
	return dataCloudwatchEventConnectionAttributes{ref: terra.ReferenceDataResource(cec)}
}

// DataCloudwatchEventConnectionArgs contains the configurations for aws_cloudwatch_event_connection.
type DataCloudwatchEventConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataCloudwatchEventConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_event_connection.
func (cec dataCloudwatchEventConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("arn"))
}

// AuthorizationType returns a reference to field authorization_type of aws_cloudwatch_event_connection.
func (cec dataCloudwatchEventConnectionAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("authorization_type"))
}

// Id returns a reference to field id of aws_cloudwatch_event_connection.
func (cec dataCloudwatchEventConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudwatch_event_connection.
func (cec dataCloudwatchEventConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("name"))
}

// SecretArn returns a reference to field secret_arn of aws_cloudwatch_event_connection.
func (cec dataCloudwatchEventConnectionAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("secret_arn"))
}
