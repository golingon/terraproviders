// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectquickconnect "github.com/golingon/terraproviders/aws/5.7.0/dataconnectquickconnect"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectQuickConnect creates a new instance of [DataConnectQuickConnect].
func NewDataConnectQuickConnect(name string, args DataConnectQuickConnectArgs) *DataConnectQuickConnect {
	return &DataConnectQuickConnect{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectQuickConnect)(nil)

// DataConnectQuickConnect represents the Terraform data resource aws_connect_quick_connect.
type DataConnectQuickConnect struct {
	Name string
	Args DataConnectQuickConnectArgs
}

// DataSource returns the Terraform object type for [DataConnectQuickConnect].
func (cqc *DataConnectQuickConnect) DataSource() string {
	return "aws_connect_quick_connect"
}

// LocalName returns the local name for [DataConnectQuickConnect].
func (cqc *DataConnectQuickConnect) LocalName() string {
	return cqc.Name
}

// Configuration returns the configuration (args) for [DataConnectQuickConnect].
func (cqc *DataConnectQuickConnect) Configuration() interface{} {
	return cqc.Args
}

// Attributes returns the attributes for [DataConnectQuickConnect].
func (cqc *DataConnectQuickConnect) Attributes() dataConnectQuickConnectAttributes {
	return dataConnectQuickConnectAttributes{ref: terra.ReferenceDataResource(cqc)}
}

// DataConnectQuickConnectArgs contains the configurations for aws_connect_quick_connect.
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

// Arn returns a reference to field arn of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("description"))
}

// Id returns a reference to field id of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("name"))
}

// QuickConnectId returns a reference to field quick_connect_id of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) QuickConnectId() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("quick_connect_id"))
}

// Tags returns a reference to field tags of aws_connect_quick_connect.
func (cqc dataConnectQuickConnectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cqc.ref.Append("tags"))
}

func (cqc dataConnectQuickConnectAttributes) QuickConnectConfig() terra.ListValue[dataconnectquickconnect.QuickConnectConfigAttributes] {
	return terra.ReferenceAsList[dataconnectquickconnect.QuickConnectConfigAttributes](cqc.ref.Append("quick_connect_config"))
}
