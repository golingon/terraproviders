// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datakeyvaultsecret "github.com/golingon/terraproviders/azurerm/3.49.0/datakeyvaultsecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataKeyVaultSecret(name string, args DataKeyVaultSecretArgs) *DataKeyVaultSecret {
	return &DataKeyVaultSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKeyVaultSecret)(nil)

type DataKeyVaultSecret struct {
	Name string
	Args DataKeyVaultSecretArgs
}

func (kvs *DataKeyVaultSecret) DataSource() string {
	return "azurerm_key_vault_secret"
}

func (kvs *DataKeyVaultSecret) LocalName() string {
	return kvs.Name
}

func (kvs *DataKeyVaultSecret) Configuration() interface{} {
	return kvs.Args
}

func (kvs *DataKeyVaultSecret) Attributes() dataKeyVaultSecretAttributes {
	return dataKeyVaultSecretAttributes{ref: terra.ReferenceDataResource(kvs)}
}

type DataKeyVaultSecretArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datakeyvaultsecret.Timeouts `hcl:"timeouts,block"`
}
type dataKeyVaultSecretAttributes struct {
	ref terra.Reference
}

func (kvs dataKeyVaultSecretAttributes) ContentType() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("content_type"))
}

func (kvs dataKeyVaultSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("id"))
}

func (kvs dataKeyVaultSecretAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("key_vault_id"))
}

func (kvs dataKeyVaultSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("name"))
}

func (kvs dataKeyVaultSecretAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("resource_id"))
}

func (kvs dataKeyVaultSecretAttributes) ResourceVersionlessId() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("resource_versionless_id"))
}

func (kvs dataKeyVaultSecretAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](kvs.ref.Append("tags"))
}

func (kvs dataKeyVaultSecretAttributes) Value() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("value"))
}

func (kvs dataKeyVaultSecretAttributes) Version() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("version"))
}

func (kvs dataKeyVaultSecretAttributes) VersionlessId() terra.StringValue {
	return terra.ReferenceString(kvs.ref.Append("versionless_id"))
}

func (kvs dataKeyVaultSecretAttributes) Timeouts() datakeyvaultsecret.TimeoutsAttributes {
	return terra.ReferenceSingle[datakeyvaultsecret.TimeoutsAttributes](kvs.ref.Append("timeouts"))
}
