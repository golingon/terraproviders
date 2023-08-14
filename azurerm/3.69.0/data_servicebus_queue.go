// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataservicebusqueue "github.com/golingon/terraproviders/azurerm/3.69.0/dataservicebusqueue"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicebusQueue creates a new instance of [DataServicebusQueue].
func NewDataServicebusQueue(name string, args DataServicebusQueueArgs) *DataServicebusQueue {
	return &DataServicebusQueue{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicebusQueue)(nil)

// DataServicebusQueue represents the Terraform data resource azurerm_servicebus_queue.
type DataServicebusQueue struct {
	Name string
	Args DataServicebusQueueArgs
}

// DataSource returns the Terraform object type for [DataServicebusQueue].
func (sq *DataServicebusQueue) DataSource() string {
	return "azurerm_servicebus_queue"
}

// LocalName returns the local name for [DataServicebusQueue].
func (sq *DataServicebusQueue) LocalName() string {
	return sq.Name
}

// Configuration returns the configuration (args) for [DataServicebusQueue].
func (sq *DataServicebusQueue) Configuration() interface{} {
	return sq.Args
}

// Attributes returns the attributes for [DataServicebusQueue].
func (sq *DataServicebusQueue) Attributes() dataServicebusQueueAttributes {
	return dataServicebusQueueAttributes{ref: terra.ReferenceDataResource(sq)}
}

// DataServicebusQueueArgs contains the configurations for azurerm_servicebus_queue.
type DataServicebusQueueArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, optional
	NamespaceId terra.StringValue `hcl:"namespace_id,attr"`
	// NamespaceName: string, optional
	NamespaceName terra.StringValue `hcl:"namespace_name,attr"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Timeouts: optional
	Timeouts *dataservicebusqueue.Timeouts `hcl:"timeouts,block"`
}
type dataServicebusQueueAttributes struct {
	ref terra.Reference
}

// AutoDeleteOnIdle returns a reference to field auto_delete_on_idle of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) AutoDeleteOnIdle() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("auto_delete_on_idle"))
}

// DeadLetteringOnMessageExpiration returns a reference to field dead_lettering_on_message_expiration of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) DeadLetteringOnMessageExpiration() terra.BoolValue {
	return terra.ReferenceAsBool(sq.ref.Append("dead_lettering_on_message_expiration"))
}

// DefaultMessageTtl returns a reference to field default_message_ttl of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) DefaultMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("default_message_ttl"))
}

// DuplicateDetectionHistoryTimeWindow returns a reference to field duplicate_detection_history_time_window of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) DuplicateDetectionHistoryTimeWindow() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("duplicate_detection_history_time_window"))
}

// EnableBatchedOperations returns a reference to field enable_batched_operations of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) EnableBatchedOperations() terra.BoolValue {
	return terra.ReferenceAsBool(sq.ref.Append("enable_batched_operations"))
}

// EnableExpress returns a reference to field enable_express of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) EnableExpress() terra.BoolValue {
	return terra.ReferenceAsBool(sq.ref.Append("enable_express"))
}

// EnablePartitioning returns a reference to field enable_partitioning of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) EnablePartitioning() terra.BoolValue {
	return terra.ReferenceAsBool(sq.ref.Append("enable_partitioning"))
}

// ForwardDeadLetteredMessagesTo returns a reference to field forward_dead_lettered_messages_to of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) ForwardDeadLetteredMessagesTo() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("forward_dead_lettered_messages_to"))
}

// ForwardTo returns a reference to field forward_to of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) ForwardTo() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("forward_to"))
}

// Id returns a reference to field id of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("id"))
}

// LockDuration returns a reference to field lock_duration of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) LockDuration() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("lock_duration"))
}

// MaxDeliveryCount returns a reference to field max_delivery_count of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) MaxDeliveryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sq.ref.Append("max_delivery_count"))
}

// MaxSizeInMegabytes returns a reference to field max_size_in_megabytes of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) MaxSizeInMegabytes() terra.NumberValue {
	return terra.ReferenceAsNumber(sq.ref.Append("max_size_in_megabytes"))
}

// Name returns a reference to field name of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("namespace_id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("namespace_name"))
}

// RequiresDuplicateDetection returns a reference to field requires_duplicate_detection of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) RequiresDuplicateDetection() terra.BoolValue {
	return terra.ReferenceAsBool(sq.ref.Append("requires_duplicate_detection"))
}

// RequiresSession returns a reference to field requires_session of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) RequiresSession() terra.BoolValue {
	return terra.ReferenceAsBool(sq.ref.Append("requires_session"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_servicebus_queue.
func (sq dataServicebusQueueAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sq.ref.Append("status"))
}

func (sq dataServicebusQueueAttributes) Timeouts() dataservicebusqueue.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicebusqueue.TimeoutsAttributes](sq.ref.Append("timeouts"))
}
