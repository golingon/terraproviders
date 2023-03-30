// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataIamOpenidConnectProvider(name string, args DataIamOpenidConnectProviderArgs) *DataIamOpenidConnectProvider {
	return &DataIamOpenidConnectProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamOpenidConnectProvider)(nil)

type DataIamOpenidConnectProvider struct {
	Name string
	Args DataIamOpenidConnectProviderArgs
}

func (iocp *DataIamOpenidConnectProvider) DataSource() string {
	return "aws_iam_openid_connect_provider"
}

func (iocp *DataIamOpenidConnectProvider) LocalName() string {
	return iocp.Name
}

func (iocp *DataIamOpenidConnectProvider) Configuration() interface{} {
	return iocp.Args
}

func (iocp *DataIamOpenidConnectProvider) Attributes() dataIamOpenidConnectProviderAttributes {
	return dataIamOpenidConnectProviderAttributes{ref: terra.ReferenceDataResource(iocp)}
}

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

func (iocp dataIamOpenidConnectProviderAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(iocp.ref.Append("arn"))
}

func (iocp dataIamOpenidConnectProviderAttributes) ClientIdList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](iocp.ref.Append("client_id_list"))
}

func (iocp dataIamOpenidConnectProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceString(iocp.ref.Append("id"))
}

func (iocp dataIamOpenidConnectProviderAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](iocp.ref.Append("tags"))
}

func (iocp dataIamOpenidConnectProviderAttributes) ThumbprintList() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](iocp.ref.Append("thumbprint_list"))
}

func (iocp dataIamOpenidConnectProviderAttributes) Url() terra.StringValue {
	return terra.ReferenceString(iocp.ref.Append("url"))
}
