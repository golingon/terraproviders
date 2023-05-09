// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataservicebusnamespaceauthorizationrule "github.com/golingon/terraproviders/azurerm/3.55.0/dataservicebusnamespaceauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicebusNamespaceAuthorizationRule creates a new instance of [DataServicebusNamespaceAuthorizationRule].
func NewDataServicebusNamespaceAuthorizationRule(name string, args DataServicebusNamespaceAuthorizationRuleArgs) *DataServicebusNamespaceAuthorizationRule {
	return &DataServicebusNamespaceAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicebusNamespaceAuthorizationRule)(nil)

// DataServicebusNamespaceAuthorizationRule represents the Terraform data resource azurerm_servicebus_namespace_authorization_rule.
type DataServicebusNamespaceAuthorizationRule struct {
	Name string
	Args DataServicebusNamespaceAuthorizationRuleArgs
}

// DataSource returns the Terraform object type for [DataServicebusNamespaceAuthorizationRule].
func (snar *DataServicebusNamespaceAuthorizationRule) DataSource() string {
	return "azurerm_servicebus_namespace_authorization_rule"
}

// LocalName returns the local name for [DataServicebusNamespaceAuthorizationRule].
func (snar *DataServicebusNamespaceAuthorizationRule) LocalName() string {
	return snar.Name
}

// Configuration returns the configuration (args) for [DataServicebusNamespaceAuthorizationRule].
func (snar *DataServicebusNamespaceAuthorizationRule) Configuration() interface{} {
	return snar.Args
}

// Attributes returns the attributes for [DataServicebusNamespaceAuthorizationRule].
func (snar *DataServicebusNamespaceAuthorizationRule) Attributes() dataServicebusNamespaceAuthorizationRuleAttributes {
	return dataServicebusNamespaceAuthorizationRuleAttributes{ref: terra.ReferenceDataResource(snar)}
}

// DataServicebusNamespaceAuthorizationRuleArgs contains the configurations for azurerm_servicebus_namespace_authorization_rule.
type DataServicebusNamespaceAuthorizationRuleArgs struct {
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
	Timeouts *dataservicebusnamespaceauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type dataServicebusNamespaceAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("namespace_id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_servicebus_namespace_authorization_rule.
func (snar dataServicebusNamespaceAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(snar.ref.Append("secondary_key"))
}

func (snar dataServicebusNamespaceAuthorizationRuleAttributes) Timeouts() dataservicebusnamespaceauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicebusnamespaceauthorizationrule.TimeoutsAttributes](snar.ref.Append("timeouts"))
}
