// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_servicebus_topic

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_servicebus_topic.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ast *DataSource) DataSource() string {
	return "azurerm_servicebus_topic"
}

// LocalName returns the local name for [DataSource].
func (ast *DataSource) LocalName() string {
	return ast.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ast *DataSource) Configuration() interface{} {
	return ast.Args
}

// Attributes returns the attributes for [DataSource].
func (ast *DataSource) Attributes() dataAzurermServicebusTopicAttributes {
	return dataAzurermServicebusTopicAttributes{ref: terra.ReferenceDataSource(ast)}
}

// DataArgs contains the configurations for azurerm_servicebus_topic.
type DataArgs struct {
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
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermServicebusTopicAttributes struct {
	ref terra.Reference
}

// AutoDeleteOnIdle returns a reference to field auto_delete_on_idle of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) AutoDeleteOnIdle() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("auto_delete_on_idle"))
}

// DefaultMessageTtl returns a reference to field default_message_ttl of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) DefaultMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("default_message_ttl"))
}

// DuplicateDetectionHistoryTimeWindow returns a reference to field duplicate_detection_history_time_window of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) DuplicateDetectionHistoryTimeWindow() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("duplicate_detection_history_time_window"))
}

// EnableBatchedOperations returns a reference to field enable_batched_operations of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) EnableBatchedOperations() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("enable_batched_operations"))
}

// EnableExpress returns a reference to field enable_express of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) EnableExpress() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("enable_express"))
}

// EnablePartitioning returns a reference to field enable_partitioning of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) EnablePartitioning() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("enable_partitioning"))
}

// Id returns a reference to field id of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("id"))
}

// MaxSizeInMegabytes returns a reference to field max_size_in_megabytes of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) MaxSizeInMegabytes() terra.NumberValue {
	return terra.ReferenceAsNumber(ast.ref.Append("max_size_in_megabytes"))
}

// Name returns a reference to field name of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("namespace_id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("namespace_name"))
}

// RequiresDuplicateDetection returns a reference to field requires_duplicate_detection of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) RequiresDuplicateDetection() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("requires_duplicate_detection"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("resource_group_name"))
}

// Status returns a reference to field status of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ast.ref.Append("status"))
}

// SupportOrdering returns a reference to field support_ordering of azurerm_servicebus_topic.
func (ast dataAzurermServicebusTopicAttributes) SupportOrdering() terra.BoolValue {
	return terra.ReferenceAsBool(ast.ref.Append("support_ordering"))
}

func (ast dataAzurermServicebusTopicAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](ast.ref.Append("timeouts"))
}
