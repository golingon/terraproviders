// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultsecret "github.com/golingon/terraproviders/azurerm/3.58.0/datakeyvaultsecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultSecret creates a new instance of [DataKeyVaultSecret].
func NewDataKeyVaultSecret(name string, args DataKeyVaultSecretArgs) *DataKeyVaultSecret {
	return &DataKeyVaultSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultSecret)(nil)

// DataKeyVaultSecret represents the Terraform data resource azurerm_key_vault_secret.
type DataKeyVaultSecret struct {
	Name string
	Args DataKeyVaultSecretArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultSecret].
func (kvs *DataKeyVaultSecret) DataSource() string {
	return "azurerm_key_vault_secret"
}

// LocalName returns the local name for [DataKeyVaultSecret].
func (kvs *DataKeyVaultSecret) LocalName() string {
	return kvs.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultSecret].
func (kvs *DataKeyVaultSecret) Configuration() interface{} {
	return kvs.Args
}

// Attributes returns the attributes for [DataKeyVaultSecret].
func (kvs *DataKeyVaultSecret) Attributes() dataKeyVaultSecretAttributes {
	return dataKeyVaultSecretAttributes{ref: terra.ReferenceDataResource(kvs)}
}

// DataKeyVaultSecretArgs contains the configurations for azurerm_key_vault_secret.
type DataKeyVaultSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Timeouts: optional
	Timeouts *datakeyvaultsecret.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultSecretAttributes struct {
	ref terra.Reference
}

// ContentType returns a reference to field content_type of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("content_type"))
}

// ExpirationDate returns a reference to field expiration_date of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("expiration_date"))
}

// Id returns a reference to field id of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("id"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("key_vault_id"))
}

// Name returns a reference to field name of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("name"))
}

// NotBeforeDate returns a reference to field not_before_date of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) NotBeforeDate() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("not_before_date"))
}

// ResourceId returns a reference to field resource_id of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("resource_id"))
}

// ResourceVersionlessId returns a reference to field resource_versionless_id of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) ResourceVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("resource_versionless_id"))
}

// Tags returns a reference to field tags of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvs.ref.Append("tags"))
}

// Value returns a reference to field value of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("value"))
}

// Version returns a reference to field version of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("version"))
}

// VersionlessId returns a reference to field versionless_id of azurerm_key_vault_secret.
func (kvs dataKeyVaultSecretAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvs.ref.Append("versionless_id"))
}

func (kvs dataKeyVaultSecretAttributes) Timeouts() datakeyvaultsecret.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultsecret.TimeoutsAttributes](kvs.ref.Append("timeouts"))
}
