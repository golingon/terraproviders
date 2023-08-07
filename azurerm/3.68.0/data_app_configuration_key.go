// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappconfigurationkey "github.com/golingon/terraproviders/azurerm/3.68.0/dataappconfigurationkey"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppConfigurationKey creates a new instance of [DataAppConfigurationKey].
func NewDataAppConfigurationKey(name string, args DataAppConfigurationKeyArgs) *DataAppConfigurationKey {
	return &DataAppConfigurationKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppConfigurationKey)(nil)

// DataAppConfigurationKey represents the Terraform data resource azurerm_app_configuration_key.
type DataAppConfigurationKey struct {
	Name string
	Args DataAppConfigurationKeyArgs
}

// DataSource returns the Terraform object type for [DataAppConfigurationKey].
func (ack *DataAppConfigurationKey) DataSource() string {
	return "azurerm_app_configuration_key"
}

// LocalName returns the local name for [DataAppConfigurationKey].
func (ack *DataAppConfigurationKey) LocalName() string {
	return ack.Name
}

// Configuration returns the configuration (args) for [DataAppConfigurationKey].
func (ack *DataAppConfigurationKey) Configuration() interface{} {
	return ack.Args
}

// Attributes returns the attributes for [DataAppConfigurationKey].
func (ack *DataAppConfigurationKey) Attributes() dataAppConfigurationKeyAttributes {
	return dataAppConfigurationKeyAttributes{ref: terra.ReferenceDataResource(ack)}
}

// DataAppConfigurationKeyArgs contains the configurations for azurerm_app_configuration_key.
type DataAppConfigurationKeyArgs struct {
	// ConfigurationStoreId: string, required
	ConfigurationStoreId terra.StringValue `hcl:"configuration_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Timeouts: optional
	Timeouts *dataappconfigurationkey.Timeouts `hcl:"timeouts,block"`
}
type dataAppConfigurationKeyAttributes struct {
	ref terra.Reference
}

// ConfigurationStoreId returns a reference to field configuration_store_id of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) ConfigurationStoreId() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("configuration_store_id"))
}

// ContentType returns a reference to field content_type of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("content_type"))
}

// Etag returns a reference to field etag of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("etag"))
}

// Id returns a reference to field id of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("id"))
}

// Key returns a reference to field key of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("key"))
}

// Label returns a reference to field label of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("label"))
}

// Locked returns a reference to field locked of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Locked() terra.BoolValue {
	return terra.ReferenceAsBool(ack.ref.Append("locked"))
}

// Tags returns a reference to field tags of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ack.ref.Append("tags"))
}

// Type returns a reference to field type of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("type"))
}

// Value returns a reference to field value of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("value"))
}

// VaultKeyReference returns a reference to field vault_key_reference of azurerm_app_configuration_key.
func (ack dataAppConfigurationKeyAttributes) VaultKeyReference() terra.StringValue {
	return terra.ReferenceAsString(ack.ref.Append("vault_key_reference"))
}

func (ack dataAppConfigurationKeyAttributes) Timeouts() dataappconfigurationkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappconfigurationkey.TimeoutsAttributes](ack.ref.Append("timeouts"))
}
