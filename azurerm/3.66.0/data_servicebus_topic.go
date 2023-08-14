// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataservicebustopic "github.com/golingon/terraproviders/azurerm/3.66.0/dataservicebustopic"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicebusTopic creates a new instance of [DataServicebusTopic].
func NewDataServicebusTopic(name string, args DataServicebusTopicArgs) *DataServicebusTopic {
	return &DataServicebusTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicebusTopic)(nil)

// DataServicebusTopic represents the Terraform data resource azurerm_servicebus_topic.
type DataServicebusTopic struct {
	Name string
	Args DataServicebusTopicArgs
}

// DataSource returns the Terraform object type for [DataServicebusTopic].
func (st *DataServicebusTopic) DataSource() string {
	return "azurerm_servicebus_topic"
}

// LocalName returns the local name for [DataServicebusTopic].
func (st *DataServicebusTopic) LocalName() string {
	return st.Name
}

// Configuration returns the configuration (args) for [DataServicebusTopic].
func (st *DataServicebusTopic) Configuration() interface{} {
	return st.Args
}

// Attributes returns the attributes for [DataServicebusTopic].
func (st *DataServicebusTopic) Attributes() dataServicebusTopicAttributes {
	return dataServicebusTopicAttributes{ref: terra.ReferenceDataResource(st)}
}

// DataServicebusTopicArgs contains the configurations for azurerm_servicebus_topic.
type DataServicebusTopicArgs struct {
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
	Timeouts *dataservicebustopic.Timeouts `hcl:"timeouts,block"`
}
type dataServicebusTopicAttributes struct {
	ref terra.Reference
}

// AutoDeleteOnIdle returns a reference to field auto_delete_on_idle of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) AutoDeleteOnIdle() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("auto_delete_on_idle"))
}

// DefaultMessageTtl returns a reference to field default_message_ttl of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) DefaultMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("default_message_ttl"))
}

// DuplicateDetectionHistoryTimeWindow returns a reference to field duplicate_detection_history_time_window of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) DuplicateDetectionHistoryTimeWindow() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("duplicate_detection_history_time_window"))
}

// EnableBatchedOperations returns a reference to field enable_batched_operations of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) EnableBatchedOperations() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("enable_batched_operations"))
}

// EnableExpress returns a reference to field enable_express of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) EnableExpress() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("enable_express"))
}

// EnablePartitioning returns a reference to field enable_partitioning of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) EnablePartitioning() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("enable_partitioning"))
}

// Id returns a reference to field id of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("id"))
}

// MaxSizeInMegabytes returns a reference to field max_size_in_megabytes of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) MaxSizeInMegabytes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("max_size_in_megabytes"))
}

// Name returns a reference to field name of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("namespace_id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("namespace_name"))
}

// RequiresDuplicateDetection returns a reference to field requires_duplicate_detection of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) RequiresDuplicateDetection() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("requires_duplicate_detection"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("status"))
}

// SupportOrdering returns a reference to field support_ordering of azurerm_servicebus_topic.
func (st dataServicebusTopicAttributes) SupportOrdering() terra.BoolValue {
	return terra.ReferenceAsBool(st.ref.Append("support_ordering"))
}

func (st dataServicebusTopicAttributes) Timeouts() dataservicebustopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicebustopic.TimeoutsAttributes](st.ref.Append("timeouts"))
}