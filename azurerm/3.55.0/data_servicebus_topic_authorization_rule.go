// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataservicebustopicauthorizationrule "github.com/golingon/terraproviders/azurerm/3.55.0/dataservicebustopicauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicebusTopicAuthorizationRule creates a new instance of [DataServicebusTopicAuthorizationRule].
func NewDataServicebusTopicAuthorizationRule(name string, args DataServicebusTopicAuthorizationRuleArgs) *DataServicebusTopicAuthorizationRule {
	return &DataServicebusTopicAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicebusTopicAuthorizationRule)(nil)

// DataServicebusTopicAuthorizationRule represents the Terraform data resource azurerm_servicebus_topic_authorization_rule.
type DataServicebusTopicAuthorizationRule struct {
	Name string
	Args DataServicebusTopicAuthorizationRuleArgs
}

// DataSource returns the Terraform object type for [DataServicebusTopicAuthorizationRule].
func (star *DataServicebusTopicAuthorizationRule) DataSource() string {
	return "azurerm_servicebus_topic_authorization_rule"
}

// LocalName returns the local name for [DataServicebusTopicAuthorizationRule].
func (star *DataServicebusTopicAuthorizationRule) LocalName() string {
	return star.Name
}

// Configuration returns the configuration (args) for [DataServicebusTopicAuthorizationRule].
func (star *DataServicebusTopicAuthorizationRule) Configuration() interface{} {
	return star.Args
}

// Attributes returns the attributes for [DataServicebusTopicAuthorizationRule].
func (star *DataServicebusTopicAuthorizationRule) Attributes() dataServicebusTopicAuthorizationRuleAttributes {
	return dataServicebusTopicAuthorizationRuleAttributes{ref: terra.ReferenceDataResource(star)}
}

// DataServicebusTopicAuthorizationRuleArgs contains the configurations for azurerm_servicebus_topic_authorization_rule.
type DataServicebusTopicAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, optional
	NamespaceName terra.StringValue `hcl:"namespace_name,attr"`
	// QueueName: string, optional
	QueueName terra.StringValue `hcl:"queue_name,attr"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// TopicId: string, optional
	TopicId terra.StringValue `hcl:"topic_id,attr"`
	// TopicName: string, optional
	TopicName terra.StringValue `hcl:"topic_name,attr"`
	// Timeouts: optional
	Timeouts *dataservicebustopicauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type dataServicebusTopicAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(star.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(star.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("primary_key"))
}

// QueueName returns a reference to field queue_name of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) QueueName() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("queue_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(star.ref.Append("send"))
}

// TopicId returns a reference to field topic_id of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) TopicId() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("topic_id"))
}

// TopicName returns a reference to field topic_name of azurerm_servicebus_topic_authorization_rule.
func (star dataServicebusTopicAuthorizationRuleAttributes) TopicName() terra.StringValue {
	return terra.ReferenceAsString(star.ref.Append("topic_name"))
}

func (star dataServicebusTopicAuthorizationRuleAttributes) Timeouts() dataservicebustopicauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicebustopicauthorizationrule.TimeoutsAttributes](star.ref.Append("timeouts"))
}
