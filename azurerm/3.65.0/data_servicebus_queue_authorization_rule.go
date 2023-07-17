// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataservicebusqueueauthorizationrule "github.com/golingon/terraproviders/azurerm/3.65.0/dataservicebusqueueauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicebusQueueAuthorizationRule creates a new instance of [DataServicebusQueueAuthorizationRule].
func NewDataServicebusQueueAuthorizationRule(name string, args DataServicebusQueueAuthorizationRuleArgs) *DataServicebusQueueAuthorizationRule {
	return &DataServicebusQueueAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicebusQueueAuthorizationRule)(nil)

// DataServicebusQueueAuthorizationRule represents the Terraform data resource azurerm_servicebus_queue_authorization_rule.
type DataServicebusQueueAuthorizationRule struct {
	Name string
	Args DataServicebusQueueAuthorizationRuleArgs
}

// DataSource returns the Terraform object type for [DataServicebusQueueAuthorizationRule].
func (sqar *DataServicebusQueueAuthorizationRule) DataSource() string {
	return "azurerm_servicebus_queue_authorization_rule"
}

// LocalName returns the local name for [DataServicebusQueueAuthorizationRule].
func (sqar *DataServicebusQueueAuthorizationRule) LocalName() string {
	return sqar.Name
}

// Configuration returns the configuration (args) for [DataServicebusQueueAuthorizationRule].
func (sqar *DataServicebusQueueAuthorizationRule) Configuration() interface{} {
	return sqar.Args
}

// Attributes returns the attributes for [DataServicebusQueueAuthorizationRule].
func (sqar *DataServicebusQueueAuthorizationRule) Attributes() dataServicebusQueueAuthorizationRuleAttributes {
	return dataServicebusQueueAuthorizationRuleAttributes{ref: terra.ReferenceDataResource(sqar)}
}

// DataServicebusQueueAuthorizationRuleArgs contains the configurations for azurerm_servicebus_queue_authorization_rule.
type DataServicebusQueueAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, optional
	NamespaceName terra.StringValue `hcl:"namespace_name,attr"`
	// QueueId: string, optional
	QueueId terra.StringValue `hcl:"queue_id,attr"`
	// QueueName: string, optional
	QueueName terra.StringValue `hcl:"queue_name,attr"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// Timeouts: optional
	Timeouts *dataservicebusqueueauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type dataServicebusQueueAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(sqar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(sqar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("primary_key"))
}

// QueueId returns a reference to field queue_id of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) QueueId() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("queue_id"))
}

// QueueName returns a reference to field queue_name of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) QueueName() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("queue_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(sqar.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_servicebus_queue_authorization_rule.
func (sqar dataServicebusQueueAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(sqar.ref.Append("send"))
}

func (sqar dataServicebusQueueAuthorizationRuleAttributes) Timeouts() dataservicebusqueueauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicebusqueueauthorizationrule.TimeoutsAttributes](sqar.ref.Append("timeouts"))
}
