// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataservicebusnamespacedisasterrecoveryconfig "github.com/golingon/terraproviders/azurerm/3.58.0/dataservicebusnamespacedisasterrecoveryconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServicebusNamespaceDisasterRecoveryConfig creates a new instance of [DataServicebusNamespaceDisasterRecoveryConfig].
func NewDataServicebusNamespaceDisasterRecoveryConfig(name string, args DataServicebusNamespaceDisasterRecoveryConfigArgs) *DataServicebusNamespaceDisasterRecoveryConfig {
	return &DataServicebusNamespaceDisasterRecoveryConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServicebusNamespaceDisasterRecoveryConfig)(nil)

// DataServicebusNamespaceDisasterRecoveryConfig represents the Terraform data resource azurerm_servicebus_namespace_disaster_recovery_config.
type DataServicebusNamespaceDisasterRecoveryConfig struct {
	Name string
	Args DataServicebusNamespaceDisasterRecoveryConfigArgs
}

// DataSource returns the Terraform object type for [DataServicebusNamespaceDisasterRecoveryConfig].
func (sndrc *DataServicebusNamespaceDisasterRecoveryConfig) DataSource() string {
	return "azurerm_servicebus_namespace_disaster_recovery_config"
}

// LocalName returns the local name for [DataServicebusNamespaceDisasterRecoveryConfig].
func (sndrc *DataServicebusNamespaceDisasterRecoveryConfig) LocalName() string {
	return sndrc.Name
}

// Configuration returns the configuration (args) for [DataServicebusNamespaceDisasterRecoveryConfig].
func (sndrc *DataServicebusNamespaceDisasterRecoveryConfig) Configuration() interface{} {
	return sndrc.Args
}

// Attributes returns the attributes for [DataServicebusNamespaceDisasterRecoveryConfig].
func (sndrc *DataServicebusNamespaceDisasterRecoveryConfig) Attributes() dataServicebusNamespaceDisasterRecoveryConfigAttributes {
	return dataServicebusNamespaceDisasterRecoveryConfigAttributes{ref: terra.ReferenceDataResource(sndrc)}
}

// DataServicebusNamespaceDisasterRecoveryConfigArgs contains the configurations for azurerm_servicebus_namespace_disaster_recovery_config.
type DataServicebusNamespaceDisasterRecoveryConfigArgs struct {
	// AliasAuthorizationRuleId: string, optional
	AliasAuthorizationRuleId terra.StringValue `hcl:"alias_authorization_rule_id,attr"`
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
	Timeouts *dataservicebusnamespacedisasterrecoveryconfig.Timeouts `hcl:"timeouts,block"`
}
type dataServicebusNamespaceDisasterRecoveryConfigAttributes struct {
	ref terra.Reference
}

// AliasAuthorizationRuleId returns a reference to field alias_authorization_rule_id of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) AliasAuthorizationRuleId() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("alias_authorization_rule_id"))
}

// DefaultPrimaryKey returns a reference to field default_primary_key of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) DefaultPrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("default_primary_key"))
}

// DefaultSecondaryKey returns a reference to field default_secondary_key of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) DefaultSecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("default_secondary_key"))
}

// Id returns a reference to field id of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("namespace_id"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("namespace_name"))
}

// PartnerNamespaceId returns a reference to field partner_namespace_id of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) PartnerNamespaceId() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("partner_namespace_id"))
}

// PrimaryConnectionStringAlias returns a reference to field primary_connection_string_alias of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("primary_connection_string_alias"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("resource_group_name"))
}

// SecondaryConnectionStringAlias returns a reference to field secondary_connection_string_alias of azurerm_servicebus_namespace_disaster_recovery_config.
func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceAsString(sndrc.ref.Append("secondary_connection_string_alias"))
}

func (sndrc dataServicebusNamespaceDisasterRecoveryConfigAttributes) Timeouts() dataservicebusnamespacedisasterrecoveryconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataservicebusnamespacedisasterrecoveryconfig.TimeoutsAttributes](sndrc.ref.Append("timeouts"))
}
