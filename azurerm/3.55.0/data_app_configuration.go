// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappconfiguration "github.com/golingon/terraproviders/azurerm/3.55.0/dataappconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppConfiguration creates a new instance of [DataAppConfiguration].
func NewDataAppConfiguration(name string, args DataAppConfigurationArgs) *DataAppConfiguration {
	return &DataAppConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppConfiguration)(nil)

// DataAppConfiguration represents the Terraform data resource azurerm_app_configuration.
type DataAppConfiguration struct {
	Name string
	Args DataAppConfigurationArgs
}

// DataSource returns the Terraform object type for [DataAppConfiguration].
func (ac *DataAppConfiguration) DataSource() string {
	return "azurerm_app_configuration"
}

// LocalName returns the local name for [DataAppConfiguration].
func (ac *DataAppConfiguration) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [DataAppConfiguration].
func (ac *DataAppConfiguration) Configuration() interface{} {
	return ac.Args
}

// Attributes returns the attributes for [DataAppConfiguration].
func (ac *DataAppConfiguration) Attributes() dataAppConfigurationAttributes {
	return dataAppConfigurationAttributes{ref: terra.ReferenceDataResource(ac)}
}

// DataAppConfigurationArgs contains the configurations for azurerm_app_configuration.
type DataAppConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Encryption: min=0
	Encryption []dataappconfiguration.Encryption `hcl:"encryption,block" validate:"min=0"`
	// Identity: min=0
	Identity []dataappconfiguration.Identity `hcl:"identity,block" validate:"min=0"`
	// PrimaryReadKey: min=0
	PrimaryReadKey []dataappconfiguration.PrimaryReadKey `hcl:"primary_read_key,block" validate:"min=0"`
	// PrimaryWriteKey: min=0
	PrimaryWriteKey []dataappconfiguration.PrimaryWriteKey `hcl:"primary_write_key,block" validate:"min=0"`
	// SecondaryReadKey: min=0
	SecondaryReadKey []dataappconfiguration.SecondaryReadKey `hcl:"secondary_read_key,block" validate:"min=0"`
	// SecondaryWriteKey: min=0
	SecondaryWriteKey []dataappconfiguration.SecondaryWriteKey `hcl:"secondary_write_key,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappconfiguration.Timeouts `hcl:"timeouts,block"`
}
type dataAppConfigurationAttributes struct {
	ref terra.Reference
}

// Endpoint returns a reference to field endpoint of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

// PublicNetworkAccess returns a reference to field public_network_access of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) PublicNetworkAccess() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("public_network_access"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("public_network_access_enabled"))
}

// PurgeProtectionEnabled returns a reference to field purge_protection_enabled of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("purge_protection_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("sku"))
}

// SoftDeleteRetentionDays returns a reference to field soft_delete_retention_days of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("soft_delete_retention_days"))
}

// Tags returns a reference to field tags of azurerm_app_configuration.
func (ac dataAppConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

func (ac dataAppConfigurationAttributes) Encryption() terra.ListValue[dataappconfiguration.EncryptionAttributes] {
	return terra.ReferenceAsList[dataappconfiguration.EncryptionAttributes](ac.ref.Append("encryption"))
}

func (ac dataAppConfigurationAttributes) Identity() terra.ListValue[dataappconfiguration.IdentityAttributes] {
	return terra.ReferenceAsList[dataappconfiguration.IdentityAttributes](ac.ref.Append("identity"))
}

func (ac dataAppConfigurationAttributes) PrimaryReadKey() terra.ListValue[dataappconfiguration.PrimaryReadKeyAttributes] {
	return terra.ReferenceAsList[dataappconfiguration.PrimaryReadKeyAttributes](ac.ref.Append("primary_read_key"))
}

func (ac dataAppConfigurationAttributes) PrimaryWriteKey() terra.ListValue[dataappconfiguration.PrimaryWriteKeyAttributes] {
	return terra.ReferenceAsList[dataappconfiguration.PrimaryWriteKeyAttributes](ac.ref.Append("primary_write_key"))
}

func (ac dataAppConfigurationAttributes) SecondaryReadKey() terra.ListValue[dataappconfiguration.SecondaryReadKeyAttributes] {
	return terra.ReferenceAsList[dataappconfiguration.SecondaryReadKeyAttributes](ac.ref.Append("secondary_read_key"))
}

func (ac dataAppConfigurationAttributes) SecondaryWriteKey() terra.ListValue[dataappconfiguration.SecondaryWriteKeyAttributes] {
	return terra.ReferenceAsList[dataappconfiguration.SecondaryWriteKeyAttributes](ac.ref.Append("secondary_write_key"))
}

func (ac dataAppConfigurationAttributes) Timeouts() dataappconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappconfiguration.TimeoutsAttributes](ac.ref.Append("timeouts"))
}
