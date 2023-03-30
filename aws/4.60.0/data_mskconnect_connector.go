// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataMskconnectConnector(name string, args DataMskconnectConnectorArgs) *DataMskconnectConnector {
	return &DataMskconnectConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMskconnectConnector)(nil)

type DataMskconnectConnector struct {
	Name string
	Args DataMskconnectConnectorArgs
}

func (mc *DataMskconnectConnector) DataSource() string {
	return "aws_mskconnect_connector"
}

func (mc *DataMskconnectConnector) LocalName() string {
	return mc.Name
}

func (mc *DataMskconnectConnector) Configuration() interface{} {
	return mc.Args
}

func (mc *DataMskconnectConnector) Attributes() dataMskconnectConnectorAttributes {
	return dataMskconnectConnectorAttributes{ref: terra.ReferenceDataResource(mc)}
}

type DataMskconnectConnectorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataMskconnectConnectorAttributes struct {
	ref terra.Reference
}

func (mc dataMskconnectConnectorAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(mc.ref.Append("arn"))
}

func (mc dataMskconnectConnectorAttributes) Description() terra.StringValue {
	return terra.ReferenceString(mc.ref.Append("description"))
}

func (mc dataMskconnectConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mc.ref.Append("id"))
}

func (mc dataMskconnectConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mc.ref.Append("name"))
}

func (mc dataMskconnectConnectorAttributes) Version() terra.StringValue {
	return terra.ReferenceString(mc.ref.Append("version"))
}
