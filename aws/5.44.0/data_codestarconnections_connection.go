// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataCodestarconnectionsConnection creates a new instance of [DataCodestarconnectionsConnection].
func NewDataCodestarconnectionsConnection(name string, args DataCodestarconnectionsConnectionArgs) *DataCodestarconnectionsConnection {
	return &DataCodestarconnectionsConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCodestarconnectionsConnection)(nil)

// DataCodestarconnectionsConnection represents the Terraform data resource aws_codestarconnections_connection.
type DataCodestarconnectionsConnection struct {
	Name string
	Args DataCodestarconnectionsConnectionArgs
}

// DataSource returns the Terraform object type for [DataCodestarconnectionsConnection].
func (cc *DataCodestarconnectionsConnection) DataSource() string {
	return "aws_codestarconnections_connection"
}

// LocalName returns the local name for [DataCodestarconnectionsConnection].
func (cc *DataCodestarconnectionsConnection) LocalName() string {
	return cc.Name
}

// Configuration returns the configuration (args) for [DataCodestarconnectionsConnection].
func (cc *DataCodestarconnectionsConnection) Configuration() interface{} {
	return cc.Args
}

// Attributes returns the attributes for [DataCodestarconnectionsConnection].
func (cc *DataCodestarconnectionsConnection) Attributes() dataCodestarconnectionsConnectionAttributes {
	return dataCodestarconnectionsConnectionAttributes{ref: terra.ReferenceDataResource(cc)}
}

// DataCodestarconnectionsConnectionArgs contains the configurations for aws_codestarconnections_connection.
type DataCodestarconnectionsConnectionArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataCodestarconnectionsConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("arn"))
}

// ConnectionStatus returns a reference to field connection_status of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) ConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("connection_status"))
}

// HostArn returns a reference to field host_arn of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) HostArn() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("host_arn"))
}

// Id returns a reference to field id of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("id"))
}

// Name returns a reference to field name of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

// ProviderType returns a reference to field provider_type of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) ProviderType() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("provider_type"))
}

// Tags returns a reference to field tags of aws_codestarconnections_connection.
func (cc dataCodestarconnectionsConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cc.ref.Append("tags"))
}
