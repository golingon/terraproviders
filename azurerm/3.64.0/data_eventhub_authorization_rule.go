// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventhubauthorizationrule "github.com/golingon/terraproviders/azurerm/3.64.0/dataeventhubauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventhubAuthorizationRule creates a new instance of [DataEventhubAuthorizationRule].
func NewDataEventhubAuthorizationRule(name string, args DataEventhubAuthorizationRuleArgs) *DataEventhubAuthorizationRule {
	return &DataEventhubAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventhubAuthorizationRule)(nil)

// DataEventhubAuthorizationRule represents the Terraform data resource azurerm_eventhub_authorization_rule.
type DataEventhubAuthorizationRule struct {
	Name string
	Args DataEventhubAuthorizationRuleArgs
}

// DataSource returns the Terraform object type for [DataEventhubAuthorizationRule].
func (ear *DataEventhubAuthorizationRule) DataSource() string {
	return "azurerm_eventhub_authorization_rule"
}

// LocalName returns the local name for [DataEventhubAuthorizationRule].
func (ear *DataEventhubAuthorizationRule) LocalName() string {
	return ear.Name
}

// Configuration returns the configuration (args) for [DataEventhubAuthorizationRule].
func (ear *DataEventhubAuthorizationRule) Configuration() interface{} {
	return ear.Args
}

// Attributes returns the attributes for [DataEventhubAuthorizationRule].
func (ear *DataEventhubAuthorizationRule) Attributes() dataEventhubAuthorizationRuleAttributes {
	return dataEventhubAuthorizationRuleAttributes{ref: terra.ReferenceDataResource(ear)}
}

// DataEventhubAuthorizationRuleArgs contains the configurations for azurerm_eventhub_authorization_rule.
type DataEventhubAuthorizationRuleArgs struct {
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *dataeventhubauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type dataEventhubAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// EventhubName returns a reference to field eventhub_name of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("eventhub_name"))
}

// Id returns a reference to field id of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(ear.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(ear.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("namespace_name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("primary_connection_string"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("primary_connection_string_alias"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("secondary_connection_string"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("secondary_connection_string_alias"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("secondary_key"))
}

// Send returns a reference to field send of azurerm_eventhub_authorization_rule.
func (ear dataEventhubAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(ear.ref.Append("send"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Timeouts() dataeventhubauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventhubauthorizationrule.TimeoutsAttributes](ear.ref.Append("timeouts"))
}
