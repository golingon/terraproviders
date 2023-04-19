// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultkey "github.com/golingon/terraproviders/azurerm/3.52.0/datakeyvaultkey"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKeyVaultKey creates a new instance of [DataKeyVaultKey].
func NewDataKeyVaultKey(name string, args DataKeyVaultKeyArgs) *DataKeyVaultKey {
	return &DataKeyVaultKey{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultKey)(nil)

// DataKeyVaultKey represents the Terraform data resource azurerm_key_vault_key.
type DataKeyVaultKey struct {
	Name string
	Args DataKeyVaultKeyArgs
}

// DataSource returns the Terraform object type for [DataKeyVaultKey].
func (kvk *DataKeyVaultKey) DataSource() string {
	return "azurerm_key_vault_key"
}

// LocalName returns the local name for [DataKeyVaultKey].
func (kvk *DataKeyVaultKey) LocalName() string {
	return kvk.Name
}

// Configuration returns the configuration (args) for [DataKeyVaultKey].
func (kvk *DataKeyVaultKey) Configuration() interface{} {
	return kvk.Args
}

// Attributes returns the attributes for [DataKeyVaultKey].
func (kvk *DataKeyVaultKey) Attributes() dataKeyVaultKeyAttributes {
	return dataKeyVaultKeyAttributes{ref: terra.ReferenceDataResource(kvk)}
}

// DataKeyVaultKeyArgs contains the configurations for azurerm_key_vault_key.
type DataKeyVaultKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datakeyvaultkey.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultKeyAttributes struct {
	ref terra.Reference
}

// Curve returns a reference to field curve of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) Curve() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("curve"))
}

// E returns a reference to field e of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) E() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("e"))
}

// Id returns a reference to field id of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("id"))
}

// KeyOpts returns a reference to field key_opts of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) KeyOpts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](kvk.ref.Append("key_opts"))
}

// KeySize returns a reference to field key_size of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) KeySize() terra.NumberValue {
	return terra.ReferenceAsNumber(kvk.ref.Append("key_size"))
}

// KeyType returns a reference to field key_type of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) KeyType() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("key_type"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("key_vault_id"))
}

// N returns a reference to field n of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) N() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("n"))
}

// Name returns a reference to field name of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("name"))
}

// PublicKeyOpenssh returns a reference to field public_key_openssh of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) PublicKeyOpenssh() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("public_key_openssh"))
}

// PublicKeyPem returns a reference to field public_key_pem of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) PublicKeyPem() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("public_key_pem"))
}

// ResourceId returns a reference to field resource_id of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("resource_id"))
}

// ResourceVersionlessId returns a reference to field resource_versionless_id of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) ResourceVersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("resource_versionless_id"))
}

// Tags returns a reference to field tags of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kvk.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("version"))
}

// VersionlessId returns a reference to field versionless_id of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("versionless_id"))
}

// X returns a reference to field x of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) X() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("x"))
}

// Y returns a reference to field y of azurerm_key_vault_key.
func (kvk dataKeyVaultKeyAttributes) Y() terra.StringValue {
	return terra.ReferenceAsString(kvk.ref.Append("y"))
}

func (kvk dataKeyVaultKeyAttributes) Timeouts() datakeyvaultkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakeyvaultkey.TimeoutsAttributes](kvk.ref.Append("timeouts"))
}
