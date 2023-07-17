// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultencryptedvalue "github.com/golingon/terraproviders/azurerm/3.65.0/datakeyvaultencryptedvalue"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultEncryptedValue creates a new instance of [DataKeyVaultEncryptedValue].
func NewDataKeyVaultEncryptedValue(name string, args DataKeyVaultEncryptedValueArgs) *DataKeyVaultEncryptedValue {
	return &DataKeyVaultEncryptedValue{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultEncryptedValue)(nil)

// DataKeyVaultEncryptedValue represents the Terraform data resource azurerm_key_vault_encrypted_value.
type DataKeyVaultEncryptedValue struct {
	Name string
	Args DataKeyVaultEncryptedValueArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultEncryptedValue].
func (kvev *DataKeyVaultEncryptedValue) DataSource() string {
	return "azurerm_key_vault_encrypted_value"
}

// LocalName returns the local name for [DataKeyVaultEncryptedValue].
func (kvev *DataKeyVaultEncryptedValue) LocalName() string {
	return kvev.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultEncryptedValue].
func (kvev *DataKeyVaultEncryptedValue) Configuration() interface{} {
	return kvev.Args
}

// Attributes returns the attributes for [DataKeyVaultEncryptedValue].
func (kvev *DataKeyVaultEncryptedValue) Attributes() dataKeyVaultEncryptedValueAttributes {
	return dataKeyVaultEncryptedValueAttributes{ref: terra.ReferenceDataResource(kvev)}
}

// DataKeyVaultEncryptedValueArgs contains the configurations for azurerm_key_vault_encrypted_value.
type DataKeyVaultEncryptedValueArgs struct {
	// Algorithm: string, required
	Algorithm terra.StringValue `hcl:"algorithm,attr" validate:"required"`
	// EncryptedData: string, optional
	EncryptedData terra.StringValue `hcl:"encrypted_data,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultKeyId: string, required
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr" validate:"required"`
	// PlainTextValue: string, optional
	PlainTextValue terra.StringValue `hcl:"plain_text_value,attr"`
	// Timeouts: optional
	Timeouts *datakeyvaultencryptedvalue.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultEncryptedValueAttributes struct {
	ref terra.Reference
}

// Algorithm returns a reference to field algorithm of azurerm_key_vault_encrypted_value.
func (kvev dataKeyVaultEncryptedValueAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceAsString(kvev.ref.Append("algorithm"))
}

// DecodedPlainTextValue returns a reference to field decoded_plain_text_value of azurerm_key_vault_encrypted_value.
func (kvev dataKeyVaultEncryptedValueAttributes) DecodedPlainTextValue() terra.StringValue {
	return terra.ReferenceAsString(kvev.ref.Append("decoded_plain_text_value"))
}

// EncryptedData returns a reference to field encrypted_data of azurerm_key_vault_encrypted_value.
func (kvev dataKeyVaultEncryptedValueAttributes) EncryptedData() terra.StringValue {
	return terra.ReferenceAsString(kvev.ref.Append("encrypted_data"))
}

// Id returns a reference to field id of azurerm_key_vault_encrypted_value.
func (kvev dataKeyVaultEncryptedValueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvev.ref.Append("id"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_key_vault_encrypted_value.
func (kvev dataKeyVaultEncryptedValueAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(kvev.ref.Append("key_vault_key_id"))
}

// PlainTextValue returns a reference to field plain_text_value of azurerm_key_vault_encrypted_value.
func (kvev dataKeyVaultEncryptedValueAttributes) PlainTextValue() terra.StringValue {
	return terra.ReferenceAsString(kvev.ref.Append("plain_text_value"))
}

func (kvev dataKeyVaultEncryptedValueAttributes) Timeouts() datakeyvaultencryptedvalue.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultencryptedvalue.TimeoutsAttributes](kvev.ref.Append("timeouts"))
}
