// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnecthoursofoperation "github.com/golingon/terraproviders/aws/4.60.0/dataconnecthoursofoperation"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectHoursOfOperation creates a new instance of [DataConnectHoursOfOperation].
func NewDataConnectHoursOfOperation(name string, args DataConnectHoursOfOperationArgs) *DataConnectHoursOfOperation {
	return &DataConnectHoursOfOperation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectHoursOfOperation)(nil)

// DataConnectHoursOfOperation represents the Terraform data resource aws_connect_hours_of_operation.
type DataConnectHoursOfOperation struct {
	Name string
	Args DataConnectHoursOfOperationArgs
}

// DataSource returns the Terraform object type for [DataConnectHoursOfOperation].
func (choo *DataConnectHoursOfOperation) DataSource() string {
	return "aws_connect_hours_of_operation"
}

// LocalName returns the local name for [DataConnectHoursOfOperation].
func (choo *DataConnectHoursOfOperation) LocalName() string {
	return choo.Name
}

// Configuration returns the configuration (args) for [DataConnectHoursOfOperation].
func (choo *DataConnectHoursOfOperation) Configuration() interface{} {
	return choo.Args
}

// Attributes returns the attributes for [DataConnectHoursOfOperation].
func (choo *DataConnectHoursOfOperation) Attributes() dataConnectHoursOfOperationAttributes {
	return dataConnectHoursOfOperationAttributes{ref: terra.ReferenceDataResource(choo)}
}

// DataConnectHoursOfOperationArgs contains the configurations for aws_connect_hours_of_operation.
type DataConnectHoursOfOperationArgs struct {
	// HoursOfOperationId: string, optional
	HoursOfOperationId terra.StringValue `hcl:"hours_of_operation_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Config: min=0
	Config []dataconnecthoursofoperation.Config `hcl:"config,block" validate:"min=0"`
}
type dataConnectHoursOfOperationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("arn"))
}

// Description returns a reference to field description of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("description"))
}

// HoursOfOperationArn returns a reference to field hours_of_operation_arn of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) HoursOfOperationArn() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("hours_of_operation_arn"))
}

// HoursOfOperationId returns a reference to field hours_of_operation_id of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) HoursOfOperationId() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("hours_of_operation_id"))
}

// Id returns a reference to field id of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](choo.ref.Append("tags"))
}

// TimeZone returns a reference to field time_zone of aws_connect_hours_of_operation.
func (choo dataConnectHoursOfOperationAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("time_zone"))
}

func (choo dataConnectHoursOfOperationAttributes) Config() terra.SetValue[dataconnecthoursofoperation.ConfigAttributes] {
	return terra.ReferenceAsSet[dataconnecthoursofoperation.ConfigAttributes](choo.ref.Append("config"))
}
