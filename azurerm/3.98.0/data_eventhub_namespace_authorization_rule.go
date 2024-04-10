// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	dataeventhubnamespaceauthorizationrule "github.com/golingon/terraproviders/azurerm/3.98.0/dataeventhubnamespaceauthorizationrule"
)

// NewDataEventhubNamespaceAuthorizationRule creates a new instance of [DataEventhubNamespaceAuthorizationRule].
func NewDataEventhubNamespaceAuthorizationRule(name string, args DataEventhubNamespaceAuthorizationRuleArgs) *DataEventhubNamespaceAuthorizationRule {
	return &DataEventhubNamespaceAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventhubNamespaceAuthorizationRule)(nil)

// DataEventhubNamespaceAuthorizationRule represents the Terraform data resource azurerm_eventhub_namespace_authorization_rule.
type DataEventhubNamespaceAuthorizationRule struct {
	Name string
	Args DataEventhubNamespaceAuthorizationRuleArgs
}

// DataSource returns the Terraform object type for [DataEventhubNamespaceAuthorizationRule].
func (enar *DataEventhubNamespaceAuthorizationRule) DataSource() string {
	return "azurerm_eventhub_namespace_authorization_rule"
}

// LocalName returns the local name for [DataEventhubNamespaceAuthorizationRule].
func (enar *DataEventhubNamespaceAuthorizationRule) LocalName() string {
	return enar.Name
}

// Configuration returns the configuration (args) for [DataEventhubNamespaceAuthorizationRule].
func (enar *DataEventhubNamespaceAuthorizationRule) Configuration() interface{} {
	return enar.Args
}

// Attributes returns the attributes for [DataEventhubNamespaceAuthorizationRule].
func (enar *DataEventhubNamespaceAuthorizationRule) Attributes() dataEventhubNamespaceAuthorizationRuleAttributes {
	return dataEventhubNamespaceAuthorizationRuleAttributes{ref: terra.ReferenceDataResource(enar)}
}

// DataEventhubNamespaceAuthorizationRuleArgs contains the configurations for azurerm_eventhub_namespace_authorization_rule.
type DataEventhubNamespaceAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataeventhubnamespaceauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type dataEventhubNamespaceAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(enar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(enar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(enar.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_eventhub_namespace_authorization_rule.
func (enar dataEventhubNamespaceAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(enar.ref.Append("send"))
}

func (enar dataEventhubNamespaceAuthorizationRuleAttributes) Timeouts() dataeventhubnamespaceauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventhubnamespaceauthorizationrule.TimeoutsAttributes](enar.ref.Append("timeouts"))
}
