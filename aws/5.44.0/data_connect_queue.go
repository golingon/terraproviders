// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	dataconnectqueue "github.com/golingon/terraproviders/aws/5.44.0/dataconnectqueue"
)

// NewDataConnectQueue creates a new instance of [DataConnectQueue].
func NewDataConnectQueue(name string, args DataConnectQueueArgs) *DataConnectQueue {
	return &DataConnectQueue{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectQueue)(nil)

// DataConnectQueue represents the Terraform data resource aws_connect_queue.
type DataConnectQueue struct {
	Name string
	Args DataConnectQueueArgs
}

// DataSource returns the Terraform object type for [DataConnectQueue].
func (cq *DataConnectQueue) DataSource() string {
	return "aws_connect_queue"
}

// LocalName returns the local name for [DataConnectQueue].
func (cq *DataConnectQueue) LocalName() string {
	return cq.Name
}

// Configuration returns the configuration (args) for [DataConnectQueue].
func (cq *DataConnectQueue) Configuration() interface{} {
	return cq.Args
}

// Attributes returns the attributes for [DataConnectQueue].
func (cq *DataConnectQueue) Attributes() dataConnectQueueAttributes {
	return dataConnectQueueAttributes{ref: terra.ReferenceDataResource(cq)}
}

// DataConnectQueueArgs contains the configurations for aws_connect_queue.
type DataConnectQueueArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// QueueId: string, optional
	QueueId terra.StringValue `hcl:"queue_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// OutboundCallerConfig: min=0
	OutboundCallerConfig []dataconnectqueue.OutboundCallerConfig `hcl:"outbound_caller_config,block" validate:"min=0"`
}
type dataConnectQueueAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_queue.
func (cq dataConnectQueueAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("arn"))
}

// Description returns a reference to field description of aws_connect_queue.
func (cq dataConnectQueueAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("description"))
}

// HoursOfOperationId returns a reference to field hours_of_operation_id of aws_connect_queue.
func (cq dataConnectQueueAttributes) HoursOfOperationId() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("hours_of_operation_id"))
}

// Id returns a reference to field id of aws_connect_queue.
func (cq dataConnectQueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_queue.
func (cq dataConnectQueueAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("instance_id"))
}

// MaxContacts returns a reference to field max_contacts of aws_connect_queue.
func (cq dataConnectQueueAttributes) MaxContacts() terra.NumberValue {
	return terra.ReferenceAsNumber(cq.ref.Append("max_contacts"))
}

// Name returns a reference to field name of aws_connect_queue.
func (cq dataConnectQueueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("name"))
}

// QueueId returns a reference to field queue_id of aws_connect_queue.
func (cq dataConnectQueueAttributes) QueueId() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("queue_id"))
}

// Status returns a reference to field status of aws_connect_queue.
func (cq dataConnectQueueAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(cq.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_connect_queue.
func (cq dataConnectQueueAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cq.ref.Append("tags"))
}

func (cq dataConnectQueueAttributes) OutboundCallerConfig() terra.ListValue[dataconnectqueue.OutboundCallerConfigAttributes] {
	return terra.ReferenceAsList[dataconnectqueue.OutboundCallerConfigAttributes](cq.ref.Append("outbound_caller_config"))
}
