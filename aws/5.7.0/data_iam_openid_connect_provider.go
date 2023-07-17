// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamOpenidConnectProvider creates a new instance of [DataIamOpenidConnectProvider].
func NewDataIamOpenidConnectProvider(name string, args DataIamOpenidConnectProviderArgs) *DataIamOpenidConnectProvider {
	return &DataIamOpenidConnectProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamOpenidConnectProvider)(nil)

// DataIamOpenidConnectProvider represents the Terraform data resource aws_iam_openid_connect_provider.
type DataIamOpenidConnectProvider struct {
	Name string
	Args DataIamOpenidConnectProviderArgs
}

// DataSource returns the Terraform object type for [DataIamOpenidConnectProvider].
func (iocp *DataIamOpenidConnectProvider) DataSource() string {
	return "aws_iam_openid_connect_provider"
}

// LocalName returns the local name for [DataIamOpenidConnectProvider].
func (iocp *DataIamOpenidConnectProvider) LocalName() string {
	return iocp.Name
}

// Configuration returns the configuration (args) for [DataIamOpenidConnectProvider].
func (iocp *DataIamOpenidConnectProvider) Configuration() interface{} {
	return iocp.Args
}

// Attributes returns the attributes for [DataIamOpenidConnectProvider].
func (iocp *DataIamOpenidConnectProvider) Attributes() dataIamOpenidConnectProviderAttributes {
	return dataIamOpenidConnectProviderAttributes{ref: terra.ReferenceDataResource(iocp)}
}

// DataIamOpenidConnectProviderArgs contains the configurations for aws_iam_openid_connect_provider.
type DataIamOpenidConnectProviderArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
}
type dataIamOpenidConnectProviderAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_openid_connect_provider.
func (iocp dataIamOpenidConnectProviderAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iocp.ref.Append("arn"))
}

// ClientIdList returns a reference to field client_id_list of aws_iam_openid_connect_provider.
func (iocp dataIamOpenidConnectProviderAttributes) ClientIdList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](iocp.ref.Append("client_id_list"))
}

// Id returns a reference to field id of aws_iam_openid_connect_provider.
func (iocp dataIamOpenidConnectProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iocp.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_iam_openid_connect_provider.
func (iocp dataIamOpenidConnectProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iocp.ref.Append("tags"))
}

// ThumbprintList returns a reference to field thumbprint_list of aws_iam_openid_connect_provider.
func (iocp dataIamOpenidConnectProviderAttributes) ThumbprintList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](iocp.ref.Append("thumbprint_list"))
}

// Url returns a reference to field url of aws_iam_openid_connect_provider.
func (iocp dataIamOpenidConnectProviderAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(iocp.ref.Append("url"))
}