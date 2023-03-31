// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappconfiguration "github.com/golingon/terraproviders/azurerm/3.49.0/dataappconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataAppConfiguration(name string, args DataAppConfigurationArgs) *DataAppConfiguration {
	return &DataAppConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppConfiguration)(nil)

type DataAppConfiguration struct {
	Name string
	Args DataAppConfigurationArgs
}

func (ac *DataAppConfiguration) DataSource() string {
	return "azurerm_app_configuration"
}

func (ac *DataAppConfiguration) LocalName() string {
	return ac.Name
}

func (ac *DataAppConfiguration) Configuration() interface{} {
	return ac.Args
}

func (ac *DataAppConfiguration) Attributes() dataAppConfigurationAttributes {
	return dataAppConfigurationAttributes{ref: terra.ReferenceDataResource(ac)}
}

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

func (ac dataAppConfigurationAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("endpoint"))
}

func (ac dataAppConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("id"))
}

func (ac dataAppConfigurationAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceBool(ac.ref.Append("local_auth_enabled"))
}

func (ac dataAppConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("location"))
}

func (ac dataAppConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("name"))
}

func (ac dataAppConfigurationAttributes) PublicNetworkAccess() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("public_network_access"))
}

func (ac dataAppConfigurationAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(ac.ref.Append("public_network_access_enabled"))
}

func (ac dataAppConfigurationAttributes) PurgeProtectionEnabled() terra.BoolValue {
	return terra.ReferenceBool(ac.ref.Append("purge_protection_enabled"))
}

func (ac dataAppConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("resource_group_name"))
}

func (ac dataAppConfigurationAttributes) Sku() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("sku"))
}

func (ac dataAppConfigurationAttributes) SoftDeleteRetentionDays() terra.NumberValue {
	return terra.ReferenceNumber(ac.ref.Append("soft_delete_retention_days"))
}

func (ac dataAppConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ac.ref.Append("tags"))
}

func (ac dataAppConfigurationAttributes) Encryption() terra.ListValue[dataappconfiguration.EncryptionAttributes] {
	return terra.ReferenceList[dataappconfiguration.EncryptionAttributes](ac.ref.Append("encryption"))
}

func (ac dataAppConfigurationAttributes) Identity() terra.ListValue[dataappconfiguration.IdentityAttributes] {
	return terra.ReferenceList[dataappconfiguration.IdentityAttributes](ac.ref.Append("identity"))
}

func (ac dataAppConfigurationAttributes) PrimaryReadKey() terra.ListValue[dataappconfiguration.PrimaryReadKeyAttributes] {
	return terra.ReferenceList[dataappconfiguration.PrimaryReadKeyAttributes](ac.ref.Append("primary_read_key"))
}

func (ac dataAppConfigurationAttributes) PrimaryWriteKey() terra.ListValue[dataappconfiguration.PrimaryWriteKeyAttributes] {
	return terra.ReferenceList[dataappconfiguration.PrimaryWriteKeyAttributes](ac.ref.Append("primary_write_key"))
}

func (ac dataAppConfigurationAttributes) SecondaryReadKey() terra.ListValue[dataappconfiguration.SecondaryReadKeyAttributes] {
	return terra.ReferenceList[dataappconfiguration.SecondaryReadKeyAttributes](ac.ref.Append("secondary_read_key"))
}

func (ac dataAppConfigurationAttributes) SecondaryWriteKey() terra.ListValue[dataappconfiguration.SecondaryWriteKeyAttributes] {
	return terra.ReferenceList[dataappconfiguration.SecondaryWriteKeyAttributes](ac.ref.Append("secondary_write_key"))
}

func (ac dataAppConfigurationAttributes) Timeouts() dataappconfiguration.TimeoutsAttributes {
	return terra.ReferenceSingle[dataappconfiguration.TimeoutsAttributes](ac.ref.Append("timeouts"))
}
