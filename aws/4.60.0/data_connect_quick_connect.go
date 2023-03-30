// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectquickconnect "github.com/golingon/terraproviders/aws/4.60.0/dataconnectquickconnect"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataConnectQuickConnect(name string, args DataConnectQuickConnectArgs) *DataConnectQuickConnect {
	return &DataConnectQuickConnect{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectQuickConnect)(nil)

type DataConnectQuickConnect struct {
	Name string
	Args DataConnectQuickConnectArgs
}

func (cqc *DataConnectQuickConnect) DataSource() string {
	return "aws_connect_quick_connect"
}

func (cqc *DataConnectQuickConnect) LocalName() string {
	return cqc.Name
}

func (cqc *DataConnectQuickConnect) Configuration() interface{} {
	return cqc.Args
}

func (cqc *DataConnectQuickConnect) Attributes() dataConnectQuickConnectAttributes {
	return dataConnectQuickConnectAttributes{ref: terra.ReferenceDataResource(cqc)}
}

type DataConnectQuickConnectArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// QuickConnectId: string, optional
	QuickConnectId terra.StringValue `hcl:"quick_connect_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// QuickConnectConfig: min=0
	QuickConnectConfig []dataconnectquickconnect.QuickConnectConfig `hcl:"quick_connect_config,block" validate:"min=0"`
}
type dataConnectQuickConnectAttributes struct {
	ref terra.Reference
}

func (cqc dataConnectQuickConnectAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cqc.ref.Append("arn"))
}

func (cqc dataConnectQuickConnectAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cqc.ref.Append("description"))
}

func (cqc dataConnectQuickConnectAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cqc.ref.Append("id"))
}

func (cqc dataConnectQuickConnectAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(cqc.ref.Append("instance_id"))
}

func (cqc dataConnectQuickConnectAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cqc.ref.Append("name"))
}

func (cqc dataConnectQuickConnectAttributes) QuickConnectId() terra.StringValue {
	return terra.ReferenceString(cqc.ref.Append("quick_connect_id"))
}

func (cqc dataConnectQuickConnectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cqc.ref.Append("tags"))
}

func (cqc dataConnectQuickConnectAttributes) QuickConnectConfig() terra.ListValue[dataconnectquickconnect.QuickConnectConfigAttributes] {
	return terra.ReferenceList[dataconnectquickconnect.QuickConnectConfigAttributes](cqc.ref.Append("quick_connect_config"))
}
